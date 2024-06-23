package module

import (
	"testing"

	"github.com/asepnur/sharing/example/testing/combine/account/core/model"
	"github.com/asepnur/sharing/example/testing/combine/account/core/repo"
	"github.com/golang/mock/gomock"
)

func Test_accountModuleDependency_Register(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repo.NewMockAccountRepo(ctrl)

	type fields struct {
		repo repo.AccountRepo
	}
	type args struct {
		account *model.Account
	}
	tests := []struct {
		name        string
		description string
		fields      fields
		args        args
		wantErr     bool
	}{
		{
			name: "successful registration",
			fields: fields{
				repo: mockRepo,
			},
			args: args{
				account: &model.Account{Username: "new_user", Password: "password"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &accountModuleDependency{
				repo: tt.fields.repo,
			}
			if tt.name == "username already exists" {
				mockRepo.EXPECT().FindByUsername(tt.args.account.Username).Return(&model.Account{Username: "existing_user", Password: "password"}, nil)
			} else {
				mockRepo.EXPECT().FindByUsername(tt.args.account.Username).Return(nil, nil)
				mockRepo.EXPECT().Insert(gomock.Any()).Return(nil)
			}
			if err := a.Register(tt.args.account); (err != nil) != tt.wantErr {
				t.Errorf("accountModuleDependency.Register() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
