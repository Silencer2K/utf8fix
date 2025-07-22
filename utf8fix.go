package utf8fix

import (
	"unicode/utf8"
)

// RFC 3629
// https://datatracker.ietf.org/doc/html/rfc3629
//
// Char. number range  |        UTF-8 octet sequence
//    (hexadecimal)    |              (binary)
// --------------------+---------------------------------------------
// 0000 0000-0000 007F | 0xxxxxxx
// 0000 0080-0000 07FF | 110xxxxx 10xxxxxx
// 0000 0800-0000 FFFF | 1110xxxx 10xxxxxx 10xxxxxx
// 0001 0000-0010 FFFF | 11110xxx 10xxxxxx 10xxxxxx 10xxxxxx

const (
	headByte  = 0xC0
	tailByte  = 0x80
	checkMask = headByte
)

// TrimIncomplete trims an incomplete UTF-8 sequence at the end of the string
func TrimIncomplete[S ~string | ~[]byte](s S) (S, bool) {
	size := len(s)
	for size > 0 && s[size-1]&checkMask == tailByte {
		size--
	}
	if size > 0 && s[size-1]&checkMask == headByte {
		size--
	} else if size != len(s) {
		return s, false
	}
	if size == len(s) {
		return s, true
	}
	if len(s)-size > utf8.UTFMax {
		return s, false
	}
	if utf8.Valid([]byte(s)[size:]) {
		return s, true
	}
	return s[:size], true
}
