package relatedsearch_test

import (
	"fmt"
	"os"
	"search/relatedsearch"
	"testing"
)

func TestMain(m *testing.M) {
	relatedsearch.Init()
	code := m.Run()
	os.Exit(code)
}

func TestAddEnAndSearchEn(t *testing.T) {
	relatedsearch.AddEn("abc")
	relatedsearch.AddEn("abea")
	relatedsearch.AddEn("abeb")
	relatedsearch.AddEn("abcc")
	relatedsearch.AddEn("abcd")
	ans := relatedsearch.SearchAllEn("abe")
	fmt.Println(ans)
}

func TestGetPinYin(t *testing.T) {
	ans := relatedsearch.GetPinYin("中国上海")
	fmt.Println(ans)
	for _, c := range ans[0][0] {
		fmt.Println(c - 'a')
	}
}
