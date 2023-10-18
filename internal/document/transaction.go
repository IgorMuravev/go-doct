package document

type ITransaction interface {
	Apply(doc *Document)
	Validate(doc *Document) bool
}
