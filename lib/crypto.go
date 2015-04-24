package crypto

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"time"
)

var cipher = "密钥"

var h = md5.New()

func CipherEncode(sourceText string) string {

	h.Write([]byte(cipher))

	cipherHash := fmt.Sprintf("%x", h.Sum(nil))

	h.Reset()

	inputData := []byte(sourceText)

	loopCount := len(inputData)

	outData := make([]byte, loopCount)

	for i := 0; i < loopCount; i++ {

		outData[i] = inputData[i] ^ cipherHash[i%32]

	}

	return fmt.Sprintf("%s", outData)

}
func Encode(sourceText string) string {

	h.Write([]byte(time.Now().Format("2006-01-02 15:04:05")))

	noise := fmt.Sprintf("%x", h.Sum(nil))

	h.Reset()

	inputData := []byte(sourceText)

	loopCount := len(inputData)

	outData := make([]byte, loopCount*2)

	for i, j := 0, 0; i < loopCount; i, j = i+1, j+1 {

		outData[j] = noise[i%32]

		j++

		outData[j] = inputData[i] ^ noise[i%32]

	}

	return base64.StdEncoding.EncodeToString([]byte(CipherEncode(fmt.Sprintf("%s", outData))))

}
func Decode(sourceText string) string {

	buf, err := base64.StdEncoding.DecodeString(sourceText)

	if err != nil {

		fmt.Println("Decode(%q) failed: %v", sourceText, err)

		return ""

	}

	inputData := []byte(CipherEncode(fmt.Sprintf("%s", buf)))

	loopCount := len(inputData)

	outData := make([]byte, loopCount)

	for i, j := 0, 0; i < loopCount; i, j = i+2, j+1 {

		outData[j] = inputData[i] ^ inputData[i+1]

	}

	return fmt.Sprintf("%s", outData)

}

// func main() {

// 	s := encode("张学友")

// 	fmt.Println(s)

// 	fmt.Println(decode(s))

// }
