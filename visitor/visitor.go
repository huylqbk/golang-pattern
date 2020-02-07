package main

import "fmt"

//double dispatch include accept and visit
func main() {
	var lang Language
	lang = English{}
	visitor := TextLanguage{
		Text: "Huy",
	}
	lang.accept(visitor)
	lang = Vietnamese{}
	lang.accept(visitor)
}

type Language interface {
	accept(LangVisitor)
}

type LangVisitor interface {
	visit(lang Language)
}

type English struct{}
type Vietnamese struct{}

func (eng English) accept(v LangVisitor) {
	v.visit(eng)
}

func showTextEnglish(t string) {
	fmt.Println("I love you", t)
}

func (vn Vietnamese) accept(v LangVisitor) {
	v.visit(vn)
}

func showTextVietNam(t string) {
	fmt.Println("Toi yeu ban,", t)
}

type TextLanguage struct {
	Text string
}

//alter for override multiple implement func visit
func (t TextLanguage) visit(lang Language) {
	switch lang.(type) {
	case English:
		showTextEnglish(t.Text)
		break
	case Vietnamese:
		showTextVietNam(t.Text)
		break
	}
}
