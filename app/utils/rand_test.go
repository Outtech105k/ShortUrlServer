package utils_test

import (
	"strings"
	"testing"

	"github.com/Outtech105k/ShortUrlServer/app/utils"
	"github.com/stretchr/testify/assert"
)

func TestMakeRandomStr(t *testing.T) {
	// 基本的な正常系と長さの検証
	t.Run("LengthCheck", func(t *testing.T) {
		lengths := []uint32{0, 1, 10, 100}
		for _, n := range lengths {
			val, err := utils.MakeRandomStr(n, true, true, true)
			assert.NoError(t, err)
			assert.Equal(t, int(n), len(val))
		}
	})

	// 文字セットの制限が正しく機能しているかの検証
	t.Run("CharacterSetValidation", func(t *testing.T) {
		const (
			lower = "abcdefghijklmnopqrstuvwxyz"
			upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
			num   = "0123456789"
		)

		tests := []struct {
			name      string
			digit     uint32
			up, lo, n bool
			allowed   string
		}{
			{"OnlyLower", 100, false, true, false, lower},
			{"OnlyUpper", 100, true, false, false, upper},
			{"OnlyNumbers", 100, false, false, true, num},
			{"UpperAndLower", 100, true, true, false, lower + upper},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				val, err := utils.MakeRandomStr(tt.digit, tt.up, tt.lo, tt.n)
				assert.NoError(t, err)
				for _, r := range val {
					assert.Contains(t, tt.allowed, string(r), "許可されていない文字が含まれています")
				}
			})
		}
	})

	// エラーケースの検証
	t.Run("NoCharacterSetError", func(t *testing.T) {
		_, err := utils.MakeRandomStr(10, false, false, false)
		assert.ErrorIs(t, err, utils.ErrNoCharacterSet)
	})

	// 確率的・統計的なテスト
	t.Run("ProbabilisticTests", func(t *testing.T) {
		// 一意性の検証: 1000回生成して重複がないか
		t.Run("Uniqueness", func(t *testing.T) {
			iterations := 1000
			results := make(map[string]bool)
			for i := 0; i < iterations; i++ {
				val, _ := utils.MakeRandomStr(16, true, true, true)
				assert.False(t, results[val], "ランダム文字列が衝突しました")
				results[val] = true
			}
		})

		// 出現範囲の検証: 十分に長い文字列を生成したとき、全文字が少なくとも一度は出現するか
		t.Run("CharacterCoverage", func(t *testing.T) {
			val, _ := utils.MakeRandomStr(1000, false, true, false)
			alphabet := "abcdefghijklmnopqrstuvwxyz"
			for _, char := range alphabet {
				assert.True(t, strings.Contains(val, string(char)), "文字 '%c' が一度も出現しませんでした", char)
			}
		})
	})
}
