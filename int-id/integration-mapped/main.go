package main

import intid "github.com/jonnyorman/firead/int-id"

type TestDocument struct {
	Prop1 string
	Prop2 int
	Prop3 string
}

type MessageData struct {
	Prop1 string `json:"prop1"`
	Prop2 int    `json:"prop2"`
}

type TestDocumentToMessageDataMapper struct{}

func (this TestDocumentToMessageDataMapper) Map(from TestDocument) MessageData {
	return MessageData{
		Prop1: from.Prop1,
		Prop2: from.Prop2,
	}
}

func main() {
	mapper := TestDocumentToMessageDataMapper{}

	intid.RunMappedIntId[TestDocument, MessageData](mapper)
}
