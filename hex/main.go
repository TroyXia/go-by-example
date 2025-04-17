package main

import (
	"encoding/hex"
	"fmt"
	"strings"
)

func encodingString(name string) string {
	if name == "" {
		return ""
	}

	hexStr := hex.EncodeToString([]byte(name))
	if len(hexStr) >= 5 {
		hexStr = hexStr[:5]
	} else {
		lastChar := hexStr[len(hexStr)-1]
		paddingLen := 5 - len(hexStr)
		hexStr += strings.Repeat(string(lastChar), paddingLen)
	}

	return hexStr
}

func main() {
	name := "p"
	fmt.Println(encodingString(name))
}
