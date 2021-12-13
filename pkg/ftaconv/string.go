package ftaconv

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"time"
)

func ToString(src interface{}) string {
	switch v := src.(type) {
	case string:
		return v
	case []byte:
		return string(v)
	case int:
		return strconv.Itoa(v)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int64:
		return strconv.FormatInt(v, 10)
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 64)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case time.Time:
		return time.Time.Format(v, "2006-01-02 15:04:05")
	case bool:
		return strconv.FormatBool(v)
	default:
		{
			b, _ := json.Marshal(v)
			return string(b)
		}
	}
	return fmt.Sprintf("%v", src)
}

// Reverse 反转字符串
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// IsContainCN 判断字符串是否为中文[精确度需要反复试验]
func IsContainCN(str string) bool {
	var hzRegexp = regexp.MustCompile("[\u4e00-\u9fa5]+")
	return hzRegexp.MatchString(str)
}
