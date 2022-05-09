package http_test

import (
	"encoding/json"
	"go-boiler-plate/cmd/router"
	cmdutil "go-boiler-plate/cmd/util"
	tokenhttp "go-boiler-plate/internal/app/domain/token/delivery/http"
	"go-boiler-plate/internal/app/payload"
	"go-boiler-plate/internal/pkg/database"
	"go-boiler-plate/test"
	"net/http"

	"github.com/golang-migrate/migrate/v4"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"repo.pegadaian.co.id/ms-pds/modules/pgdutil"
)

var (
	db           *database.Db
	migrator     *migrate.Migrate
	pl           interface{}
	response     pgdutil.Response
	expectResp   pgdutil.Response
	responseData map[string]interface{}
	e            pgdutil.DummyEcho
	rtr          router.Router
	handler      tokenhttp.TokenHandler

	_ = BeforeSuite(func() {
		cmdutil.LoadEnv()
		cmdutil.LoadTestData()
		db, migrator = test.NewTestDb()
	})

	_ = AfterSuite(func() {
		_ = migrator.Drop()

		db.Sqlx.Close()
		migrator.Close()
	})

	_ = BeforeEach(func() {
		pl = nil
		response = pgdutil.Response{}
		responseData = map[string]interface{}{}
		expectResp = pgdutil.Response{}
		e = pgdutil.DummyEcho{}
		rtr = router.NewRoutes(db)
		handler = tokenhttp.TokenHandler{
			Ihandler:      pgdutil.NewHandler(&pgdutil.Handler{}),
			ITokenUsecase: rtr.Usecases.ITokenUsecase,
		}
	})
)

var _ = Describe("TokenHandler", func() {
	Describe("HCreateToken", func() {
		var reqpl payload.TokenRequest

		JustBeforeEach(func() {
			reqpl = payload.TokenRequest{
				Username: "username123",
				Password: "password123",
			}
			pl = reqpl
			e = pgdutil.NewDummyEcho(http.MethodPost, "/", pl)
			_ = handler.HCreateToken(e.Context)
			_ = json.Unmarshal(e.Response.Body.Bytes(), &response)
			newExpectedResp()
		})

		It("expect response to return as expected", func() {
			responseData = response.Data.(map[string]interface{})
			Expect(responseData["username"]).To(Equal("username123"))
		})
	})
})

func newExpectedResp(isError ...bool) {
	if isError == nil {
		isError = append(isError, false)
	}

	expectResp = pgdutil.Response{
		Code:    "00",
		Status:  "Success",
		Message: "Data Berhasil Dikirim",
	}

	if isError[0] {
		response.Data = nil
		expectResp.Code = "99"
		expectResp.Status = "Error"
		expectResp.Message = ""

		return
	}

	if response.Data != nil {
		responseData = response.Data.(map[string]interface{})
	}
}
