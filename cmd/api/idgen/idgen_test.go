package idgen_test

import (
	"fmt"
	"os"
	"searchengine3090ti/cmd/api/idgen"
	"testing"
)

func TestMain(m *testing.M) {
	idgen.Init()
	code := m.Run()
	os.Exit(code)
}

func TestGetID(t *testing.T) {
	ret := idgen.GetID()
	fmt.Printf("ID = %b", ret)
}
