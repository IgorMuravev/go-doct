package document

type Document struct {
	Name string
	Data []byte
}

func (doc *Document) Apply(tr ITransaction) *Document {
	doc = tr.Apply(doc)
	return doc
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
