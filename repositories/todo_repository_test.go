package repositories

import (
	"database/sql"
	"regexp"
	"testing"

	"github.com/go-test/deep"
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
		AddRow(1, "Make some program", "Create TODO app for testing")

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `todo`")).WillReturnRows(rows)
	result := []models.Todo{}

	err := s.repository.GetAllTodo(&result)
	require.NoError(s.T(), err)
}

func (s *Suite) TestCreateATodo() {
	var (
		id    = 1
		title = "testing"
		desc  = "testing description"
	)

	s.mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "todo" ("id","title","description") 
			VALUES ($1,$2,$3) RETURNING "todo"."id"`)).
		WithArgs(id, title, desc).
		WillReturnRows(
			sqlmock.NewRows([]string{"id"}).AddRow(id))

	var result models.Todo
	err := s.repository.CreateATodo(&result)

	require.NoError(s.T(), err)
}

func (s *Suite) TestGetATodo() {

	var (
		id          = 1
		title       = "Make some program"
		description = "Create TODO app for testing"
	)

	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "todo" WHERE (id = $1)`)).
		WithArgs(id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "description"}).
			AddRow(id, title, description))

	result := models.Todo{}

	err := s.repository.GetATodo(&result, id)
	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(&models.Todo{ID: id, Title: title, Description: description}, &result))
}
