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
	Found int         `json:"found"`
	Page  int         `json:"page"`
	Hits  []SearchHit `json:"hits"`
}

type SearchHit struct {
	Document   ProductSearchDocument `json:"document"`
	Highlights []Highlight           `json:"highlights"`
}

type Highlight struct {
	Field         string   `json:"field"`
	Snippet       string   `json:"snippet"`
	MatchedTokens []string `json:"matched_tokens"`
}

//=========================================================
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

//====================================================================
// PART BASED SUGGESTION
//=====================================================================

type PartSuggestionResponse struct {
	PartNo       string `json:"part_no"`
	MatchedValue string `json:"matched_value"`
}
