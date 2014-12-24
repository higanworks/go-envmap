package envmap

import (
	"os"
	"strings"
)

func All() map[string]string {
	data := os.Environ()
	items := make(map[string]string)
	for _, val := range data {
		splits := strings.Split(val, "=")
		key := splits[0]
		value := strings.Join(splits[1:], "=")
		items[key] = value
	}
	return items
}

func ListKeys() []string {
	data := os.Environ()
	var keys []string
	for _, val := range data {
		splits := strings.Split(val, "=")
		keys = append(keys, splits[0])
	}
	return keys
}
