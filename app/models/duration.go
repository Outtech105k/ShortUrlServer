package models

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"time"
)

type Duration struct {
	time.Duration
}

var durationRegex = regexp.MustCompile(`^(\d+)([smhd])$`)

// "5m", "10h", "3d"のような形式の文字列を解析し、Duration型に変換する
func (d *Duration) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return fmt.Errorf("stringParse: %w", err)
	}

	matches := durationRegex.FindStringSubmatch(s)
	if len(matches) != 3 {
		return fmt.Errorf("invalid duration format: %s", s)
	}

	value, err := strconv.Atoi(matches[1])
	if err != nil {
		return fmt.Errorf("atoi: %w", err)
	}

	unit := matches[2]
	switch unit {
	case "s":
		d.Duration = time.Duration(value) * time.Second
	case "m":
		d.Duration = time.Duration(value) * time.Minute
	case "h":
		d.Duration = time.Duration(value) * time.Hour
	case "d":
		d.Duration = time.Duration(value) * 24 * time.Hour
	default:
		return fmt.Errorf("invalid duration unit: %s", unit)
	}

	return nil
}
