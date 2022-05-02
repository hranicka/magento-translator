package main

import (
	"github.com/bregydoc/gtranslate"
)

type Translator struct {
	FromLang string
	ToLang   string
}

func (t Translator) Translate(s string) (string, error) {
	return gtranslate.TranslateWithParams(
		s,
		gtranslate.TranslationParams{
			From: t.FromLang,
			To:   t.ToLang,
		},
	)
}
