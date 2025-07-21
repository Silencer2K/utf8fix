package utf8fix

import (
	"testing"
)

func TestTrimIncomplete(t *testing.T) {
	cases := []struct {
		name string
		in   string
		out  string
		ok   bool
	}{
		{
			"ascii",
			"test",
			"test",
			true,
		},
		{
			"text_valid",
			"testтест",
			"testтест",
			true,
		},
		{
			"text_partial",
			"testтест"[:11],
			"testтес",
			true,
		},
		{
			"emoji_valid",
			"test\U0001F600\U0001F600",
			"test\U0001F600\U0001F600",
			true,
		},
		{
			"emoji_partial",
			"test\U0001F600\U0001F600"[:11],
			"test\U0001F600",
			true,
		},
		{
			"no_head_byte",
			"test\x80",
			"test\x80",
			false,
		},
		{
			"too_long",
			"test\xC0\x80\x80\x80\x80",
			"test\xC0\x80\x80\x80\x80",
			false,
		},
	}

	for _, c := range cases {
		out, ok := TrimIncomplete(c.in)
		if out != c.out || ok != c.ok {
			t.Errorf("%s: got (%v, %v), expected (%v, %v)",
				c.name, out, ok, c.out, c.ok,
			)
		}
	}
}
