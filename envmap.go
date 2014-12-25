package envmap

import (
	"os"
	"strings"
)

var OsEnv = func() []string {
	return os.Environ()
}

func All() map[string]string {
	data := OsEnv()
	items := make(map[string]string)
	for _, val := range data {
		splits := strings.SplitN(val, "=", 2)
		key := splits[0]
		value := splits[1]
		items[key] = value
	}
	return items
}

func ListKeys() []string {
	data := OsEnv()
	var keys []string
	for _, val := range data {
		splits := strings.SplitN(val, "=", 2)
		keys = append(keys, splits[0])
	}
	return keys
}
