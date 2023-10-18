package document

import doct "doct/internal/document"

// EraseTransaction is a transaction which blanks existing document
type EraseTransaction struct {
}

func (tr *EraseTransaction) Apply(doc *doct.Document) {
	doc.Data = make([]byte, 0)
}

func (tr *EraseTransaction) Validate(doc *doct.Document) bool {
	return doc != nil &&
		doc.GetSize() > 0
}
