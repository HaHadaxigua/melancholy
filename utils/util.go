package utils

import (
	"fmt"
	"github.com/google/uuid"
	"reflect"
	"strings"
	"time"
)

const (
	kib = 1024
	mib = 1048576
	gib = 1073741824
	tib = 1099511627776
)

func FormatBytes(i int64) (result string) {
	switch {
	case i >= tib:
		result = fmt.Sprintf("%6.2fTB", float64(i)/tib)
	case i >= gib:
		result = fmt.Sprintf("%6.2fGB", float64(i)/gib)
	case i >= mib:
		result = fmt.Sprintf("%6.2fMB", float64(i)/mib)
	case i >= kib:
		result = fmt.Sprintf("%6.2fKB", float64(i)/kib)
	default:
		result = fmt.Sprintf("%7dB", i)
	}

	if len(result) > 8 {
		result = strings.Join([]string{result[:6], result[7:]}, "")
	}

	return
}

func FormatTime(i int64) string {
	if i < 60 {
		return fmt.Sprintf("%2ds", i)
	} else if i < 3600 {
		s := i % 60
		m := i / 60
		if s == 0 {
			return fmt.Sprintf("%2dm", m)
		} else {
			return fmt.Sprintf("%2dm", m) + FormatTime(s)
		}

	} else {
		s := i % 3600
		h := i / 3600
		if s == 0 {
			return fmt.Sprintf("%2dh", h)
		} else {
			return fmt.Sprintf("%2dh", h) + FormatTime(s)
		}
	}
}

func FormatStringToTime(timeStr string) time.Time {
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr, loc)
	return theTime
}

func GenUUID() string {
	return uuid.New().String()
}

// StructToMap 结构体转map
func StructToMap(in interface{}, tag string) (map[string]interface{}, error) {
	out := make(map[string]interface{})

	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct { // 非结构体返回错误提示
		return nil, fmt.Errorf("ToMap only accepts struct or struct pointer; got %T", v)
	}

	t := v.Type()
	// 遍历结构体字段
	// 指定tagName值为map中key;字段值为map中value
	for i := 0; i < v.NumField(); i++ {
		fi := t.Field(i)
		if tagValue := fi.Tag.Get(tag); tagValue != "" {
			if !v.Field(i).IsZero() {
				out[tagValue] = v.Field(i).Interface()
			}
		}
	}
	return out, nil
}
