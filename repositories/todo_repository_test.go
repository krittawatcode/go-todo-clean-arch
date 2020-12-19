package repositories

import (
	"database/sql"
	"regexp"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/krittawatcode/go-todo-clean-arch/domains"
	"github.com/krittawatcode/go-todo-clean-arch/models"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

type Suite struct {
	suite.Suite
	db   *gorm.DB
	mock sqlmock.Sqlmock

	repository domains.ToDoRepository
	todo       *models.Todo
}

func (s *Suite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.db, err = gorm.Open("mysql", db)
	require.NoError(s.T(), err)

	s.db.LogMode(true)

	if err != nil {
		s.T().Error(err)
	}
	// Explicitly clean up GORM after the test
	s.T().Cleanup(func() {
		s.db.Close()
	})
	// if err := s.db.AutoMigrate(&models.Todo{}); err != nil {
	// 	s.T().Error(err)
	// }
	s.repository = NewToDoRepository(s.db)
}

func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) TestGetAllTodo() {

	rows := sqlmock.NewRows([]string{"id", "title", "description"}).
		AddRow(1, "Make some program", "Create TODO app for testing").
		AddRow(2, "Make some program", "Create TODO app for testing")

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `todo`")).WillReturnRows(rows)
	result := []models.Todo{}

	err := s.repository.GetAllTodo(&result)
	require.NoError(s.T(), err)
}

func (s *Suite) TestCreateATodo() {
	var (
		title = "testing"
		desc  = "testing description"
	)

	expectResult := sqlmock.NewRows([]string{"id", "title", "description"}).AddRow(1, "testing", "testing description")

	s.mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "todo" ("title", "description") VALUES ($1, $2)`)).
		WithArgs(title, desc).
		WillReturnRows(expectResult)

	var result models.Todo
	err := s.repository.CreateATodo(&result)
	require.NoError(s.T(), err)
}
