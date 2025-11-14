package member_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"example.com/member/internal/member"
	"github.com/agussyahrilmubarok/gox/pkg/xexception"
	"github.com/labstack/echo/v4"
	mock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type HandlerTestSuite struct {
	suite.Suite
	mockIService *MockIService
	handler      *member.Handler
	e            *echo.Echo
}

func (ts *HandlerTestSuite) SetupTest() {
	ts.mockIService = NewMockIService(ts.T())
	ts.handler = member.NewHandler(ts.mockIService)
	ts.e = echo.New()
	ts.e.GET("/api/v1/members", ts.handler.FindAll)
	ts.e.GET("/api/v1/members/find", ts.handler.Find)
	ts.e.GET("/api/v1/members/inid-dummy", ts.handler.InitDummy)
	ts.e.GET("/api/v1/members/clean-dummy", ts.handler.CleanDummy)
}

func TestHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(HandlerTestSuite))
}

// -- FindAll --

func (ts *HandlerTestSuite) TestFindAllSuccess() {
	members := []member.Member{
		{ID: "1", Code: "MEM-001", Name: "John Doe"},
		{ID: "2", Code: "MEM-002", Name: "Jane Doe"},
	}

	ts.mockIService.On("FindAll", mock.Anything).Return(members, nil).Once()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/members", nil)
	rec := httptest.NewRecorder()
	c := ts.e.NewContext(req, rec)

	err := ts.handler.FindAll(c)

	ts.NoError(err)
	ts.Equal(http.StatusOK, rec.Code)
	ts.mockIService.AssertExpectations(ts.T())
}

func (ts *HandlerTestSuite) TestFindAllHttpError() {
	httpErr := &xexception.Http{
		Code:    http.StatusBadRequest,
		Message: "custom bad request",
	}

	ts.mockIService.On("FindAll", mock.Anything).Return(nil, httpErr).Once()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/members", nil)
	rec := httptest.NewRecorder()
	c := ts.e.NewContext(req, rec)

	err := ts.handler.FindAll(c)

	ts.NoError(err)
	ts.Equal(http.StatusBadRequest, rec.Code)
	expectedBody := `{"error":"custom bad request"}`
	ts.Equal(expectedBody, strings.TrimSpace(rec.Body.String()))
	ts.mockIService.AssertExpectations(ts.T())
}

func (ts *HandlerTestSuite) TestFindAllFail() {
	ts.mockIService.On("FindAll", mock.Anything).Return(nil, errors.New("db error")).Once()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/members", nil)
	rec := httptest.NewRecorder()
	c := ts.e.NewContext(req, rec)

	err := ts.handler.FindAll(c)

	ts.NoError(err)
	ts.Equal(http.StatusInternalServerError, rec.Code)
	ts.mockIService.AssertExpectations(ts.T())
}

// -- Find --

func (ts *HandlerTestSuite) TestFindParamCodeSuccess() {
	memberData := &member.Member{ID: "1", Code: "MEM-001", Name: "John Doe"}

	ts.mockIService.On("FindByCode", mock.Anything, "MEM-001").Return(memberData, nil).Once()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/members/find?code=MEM-001", nil)
	rec := httptest.NewRecorder()
	c := ts.e.NewContext(req, rec)

	err := ts.handler.Find(c)

	ts.NoError(err)
	ts.Equal(http.StatusOK, rec.Code)
	ts.mockIService.AssertExpectations(ts.T())
}

func (ts *HandlerTestSuite) TestFindHttpErrorByCode() {
	httpErr := &xexception.Http{
		Code:    http.StatusNotFound,
		Message: "member not found",
	}

	ts.mockIService.On("FindByCode", mock.Anything, "MEM-404").Return(nil, httpErr).Once()

	req := httptest.NewRequest(http.MethodGet, "/members/find?code=MEM-404", nil)
	rec := httptest.NewRecorder()
	c := ts.e.NewContext(req, rec)

	err := ts.handler.Find(c)

	ts.NoError(err)
	ts.Equal(http.StatusNotFound, rec.Code)
	expectedBody := `{"error":"member not found"}`
	ts.Equal(expectedBody, strings.TrimSpace(rec.Body.String()))
	ts.mockIService.AssertExpectations(ts.T())
}

func (ts *HandlerTestSuite) TestFindParamCodeFail() {
	req := httptest.NewRequest(http.MethodGet, "/api/v1/members/find", nil)
	rec := httptest.NewRecorder()

	c := ts.e.NewContext(req, rec)
	err := ts.handler.Find(c)
	ts.NoError(err)
	ts.Equal(http.StatusBadRequest, rec.Code)

	ts.mockIService.On("FindByCode", mock.Anything, "MEM-001").Return(nil, errors.New("db error")).Once()

	req2 := httptest.NewRequest(http.MethodGet, "/members/find?code=MEM-001", nil)
	rec2 := httptest.NewRecorder()
	c2 := ts.e.NewContext(req2, rec2)

	err = ts.handler.Find(c2)

	ts.NoError(err)
	ts.Equal(http.StatusInternalServerError, rec2.Code)
	ts.mockIService.AssertExpectations(ts.T())
}

// -- InitDummy --

func (ts *HandlerTestSuite) TestInitDummySuccess() {
	ts.mockIService.On("Save", mock.Anything, mock.AnythingOfType("*member.Member")).
		Return(nil, nil).Twice()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/members/init-dummy", nil)
	rec := httptest.NewRecorder()
	c := ts.e.NewContext(req, rec)

	err := ts.handler.InitDummy(c)

	ts.NoError(err)
	ts.Equal(http.StatusOK, rec.Code)
	ts.Contains(rec.Body.String(), "MEMBER-1000")
	ts.Contains(rec.Body.String(), "MEMBER-1001")
	ts.mockIService.AssertExpectations(ts.T())
}

func (ts *HandlerTestSuite) TestInitDummyHttpError() {
	httpErr := &xexception.Http{
		Code:    http.StatusBadRequest,
		Message: "custom bad request",
	}

	ts.mockIService.On("Save", mock.Anything, mock.AnythingOfType("*member.Member")).
		Return(nil, httpErr).Once()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/members/init-dummy", nil)
	rec := httptest.NewRecorder()
	c := ts.e.NewContext(req, rec)

	err := ts.handler.InitDummy(c)

	ts.NoError(err)
	ts.Equal(http.StatusBadRequest, rec.Code)
	expectedBody := `{"error":"custom bad request"}`
	ts.Equal(expectedBody, strings.TrimSpace(rec.Body.String()))
	ts.mockIService.AssertExpectations(ts.T())
}

// -- CleanDummy --

func (ts *HandlerTestSuite) TestCleanDummySuccess() {
	members := []member.Member{
		{ID: "1", Code: "MEM-1000", Name: "John Doe"},
		{ID: "2", Code: "MEM-1001", Name: "Jane Smith"},
	}

	ts.mockIService.On("FindAll", mock.Anything).Return(members, nil).Once()
	ts.mockIService.On("DeleteByID", mock.Anything, "1").Return(nil).Once()
	ts.mockIService.On("DeleteByID", mock.Anything, "2").Return(nil).Once()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/members/clean-dummy", nil)
	rec := httptest.NewRecorder()
	c := ts.e.NewContext(req, rec)

	err := ts.handler.CleanDummy(c)

	ts.NoError(err)
	ts.Equal(http.StatusOK, rec.Code)
	ts.Contains(rec.Body.String(), "dummy data cleaned")
	ts.mockIService.AssertExpectations(ts.T())
}

func (ts *HandlerTestSuite) TestCleanDummyHttpError() {
	members := []member.Member{
		{ID: "1", Code: "MEM-1000", Name: "John Doe"},
		{ID: "2", Code: "MEM-1001", Name: "Jane Smith"},
	}
	httpErr := &xexception.Http{
		Code:    http.StatusInternalServerError,
		Message: "custom internal error",
	}

	ts.mockIService.On("FindAll", mock.Anything).Return(members, nil).Once()
	ts.mockIService.On("DeleteByID", mock.Anything, "1").Return(httpErr).Once()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/members/clean-dummy", nil)
	rec := httptest.NewRecorder()
	c := ts.e.NewContext(req, rec)

	err := ts.handler.CleanDummy(c)

	ts.NoError(err)
	ts.Equal(http.StatusInternalServerError, rec.Code)
	expectedBody := `{"error":"custom internal error"}`
	ts.Equal(expectedBody, strings.TrimSpace(rec.Body.String()))
	ts.mockIService.AssertExpectations(ts.T())
}
