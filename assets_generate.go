// +build ignore

/*
 * Copyright (c) 2020. Adam Presley All Rights Reserved
 */

package main

import (
	"log"
	"regexp"

	"github.com/shurcooL/vfsgen"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/js"

	"github.com/adampresley/random-dad-joke/assets"
)

func main() {
	var err error

	m := minify.New()
	m.AddFunc("text/css", css.Minify)
	m.AddFuncRegexp(regexp.MustCompile("^(application|text)/(x-)?(java|ecma)script$"), js.Minify)

	err = vfsgen.Generate(assets.Assets, vfsgen.Options{
		Filename:     "./assets/admincode.go",
		PackageName:  "assets",
		BuildTags:    "!dev",
		VariableName: "Assets",
	})

	if err != nil {
		log.Fatal(err)
	}
}
