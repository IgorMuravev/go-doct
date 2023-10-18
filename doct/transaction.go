package doct

type ITransaction interface {
	Apply(doc *Document)
	Validate(doc *Document) bool
}
