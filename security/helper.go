package security

import (
	"errors"
	"fmt"
	"github.com/XiBao/topsdk/security/util"
	"strconv"
)

type Helper struct {
	secret    string
	version   int
	appConfig map[string]string
}

func NewHelper(secret string, version int, appConfig map[string]string) *Helper {
	return &Helper{
		secret:    secret,
		version:   version,
		appConfig: appConfig,
	}
}

func (h *Helper) GetAppConfigVal(key string, defaultVal string) string {
	if h.appConfig == nil {
		return defaultVal
	}
	v, ok := h.appConfig[key]
	if !ok {
		v = defaultVal
	}
	return v
}

/*
* 加密逻辑
 */
func (h *Helper) Encrypt(data string, typ string, version int) (string, error) {
	sep, err := get_sep_by_type(typ)
	if err != nil {
		return "", err
	}
	isIndexEncrypt := h.isIndexEncrypt(typ, version)
	if isIndexEncrypt || typ == CRYPT_TYPE_SEARCH {
		if typ == CRYPT_TYPE_PHONE {
			return h.EncryptPhoneIndex(data, sep)
		} else {
			compressLen, _ := strconv.ParseInt(h.GetAppConfigVal("encrypt_index_compress_len", "3"), 10, 64)
			slideSize, _ := strconv.ParseInt(h.GetAppConfigVal("encrypt_slide_size", "4"), 10, 64)
			return h.EncryptNormalIndex(data, int(compressLen), int(slideSize), sep)
		}
	} else {
		if typ == CRYPT_TYPE_PHONE {
			return h.EncryptPhone(data, sep)
		} else {
			return h.EncryptNormal(data, sep)
		}
	}
}

/*
* 解密逻辑
 */
func (h *Helper) Decrypt(data string, cryptType string) (string, error) {
	if !isEncryptData(data, cryptType) {
		return "", fmt.Errorf("数据[%s]不是类型为[%s]的加密数据", data, cryptType)
	}
	dataLen := len(data)
	sep, err := get_sep_by_type(cryptType)
	if err != nil {
		return "", err
	}

	var secretData *SecretData
	if []rune(data)[dataLen-2] == sep {
		secretData = getIndexSecretData(data, sep)
	} else {
		secretData = getSecretData(data, sep)
	}
	if secretData == nil {
		return data, nil
	}

	result, err := util.Decrypt(secretData.Base64Value, h.secret)
	if err != nil {
		return "", err
	}

	if sep == SEP_CHAR_PHONE && !secretData.Search {
		return secretData.OriValue + result, nil
	}

	return result, nil
}

func (h *Helper) Search(data, cryptType string) (string, error) {
	sep, err := get_sep_by_type(cryptType)
	if err != nil {
		return "", err
	}

	if sep == SEP_CHAR_PHONE {
		if len(data) != 4 {
			return "", errors.New("phoneNumber error")
		}
		encryptRet, err := util.HmacMD5EncryptToBase64(data, h.secret, 0)
		if err != nil {
			return "", err
		}
		return string(sep) + encryptRet + string(sep), nil
	} else {
		compressLen, _ := strconv.ParseInt(h.GetAppConfigVal("encrypt_index_compress_len", "3"), 10, 64)
		slideSize, _ := strconv.ParseInt(h.GetAppConfigVal("encrypt_slide_size", "4"), 10, 64)
		slideList := getSlideWindows(data, int(slideSize))

		var builder string
		for _, slide := range slideList {
			v, err := util.HmacMD5EncryptToBase64(slide, h.secret, int(compressLen))
			if err != nil {
				return "", err
			}
			builder += v
		}
		return builder, nil
	}
}

func (h *Helper) isIndexEncrypt(key string, version int) bool {
	if version < 0 {
		key = "previous_" + key
	} else {
		key = "current_" + key
	}

	return h.GetAppConfigVal(key, "") == "2"
}

func (h *Helper) EncryptPhoneIndex(data string, sep rune) (string, error) {
	dataLen := len(data)
	if dataLen < 11 {
		return data, nil
	}

	hMacEncryptRet, err := util.HmacMD5EncryptToBase64(data[dataLen-4:], h.secret, 0)
	if err != nil {
		return "", err
	}

	encryptRet, err := util.Encrypt(data, h.secret)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%c%s%c%s%c%d%c%c", sep, hMacEncryptRet, sep, encryptRet, sep, h.version, sep, sep), nil
}

func (h *Helper) EncryptNormalIndex(data string, compressLen, slideSize int, sep rune) (string, error) {
	slideList := getSlideWindows(data, slideSize)
	builder := ""
	for _, slide := range slideList {
		v, err := util.HmacMD5EncryptToBase64(slide, h.secret, compressLen)
		if err != nil {
			return "", err
		}
		builder += v
	}
	encryptRet, err := util.Encrypt(data, h.secret)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%c%s%c%s%c%d%c%c", sep, encryptRet, sep, builder, sep, h.version, sep, sep), nil
}

// 加密逻辑,手机号码格式
func (h *Helper) EncryptPhone(data string, sep rune) (string, error) {
	dataLen := len(data)
	if dataLen < 11 {
		return data, nil
	}

	prefixNumber := data[0 : dataLen-8]
	last8Number := data[dataLen-8:]
	encryptRet, err := util.Encrypt(last8Number, h.secret)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%c%s%c%s%c%d%c", sep, prefixNumber, sep, encryptRet, sep, h.version, sep), nil
}

// 加密逻辑,非手机号码格式
func (h *Helper) EncryptNormal(data string, sep rune) (string, error) {
	encryptRet, err := util.Encrypt(data, h.secret)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%c%s%c%d%c", sep, encryptRet, sep, h.version, sep), nil
}
