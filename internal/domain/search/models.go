package search

type IndexRequest struct {
	Collection string
	ID         string
	Payload    any
}

type ProductSearchDocument struct {
	ID       string `json:"id"`
	Company  string `json:"company"`
	Model    string `json:"model"`
	PartNo   string `json:"part_no"`
	Brand    string `json:"brand"`
	Category string `json:"category"`
}
