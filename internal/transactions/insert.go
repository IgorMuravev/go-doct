package document

import doct "doct/internal/document"

// InsertTransaction is a transaction which inserts some data into a specified position
type InsertTransaction struct {
	Position int
	Data     []byte
}

func (tr *InsertTransaction) Apply(doc *doct.Document) {
	data := make([]byte, doc.GetSize()+len(tr.Data))
	copy(data[0:], doc.Data[:tr.Position])
	copy(data[tr.Position:], tr.Data)
	copy(data[tr.Position+len(tr.Data):], doc.Data[tr.Position:])
	doc.Data = data
}

func (tr *InsertTransaction) Validate(doc *doct.Document) bool {
	return doc != nil &&
		tr.Data != nil &&
		len(tr.Data) > 0 &&
		tr.Position >= 0 &&
		tr.Position <= doc.GetSize()
}
