package ftaconv

import (
	"math/rand"
	"reflect"
	"time"
	"unsafe"
)

type sliceT struct {
	array unsafe.Pointer
	len   int
	cap   int
}

type stringStruct struct {
	str unsafe.Pointer
	len int
}

func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func StringToBytes(s string) []byte {
	//str := (*stringStruct)(unsafe.Pointer(&s))
	//ret := sliceT{array: unsafe.Pointer(str.str), len: str.len, cap: str.len}
	//return *(*[]byte)(unsafe.Pointer(&ret))
	return S2B(s)
}

// B2S converts byte slice to a string without memory allocation.
// See https://groups.google.com/forum/#!msg/Golang-Nuts/ENgbUzYvCuU/90yGx7GUAgAJ .
func B2S(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// S2B converts string to a byte slice without memory allocation.
//
// Note it may break if string and/or slice header will change
// in the future go versions.
func S2B(s string) (b []byte) {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	bh.Data = sh.Data
	bh.Cap = sh.Len
	bh.Len = sh.Len
	return b
}

const (
	charset        = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	charsetIdxBits = 6                     // 6 bits to represent a charset index
	charsetIdxMask = 1<<charsetIdxBits - 1 // All 1-bits, as many as charsetIdxBits
	charsetIdxMax  = 63 / charsetIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

// RandBytes returns dst with a string random bytes
// Make sure that dst has the length you need.
func RandBytes(dst []byte) []byte {
	n := len(dst)

	for i, cache, remain := n-1, src.Int63(), charsetIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), charsetIdxMax
		}

		if idx := int(cache & charsetIdxMask); idx < len(charset) {
			dst[i] = charset[idx]
			i--
		}

		cache >>= charsetIdxBits
		remain--
	}

	return dst
}

// CopyBytes returns a copy of byte slice in a new pointer.
func CopyBytes(b []byte) []byte {
	return []byte(B2S(b))
}

// EqualBytes reports whether a and b
// are the same length and contain the same bytes.
// A nil argument is equivalent to an empty slice.
func EqualBytes(a, b []byte) bool {
	return B2S(a) == B2S(b)
}

// ExtendBytes extends b to needLen bytes.
func ExtendBytes(b []byte, needLen int) []byte {
	b = b[:cap(b)]
	if n := needLen - cap(b); n > 0 {
		b = append(b, make([]byte, n)...)
	}

	return b[:needLen]
}

// PrependBytes prepends bytes into a given byte slice.
func PrependBytes(dst []byte, src ...byte) []byte {
	dstLen := len(dst)
	srcLen := len(src)

	dst = ExtendBytes(dst, dstLen+srcLen)
	copy(dst[srcLen:], dst[:dstLen])
	copy(dst[:srcLen], src)

	return dst
}

// PrependStringBytes prepends a string into a given byte slice.
func PrependStringBytes(dst []byte, src string) []byte {
	return PrependBytes(dst, S2B(src)...)
}
