package security

import (
	"github.com/XiBao/topsdk/security/util"
	"regexp"
	"strconv"
	"strings"
)

type SecretData struct {
	OriValue    string
	Base64Value string
	Version     int
	Search      bool
}

func getSecretData(data string, sep rune) *SecretData {
	sdata := &SecretData{}
	vals := strings.Split(strings.Trim(data, string(sep)), string(sep))
	if sep == SEP_CHAR_PHONE {
		if len(vals) != 3 {
			return nil
		}
		version := vals[2]
		v, err := strconv.ParseInt(version, 10, 64)
		if err == nil {
			sdata.OriValue = vals[0]
			sdata.Base64Value = vals[1]
			sdata.Version = int(v)
		}
	} else {
		if len(vals) != 2 {
			return nil
		}
		version := vals[1]
		v, err := strconv.ParseInt(version, 10, 64)
		if err == nil {
			sdata.Base64Value = vals[0]
			sdata.Version = int(v)
		}
	}

	return sdata
}

func getIndexSecretData(data string, sep rune) *SecretData {
	sdata := &SecretData{}
	vals := strings.Split(strings.Trim(data, string(sep)), string(sep))
	if len(vals) != 3 {
		return nil
	}
	version := vals[2]
	v, err := strconv.ParseInt(version, 10, 64)
	if err == nil {
		if sep == SEP_CHAR_PHONE {
			sdata.OriValue = vals[0]
			sdata.Base64Value = vals[1]
		} else {
			sdata.OriValue = vals[1]
			sdata.Base64Value = vals[0]
		}
		sdata.Version = int(v)
	}
	sdata.Search = true
	return sdata
}

func checkEncryptData(data []string) bool {
	if len(data) == 2 {
		return util.IsBase64(data[0])
	} else {
		return util.IsBase64(data[0]) && util.IsBase64(data[1])
	}
}

/*
* 判断是否是加密数据
 */
func isEncryptData(data string, typ string) bool {
	dataLen := len(data)
	if dataLen < 4 {
		return false
	}

	sep, err := get_sep_by_type(typ)
	if err != nil {
		return false
	}
	if []rune(data)[0] != sep || []rune(data)[dataLen-1] != sep {
		return false
	}

	vals := strings.Split(strings.Trim(data, string(sep)), string(sep))
	valCnt := len(vals)
	if sep == SEP_CHAR_PHONE {
		if valCnt != 3 {
			return false
		}

		if []rune(data)[dataLen-2] == sep {
			return checkEncryptData(vals)
		} else {
			version := vals[valCnt-1]
			_, err := strconv.ParseInt(version, 10, 64)
			if err != nil {
				return false
			}
			return util.IsBase64(vals[valCnt-2])
		}
	}
	return (([]rune(data)[dataLen-2] == sep && valCnt == 3) || valCnt == 2) && checkEncryptData(vals)
}

func getDataByType(data string, typ string) *SecretData {
	dataLen := len(data)
	sep, err := get_sep_by_type(typ)
	if err != nil {
		return nil
	}
	if []rune(data)[dataLen-2] == sep {
		return getIndexSecretData(data, sep)
	} else {
		return getSecretData(data, sep)
	}
}

func isPublicData(data string, typ string) bool {
	sdata := getDataByType(data, typ)
	if sdata == nil {
		return false
	}
	return sdata.Version < 0
}

func getSlideWindows(data string, slideSize int) (windows []string) {
	endIndex := 0
	startIndex := 0
	currentWindowSize := 0
	currentWindow := []string{}

	reg := regexp.MustCompile(`(?ms).`)
	utf8_strs := reg.FindAllString(data, -1)

	dataLen := len(utf8_strs)

	for {
		if !(endIndex < dataLen || currentWindowSize > slideSize) {
			break
		}

		var startsWithLetterOrDigit bool
		if len(currentWindow) > 0 {
			startsWithLetterOrDigit = util.IsLetterOrDigit(currentWindow[0])
		}
		if endIndex == dataLen && startsWithLetterOrDigit == false {
			break
		}

		if currentWindowSize == slideSize && startsWithLetterOrDigit == false && util.IsLetterOrDigit(utf8_strs[endIndex]) {
			endIndex++
			currentWindow = utf8_strs[startIndex:endIndex]
			currentWindowSize = 5
		} else {
			if endIndex != 0 {
				if startsWithLetterOrDigit {
					currentWindowSize -= 1
				} else {
					currentWindowSize -= 2
				}
				startIndex++
			}
			for {
				if !(currentWindowSize < slideSize && endIndex < dataLen) {
					break
				}
				currentChar := utf8_strs[endIndex]
				if util.IsLetterOrDigit(currentChar) {
					currentWindowSize += 1
				} else {
					currentWindowSize += 2
				}
				endIndex++
			}

			currentWindow = utf8_strs[startIndex:endIndex]
		}
		windows = append(windows, strings.Join(currentWindow, ""))
	}
	return
}
