package accounts

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/art-injener/otus-homework/internal/models"
	"github.com/art-injener/otus-homework/internal/models/request"
	"github.com/art-injener/otus-homework/internal/repository"
	"github.com/art-injener/otus-homework/internal/repository/mysql"
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
	rows, err := r.db.QueryContext(ctx, mysql.QueryGetAllAccounts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	accounts := make([]*models.Account, 0, defaultCountAccounts)
	for rows.Next() {
		var account models.Account
		err = rows.Scan(
			&account.ID,
			&account.LoginID,
			&account.Name,
			&account.Surname,
			&account.Birthday,
			&account.Sex,
			&account.Hobby,
			&account.City,
			&account.Avatar,
		)

		if err != nil {
			return nil, err
		}

		if account.Name != "" && account.Surname != "" {
			accounts = append(accounts, &account)
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return accounts, nil
}

func (r *accountsRepositoryImpl) GetAccountByID(ctx context.Context, id int) (*models.Account, error) {
	return r.getAccountByParam(ctx, "id", id)
}

func (r *accountsRepositoryImpl) AddAccount(ctx context.Context, account *models.Account) error {
	stmt, err := r.db.PrepareContext(ctx, mysql.QueryAddAccount)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(
		&account.LoginID,
		&account.Name,
		&account.Surname,
		&account.Birthday,
		&account.Sex,
		&account.Hobby,
		&account.City,
		&account.Avatar,
	)
	if err != nil {
		return err
	}

	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return err
	}
	account.ID = int(lastInsertID)

	return nil
}

func (r *accountsRepositoryImpl) GetAccountByUserID(ctx context.Context, userID int) (*models.Account, error) {
	return r.getAccountByParam(ctx, "login_id", userID)
}

func (r *accountsRepositoryImpl) UpdateAccount(ctx context.Context, acc *models.Account) error {
	_, err := r.db.ExecContext(ctx, fmt.Sprintf(mysql.QueryUpdateAccount,
		acc.Name,
		acc.Surname,
		acc.Birthday,
		acc.Sex,
		acc.Hobby,
		acc.City,
		acc.Avatar,
		acc.LoginID))
	if err != nil {
		return err
	}

	return nil
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
	stmt, err := r.db.PrepareContext(ctx, mysql.QueryAddUser)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, user.Email, user.EncryptedPassword)
	if err != nil {
		return err
	}

	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return err
	}
	user.ID = int(lastInsertID)
	return nil
}

func (r *accountsRepositoryImpl) getUserWithCheckParam(
	ctx context.Context,
	checkUserParamFunc checkUserParamFunc) (*request.User, error) {
	rows, err := r.db.QueryContext(ctx, mysql.QueryGetLoginsInfo)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var user request.User
		err = rows.Scan(&user.ID, &user.Email, &user.EncryptedPassword)
		if err != nil {
			return nil, err
		}
		if checkUserParamFunc(&user) {
			return &user, nil
		}
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return nil, repository.ErrUserNotFound
}

func (r *accountsRepositoryImpl) getAccountByParam(
	ctx context.Context,
	param string,
	value interface{}) (*models.Account, error) {
	row := r.db.QueryRowContext(ctx, fmt.Sprintf(mysql.QueryGetAccountByParam, param), value)

	var account models.Account
	err := row.Scan(
		&account.ID,
		&account.LoginID,
		&account.Name,
		&account.Surname,
		&account.Birthday,
		&account.Sex,
		&account.Hobby,
		&account.City,
		&account.Avatar,
	)

	if err != nil {
		return nil, err
	}
	account.FormattedBirthday("02 January 2006")

	return &account, nil
}

func (r *accountsRepositoryImpl) MakeFriends(ctx context.Context, currentUserID, friendID int) error {

	if isFriends, err := r.IsFriends(ctx, currentUserID, friendID); err != nil {
		return err
	} else if isFriends {
		return nil
	}

	insert, err := r.db.QueryContext(ctx, mysql.QueryMakeFriends, currentUserID, friendID)
	if err != nil {
		return err
	}
	defer insert.Close()

	return nil
}

func (r *accountsRepositoryImpl) GetFriends(ctx context.Context, accountID int) ([]*models.Account, error) {
	rows, err := r.db.QueryContext(ctx, mysql.QueryGetFriends, accountID)
	if err != nil {
		return nil, err
	}

	accounts := make([]*models.Account, 0)
	for rows.Next() {
		var account models.Account
		err := rows.Scan(&account.ID, &account.Name, &account.Surname)

		if err != nil {
			return nil, err
		}
		accounts = append(accounts, &account)
	}
	defer rows.Close()

	return accounts, nil
}

func (r *accountsRepositoryImpl) IsFriends(ctx context.Context, currentUserID, friendID int) (bool, error) {
	var count int

	if err := r.db.QueryRowContext(ctx, mysql.QueryCheckFriends, currentUserID, friendID).Scan(&count); err != nil {
		return false, err
	}
	return count > 0, nil
}
