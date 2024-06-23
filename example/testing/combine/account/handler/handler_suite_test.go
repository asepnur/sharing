package handler_test

import (
	"context"
	"testing"

	"github.com/asepnur/sharing/example/testing/combine/account/core/model"
	"github.com/asepnur/sharing/example/testing/combine/account/handler"

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
		accountHandler = handler.NewAccountHandler()
	})

	Describe("Register", func() {
		It("should register a new account", func() {
			err := accountHandler.Register(context.Background(), &model.Account{
				Username: "testuser",
				Password: "password",
			})
			Expect(err).To(BeNil())
		})

		It("should return an error if the user already exists", func() {
			err := accountHandler.Register(context.Background(), &model.Account{
				Username: "same_username",
				Password: "password",
			})
			Expect(err).To(BeNil())
			err = accountHandler.Register(context.Background(), &model.Account{
				Username: "same_username",
				Password: "password",
			})
			Expect(err).ToNot(BeNil())
		})
	})

	Describe("Authenticate", func() {
		It("should authenticate a registered user", func() {
			account := &model.Account{
				Username: "authenticate_user",
				Password: "password",
			}
			err := accountHandler.Register(context.Background(), account)
			Expect(err).To(BeNil())
			err = accountHandler.Authenticate(context.Background(), account)
			Expect(err).To(BeNil())
		})

		It("should return an error if the credentials are incorrect", func() {
			account := &model.Account{
				Username: "authenticate_user_2",
				Password: "wrongpass",
			}
			err := accountHandler.Authenticate(context.Background(), account)
			Expect(err).ToNot(BeNil())
		})
	})
})
