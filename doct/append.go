package doct

// AppendTransaction is a transaction which append data to the end of a document
type AppendTransaction struct {
	Data []byte
}

func (tr *AppendTransaction) Apply(doc *Document) {
	doc.Data = append(doc.Data, tr.Data...)
}

func (tr *AppendTransaction) Validate(doc *Document) bool {
	return doc != nil &&
		tr.Data != nil &&
		len(tr.Data) > 0
}
