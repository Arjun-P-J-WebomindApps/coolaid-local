package searchrepo

import (
	"context"

	domain_search "github.com/webomindapps-dev/coolaid-backend/internal/domain/search"
	"github.com/webomindapps-dev/coolaid-backend/internal/shared/sqlnull"
)

func (s *searchQueries) GetSimilarModels(
	ctx context.Context,
	token string,
) ([]string, error) {

	rows, err := s.q.GetModelSimilar(ctx, token)
	if err != nil {
		return nil, err
	}

	result := make([]string, 0, len(rows))
	for _, r := range rows {
		result = append(result, r.Name)
	}

	return result, nil
}

func (s *searchQueries) GetPartSuggestions(
	ctx context.Context,
	token string,
) ([]domain_search.PartSuggestionResponse, error) {
	rows, err := s.q.GetProductPartNos(ctx, sqlnull.String(&token))

	if err != nil {
		return nil, err
	}

	result := make([]domain_search.PartSuggestionResponse, 0, len(rows))
	for _, r := range rows {
		result = append(result, domain_search.PartSuggestionResponse{
			PartNo:       r,
			MatchedValue: r,
		})
	}

	return result, nil
}

func (s *searchQueries) GetOemSuggestions(
	ctx context.Context,
	token string,
) ([]domain_search.PartSuggestionResponse, error) {

	rows, err := s.q.SearchOemNumbersByPartialMatch(ctx, sqlnull.String(&token))
	if err != nil {
		return nil, err
	}

	result := make([]domain_search.PartSuggestionResponse, 0, len(rows))
	for _, r := range rows {
		result = append(result, domain_search.PartSuggestionResponse{
			PartNo:       r.PartNo,
			MatchedValue: r.OemNumber,
		})
	}

	return result, nil
}

func (s *searchQueries) GetVendorSuggestions(
	ctx context.Context,
	token string,
) ([]domain_search.PartSuggestionResponse, error) {

	rows, err := s.q.SearchVendorListingsByPartialMatch(ctx, sqlnull.String(&token))
	if err != nil {
		return nil, err
	}

	result := make([]domain_search.PartSuggestionResponse, 0, len(rows))
	for _, r := range rows {
		result = append(result, domain_search.PartSuggestionResponse{
			PartNo:       r.ProductPartNo,
			MatchedValue: r.VendorPartNo,
		})
	}

	return result, nil
}
