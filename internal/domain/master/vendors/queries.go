package vendor

import (
	"context"
	"sort"
	"strings"

	"github.com/webomindapps-dev/coolaid-backend/internal/domain/master/shared"
)

func (s *Service) List(
	ctx context.Context,
	name *string,
) ([]Vendor, error) {

	rows, err := s.DB.Queries().GetVendorsWithContacts(
		ctx,
		shared.ValueOrEmpty(name),
	)
	if err != nil {
		return nil, ErrInternal
	}

	grouped := make(map[string]*Vendor)

	for _, r := range rows {

		v, ok := grouped[r.VendorID]
		if !ok {
			v = &Vendor{
				ID:             r.VendorID,
				CompanyName:    r.CompanyName,
				VendorContacts: []VendorContact{},
			}
			grouped[r.VendorID] = v
		}

		// LEFT JOIN safety
		if r.ContactPerson == nil {
			continue
		}

		v.VendorContacts = append(v.VendorContacts, VendorContact{
			ContactPerson: *r.ContactPerson,
			MobileNumber:  shared.Deref(r.MobileNumber),
			EmailID:       shared.Deref(r.EmailID),
		})
	}

	out := make([]Vendor, 0, len(grouped))
	for _, v := range grouped {
		if v.VendorContacts == nil {
			v.VendorContacts = []VendorContact{}
		}

		out = append(out, *v)
	}

	// Deterministic ordering
	sort.Slice(out, func(i, j int) bool {
		return strings.ToLower(out[i].CompanyName) <
			strings.ToLower(out[j].CompanyName)
	})

	return out, nil
}
