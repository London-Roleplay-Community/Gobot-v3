package components

import (
	"fmt"
	"strconv"
	"strings"
)

func HexToBase10(hex string) int {
	hex = strings.TrimPrefix(hex, "#")
	decimal, err := strconv.ParseInt(hex, 16, 64)
	if err != nil {
		fmt.Printf("Could not convert.. %v", err)
	}

	return int(decimal)
}
