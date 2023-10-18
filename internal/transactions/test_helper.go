package document

import doct "doct/internal/document"

func getTestDoct(content string) *doct.Document {
	doc := doct.NewDocument("test.doct")
	doc.Data = []byte(content)
	return doc
}
