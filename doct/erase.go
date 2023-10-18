package doct

// EraseTransaction is a transaction which blanks existing document
type EraseTransaction struct {
}

func (tr *EraseTransaction) Apply(doc *Document) {
	doc.Data = make([]byte, 0)
}

func (tr *EraseTransaction) Validate(doc *Document) bool {
	return doc != nil &&
		doc.GetSize() > 0
}
