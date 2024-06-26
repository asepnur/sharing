package repo

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func Test_accountRepoImpl_FindByUsername(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open("mysql", db)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database connection", err)
	}
	defer gormDB.Close()
	repo := NewAccountRepo(gormDB)
	rows := sqlmock.NewRows([]string{"Serial", "Username", "Password", "Type"}).
		AddRow("1234", "testuser", "password123", "admin")

	mock.ExpectQuery("^SELECT \\* FROM account WHERE username = \\?").
		WithArgs("testuser").
		WillReturnRows(rows)

	account, err := repo.FindByUsername("testuser")

	assert.NoError(t, err)
	assert.NotNil(t, account)
	assert.Equal(t, "1234", account.Serial)
	assert.Equal(t, "testuser", account.Username)
	assert.Equal(t, "password123", account.Password)
	assert.Equal(t, "admin", account.Type)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
