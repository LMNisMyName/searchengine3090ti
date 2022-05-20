package main

import (
	"log"
	"search/dal"
	searchapi "search/kitex_gen/SearchApi/search"
	"search/tokenizer"
)

func init() {
	dal.Init()
	tokenizer.Init()
}

func main() {
	svr := searchapi.NewServer(new(SearchImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
