// Filename: internal/data/filters.go

import (
	"github.com/spector-asael/banking/internal/validator"
)

type Filters struct {
	Page int 
	PageSize int 
}

// ValidateFilters to check to see if the data provided for the filters is valid. 
// We want to make sure that the page number is greater than zero and that the page size is between 1 and 100 (inclusive).
func ValidateFilters(v *validator.Validator, f Filters) {
   v.Check(f.Page > 0, "page", "must be greater than zero")
   v.Check(f.Page <= 500, "page", "must be a maximum of 500")
   v.Check(f.PageSize > 0, "page_size", "must be greater than zero")
   v.Check(f.PageSize <= 100, "page_size", "must be a maximum of 100")
}

// calculate how many records to send back
func (f Filters) limit() int {
    return f.PageSize
}


// calculate the offset so that we remember how many records have
// been sent and how many remain to be sent
func (f Filters) offset() int {
    return (f.Page - 1) * f.PageSize
}
