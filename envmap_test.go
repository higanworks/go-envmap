package envmap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAll(t *testing.T) {
	oldosEnv := osEnv
	osEnv = func() []string {
		var hoge []string
		hoge = append(hoge, "HOGE=hogehoge")
		hoge = append(hoge, "PIYO=piyopiyo")
		return hoge
	}

	// assert equality
	assert.Equal(t, All(), map[string]string{"HOGE": "hogehoge", "PIYO": "piyopiyo"}, "they should be equal")
	osEnv = oldosEnv
}

func TestMatched(t *testing.T) {
	oldosEnv := osEnv
	osEnv = func() []string {
		var hoge []string
		hoge = append(hoge, "HOGE=hogehoge")
		hoge = append(hoge, "PIYO=piyopiyo")
		return hoge
	}

	// assert equality
	assert.Equal(t, Matched("^HO"), map[string]string{"HOGE": "hogehoge"}, "they should be equal")
	assert.Equal(t, Matched("YO$"), map[string]string{"PIYO": "piyopiyo"}, "they should be equal")
	assert.Equal(t, Matched(".*O.*"), map[string]string{"HOGE": "hogehoge", "PIYO": "piyopiyo"}, "they should be equal")
	osEnv = oldosEnv
}

func TestListKeys(t *testing.T) {
	oldosEnv := osEnv
	osEnv = func() []string {
		var hoge []string
		hoge = append(hoge, "HOGE=hogehoge")
		hoge = append(hoge, "PIYO=piyopiyo")
		return hoge
	}

	// assert equality
	assert.Equal(t, ListKeys(), []string{"HOGE", "PIYO"}, "they should be equal")
	osEnv = oldosEnv
}
