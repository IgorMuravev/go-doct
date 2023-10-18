package doct

import (
	"testing"
)

func Test_BlackBox(t *testing.T) {

	doc := NewDocument("test.doct")

	transactions := []ITransaction{
		&AppendTransaction{
			Data: []byte("Hello world [bad]"),
		},

		&EraseTransaction{},

		&AppendTransaction{
			Data: []byte("abc"),
		},

		&AppendTransaction{
			Data: []byte("def"),
		},

		&InsertTransaction{
			Position: 3,
			Data:     []byte("xyz"),
		},

		&RemoveTransaction{
			Position: 4,
			Count:    2,
		},

		&RemoveTransaction{
			Position: 6,
			Count:    0,
		},

		&InsertTransaction{
			Position: 7,
			Data:     []byte("123"),
		},
	}

	want := "abcxdef123"

	for _, t := range transactions {
		doc.Apply(t)
	}

	docstring := string(doc.Data)
	if docstring != want {
		t.Errorf("Result of applies transaction`s list is %v want %v ", docstring, want)
	}
}
