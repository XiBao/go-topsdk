package security

import (
	"fmt"
	"github.com/XiBao/taobao"
	"github.com/XiBao/topsdk/security/request"
	"strings"
	"sync"
	"time"
)

type Context struct {
	Secret         string
	Version        int
	InvalidTime    time.Time
	MaxInvalidTime time.Time
	AppConfig      map[string]string

	CacheKey string
	Session  string

	ApiKey *taobao.ApiKey
	Helper *Helper

	// 需要加锁的字段
	LastUploadTime time.Time
	counters       map[string]int
	m              *sync.RWMutex
}

func NewContext(apiKey *taobao.ApiKey) *Context {
	c := &Context{
		AppConfig: make(map[string]string),

		LastUploadTime: time.Now(),
		counters:       make(map[string]int),
		m:              new(sync.RWMutex),
	}
	if apiKey != nil {
		c.ApiKey = apiKey
	}

	c.Helper = NewHelper(c.Secret, c.Version, c.AppConfig)
	return c
}

func NewContextWithTopSecret(apiKey *taobao.ApiKey, session string, topSecret *request.TopSecret) *Context {
	now := time.Now()
	ctx := NewContext(apiKey)
	dd, _ := time.ParseDuration(fmt.Sprintf("%ds", topSecret.MaxInterval))
	ctx.MaxInvalidTime = now.Add(dd)
	dd2, _ := time.ParseDuration(fmt.Sprintf("%ds", topSecret.Interval))
	ctx.InvalidTime = now.Add(dd2)
	ctx.Secret = topSecret.Secret
	ctx.Session = session
	ctx.AppConfig = topSecret.AppConfig
	if session == "" {
		ctx.Version = (-1) * topSecret.Version
	} else {
		ctx.Version = topSecret.Version
	}
	ctx.Helper = NewHelper(ctx.Secret, ctx.Version, ctx.AppConfig)

	return ctx
}
func (c *Context) ToLongString() string {
	ret := []string{
		c.Session,
	}
	c.m.RLock()
	for _, op := range []uint8{OP_ENCRYPT, OP_DECRYPT, OP_SEARCH} {
		for _, t := range []string{CRYPT_TYPE_PHONE, CRYPT_TYPE_NICK, CRYPT_TYPE_RECEIVER_NAME, CRYPT_TYPE_SIMPLE, CRYPT_TYPE_SEARCH} {
			v, ok := c.counters[c.get_stat_key(op, t)]
			if !ok {
				v = 0
			}
			ret = append(ret, fmt.Sprintf("%d", v))
		}
	}
	c.m.RUnlock()
	return strings.Join(ret, ",")
}

func (c *Context) get_stat_key(op uint8, cryptType string) string {
	return fmt.Sprintf("%d_%s", op, cryptType)
}

func (c *Context) IncrCounter(op uint8, cryptType string) {
	key := c.get_stat_key(op, cryptType)
	c.m.Lock()
	v, ok := c.counters[key]
	if !ok {
		v = 1
	} else {
		v++
	}
	c.counters[key] = v
	c.m.Unlock()
}
func (c *Context) CheckUpload() error {
	c.m.RLock()
	lastUploadTime := c.LastUploadTime
	c.m.RUnlock()
	if time.Now().Sub(lastUploadTime) > time.Duration(300*time.Second) {
		if err := c.report(); err != nil {
			return err
		}

		c.clearReport()
	}

	return nil
}

/*
* 上报信息
 */
const (
	APP_SECRET_TYPE      = "2"
	APP_USER_SECRET_TYPE = "3"
)

func (c *Context) report() error {
	req := request.NewTopSdkFeedbackUploadRequest(c.ApiKey)
	req.SetContext(c.ToLongString())
	if c.Session == "" {
		req.SetUploadType(APP_SECRET_TYPE)
	} else {
		req.SetUploadType(APP_USER_SECRET_TYPE)
	}
	_, err := request.TopSdkFeedbackUpload(req)
	return err
}

func (c *Context) clearReport() {
	c.m.Lock()
	c.counters = make(map[string]int)
	c.LastUploadTime = time.Now()
	c.m.Unlock()
}
