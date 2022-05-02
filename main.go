package main

import (
	"flag"
	"fmt"
	"regexp"
)

var fileNameRegExp = regexp.MustCompile(`([a-z]{2})_[A-Z]{2}\.csv$`)

func main() {
	// load flags
	fromFile := flag.String("from", "en_US.csv", "Source file path. Must specify language, eg. en_US.csv")
	toFile := flag.String("to", "cs_CZ.csv", "Translation file path. Existing file will be overwritten. Must specify language, eg. cs_CZ.csv")
	flag.Parse()

	// validate inputs
	fromLangs := fileNameRegExp.FindStringSubmatch(*fromFile)
	if len(fromLangs) == 0 {
		panic("source (from) file name is not valid")
	}
	toLangs := fileNameRegExp.FindStringSubmatch(*toFile)
	if len(toLangs) == 0 {
		panic("target (to) file name is not valid")
	}

	// translate
	texts, err := ReadCSV(*fromFile)
	if err != nil {
		panic(err)
	}
	translations := make([][]string, 0, len(texts))

	tr := &Translator{
		FromLang: fromLangs[1],
		ToLang:   toLangs[1],
	}

	for _, st := range texts {
		tt, err := tr.Translate(st)
		if err != nil {
			panic(err)
		}
		translations = append(translations, []string{st, tt})
		fmt.Printf("%s -> %s\n", st, tt)
	}

	err = WriteCSV(*toFile, translations)
	if err != nil {
		panic(err)
	}

	// finish
	fmt.Println()
	fmt.Println("Done.")
}
