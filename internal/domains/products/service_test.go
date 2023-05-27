package products

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

// TestCategoryService_GetAllCategories functionality
func TestCategoryService_GetAllCategories(t *testing.T) {
	conn, err := setupDbConnection()
	assert.NoError(t, err, "Setting database connection up failed")

	sv := createService(conn)

	_, err = sv.GetAllCategories()
	assert.Error(t, err, "Expected categories not found error")
	assert.ErrorIs(t, err, NoCategoriesFound, "Expected categories not found error")

	createdCats := mockAndInsertCategories(conn, 2)

	fetchedCategories, err := sv.GetAllCategories()
	assert.NoError(t, err, "Fetching Categories from db failed")
	assertCategories(t, createdCats, fetchedCategories)
}

// createService and return it for testing purpose
func createService(db *gorm.DB) CategoryServiceInterface {
	return NewCategoryService(NewCategoryRepo(db))
}