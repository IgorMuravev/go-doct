package document

import doct "doct/internal/document"

// RemoveTransaction is a transaction which removes some data in specified range
type RemoveTransaction struct {
	Position int
	Count    int
}

func (tr *RemoveTransaction) Apply(doc *doct.Document) *doct.Document {
	data := make([]byte, 0, doc.GetSize()-tr.Count)
	copy(data, doc.Data[:tr.Position])
	copy(data[tr.Position:], doc.Data[tr.Position+tr.Count:])
	doc.Data = data
	return doc
}

func (tr *RemoveTransaction) Validate(doc *doct.Document) bool {
	return doc.GetSize() >= tr.Position+tr.Count
}
