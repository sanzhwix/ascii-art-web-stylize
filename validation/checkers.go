package validators

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
)

const bannerHash = "c3ec7584fb7ecfbd739e6b3f6f63fd1fe557d2ae3e24f870730d9cf8b2559e94"

func BannerValidity() bool {
	hash := sha256.New()

	content, err := os.ReadFile("banners/standard.txt")
	if err != nil {
		fmt.Println("could not read the banner file")
	}

	hash.Write(content)
	hashCalculated := hex.EncodeToString(hash.Sum(nil))
	return hashCalculated == bannerHash
}

func AsciiCharValidation(text string) bool {
	for _, char := range text {
		if char > 127 {
			fmt.Printf("not valid ASCII character: %q", char)
			return false
		}
	}
	return true
}
