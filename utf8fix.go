package utf8fix

import (
	"unicode/utf8"
)

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
