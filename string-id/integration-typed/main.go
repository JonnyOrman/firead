package main

import stringid "github.com/jonnyorman/firead/string-id"

type TestDocument struct {
	Prop1 string `json:"prop1"`
	Prop2 int    `json:"prop2"`
}

func main() {
	stringid.RunTypedStringId[TestDocument]()
}
