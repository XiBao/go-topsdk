package request

import (
	"strconv"
)

// IntAssert 转化数据为int类型
func IntAssert(val interface{}) int {
	switch val.(type) {
	case uint:
		return int(val.(uint))
	case int:
		return val.(int)
	case int64:
		return int(val.(int64))
	case int32:
		return int(val.(int32))
	case int16:
		return int(val.(int16))
	case int8:
		return int(val.(int8))
	case uint8:
		return int(val.(uint8))
	case uint16:
		return int(val.(uint16))
	case uint32:
		return int(val.(uint32))
	case uint64:
		return int(val.(uint64))
	case float64:
		return int(val.(float64))
	case float32:
		return int(val.(float32))
	case nil:
		return int(0)
	case string:
		n, e := strconv.ParseInt(val.(string), 10, 32)
		if e != nil {
			n, e := strconv.ParseFloat(val.(string), 64)
			if e != nil {
				return int(0)
			}
			return int(n)
		}
		return int(n)
	case bool:
		switch val.(bool) {
		case true:
			return int(1)
		case false:
			return int(0)
		}
	}
	return val.(int)
}

// StringAssert 转化数据为string类型
func StringAssert(val interface{}) string {
	switch val.(type) {
	case uint:
		return strconv.FormatUint(uint64(val.(uint)), 10)
	case int:
		return strconv.FormatInt(int64(val.(int)), 10)
	case int64:
		return strconv.FormatInt(val.(int64), 10)
	case int32:
		return strconv.FormatInt(int64(val.(int32)), 10)
	case int16:
		return strconv.FormatInt(int64(val.(int16)), 10)
	case int8:
		return strconv.FormatInt(int64(val.(int8)), 10)
	case uint8:
		return strconv.FormatUint(uint64(val.(uint8)), 10)
	case uint16:
		return strconv.FormatUint(uint64(val.(uint16)), 10)
	case uint32:
		return strconv.FormatUint(uint64(val.(uint32)), 10)
	case uint64:
		return strconv.FormatUint(val.(uint64), 10)
	case float64:
		return strconv.FormatFloat(val.(float64), 'f', 'f', 64)
	case float32:
		return strconv.FormatFloat(float64(val.(float32)), 'f', 'f', 32)
	case nil:
		return ""
	case []byte:
		return string(val.([]byte))
	case string:
		return val.(string)
	case bool:
		return strconv.FormatBool(val.(bool))
	}
	return val.(string)
}
