package validators

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
)

const (
	standardHash   = "56d0071a1d7439793953dae6ab3086e1ba4f2947028bc3d6ac4ec475956dff62"
	shadowHash     = "617c65ecec88bb54eeacb75aa6d8920c73bddb15a8d49f1b58cca1d63897034f"
	thinkertoyHash = "64285e4960d199f4819323c4dc6319ba34f1f0dd9da14d07111345f5d76c3fa3"
)

func BannerValidity(name string) bool {
	hash := sha256.New()
	// if name == "" {
	// 	name = "standard"
	// }

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

// func BannerValidity(name string) bool {
// 	hash := sha256.New()

// 	if name == "standard" || name == "shadow" || name == "thinkertoy" {
// 		content, err := os.ReadFile("banners/" + name + ".txt")
// 		if err != nil {
// 			fmt.Println("❌ could not read the banner file:", err)
// 			return false
// 		}
// 		hash.Write(content)
// 		hashCalculated := hex.EncodeToString(hash.Sum(nil))
// 		fmt.Println("🔎 HASH CALCULATED:", hashCalculated)
// 		fmt.Println("🔐 HASH EXPECTED:", returnHash(name))
// 		return hashCalculated == returnHash(name)
// 	}
// 	fmt.Println("⚠️ Unknown banner:", name)
// 	return false
// }

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
