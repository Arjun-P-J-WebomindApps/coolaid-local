package search

import "context"

func (s *Service) Index(ctx context.Context, req IndexRequest) error {
	if req.Collection == "" {
		return ErrMissingCollection
	}
	if req.ID == "" {
		return ErrMissingID
	}

	if req.Payload == nil {
		return ErrMissingPayload
	}

	return s.port.Index(ctx, req)
}

func (s *Service) Update(ctx context.Context, req IndexRequest) error {
	if req.Collection == "" {
		return ErrMissingCollection
	}
	if req.ID == "" {
		return ErrMissingID
	}
	if req.Payload == nil {
		return ErrMissingPayload
	}
	return s.port.Update(ctx, req)
}

func (s *Service) Delete(ctx context.Context, collection string, id string) error {
	if collection == "" {
		return ErrMissingCollection
	}
	if id == "" {
		return ErrMissingID
	}

	return s.port.Delete(ctx, collection, id)
}
