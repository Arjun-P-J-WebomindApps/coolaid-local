package product

import (
	"context"
	"strings"
)

// isUnicodePart determines whether a row should be classified
// into the "Unicode" bucket based on the Unicode filter.
//
// Business Rule:
// - If Unicode filter is NOT provided → nothing is considered Unicode.
// - If ModelName matches Unicode → goes to Self (NOT Unicode).
// - If ModelName does NOT match Unicode → goes to Unicode.
func isUnicodePart(r FilteredRow, filterUnicode string) bool {

	// No Unicode filter provided → no Unicode classification
	if filterUnicode == "" {
		return false
	}

	// If ModelName does NOT match Unicode → classify as Unicode
	return !strings.EqualFold(
		strings.TrimSpace(r.ModelName),
		filterUnicode,
	)
}

// GetFilterPageDetails retrieves filtered rows,
// separates them into Self and Unicode buckets,
// and applies field-based deduplication.
//
// Unicode classification and deduplication are
// handled at service layer to preserve flexible
// business rules and avoid complex SQL branching.
func (s *Service) GetFilterPageDetails(
	ctx context.Context,
	params FilterSelectionParams,
	self_QueryFields []string,
	unicode_QueryFields []string,
) (FilterSelection, error) {

	rows, err := s.DB.Queries().GetFilteredProducts(ctx, params)
	if err != nil {
		return FilterSelection{
			Self:    []FilterItem{},
			Unicode: []FilterItem{},
		}, ErrInternal
	}

	if len(rows) == 0 {
		return FilterSelection{
			Self:    []FilterItem{},
			Unicode: []FilterItem{},
		}, nil
	}

	self := make([]FilterItem, 0, len(rows))
	unicode := make([]FilterItem, 0, len(rows))

	unicodeVal := ""
	if params.Unicode != nil {
		unicodeVal = strings.TrimSpace(*params.Unicode)
	}

	for _, r := range rows {

		item := FilterItem{
			CompanyName:   r.CompanyName,
			CompanyImage:  r.CompanyImage,
			ModelName:     r.ModelName,
			ModelImage:    r.ModelImage,
			BrandName:     r.BrandName,
			BrandImage:    r.BrandImage,
			CategoryName:  r.CategoryName,
			CategoryImage: r.CategoryImage,
			PartNo:        r.PartNo,
		}

		// Separate rows into "Self" and "Unicode" buckets
		// based on Unicode filter classification rules.
		if isUnicodePart(r, unicodeVal) {
			unicode = append(unicode, item)
		} else {
			self = append(self, item)
		}

	}

	// normalize standardizes strings for stable dedup keys:
	// - trims leading/trailing spaces
	// - collapses internal whitespace
	// - converts to lowercase
	normalize := func(ss []string) []string {
		out := make([]string, 0, len(ss))
		for _, s := range ss {
			out = append(out, strings.ToLower(strings.TrimSpace(s)))
		}
		return out
	}

	selfFields := normalize(self_QueryFields)
	unicodeFields := normalize(unicode_QueryFields)

	// Deduplicate results based on requested GraphQL fields.
	// Only selected fields contribute to the composite dedup key.
	// First occurrence wins.
	self_Deduped, _ := DedupByFieldsFO(self, selfFields)
	unicode_Deduped, _ := DedupByFieldsFO(unicode, unicodeFields)

	return FilterSelection{
		Self:    self_Deduped,
		Unicode: unicode_Deduped,
	}, nil
}

// func (s *Service) GetProductAggregateByPartNo(
// 	ctx context.Context,
// 	partNo string,
// ) (*ProductDetails, error) {

// 	if partNo == "" {
// 		return nil, ErrInvalidPartNo
// 	}

// 	Q := s.DB.Queries()

// 	product, err := s.getProductByPartNo(ctx, Q, partNo)
// 	if err != nil {
// 		return nil, err
// 	}

// 	variant, _ := s.getVariantByPartNo(ctx, Q, partNo)
// 	pricing, _ := s.getPricingByPartNo(ctx, Q, partNo)
// 	inventory, _ := s.getInventoryByPartNo(ctx, Q, partNo)
// 	offer, _ := s.getOfferByPartNo(ctx, Q, partNo)

// 	return &ProductDetails{
// 		Product:      product,
// 		ModelVariant: variant,
// 		Pricing:      pricing,
// 		Inventory:    inventory,
// 		Offer:        offer,
// 	}, nil
// }

func (s *Service) ListProducts(
	ctx context.Context,
	params ProductsFilterParams,
) ([]*ProductDetails, error) {

	Q := s.DB.Queries()

	// ------------------------------------------------
	// 1️⃣ Fetch base aggregate rows
	// ------------------------------------------------

	rows, err := Q.GetProductDetailsList(ctx, params)
	if err != nil {
		return nil, ErrInternal
	}

	if len(rows) == 0 {
		return []*ProductDetails{}, nil
	}

	// ------------------------------------------------
	// 2️⃣ Extract all variants once
	// ------------------------------------------------

	variants := make([]*ModelVariantRow, 0, len(rows))
	for i := range rows {
		variants = append(variants, &rows[i].ModelVariant)
	}

	// ------------------------------------------------
	// 3️⃣ Vendor Bulk Flow (same steps as resolver)
	// ------------------------------------------------

	// Step A: Collect unique IDs
	vendorIDs := collectUniqueVendorIDs(variants)

	// Step B: Fetch from DB
	vendorLookup, err := s.buildVendorLookup(ctx, Q, vendorIDs)
	if err != nil {
		return nil, err
	}

	// Step C: Attach per partNo
	vendorMap := attachVendorsToVariants(variants, vendorLookup)

	// ------------------------------------------------
	// 4️⃣ OEM Bulk Flow (same steps as resolver)
	// ------------------------------------------------

	oemIDs := collectUniqueOEMIDs(variants)

	oemLookup, err := s.buildOEMLookup(ctx, Q, oemIDs)
	if err != nil {
		return nil, err
	}

	oemMap := attachOEMToVariants(variants, oemLookup)

	// ------------------------------------------------
	// 5️⃣ Group by PartNo (build aggregate)
	// ------------------------------------------------

	grouped := make(map[string]*ProductDetails)
	for _, r := range rows {

		partNo := r.Product.PartNo

		if _, exists := grouped[partNo]; exists {
			continue
		}

		product := mapResolvedProductRowToModel(&r.Product)
		variant := mapModelVariantRowToModel(&r.ModelVariant)
		pricing := mapPricingRowToModel(&r.Pricing)
		inventory := mapInventoryRowToModel(&r.Inventory)
		offer := mapOfferRowToModel(&r.Offer)

		// Attach relations
		if variant != nil {
			variant.Vendors = vendorMap[partNo]
			variant.OemNumbers = oemMap[partNo]
		}

		grouped[partNo] = &ProductDetails{
			Product:      product,
			ModelVariant: variant,
			Pricing:      pricing,
			Inventory:    inventory,
			Offer:        offer,
		}
	}

	// ------------------------------------------------
	// 6️⃣ Convert map → slice
	// ------------------------------------------------

	result := make([]*ProductDetails, 0, len(grouped))
	for _, v := range grouped {
		result = append(result, v)
	}

	return result, nil
}
