package search

//==================================================
//       INDEXING MODELS
//==================================================
type IndexRequest struct {
	Collection string
	ID         string
	Payload    interface{}
}

type ProductSearchDocument struct {
	ID       string `json:"id"`
	Company  string `json:"company"`
	Model    string `json:"model"`
	PartNo   string `json:"part_no"`
	Brand    string `json:"brand"`
	Category string `json:"category"`
}

//==================================================
//			SEARCH MODELS
//==================================================
type SearchRequest struct {
	Collection string
	Query      string
	QueryBy    string
	Page       int
	PerPage    int
}

type SearchResponse struct {
	Found int
	Page  int
	Hits  []SearchHit
}

type SearchHit struct {
	Document   interface{}
	Highlights []Highlight
}

type Highlight struct {
	Field         string
	Snippet       string
	MatchedTokens []string
}

//========================================================
//          SUGGEST MODELS
//==========================================================

type SuggestRequest struct {
	Collection string
	Query      string
	QueryBy    string
}

type SuggestResponse struct {
	Hits []SearchHit
}
