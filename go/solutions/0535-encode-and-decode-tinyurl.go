package solutions

type Codec struct {
}

func CodecConstructor() Codec {
	return Codec{}
}

// Encodes a URL to a shortened URL.
func (this *Codec) Encode(longUrl string) string {
	return longUrl
}

// Decodes a shortened URL to its original URL.
func (this *Codec) Decode(shortUrl string) string {
	return shortUrl
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * url := obj.encode(longUrl);
 * ans := obj.decode(url);
 */
