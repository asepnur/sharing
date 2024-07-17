package tdd

import (
	"errors"
	"testing"
)

// Step 2: Create a test function of Validate() function.
func TestAccount_Validate(t *testing.T) {
	type fields struct {
		Username string
		Password string
	}
	tests := []struct {
		name      string
		fields    fields
		wantErr   bool
		errorType error
	}{
		// Step 3: Create a test cases.
		{
			name: "case 1: valid account",
			fields: fields{
				Username: "username",
				Password: "password",
			},
			wantErr:   false,
			errorType: nil,
		},
		{
			name: "case 2: username empty, password exist",
			fields: fields{
				Username: "",
				Password: "password",
			},
			wantErr:   true,
			errorType: ErrUsernameEmpty,
		},
		{
			name: "case 3: username exist, password empty",
			fields: fields{
				Username: "username",
				Password: "",
			},
			wantErr:   true,
			errorType: ErrPasswordEmpty,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Account{
				Username: tt.fields.Username,
				Password: tt.fields.Password,
			}
			err := a.Validate()
			if err != nil == tt.wantErr {
				t.Errorf("Account.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				if errors.Is(err, tt.errorType) {
					t.Errorf("Account.Type() error = %v, wantErr %v", err, tt.errorType)
				}
			}
		})
	}
}
