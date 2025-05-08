package utils

import (
	"crypto/rand"
	"fmt"
	"strings"
)

var (
	ErrNoCharacterSet = fmt.Errorf("no character set selected")
)

// 引数に指定した桁数のランダムな文字列を生成する関数
func MakeRandomStr(digit uint32, useUppercase, useLowercase, useNumbers bool) (string, error) {
	var builder strings.Builder
	builder.Grow(62) // 最大長（a-zA-Z0-9）を予約

	if useLowercase {
		builder.WriteString("abcdefghijklmnopqrstuvwxyz")
	}
	if useUppercase {
		builder.WriteString("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	}
	if useNumbers {
		builder.WriteString("0123456789")
	}

	letters := builder.String()
	if len(letters) == 0 {
		return "", ErrNoCharacterSet
	}

	// 乱数を生成
	b := make([]byte, digit)
	if _, err := rand.Read(b); err != nil {
		return "", fmt.Errorf("generateRandom: %w", err)
	}

	// letters からランダムに取り出して文字列を生成
	result := make([]byte, digit)
	for i, v := range b {
		result[i] = letters[int(v)%len(letters)]
	}
	return string(result), nil
}
