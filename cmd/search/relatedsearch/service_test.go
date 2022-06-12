package relatedsearch_test

import (
	"fmt"
	"os"
	"searchengine3090ti/cmd/search/relatedsearch"
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
	relatedsearch.Add("abe")
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

func TestSearchAll(t *testing.T) {
	relatedsearch.Add("abc")
	relatedsearch.Add("abea")
	relatedsearch.Add("abeb")
	relatedsearch.Add("abe上海")
	relatedsearch.Add("abcd")
	ans := relatedsearch.SearchAll("abe")
	fmt.Println(ans)
}

func TestSearchTopK(t *testing.T) {
	relatedsearch.Add("上海")
	relatedsearch.Add("上海复旦大学")
	relatedsearch.Add("上海复旦")
	relatedsearch.Add("上海交大")
	relatedsearch.Add("上海交大")
	relatedsearch.Add("上海交大")
	relatedsearch.Add("上海交通大学a")
	relatedsearch.Add("上海交通大学b")
	relatedsearch.Add("上海交通大学c")
	relatedsearch.Add("上海交通大学d")
	relatedsearch.Add("上海交通大学e")
	relatedsearch.Add("上海交通大学f")
	relatedsearch.Add("上海交通大学g")
	relatedsearch.Add("上海交通大学a")
	relatedsearch.Add("北京大学")
	ans := relatedsearch.SearchTopK("1", 10)
	fmt.Println(ans)
}

func TestGetPinYinForMix(t *testing.T) {
	s := "1"
	fmt.Println(relatedsearch.GetPinYinForMix(s))
}
