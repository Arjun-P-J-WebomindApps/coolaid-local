package mapping

import (
	"github.com/google/uuid"
	"github.com/webomindapps-dev/coolaid-backend/internal/domain/master/company"
	"github.com/webomindapps-dev/coolaid-backend/internal/generated/graphql/model"
)

// Single
func CompanyToGraphQL(c *company.Company) (*model.Company, error) {
	id, err := uuid.Parse(c.ID)
	if err != nil {
		return nil, err
	}

	return &model.Company{
		ID:     id,
		Name:   c.Name,
		Status: c.Status,
		Image:  c.ImageURL,
	}, nil
}

// List
func CompaniesToGraphQL(list []company.Company) ([]*model.Company, error) {
	out := make([]*model.Company, 0, len(list))

	for _, c := range list {
		cc := c // avoid pointer aliasing
		m, err := CompanyToGraphQL(&cc)
		if err != nil {
			return nil, err
		}
		out = append(out, m)
	}

	return out, nil
}
