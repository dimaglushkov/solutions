package main

import "strconv"

// source: https://leetcode.com/problems/encode-and-decode-tinyurl/

var count int64

type Codec struct {
	urls map[string]string
}

func Constructor() Codec {
	return Codec{urls: make(map[string]string)}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	strVal := strconv.FormatInt(count, 10)
	this.urls[strVal] = longUrl
	count++
	return strVal
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	return this.urls[shortUrl]
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * url := obj.encode(longUrl);
 * ans := obj.decode(url);
 */
