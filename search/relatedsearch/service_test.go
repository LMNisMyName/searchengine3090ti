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
	relatedsearch.Add("abc")
	relatedsearch.Add("abea")
	relatedsearch.Add("abeb")
	relatedsearch.Add("abcc")
	relatedsearch.Add("abcd")
	ans := relatedsearch.SearchAll("abe")
	fmt.Println(ans)
}

func TestAddEnAndSearchCh(t *testing.T) {
	relatedsearch.Add("上海")
	relatedsearch.Add("上海交大")
	relatedsearch.Add("上海交通大学")
	relatedsearch.Add("上海复旦大学")
	relatedsearch.Add("北京大学")
	ans := relatedsearch.SearchAll("上海")
	fmt.Println(ans)
}

func TestSearchTopK(t *testing.T) {
	relatedsearch.Add("上海")
	relatedsearch.Add("上海复旦大学")
	relatedsearch.Add("上海复旦大学")
	relatedsearch.Add("上海交大")
	relatedsearch.Add("上海交大")
	relatedsearch.Add("上海交大")
	relatedsearch.Add("上海交通大学")
	relatedsearch.Add("上海交通大学")
	relatedsearch.Add("上海交通大学")
	relatedsearch.Add("上海交通大学")
	relatedsearch.Add("北京大学")
	ans := relatedsearch.SearchTopK("上海", 4)
	fmt.Println(ans)
}

func TestGetPinYin(t *testing.T) {
	ans := relatedsearch.GetPinYin("中国上海")
	fmt.Println(ans)
}
