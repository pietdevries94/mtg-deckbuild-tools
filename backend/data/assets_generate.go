// +build ignore

package main

import (
	"log"
	"net/http"

	"github.com/shurcooL/vfsgen"
)

func main() {
	err := vfsgen.Generate(http.Dir("../../frontend/dist"), vfsgen.Options{
		PackageName:  "data",
		VariableName: "frontendFS",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
