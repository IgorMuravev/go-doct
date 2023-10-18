package document

import doct "doct/internal/document"

// AppendTransaction is a transaction which append data to the end of a document
type AppendTransaction struct {
	Data []byte
}

func (tr *AppendTransaction) Apply(doc *doct.Document) {
	doc.Data = append(doc.Data, tr.Data...)
}

func (tr *AppendTransaction) Validate(doc *doct.Document) bool {
	return doc != nil &&
		tr.Data != nil &&
		len(tr.Data) > 0
}
