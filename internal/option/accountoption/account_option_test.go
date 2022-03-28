package option

import (
	"testing"

	"github.com/art-injener/otus-homework/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestAccountOption(t *testing.T) {
	tests := []struct {
		name  string
		valid bool
		value interface{}
		test  func(value interface{}, account *models.Account) interface{}
	}{
		{
			name:  "positive age",
			valid: true,
			value: 19,
			test: func(value interface{}, account *models.Account) interface{} {
				WithAge(value.(int))(account)
				return account.Age
			},
		},
		{
			name:  "negative age",
			valid: false,
			value: -19,
			test: func(value interface{}, account *models.Account) interface{} {
				WithAge(value.(int))(account)
				return account.Age
			},
		},
		{
			name:  "correct sex",
			valid: true,
			value: "male",
			test: func(value interface{}, account *models.Account) interface{} {
				WithSex(value.(string))(account)
				return account.Sex
			},
		},
		{
			name:  "incorrect sex",
			valid: false,
			value: "man",
			test: func(value interface{}, account *models.Account) interface{} {
				WithSex(value.(string))(account)
				return account.Sex
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			account := models.NewAccount()
			res := tt.test(tt.value, account)
			if tt.valid {
				assert.True(t, tt.value == res)
			} else {
				assert.False(t, tt.value == res)
			}
		})
	}
}
