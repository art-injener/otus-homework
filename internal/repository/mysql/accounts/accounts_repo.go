package accounts

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/art-injener/otus-homework/internal/models"
	"github.com/art-injener/otus-homework/internal/models/request"
	"github.com/art-injener/otus-homework/internal/repository"
)

type checkUserParamFunc func(*request.User) bool

const (
	defaultCountAccounts = 1024
)

type accountsRepositoryImpl struct {
	db *sql.DB
}

var _ repository.AccountsRepository = &accountsRepositoryImpl{}

func NewAccountsRepo(db *sql.DB) *accountsRepositoryImpl {
	return &accountsRepositoryImpl{
		db: db,
	}
}

func (r *accountsRepositoryImpl) GetAllAccounts(ctx context.Context) ([]*models.Account, error) {
	rows, err := r.db.QueryContext(ctx, "select id,login_id,name,surname,age,hobby,city from users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	accounts := make([]*models.Account, 0, defaultCountAccounts)
	for rows.Next() {
		var account models.Account
		err := rows.Scan(&account.ID, &account.LoginID, &account.Name, &account.Surname, &account.Age, &account.Hobby, &account.City)
		//err := rows.Scan(&account.ID)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, &account)
	}
	return accounts, nil
}

func (r *accountsRepositoryImpl) GetAccountByID(ctx context.Context, id int) (*models.Account, error) {
	return r.getAccountByParam(ctx, "id", id)
}

func (r *accountsRepositoryImpl) AddAccount(ctx context.Context, account *models.Account) error {
	insert, err := r.db.QueryContext(ctx, "insert into users (login_id, name, surname, age, sex, hobby, city) values (?, ?, ?, ?, ?, ?, ?)",
		&account.LoginID, &account.Name, &account.Surname, &account.Age, &account.Sex, &account.Hobby, &account.City)
	if err != nil {
		return err
	}
	insert.Close()
	return nil
}

func (r *accountsRepositoryImpl) GetAccountByUserID(ctx context.Context, userID int) (*models.Account, error) {
	return r.getAccountByParam(ctx, "login_id", userID)
}

func (r *accountsRepositoryImpl) GetUserByEmail(ctx context.Context, email string) (*request.User, error) {
	return r.getUserWithCheckParam(ctx, func(user *request.User) bool {
		return user.Email == email
	})
}

func (r *accountsRepositoryImpl) GetUserByID(ctx context.Context, id int) (*request.User, error) {
	return r.getUserWithCheckParam(ctx, func(user *request.User) bool {
		return user.ID == id
	})
}

func (r *accountsRepositoryImpl) AddNewUser(ctx context.Context, user *request.User) error {
	insert, err := r.db.QueryContext(ctx, "insert into logins_info (email, password) values (?, ?)", user.Email, user.EncryptedPassword)
	if err != nil {
		return err
	}
	insert.Close()
	return nil
}

func (r *accountsRepositoryImpl) getUserWithCheckParam(ctx context.Context, checkUserParamFunc checkUserParamFunc) (*request.User, error) {
	rows, err := r.db.QueryContext(ctx, "select * from logins_info")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user request.User
		err := rows.Scan(&user.ID, &user.Email, &user.EncryptedPassword)
		if err != nil {
			return nil, err
		}
		if checkUserParamFunc(&user) {
			return &user, nil
		}
	}
	return nil, repository.ErrUserNotFound
}

func (r *accountsRepositoryImpl) getAccountByParam(ctx context.Context, param string, value interface{}) (*models.Account, error) {
	row := r.db.QueryRowContext(ctx, fmt.Sprintf("select * from users where `%s`=?", param), value)

	var account models.Account
	err := row.Scan(&account.ID, &account.LoginID, &account.Name, &account.Surname, &account.Age, &account.Sex, &account.Hobby, &account.City)
	if err != nil {
		return nil, err
	}
	return &account, nil
}
