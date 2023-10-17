package document

type ITransaction interface {
	Apply(doc *Document) *Document
	Validate(doc *Document) bool
}
