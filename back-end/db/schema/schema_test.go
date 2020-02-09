package schema

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repository Repository
	person     *model.Person
}

// A function of struct.
func (s *Suite) SetupSuite() {
	var (
		db *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err gorm.Open("postgres", db)
	require.NoError(s.T(), err)

	s.DB.Logmode(true)

	s.repository = CreateRepository(s.DB)
}

func TestMigrate(t *testing.T) {
	// Create a dummy DB and run Migrate(db)

	// Check that all tables migrate should create exist.
}
