package schema

import (
	"database/sql"
	"fmt"
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

// Literally this entire time, all I needed to do was
// 	make a regular test function actually start the suite,
//	and pass the suite the test context.  AHHHHHHHHH!!!
func TestRunSuite(t *testing.T) {
	suite.Run(t, new(Suite))
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
		`SELECT * FROM "users" WHERE "users"."deleted_at" IS NULL AND ((id = $1)) LIMIT 1`)).
		// This seems REALLY hacky.  We're using a print method
		// that returns the string it prints to convert a
		// uint into a string.  Revisit.  Also hilarious.
		WithArgs(fmt.Sprint(id)).
		WillReturnRows(sqlmock.NewRows([]string{
			"id",
			"firstName",
			"lastName",
			"role",
			"email",
		}).
			AddRow(fmt.Sprint(0), firstName, lastName, role, email))

	s.T().Log("=======================")
	user := Users{}
	res, err := user.FindUserByID(s.DB, 1)

	require.NoError(s.T(), err)
	require.Exactly(s.T(), &Users{
		FirstName: "Bob",
		LastName:  "Ross",
		Role:      "User",
		Email:     "bobbyrossy@ross.ross",
	}, res)
}

func TestMigrate(t *testing.T) {
	// Create a dummy DB and run Migrate(db)

	// Check that all tables migrate should create exmodule github.com/alumni-lab/sap-cloneist.
}

func TestFindUserByID(t *testing.T) {

}

func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}
