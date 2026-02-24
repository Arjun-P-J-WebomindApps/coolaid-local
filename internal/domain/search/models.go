package search

//========================================================================
// Store
//=======================================================================

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

//=====================================================================
//     Search
//====================================================================

// Raw request to search engine
type SearchRequest struct {
	Collection string
	Query      string
	QueryBy    string
	PerPage    int
	Page       int
}

// Raw response from search engine
type SearchResponse struct {
	Found int
	Page  int
	Hits  []SearchHit
}

type SearchHit struct {
	Document   map[string]any
	Highlights []Highlight
}

type Highlight struct {
	Field         string   `json:"field"`
	Snippet       string   `json:"snippet"`
	MatchedTokens []string `json:"matched_tokens,omitempty"`
}

type Results struct {
	Hits []Hit `json:"hits"`
}

type Hit struct {
	Document   Document    `json:"document"`
	Highlights []Highlight `json:"highlight,omitempty"`
	Score      float64     `json:"score,omitempty"` // optional internal usage
}

type Document struct {
	ID        string `json:"id"`
	Company   string `json:"company,omitempty"`
	Model     string `json:"model,omitempty"`
	ModelType string `json:"model_type,omitempty"`
	Category  string `json:"category,omitempty"`
	Brand     string `json:"brand,omitempty"`
	PartNo    string `json:"part_no,omitempty"`
	FuelType  string `json:"fuel_type,omitempty"`
	Gen       string `json:"gen,omitempty"`
}

type PartData struct {
	ID       string `json:"id"`
	PartNo   string `json:"part_no,omitempty"`
	OemNo    string `json:"oem_no,omitempty"`
	VendorNo string `json:"vendor_no,omitempty"`
}
