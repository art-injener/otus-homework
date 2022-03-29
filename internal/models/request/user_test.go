package request

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser_BeforeCreate(t *testing.T) {
	user := TestUser(t)

	assert.NoError(t, user.BeforeCreate())
	assert.NotEmpty(t, user.EncryptedPassword)

}

func TestUserRequest_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		test    func() *User
		isValid bool
	}{
		{
			name: "ValidUser",
			test: func() *User {
				return TestUser(t)
			},
			isValid: true,
		},
		{
			name: "invalid email",
			test: func() *User {
				user := TestUser(t)
				user.Email = "qwert"
				return user
			},
			isValid: false,
		},
		{
			name: "invalid password",
			test: func() *User {
				user := TestUser(t)
				user.Password = "pas"
				return user
			},
			isValid: false,
		},
		{
			name: "invalid repeated password",
			test: func() *User {
				user := TestUser(t)
				user.RepeatedPassword = "pas"
				return user
			},
			isValid: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.test().Validate())
			} else {
				assert.Error(t, tc.test().Validate())
			}
		})
	}
}
