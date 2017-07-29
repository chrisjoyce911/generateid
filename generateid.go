package generateid

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

const charset = "ABCDEFGHIJKLMNPQRSTUVWXYZ123456789"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func fourchar(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Int()%len(charset)]
	}
	return string(b)
}

// GenerateID : return a string looking like 9AGZ-HQ4H-16TR-9HDF
func GenerateID() string {
	var e = strings.Join([]string{fourchar(4, charset), fourchar(4, charset), fourchar(4, charset), fourchar(4, charset)}, "-")
	return string(e)
}
