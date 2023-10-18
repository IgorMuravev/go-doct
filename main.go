package main

import (
	doct "doct/internal/document"
	tr "doct/internal/transactions"
)

func main() {
	doc := doct.NewDocument("test.doct")

	transactions := []doct.ITransaction{
		&tr.AppendTransaction{
			Data: []byte("Hello world [bad]"),
		},

		&tr.EraseTransaction{},

		&tr.AppendTransaction{
			Data: []byte("abc"),
		},

		&tr.AppendTransaction{
			Data: []byte("def"),
		},

		&tr.InsertTransaction{
			Position: 3,
			Data:     []byte("xyz"),
		},

		&tr.RemoveTransaction{
			Position: 4,
			Count:    2,
		},

		&tr.RemoveTransaction{
			Position: 6,
			Count:    0,
		},

		&tr.InsertTransaction{
			Position: 7,
			Data:     []byte("123"),
		},
	}

	println(string(doc.Data))

	for _, t := range transactions {
		err := doc.Apply(t)

		if err != nil {
			panic(err)
		}

		println(string(doc.Data))
	}

	println(string(doc.Data))
}
