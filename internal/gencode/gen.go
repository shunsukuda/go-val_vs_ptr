package main

import (
	"os"

	gencode "github.com/shunsukuda/go-gencode"
)

type tmplConfing struct {
	FuncType []string
	Type     []string
	Words    []int
}

var (
	conf = tmplConfing{
		FuncType: []string{"Func", "ValueReciver", "PointerReciver"},
		Type:     []string{"Value", "Pointer"},
		Words:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	}
)

type tmplSet struct {
	Name   string
	Input  string
	Output string
	Config interface{}
}

func main() {
	set := []tmplSet{
		{Name: "Bench", Input: "tmpl_bench_test.go", Output: "gen_bench_test.go", Config: conf},
		{Name: "Bench", Input: "tmpl_bench_global_test.go", Output: "gen_bench_global_test.go", Config: conf},
	}

	PROJECT_PATH := os.Getenv("GOPATH") + "/src/github.com/shunsukuda/go-val_vs_ptr/"
	TEMPLATE_DIR := PROJECT_PATH + "internal/gencode/"
	gencode.DoGoFmt = true

	for i := range set {
		set[i].Input = TEMPLATE_DIR + set[i].Input
		set[i].Output = PROJECT_PATH + set[i].Output
		gencode.GenCode(set[i].Name, set[i].Input, set[i].Output, set[i].Config)
	}

}
