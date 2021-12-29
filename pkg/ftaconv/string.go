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

// ReverseString 反转字符串
func ReverseString(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// IsContainCNString 判断字符串是否为中文[精确度需要反复试验]
func IsContainCNString(str string) bool {
	var hzRegexp = regexp.MustCompile("[\u4e00-\u9fa5]+")
	return hzRegexp.MatchString(str)
}

// CopyString returns a copy of string in a new pointer.
func CopyString(s string) string {
	return string(S2B(s))
}

// CopySliceString returns a copy of the slice.
func CopySliceString(slice []string) []string {
	dst := make([]string, len(slice))
	copy(dst, slice)

	return dst
}

// IndexOfString returns index position in slice from given string
// If value is -1, the string does not found.
func IndexOfString(slice []string, s string) int {
	for i, v := range slice {
		if v == s {
			return i
		}
	}

	return -1
}

// IncludeString returns true or false if given string is in slice.
func IncludeString(slice []string, s string) bool {
	return IndexOfString(slice, s) != -1
}

// UniqueAppendString appends a string if not exist in the slice.
func UniqueAppendString(slice []string, s ...string) []string {
	for i := range s {
		if IndexOfString(slice, s[i]) != -1 {
			continue
		}

		slice = append(slice, s[i])
	}

	return slice
}

// EqualSlicesString checks if the slices are equal.
func EqualSlicesString(slice1, slice2 []string) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}

	return true
}

// ReverseSliceString reverses a string slice.
func ReverseSliceString(slice []string) []string {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}

	return slice
}
