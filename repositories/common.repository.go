package repositories

import (
	"gorm.io/gorm"
)

func Paginate(query *gorm.DB, limit int, page int) (*gorm.DB, error) {
	if limit == 0 && page == 0 {
		limit = 10
		page = 1
	}
	offset := (page - 1) * limit
	// Set the pagination parameters
	query = query.Offset(offset).Limit(limit)

	// Execute the query and return the result
	return query, query.Error
}
