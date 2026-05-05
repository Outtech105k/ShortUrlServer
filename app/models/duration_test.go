package models_test

import (
	"testing"
	"time"

	"github.com/Outtech105k/ShortUrlServer/app/models"
	"github.com/stretchr/testify/assert"
)

func TestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		out     time.Duration
		wantErr bool
	}{
		// 正常系
		{"10seconds", `"10s"`, 10 * time.Second, false},
		{"20minutes", `"20m"`, 20 * time.Minute, false},
		{"5hours", `"5h"`, 5 * time.Hour, false},
		{"2days", `"2d"`, 2 * 24 * time.Hour, false},

		// 異常系: フォーマット不正
		{"invalid unit", `"10x"`, 0, true},
		{"no unit", `"10"`, 0, true},
		{"no value", `"s"`, 0, true},
		{"empty string", `""`, 0, true},
		{"space included", `"10 s"`, 0, true},

		// 異常系: JSON型不正
		{"numeric value", `10`, 0, true},
		{"boolean value", `true`, 0, true},
		{"object value", `{"val":"10s"}`, 0, true},
		{"null", `null`, 0, true},

		// 異常系: 巨大な数値による Atoi 失敗
		{"huge value", `"99999999999999999999999999s"`, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := models.Duration{}
			err := d.UnmarshalJSON([]byte(tt.in))

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.out, d.Duration)
			}
		})
	}
}
