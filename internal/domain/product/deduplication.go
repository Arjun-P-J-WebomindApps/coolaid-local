package product

import (
	"strings"
)

// normalize trims, collapses internal spaces, and lowercases for stable keys.
func normalize(s string) string {
	s = strings.TrimSpace(s)
	s = strings.Join(strings.Fields(s), " ")
	return strings.ToLower(s)
}

/*
get a list of values based on available keys and make a string separated by |

compositeKeyFO builds a stable composite key usingonly the requested fields.
This enables dynamic deduplication depending on GraphQL selection.
Unknown fields contribute an empty segment to
preserve positional consistency in the key.
*/
func compositeKeyFO(fo FilterItem, fields []string) string {
	if len(fields) == 0 {
		fields = []string{"category_name", "model_name", "company_name", "brand_name"}
	}
	parts := make([]string, 0, len(fields))
	for _, f := range fields {
		switch f {
		case "category_name":
			parts = append(parts, normalize(fo.CategoryName))
		case "model_name":
			parts = append(parts, normalize(fo.ModelName))
		case "company_name":
			parts = append(parts, normalize(fo.CompanyName))
		case "brand_name":
			parts = append(parts, normalize(fo.BrandName))
		case "part_no":
			parts = append(parts, normalize(fo.PartNo))
		default:
			// keep position; unknown field contributes empty segment
			parts = append(parts, "")
		}
	}
	return strings.Join(parts, "|")
}

// DedupByFieldsFO returns a unique slice (first occurrence wins) and the set of keys.
func DedupByFieldsFO(items []FilterItem, fields []string) ([]FilterItem, map[string]struct{}) {
	/*map for elements seen*/
	seen := make(map[string]struct{}, len(items))

	//output
	out := make([]FilterItem, 0, len(items))

	for _, it := range items {
		k := compositeKeyFO(it, fields) //Create string based on available fields
		if _, ok := seen[k]; ok {       //if found go to next
			continue
		}
		seen[k] = struct{}{}  //if not found add it to seen list
		out = append(out, it) //add the data to output list
	}
	return out, seen
}
