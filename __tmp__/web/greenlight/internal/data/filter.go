package data

import (
	"math"
	"strings"

	"greenlight.id/internal/validator"
)

// `Filters` defines the parameters used for paginating, sorting, and filtering results.
type Filters struct {
	Page         int      // The current page number.
	PageSize     int      // The number of records per page.
	Sort         string   // The sorting criteria.
	SortSafeList []string // A list of allowed safe sort parameters to prevent SQL injection or unsafe sorting.
}

// `Metadata` defines the structure used for providing pagination information in the API response.
type Metadata struct {
	CurrentPage  int `json:"current_page,omitempty"`  // The current page number.
	PageSize     int `json:"page_size,omitempty"`     // The number of records per page.
	FirstPage    int `json:"first_page,omitempty"`    // The first available page (usually 1).
	LastPage     int `json:"last_page,omitempty"`     // The last available page.
	TotalRecords int `json:"total_records,omitempty"` // The total number of records available.
}

// `sortColumn` returns the column name to sort by, after verifying it is in the `SortSafeList`.
// If the sort value is not in the list, it panics to prevent unsafe sorting parameters.
func (f Filters) sortColumn() string {
	for _, safeValue := range f.SortSafeList {
		if f.Sort == safeValue {
			// Removes any leading `-` from the sort value (used to indicate descending order).
			return strings.TrimPrefix(f.Sort, "-")
		}
	}

	// Panic if the sort value is not in the safe list, to protect against unsafe sorting.
	panic("unsafe sort parameter: " + f.Sort)
}

// `sortDirection` returns "DESC" if the sort value starts with `-`, indicating descending order.
// Otherwise, it returns "ASC" for ascending order.
func (f Filters) sortDirection() string {
	if strings.HasPrefix(f.Sort, "-") {
		return "DESC"
	}

	return "ASC"
}

// `limit` returns the number of records to retrieve per page based on the `PageSize`.
func (f Filters) limit() int {
	return f.PageSize
}

// `offset` calculates the offset for pagination (i.e., how many records to skip based on the current page).
// This is used in SQL queries to implement pagination.
func (f Filters) offset() int {
	return (f.Page - 1) * f.PageSize
}

// `calculateMetadata` takes the total number of records, the current page, and page size, and returns metadata
// about the pagination. It includes the current page, total number of pages, and total records.
func calculateMetadata(totalRecords, page, pageSize int) Metadata {
	if totalRecords == 0 {
		// Return an empty `Metadata` struct if there are no records.
		return Metadata{}
	}

	// Calculate the last page using the total number of records and page size.
	return Metadata{
		CurrentPage:  page,
		PageSize:     pageSize,
		FirstPage:    1,                                                         // The first page is always 1.
		LastPage:     int(math.Ceil(float64(totalRecords) / float64(pageSize))), // Calculate last page.
		TotalRecords: totalRecords,                                              // Total number of records available.
	}
}

// `ValidateFilters` performs validation on the `Filters` struct using the `validator` package.
// It checks that the page and page size are within allowed limits and that the sort parameter is valid.
func ValidateFilters(v *validator.Validator, f Filters) {
	v.Check(f.Page > 0, "page", "must be greater than zero")                       // The page number must be greater than zero.
	v.Check(f.Page <= 10_000_000, "page", "must be a maximum of 10 million")       // The page number must not exceed 10 million.
	v.Check(f.PageSize > 0, "page_size", "must be greater than zero")              // The page size must be greater than zero.
	v.Check(f.PageSize <= 100, "page_size", "must be a maximum of 100")            // The page size must not exceed 100.
	v.Check(validator.In(f.Sort, f.SortSafeList...), "sort", "invalid sort value") // Validate that the sort value is in the safe list.
}
