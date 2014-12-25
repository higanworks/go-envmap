package envmap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAll(t *testing.T) {
	oldOsEnv := OsEnv
	OsEnv = func() []string {
		var hoge []string
		hoge = append(hoge, "HOGE=hogehoge")
		return hoge
	}

	// assert equality
	assert.Equal(t, All(), map[string]string{"HOGE": "hogehoge"}, "they should be equal")
	OsEnv = oldOsEnv
}

func TestListKeys(t *testing.T) {
	oldOsEnv := OsEnv
	OsEnv = func() []string {
		var hoge []string
		hoge = append(hoge, "HOGE=hogehoge")
		return hoge
	}

	// assert equality
	assert.Equal(t, ListKeys(), []string{"HOGE"}, "they should be equal")
	OsEnv = oldOsEnv
}
