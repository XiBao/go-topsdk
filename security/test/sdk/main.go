package main

import (
	"flag"
	"fmt"
	"github.com/XiBao/taobao"
	"github.com/XiBao/topsdk"
	"github.com/XiBao/topsdk/security"
	"github.com/bububa/goconfig/config"
	log "github.com/kdar/factorlog"
	"math/rand"
	"os"
	"runtime"
	"time"
)

const (
	_CONFIG_FILE = "/var/code/go/config.cfg"
)

var (
	logFlag = flag.String("log", "", "set log path")
)

var (
	apiKey *taobao.ApiKey
	logger *log.FactorLog
)

func init() {
	cfg, _ := config.ReadDefault(_CONFIG_FILE)

	dspKey, _ := cfg.String("dsp-api", "key")
	dspSecret, _ := cfg.String("dsp-api", "secret")
	apiKey = &taobao.ApiKey{
		Key:    dspKey,
		Secret: dspSecret,
	}

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

	randomNum := "8zGfA6dTSZwEuSe7puzFsP+b2U0PvOhjDEM3fP1Q6hE="
	session := "6101010683a841d3d716879ZZ8cb6a144c66860781f932b347560967"
	client := security.NewClient(topsdk.NewClient(apiKey.Key, apiKey.Secret), randomNum)

	typ := security.CRYPT_TYPE_PHONE
	val := "13834566786"
	fmt.Printf("原文: %s\n", val)
	encryptValue, err := client.Encrypt(val, typ, session, 0)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("加密后:%s\n", encryptValue)
	searchRet, err := client.Search("6786", typ, session, 0)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("search明文：%s  --> %s\n", val, searchRet)

	if client.IsEncryptData(encryptValue, typ) {
		originalValue, err := client.Decrypt(encryptValue, typ, session)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Printf("解密后:%s\n", originalValue)
	}

	typs := []string{
		security.CRYPT_TYPE_NORMAL,
		security.CRYPT_TYPE_NICK,
		security.CRYPT_TYPE_RECEIVER_NAME,
	}
	val2 := "啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊看哦【啊啊啊的"
	for _, typ2 := range typs {
		fmt.Println("==============================TOP================================")
		encty2, err := client.Encrypt(val2, typ2, session, 0)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Printf("%s|明文：%s ---->密文：%s \n", typ2, val2, encty2)

		if client.IsEncryptData(encty2, typ2) {
			originalValue2, err := client.Decrypt(encty2, typ2, session)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Printf("解密后: %s\n", originalValue2)
			searchRet2, err := client.Search(originalValue2, typ2, session, 0)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Printf("search明文：%s --> %s\n", originalValue2, searchRet2)
		} else {
			fmt.Println("不是加密数据")
		}
	}

	s := "xxxuxxxuxxxu"
	typ = security.CRYPT_TYPE_NICK
	session = ""
	encryptNick, err := client.Encrypt(s, typ, session, 0)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("加密后:%s\n", encryptNick)
	searchRet2, err := client.Search("xxxu", typ, session, 0)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("search明文：%s -->%s\n", s, searchRet2)
	if client.IsEncryptData(encryptNick, "nick") {
		originalNick, err := client.Decrypt(encryptNick, typ, session)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Printf("解密后: %s\n", originalNick)
	} else {
		fmt.Printf("不是加密数据: %s\n", encryptNick)
	}
}
