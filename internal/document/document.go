package document

import "errors"

type Document struct {
	Name string
	Data []byte
}

func (doc *Document) Apply(tr ITransaction) error {
	if !tr.Validate(doc) {
		return errors.New("Invalid transaction")
	}

	tr.Apply(doc)
	return nil
}

func (doc *Document) GetSize() int {
	return len(doc.Data)
}

func NewDocument(name string) *Document {
	return &Document{
		Name: name,
		Data: make([]byte, 0),
	}
}
