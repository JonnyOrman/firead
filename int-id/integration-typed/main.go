package main

import intid "github.com/jonnyorman/firead/int-id"

type TestDocument struct {
	Prop1 string `json:"prop1"`
	Prop2 int    `json:"prop2"`
}

func main() {
	intid.RunTypedIntId[TestDocument]()
}
