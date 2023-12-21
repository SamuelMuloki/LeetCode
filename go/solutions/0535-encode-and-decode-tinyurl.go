package solutions

import (
	"math/rand"
	"time"
)

var asciiCharacters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

type Codec struct {
	url2Code map[string]string
	code2Url map[string]string
}

func CodecConstructor() Codec {
	return Codec{
		url2Code: make(map[string]string),
		code2Url: make(map[string]string),
	}
}

// Encodes a URL to a shortened URL.
func (this *Codec) Encode(longUrl string) string {
	if code, ok := this.url2Code[longUrl]; ok {
		return "http://tinyurl.com/" + code
	}

	code := randStr()
	this.url2Code[longUrl] = code
	this.code2Url[code] = longUrl
	return "http://tinyurl.com/" + code
}

// Decodes a shortened URL to its original URL.
func (this *Codec) Decode(shortUrl string) string {
	code := shortUrl[len(shortUrl)-6:]
	if url, ok := this.code2Url[code]; ok {
		return url
	}

	return ""
}

func randStr() string {
	rand.Seed(time.Now().UnixNano())
	res := make([]byte, 6)
	for i := range res {
		random := rand.Intn(len(asciiCharacters))
		res[i] = asciiCharacters[random]
	}
	return string(res)
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * url := obj.encode(longUrl);
 * ans := obj.decode(url);
 */
