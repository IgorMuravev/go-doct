package document

import doct "doct/internal/document"

// InsertTransaction is a transaction which inserts some data into a specified position
type InsertTransaction struct {
	Position int
	Data     []byte
}

func (tr *InsertTransaction) Apply(doc *doct.Document) *doct.Document {
	data := make([]byte, 0, doc.GetSize()+len(tr.Data))
	copy(data, doc.Data[:tr.Position])
	copy(data[tr.Position:], tr.Data)
	copy(data[len(data):], doc.Data[tr.Position:])
	doc.Data = data
	return doc
}

func (tr *InsertTransaction) Validate(doc *doct.Document) bool {
	return tr.Position <= doc.GetSize()
}
