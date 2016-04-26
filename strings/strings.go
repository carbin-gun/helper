package strings

import (
	"strconv"
	"strings"
)

//Int64 try to parse a string val to int64 value.if the parse fail,return the defaultVal
func Int64(src string, defaultVal int64) int64 {
	if strings.TrimSpace(src) == "" {
		return defaultVal
	}

	val, err := AtoI64(src)
	if err != nil {
		return defaultVal
	}
	return val

}

//AtoI64 try to parse the src string to int64
func AtoI64(src string) (int64, error) {
	i64, err := strconv.ParseInt(src, 10, 0)
	return i64, err
}
