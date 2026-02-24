package search

import (
	"context"
	"errors"
	"sort"
	"strings"
)

type service struct {
	engine   SearchEngine
	priority PriorityRepository

	collection string
	perPage    int
}

func NewService(
	engine SearchEngine,
	priority PriorityRepository,
) Service {
	return &service{
		engine:     engine,
		priority:   priority,
		collection: "product_parts",
		perPage:    15,
	}
}

func (s *service) GetSuggestions(
	ctx context.Context,
	query string,
	queryBy string,
) (*Results, error) {

	if s.engine == nil {
		return nil, errors.New("search engine not configured")
	}

	q := strings.TrimSpace(query)
	if q == "" {
		return &Results{}, nil
	}

	token, isPart := extractPartToken(q)

	// ============================
	// 1️⃣ DB Priority Logic
	// ============================

	if isPart && s.priority != nil {

		// Check if this token matches model
		isModel, err := s.priority.GetModelSimilar(ctx, token)
		if err != nil {
			return nil, err
		}

		// If NOT model → treat as part search
		if !isModel {

			partDocs, _ := s.priority.GetPartSuggestions(ctx, token, s.perPage)
			oemDocs, _ := s.priority.GetOemSuggestions(ctx, token, s.perPage)
			vendorDocs, _ := s.priority.GetVendorSuggestions(ctx, token, s.perPage)

			docs := interleaveDocuments(s.perPage, partDocs, oemDocs, vendorDocs)

			hits := make([]Hit, 0, len(docs))
			for _, d := range docs {
				hits = append(hits, Hit{
					Document: d,
					Score:    100, // priority boost
				})
			}

			return &Results{Hits: hits}, nil
		}
	}

	// ============================
	// 2️⃣ Typesense Search
	// ============================

	if queryBy == "" {
		if isPart {
			queryBy = "part_no"
		} else {
			queryBy = strings.Join(fieldsForCloseness, ",")
		}
	}

	req := SearchRequest{
		Collection: s.collection,
		Query:      q,
		QueryBy:    queryBy,
		PerPage:    s.perPage,
		Page:       1,
	}

	resp, err := s.engine.Search(ctx, req)
	if err != nil {
		return nil, err
	}

	results := make([]Hit, 0, len(resp.Hits))

	for _, h := range resp.Hits {
		doc := mapToDocument(h.Document)

		score := markClosenessScore(h.Highlights)

		if isPart && strings.EqualFold(doc.PartNo, token) {
			score += 50
		}

		results = append(results, Hit{
			Document:   doc,
			Highlights: h.Highlights,
			Score:      score,
		})
	}

	sort.SliceStable(results, func(i, j int) bool {
		return results[i].Score > results[j].Score
	})

	return &Results{Hits: results}, nil
}
