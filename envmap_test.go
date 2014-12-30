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
		hoge = append(hoge, "PIYO=piyopiyo")
		return hoge
	}

	// assert equality
	assert.Equal(t, All(), map[string]string{"HOGE": "hogehoge", "PIYO": "piyopiyo"}, "they should be equal")
	OsEnv = oldOsEnv
}

func TestListKeys(t *testing.T) {
	oldOsEnv := OsEnv
	OsEnv = func() []string {
		var hoge []string
		hoge = append(hoge, "HOGE=hogehoge")
		hoge = append(hoge, "PIYO=piyopiyo")
		return hoge
	}

	// assert equality
	assert.Equal(t, ListKeys(), []string{"HOGE", "PIYO"}, "they should be equal")
	OsEnv = oldOsEnv
}
