//go:build gofuzz
// +build gofuzz

package slackmarkdown

// Fuzz is to be used by https://github.com/dvyukov/go-fuzz
func Fuzz(data []byte) int {
	Parse(data, nil)
	return 0
}
