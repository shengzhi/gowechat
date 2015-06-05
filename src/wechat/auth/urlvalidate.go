package auth

import(
	"sort"
	"crypto/sha1"
	"strings"
	"io"
	"fmt"
)

const(
	token = "wechat4go"
)

func makeSinature(nonce,timestamp string) string{
	params := []string{token,nonce,timestamp}
	sort.Strings(params)
	s1 := sha1.New()
	io.WriteString(s1,strings.Join(params,""))
	return fmt.Sprintf("%x",s1.Sum(nil))
}

func ValidateUrl(nonce,timestamp,clientSinature string) bool{
	serverSinature := makeSinature(nonce,timestamp)
	return serverSinature == clientSinature
}