package stringtool

import "testing"

type tests struct {
	haystack string
	needle   string
	index    int
}

// unit test
func TestStrStr(t *testing.T) {
	lists := []tests{
		{
			haystack: "hello",
			needle:   "lo",
			index:    3,
		},
		{
			haystack: "world",
			needle:   "or",
			index:    1,
		},
		{
			haystack: "golang",
			needle:   "lang",
			index:    2,
		},
	}
	for _, v := range lists {
		if actual := StrStr(v.haystack, v.needle); v.index != actual {
			t.Errorf("string find failed! haystack is : %q, needle is : %q, expected index is : %d, but got index is %d\n",
				v.haystack, v.needle, v.index, actual)
		}
	}

}
