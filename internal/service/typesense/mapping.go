package typesense

import (
	"github.com/typesense/typesense-go/typesense/api"
	"github.com/webomindapps-dev/coolaid-backend/internal/domain/search"
	"github.com/webomindapps-dev/coolaid-backend/internal/shared/ptr"
)

func mapHit(h api.SearchResultHit) search.SearchHit {
	return search.SearchHit{
		Document:   mapDocument(h.Document),
		Highlights: mapHighlights(h.Highlights),
	}
}

func mapDocument(doc *map[string]any) search.ProductSearchDocument {
	if doc == nil {
		return search.ProductSearchDocument{}
	}

	m := *doc

	return search.ProductSearchDocument{
		ID:       getString(m, "id"),
		Company:  getString(m, "company"),
		Model:    getString(m, "model"),
		PartNo:   getString(m, "part_no"),
		Brand:    getString(m, "brand"),
		Category: getString(m, "category"),
	}
}

func mapHighlights(hls *[]api.SearchHighlight) []search.Highlight {
	if hls == nil {
		return nil
	}

	out := make([]search.Highlight, 0, len(*hls))

	for _, hl := range *hls {
		field := ptr.String(hl.Field)
		if field == "" {
			continue
		}

		out = append(out, search.Highlight{
			Field: field,
		})
	}

	return out
}

func getString(m map[string]interface{}, key string) string {
	if v, ok := m[key].(string); ok {
		return v
	}
	return ""
}
