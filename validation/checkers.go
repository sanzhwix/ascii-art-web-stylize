package validators

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
)

const (
	standardHash   = "c3ec7584fb7ecfbd739e6b3f6f63fd1fe557d2ae3e24f870730d9cf8b2559e94"
	shadowHash     = "26b94d0b134b77e9fd23e0360bfd81740f80fb7f6541d1d8c5d85e73ee550f73"
	thinkertoyHash = "64285e4960d199f4819323c4dc6319ba34f1f0dd9da14d07111345f5d76c3fa3"
)

func BannerValidity(name string) bool {
	hash := sha256.New()

	if name == "standard" || name == "shadow" || name == "thinkertoy" {
		content, err := os.ReadFile("banners/" + name + ".txt")
		if err != nil {
			fmt.Println("could not read the banner file")
		}
		hash.Write(content)
		hashCalculated := hex.EncodeToString(hash.Sum(nil))
		return hashCalculated == returnHash(name)
	}
	return false
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

func returnHash(name string) string {
	switch name {
	case "standard":
		return standardHash
	case "shadow":
		return shadowHash
	case "thinkertoy":
		return thinkertoyHash
	}
	return "Hash not found"
}
