package generateid

import (
	"math/rand"
	"strings"
)

const charset = "ABCDEFGHIJKLMNPQRSTUVWXYZ123456789"

func randomchar(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Int()%len(charset)]
	}
	return string(b)
}

// GenerateID : return a string looking like 9AGZ-HQ4H-16TR-9HDF
func GenerateID() string {
	var e = strings.Join([]string{randomchar(4, charset), randomchar(4, charset), randomchar(4, charset), randomchar(4, charset)}, "-")
	return string(e)
}
