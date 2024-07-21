package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	fmt.Println("crypter?")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	in := scanner.Bytes()
	fmt.Println("key")
	scanner.Scan()
	keys := KeyParse(scanner.Text(), len(in))
	ans := base64.StdEncoding.EncodeToString(Encrypt(keys, in))
	fmt.Println(ans)

	scanner.Scan()
	base64In := scanner.Text()
	fmt.Println("key")
	scanner.Scan()
	dkey := scanner.Text()
	decodeIn, _ := base64.StdEncoding.DecodeString(base64In)
	keys = KeyParse(string(dkey), len(decodeIn))
	fmt.Println(string(Decrypt(keys, decodeIn)))
}

func Encrypt(keys []int, in []byte) []byte {
	for _, k := range keys {
		for i := 0; i < k; i++ {
			in[i] = in[i] ^ byte(k)
		}
	}
	for _, k := range keys {
		for i := len(in) - 1; i >= k; i-- {
			in[i] = in[i] ^ byte(k)
		}
	}

	return in
}

func Decrypt(keys []int, in []byte) []byte {
	var rekey []int
	for i := len(keys) - 1; i >= 0; i-- {
		rekey = append(rekey, keys[i])
	}

	for _, k := range rekey {
		for i := len(in) - 1; i >= k; i-- {
			in[i] = in[i] ^ byte(k)
		}
	}
	for _, k := range rekey {
		for i := 0; i < k; i++ {
			in[i] = in[i] ^ byte(k)
		}
	}

	return in
}

func KeyParse(key string, i int) []int {
	tkey := []byte(key)
	var keys []int
	for _, b := range tkey {
		keys = append(keys, int(b)%i)
	}
	return keys
}
