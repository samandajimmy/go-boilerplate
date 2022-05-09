package usecase_test

import (
	tknusecase "go-boiler-plate/internal/app/domain/token/usecase"
	"go-boiler-plate/internal/app/model"
	"go-boiler-plate/internal/app/payload"
	"go-boiler-plate/test"
	"go-boiler-plate/test/mock"
	"net/http"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"repo.pegadaian.co.id/ms-pds/modules/pgdutil"
)

var _ = Describe("TokenUcase", func() {
	var e pgdutil.DummyEcho
	var mockCtrl *gomock.Controller
	var mockRepos mock.MockRepositories
	var mockUtil *mock.MockIUtil
	var usecase tknusecase.TokenUsecase

	BeforeEach(func() {
		e = pgdutil.NewDummyEcho(http.MethodPost, "/")
		mockCtrl = gomock.NewController(GinkgoT())
		mockRepos, _ = test.LoadMockRepoUsecase(mockCtrl)
		mockUtil = mock.NewMockIUtil(mockCtrl)
		usecase.ITokenRepo = mockRepos.MockITokenRepository
	})

	Describe("UCreateToken", func() {
		var pl payload.TokenRequest
		var tokenResponse payload.TokenResponse

		BeforeEach(func() {
			pl = payload.TokenRequest{
				Username: "username",
				Password: "password",
			}
			usecase.IUtil = mockUtil
			mockUtil.EXPECT().BcryptHashedPassword(pl.Password).Return("password123")
			accToken := model.AccountToken{
				Username: pl.Username,
				Password: "password123",
			}
			mockRepos.MockITokenRepository.EXPECT().RCreate(e.Context, &accToken).Return(nil)
		})
		JustBeforeEach(func() {
			tokenResponse, _ = usecase.UCreateToken(e.Context, pl)
		})

		It("expect response to return as expected", func() {
			Expect(tokenResponse.Username).To(Equal(pl.Username))
		})
	})
})
