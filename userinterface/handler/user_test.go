package handler_test

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/JY8752/go-onion-architecture-sample/domain/model"
	mock_registory "github.com/JY8752/go-onion-architecture-sample/mocks/registory"
	mock_service "github.com/JY8752/go-onion-architecture-sample/mocks/service"
	"github.com/JY8752/go-onion-architecture-sample/userinterface/handler"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/sebdah/goldie/v2"
	"github.com/stretchr/testify/assert"
)

const (
	goldenDir = "../../testdata/golden/"
)

func TestCreateUserHandler(t *testing.T) {
	// given
	e := echo.New()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// エンドポイントの登録
	service := mock_service.NewMockUserService(ctrl)
	service.EXPECT().Create("user1").Return(model.UserId(1), nil)

	registory := mock_registory.NewMockRegistory(ctrl)
	registory.EXPECT().UserService().Return(service)

	handler.CreateUserHandler(e, registory)

	// リクエストの作成
	body := `{"name": "user1"}`
	w := httptest.NewRecorder()
	r := httptest.NewRequest(echo.POST, "/users", strings.NewReader(body))
	r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	defer r.Body.Close()

	// when
	e.ServeHTTP(w, r)

	// then
	assert.Equal(t, 200, w.Code)
	g := goldie.New(t, goldie.WithFixtureDir(goldenDir))
	g.Assert(t, t.Name(), w.Body.Bytes())
}
