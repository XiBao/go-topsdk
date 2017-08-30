package security

import (
	"fmt"
	"github.com/XiBao/taobao"
	"github.com/XiBao/topsdk"
	"github.com/XiBao/topsdk/security/request"
	"log"
	"strconv"
	"time"
)

type Client struct {
	topClient *topsdk.Client
	randomNum string

	cache ICache
}

func NewClient(topClient *topsdk.Client, randomNum string) *Client {
	c := &Client{
		topClient: topClient,
		randomNum: randomNum,
	}

	return c
}

/**
* 设置缓存处理器
 */
func (c *Client) SetCache(cache ICache) {
	c.cache = cache
}

/**
* 加密单条数据
 */
func (c *Client) Encrypt(data, typ, session string, version int) (string, error) {
	if data == "" || typ == "" {
		log.Println("Encrypt: empty data or empty typ")
		return data, nil
	}

	ctx := c.getCacheContext(session, 0)
	if ctx == nil {
		log.Println("Encrypt: no secret context")
		return data, nil
	}
	//log.Printf("Encrypt: found secret context: %s", util.Json(ctx))

	c.incrCounter(OP_ENCRYPT, typ, ctx, true)
	return ctx.Helper.Encrypt(data, typ, version)
}

func (c *Client) Decrypt(data, typ, session string) (string, error) {
	if data == "" || typ == "" {
		log.Println("Decrypt: empty data or empty typ")
		return data, nil
	}

	sdata := getDataByType(data, typ)
	if sdata == nil {
		log.Println("Decrypt: no secret data")
		return data, nil
	}
	//log.Printf("Decrypt: found secret data: %s", util.Json(sdata))

	var ctx *Context
	if isPublicData(data, typ) {
		ctx = c.getCacheContext("", sdata.Version)
	} else {
		ctx = c.getCacheContext(session, sdata.Version)
	}
	if ctx == nil {
		log.Println("Decrypt: no secret context")
		return data, nil
	}
	//log.Printf("Decrypt: found secret context: %s", util.Json(ctx))

	c.incrCounter(OP_DECRYPT, typ, ctx, true)

	return ctx.Helper.Decrypt(data, typ)
}

/**
 * 密文检索。 手机号码格式：$base64(H-MAC(phone后4位))$ simple格式：base64(H-MAC(滑窗))
 *
 * @param data
 *            明文数据
 * @param type
 *            加密字段类型(例如：simple\phone)
 * @param session
 *            用户身份,用户级加密必填
 * @param version
 *            秘钥历史版本
 * @return
 */
func (c *Client) Search(data, typ, session string, version int) (string, error) {
	if data == "" || typ == "" {
		return data, nil
	}

	ctx := c.getCacheContext(session, version)
	c.incrCounter(OP_SEARCH, typ, ctx, true)
	if ctx == nil || ctx.Secret == "" {
		return data, nil
	}

	return ctx.Helper.Search(data, typ)
}

/**
* 加密多条数据，必须是同一个type和用户,返回结果是 KV结果
 */
func (c *Client) EncryptBatch(dataList []string, typ, session string, version int) ([]string, error) {
	if len(dataList) == 0 || typ == "" {
		return dataList, nil
	}

	ctx := c.getCacheContext(session, 0)
	if ctx == nil {
		return dataList, nil
	}

	ret := make([]string, len(dataList))
	for k, v := range dataList {
		retV, err := ctx.Helper.Encrypt(v, typ, version)
		if err == nil {
			c.incrCounter(OP_ENCRYPT, typ, ctx, true)
			ret[k] = retV
		} else {
			ret[k] = v
		}
	}

	return ret, nil
}

/**
* 多条数据解密，必须是同一个type和用户,返回结果是 KV结果
* 非加密数据直接返回原文
 */
func (c *Client) DecryptBatch(dataList []string, typ, session string) ([]string, error) {
	if len(dataList) == 0 || typ == "" {
		return nil, nil
	}

	ret := make([]string, len(dataList))
	for k, v := range dataList {
		retV := v
		secretData := getDataByType(v, typ)
		if secretData != nil {
			var ctx *Context
			if isPublicData(v, typ) {
				ctx = c.getCacheContext("", secretData.Version)
			} else {
				ctx = c.getCacheContext(session, secretData.Version)
			}
			if ctx != nil {
				s, err := ctx.Helper.Decrypt(v, typ)
				if err == nil {
					retV = s
					c.incrCounter(OP_DECRYPT, typ, ctx, true)
				}
			}
		}
		ret[k] = retV
	}

	return ret, nil
}

func (c *Client) incrCounter(op uint8, typ string, ctx *Context, flush bool) {
	if ctx == nil {
		return
	}
	ctx.IncrCounter(op, typ)
	if flush && c.cache != nil {
		c.cache.Set(ctx.CacheKey, ctx)
	}
}

/**
* 获取秘钥，使用缓存
 */
func (c *Client) getCacheContext(session string, version int) *Context {
	//log.Printf("getCacheContext: session=%s, version=%d\n", session, version)

	var cacheKey string
	if c.cache != nil {
		cacheKey = c.buildCacheKey(session, version)
		if ctx := c.cache.Get(cacheKey); ctx != nil {
			//log.Printf("found ctx: %s", util.Json(ctx))
			ctx.CheckUpload()

			if ctx.InvalidTime.After(time.Now()) {
				return ctx
			}
		}
	}

	ctx, err := c.getContext(session, version)
	if err != nil {
		log.Printf("Error: %s", err.Error())
		return nil
	}
	if c.cache != nil {
		ctx.CacheKey = cacheKey
		c.cache.Set(cacheKey, ctx)
	}
	return ctx
}

/**
* 获取秘钥，不使用缓存
 */
func (c *Client) getContext(session string, version int) (*Context, error) {
	apiKey := &taobao.ApiKey{Key: c.topClient.AppKey, Secret: c.topClient.SecretKey}
	req := request.NewTopSecretGetRequest(apiKey, session)
	req.SetRandomNum(c.randomNum)
	if version < 0 {
		version = 0
		session = ""
	}
	if version != 0 {
		req.SetSecretVersion(version)
	}
	if session != "" && string(session[0]) == "_" {
		v, _ := strconv.ParseUint(session[1:], 10, 64)
		req.SetCustomerUserId(v)
		session = ""
	}
	req.Session = session
	topSecret, err := request.TopSecretGet(req)
	if err != nil {
		return nil, err
	}

	return NewContextWithTopSecret(apiKey, session, topSecret), nil
}

func (c *Client) InitContext(session string) *Context {
	return c.getCacheContext(session, 0)
}

func (c *Client) GenerateCustomerSession(userId uint64) string {
	return fmt.Sprintf("_%d", userId)
}

func (c *Client) buildCacheKey(session string, version int) string {
	if session == "" {
		return c.topClient.AppKey
	}

	if version <= 0 {
		return session
	}

	return fmt.Sprintf("%s_%d", session, version)
}

/*
* 判断是否是加密数据
 */
func (c *Client) IsEncryptData(data string, typ string) bool {
	return isEncryptData(data, typ)
}
