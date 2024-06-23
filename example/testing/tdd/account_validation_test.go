package tdd

import "testing"

// Step 2: Create a test function of Validate() function.
func TestAccount_Validate(t *testing.T) {
	type fields struct {
		Username string
		Password string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// Step 3: Create a test cases.
		{
			name: "case 1: valid account",
			fields: fields{
				Username: "username",
				Password: "password",
			},
			wantErr: false,
		},
		{
			name: "case 2: username empty, password exist",
			fields: fields{
				Username: "",
				Password: "password",
			},
			wantErr: true,
		},
		{
			name: "case 3: username exist, password empty",
			fields: fields{
				Username: "username",
				Password: "",
			},
			wantErr: true,
		},
		// NOTE: this is redundant case because it's already covered in the 2nd and 3rd case.
		{
			name: "case 4: username exist, password empty",
			fields: fields{
				Username: "",
				Password: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Account{
				Username: tt.fields.Username,
				Password: tt.fields.Password,
			}
			if err := a.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Account.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
