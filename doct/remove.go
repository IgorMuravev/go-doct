package doct

// RemoveTransaction is a transaction which removes some data in specified range
type RemoveTransaction struct {
	Position int
	Count    int
}

func (tr *RemoveTransaction) Apply(doc *Document) {
	data := make([]byte, doc.GetSize()-tr.Count)
	copy(data, doc.Data[:tr.Position])
	copy(data[tr.Position:], doc.Data[tr.Position+tr.Count:])
	doc.Data = data
}

func (tr *RemoveTransaction) Validate(doc *Document) bool {
	return doc != nil &&
		tr.Position >= 0 &&
		tr.Count > 0 &&
		doc.GetSize() >= tr.Position+tr.Count
}
