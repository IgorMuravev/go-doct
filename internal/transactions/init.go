package document

import doct "doct/internal/document"

// EraseTransaction is a transaction which blanks existing document
type EraseTransaction struct {
}

func (tr *EraseTransaction) Apply(doc *doct.Document) *doct.Document {
	doc = &doct.Document{
		Name: doc.Name,
		Data: make([]byte, 0),
	}

	return doc
}

func (tr *EraseTransaction) Validate(doc *doct.Document) bool {
	return true
}
