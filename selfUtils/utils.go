package selfUtils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

func padding(word *[]byte, size int) {
	last := len(*word) % size
	paddingLen := size - last
	pad := bytes.Repeat([]byte{0}, paddingLen)
	*word = append(*word, pad...)
}

func Crypto(text string) (string, error) {
	key := []byte("天王盖地虎!")
	txt := []byte(text)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	move := key[:blockSize]
	padding(&txt, blockSize)
	mode := cipher.NewCBCEncrypter(block, move)
	encrypted := make([]byte, len(txt))
	mode.CryptBlocks(encrypted, txt)
	return base64.StdEncoding.EncodeToString(encrypted), nil
}
