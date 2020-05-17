package question_531_540

import (
	"encoding/base64"
)

// 535. TinyURL 的加密与解密
// https://leetcode-cn.com/problems/encode-and-decode-tinyurl
// Topics: 哈希表 数学

type Codec struct {
	Data map[string]string
}

func Constructor() Codec {
	return Codec{Data: make(map[string]string)}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	key := base64.URLEncoding.EncodeToString([]byte(longUrl))
	this.Data[key] = longUrl
	return key
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	return this.Data[shortUrl]
}
