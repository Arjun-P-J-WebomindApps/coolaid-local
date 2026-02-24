package search

import "context"

func (s *Service) Suggest(ctx context.Context, req SuggestRequest) (*SuggestResponse, error) {

	// 1️⃣ Validate request
	if err := validateSuggestRequest(req); err != nil {
		return nil, err
	}

	// 2️⃣ Detect numeric-driven token
	token, isPriority := firstNumericDrivenToken(req.Query)

	// 3️⃣ If priority token and should do part search → build part suggestions
	if isPriority && s.shouldDoPartSearch(ctx, token) {
		return s.buildPartSuggestions(ctx, token)
	}

	// 4️⃣ Otherwise fallback to normal search
	return s.fallbackSuggestSearch(ctx, req)
}
