package accountoption

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
			name:  "correct sex",
			valid: true,
			value: 'm',
			test: func(value interface{}, account *models.Account) interface{} {
				WithSex(value.(string))(account)

				return account.Sex
			},
		},
		{
			name:  "incorrect sex",
			valid: false,
			value: 'v',
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
