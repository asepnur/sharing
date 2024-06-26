package repo

import (
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/asepnur/sharing/example/testing/combine/account/core/model"
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

func Test_accountRepoImpl_FindByUsername_ExpectError(t *testing.T) {
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

	mock.ExpectQuery("^SELECT \\* FROM account WHERE username = \\?").
		WithArgs("1").
		WillReturnError(errors.New("scan error"))

	account, err := repo.FindByUsername("1")
	assert.Error(t, err)
	assert.Nil(t, account)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

/*
	func Test_accountRepoImpl_FindByUsername_ScanError(t *testing.T) {
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

		expectedRows := sqlmock.NewRows([]string{"Serial", "Username", "Password", "Type"}).AddRow(1, 2, 3, 4)
		mock.ExpectQuery("^SELECT \\* FROM account WHERE username = \\?").
			WithArgs("1").
			WillReturnRows(expectedRows)

		account, err := repo.FindByUsername("1")
		assert.Error(t, err)
		assert.Nil(t, account)
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	}
*/
func TestInsert(t *testing.T) {
	// Create a new mock database connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Open GORM DB connection with the mock database
	gormDB, err := gorm.Open("mysql", db)
	if err != nil {
		t.Fatalf("failed to open gorm db: %v", err)
	}

	// Create an instance of accountRepoImpl with the mock connection
	repo := &accountRepoImpl{conn: gormDB}

	// Define the account to insert
	account := &model.Account{
		Username: "John Doe",
		Serial:   "1234",
		Password: "password123",
		Type:     "default",
	}

	// Set up expectations
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `account` (`serial`,`username`,`password`,`type`) VALUES (?,?,?,?)")).
		WithArgs(account.Serial, account.Username, account.Password, account.Type).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Call the Insert function
	err = repo.Insert(account)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
