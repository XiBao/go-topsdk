package main

import (
	"flag"
	"github.com/XiBao/common"
	"github.com/XiBao/dbpool"
	"github.com/XiBao/taobao"
	"github.com/XiBao/taobao/api"
	"github.com/XiBao/topsdk"
	"github.com/XiBao/topsdk/security"
	"github.com/XiBao/topsdk/security/util"
	"github.com/bububa/goconfig/config"
	"github.com/bububa/mymysql/autorc"
	_ "github.com/bububa/mymysql/thrsafe"
	log "github.com/kdar/factorlog"
	"math/rand"
	"os"
	"runtime"
	"time"
	"fmt"
)

const (
	_CONFIG_FILE = "/home/steven/var/code/go/config.cfg"
)

var (
	logFlag = flag.String("log", "", "set log path")
)

var (
	apiKey *taobao.ApiKey
	logger *log.FactorLog

	mysqlConfigMaster *dbpool.MySQLConfig

	SecurityClients = make(map[string]*security.Client)
)

func init() {
	cfg, _ := config.ReadDefault(_CONFIG_FILE)

	hostMaster, _ := cfg.String("masterdb", "host")
	userMaster, _ := cfg.String("masterdb", "user")
	passwdMaster, _ := cfg.String("masterdb", "passwd")
	dbnameMaster, _ := cfg.String("masterdb", "dbname")
	mysqlConfigMaster = &dbpool.MySQLConfig{
		Host:   hostMaster,
		User:   userMaster,
		Passwd: passwdMaster,
		DbName: dbnameMaster,
	}

	dspKey, _ := cfg.String("dsp-api", "key")
	dspSecret, _ := cfg.String("dsp-api", "secret")
	apiKey = &taobao.ApiKey{
		Key:    dspKey,
		Secret: dspSecret,
	}
	dspSecurityNum, _ := cfg.String("dsp-api", "securitynum")
	if dspSecurityNum != "" {
		SecurityClients[dspKey] = security.NewClient(topsdk.NewClient(apiKey.Key, apiKey.Secret), dspSecurityNum)
	}
	fmt.Printf("dspKey = %s, dspSecret = %s, dspSecurityNum=%s\n", dspKey, dspSecret, dspSecurityNum)
	rand.Seed(time.Now().UnixNano())
}

func SetGlobalLogger(logPath string) *log.FactorLog {
	sfmt := `%{Color "red:white" "CRITICAL"}%{Color "red" "ERROR"}%{Color "yellow" "WARN"}%{Color "green" "INFO"}%{Color "cyan" "DEBUG"}%{Color "blue" "TRACE"}[%{Date} %{Time}] [%{SEVERITY}:%{ShortFile}:%{Line}] %{Message}%{Color "reset"}`
	logger := log.New(os.Stdout, log.NewStdFormatter(sfmt))
	if len(logPath) > 0 {
		logf, err := os.OpenFile(logPath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0640)
		if err != nil {
			return logger
		}
		logger = log.New(logf, log.NewStdFormatter(sfmt))
	}
	logger.SetSeverities(log.INFO | log.WARN | log.ERROR | log.FATAL | log.CRITICAL)
	return logger
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	flag.Parse()

	logger = SetGlobalLogger(*logFlag)

	mdb := autorc.New("tcp", "", mysqlConfigMaster.Host, mysqlConfigMaster.User, mysqlConfigMaster.Passwd, mysqlConfigMaster.DbName)
	mdb.Register("set names utf8")
	service := &common.Service{
		MasterDb: mdb,
	}

	query := `SELECT
	fo.id AS order_id,
		ass.account_id AS account_id,
		ass.nick AS nick,
		ass.top_session AS top_session,
		ass.app_key AS app_key
	FROM fanli.fanli_orders AS fo
	INNER JOIN xibao_v2.shops AS s ON (s.id = fo.shop_id)
	INNER JOIN xibao_v2.account_sessions AS ass ON (
		ass.nick = s.nick AND ass.r1_expires_at>NOW()
	)
	where fo.created_at >= DATE_ADD(now(),INTERVAL -2 day)
	AND ass.nick = '品瑞商铺'
	GROUP BY ass.account_id
	ORDER BY fo.id LIMIT 1`

	rows, _, err := service.MasterDb.Query(query)
	if err != nil {
		logger.Error(err)
		return
	}
	if len(rows) == 0 {
		logger.Error("no rows")
		return
	}
	logger.Infof("rows = %s", util.Json(rows))
	orderId := rows[0].Uint64(0)
	accountId := rows[0].Uint64(1)
	//nick := rows[0].Str(2)
	session := rows[0].Str(3)

	req := &api.TradeGetRequest{
		Session:   session,
		AccountId: accountId,
		Fields:    "tid,buyer_nick,receiver_name,receiver_city,receiver_district,receiver_address,receiver_zip,receiver_mobile,receiver_phone",
		Tid:       orderId,
		Service:   service,
		AnApiKey:  apiKey,
	}

	trade, err := api.MustGetTradeFullInfo(req)
	if err != nil {
		logger.Error(err)
		trade = &api.Trade{
			Id:   req.Tid,
			IsWt: true,
		}
	}
	logger.Infof("trade = %s", util.Json(trade))

	//randomNum := "8zGfA6dTSZwEuSe7puzFsP+b2U0PvOhjDEM3fP1Q6hE="
	//client := security.NewClient(topsdk.NewClient(apiKey.Key, apiKey.Secret), randomNum)
	client, ok := SecurityClients[apiKey.Key]
	if !ok {
		logger.Errorf("no security client for apiKey: %s", apiKey.Key)
		return
	}
	buyer_nick, err := client.Decrypt(trade.BuyerNick, security.CRYPT_TYPE_NICK, session)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	receiver_name, err := client.Decrypt(trade.ReceiverName, security.CRYPT_TYPE_RECEIVER_NAME, session)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	receiver_mobile, err := client.Decrypt(trade.ReceiverMobile, security.CRYPT_TYPE_PHONE, session)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	logger.Infof("buyer_nick = %s, receiver_name = %s, receiver_mobile = %s", buyer_nick, receiver_name, receiver_mobile)
}
