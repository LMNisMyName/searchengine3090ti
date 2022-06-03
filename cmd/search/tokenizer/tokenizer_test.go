package tokenizer_test

import (
	"fmt"
	"os"
	"searchengine3090ti/cmd/search/tokenizer"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tk *tokenizer.Tokenizer

func TestMain(m *testing.M) {
	tokenizer.Init()
	tk = tokenizer.NewTokenizer("tmp/dictionary.txt")
	os.Exit(m.Run())
	//output := tk.Cut("天气真好")
	//expectOutput := []string{"天气", "真好"}
	//fmt.Println(output, expectOutput)

}

//func TestNewTokenizer(t *testing.T) {
//
//}
func TestCut(t *testing.T) {

	output := tk.Cut("天气真好")
	expectOutput := []string{"天气", "真好"}
	fmt.Println(output, expectOutput)
	assert.Equal(t, output, expectOutput)
}
