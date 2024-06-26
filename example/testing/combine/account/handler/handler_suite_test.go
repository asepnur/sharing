package handler_test

import (
	"context"
	"testing"

	"github.com/asepnur/sharing/example/testing/combine/account/core/model"
	"github.com/asepnur/sharing/example/testing/combine/account/core/module"
	"github.com/asepnur/sharing/example/testing/combine/account/handler"
	"github.com/asepnur/sharing/example/testing/combine/account/pkg/conn"
	"github.com/asepnur/sharing/example/testing/combine/account/repo"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestHandler(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Handler Suite")
}

var _ = Describe("AuthService", func() {
	var (
		accountHandler handler.AccountHandler
	)

	BeforeEach(func() {
		mysqlConn, err := conn.InitConnection()
		if err != nil {
			panic(err)
		}
		accountRepo := repo.NewAccountRepo(mysqlConn)
		accountModule := module.NewAccountModule(accountRepo)
		accountHandler = handler.NewAccountHandler(accountModule)
	})

	Describe("Register", func() {
		It("should register a new account", func() {
			err := accountHandler.Register(context.Background(), &model.Account{
				Serial:   "testuser",
				Username: "testuser",
				Password: "password",
				Type:     "default",
			})
			Expect(err).To(BeNil())
		})

		//		It("should return an error if the user already exists", func() {
		//			err := accountHandler.Register(context.Background(), &model.Account{
		//				Username: "same_username",
		//				Password: "password",
		//			})
		//			Expect(err).To(BeNil())
		//			err = accountHandler.Register(context.Background(), &model.Account{
		//				Username: "same_username",
		//				Password: "password",
		//			})
		//			Expect(err).ToNot(BeNil())
		//		})
	})

	// Describe("Authenticate", func() {
	// 	It("should authenticate a registered user", func() {
	// 		account := &model.Account{
	// 			Username: "authenticate_user",
	// 			Password: "password",
	// 		}
	// 		err := accountHandler.Register(context.Background(), account)
	// 		Expect(err).To(BeNil())
	// 		err = accountHandler.Authenticate(context.Background(), account)
	// 		Expect(err).To(BeNil())
	// 	})

	// 	It("should return an error if the credentials are incorrect", func() {
	// 		account := &model.Account{
	// 			Username: "authenticate_user_2",
	// 			Password: "wrongpass",
	// 		}
	// 		err := accountHandler.Authenticate(context.Background(), account)
	// 		Expect(err).ToNot(BeNil())
	// 	})
	// })
})
