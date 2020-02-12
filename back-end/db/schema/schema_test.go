package schema

import (
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	user Users
}

func (s *Suite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open("postgres", db)
	require.NoError(s.T(), err)

	s.DB.LogMode(true)
}

func (s *Suite) TestUserGet() {
	var (
		id        uint32 = 1
		firstName string = "Bob"
		lastName  string = "Ross"
		role      string = "User"
		email     string = "bobbyrossy@ross.ross"
	)

	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "users" WHERE (id = $1)`)).
		WithArgs(string(id)).
		WillReturnRows(sqlmock.NewRows([]string{
			"id",
			"firstName",
			"lastName",
			"role",
			"email",
		}).
			AddRow(string(1), firstName, lastName, role, email))

	user := Users{}
	res, err := user.FindUserByID(s.DB, 1)
}

func TestMigrate(t *testing.T) {
	// Create a dummy DB and run Migrate(db)

	// Check that all tables migrate should create exist.
}

func TestFinderUserByID(t *testing.T) {

}
