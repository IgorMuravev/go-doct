package doct

func getTestDoct(content string) *Document {
	doc := NewDocument("test.doct")
	doc.Data = []byte(content)
	return doc
}
