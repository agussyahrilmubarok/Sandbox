package course_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"example.com/course/internal/course"
	"github.com/agussyahrilmubarok/gox/pkg/xexception"
	"github.com/labstack/echo/v4"
	mock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type HandlerTestSuite struct {
	suite.Suite
	mockIService *MockIService
	handler      *course.Handler
	e            *echo.Echo
}

func (ts *HandlerTestSuite) SetupTest() {
	ts.mockIService = NewMockIService(ts.T())
	ts.handler = course.NewHandler(ts.mockIService)
	ts.e = echo.New()
	ts.e.GET("/api/v1/courses", ts.handler.FindAll)
	ts.e.GET("/api/v1/courses/find", ts.handler.Find)
	ts.e.POST("/api/v1/courses/reserve", ts.handler.ReserveCourse)
	ts.e.POST("/api/v1/courses/release", ts.handler.ReleaseCourse)
	ts.e.POST("/api/v1/courses/init-dummy", ts.handler.InitDummy)
	ts.e.DELETE("/api/v1/courses/clean-dummy", ts.handler.CleanDummy)
}

func TestHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(HandlerTestSuite))
}

// --- FindAll ---

func (ts *HandlerTestSuite) TestFindAllSuccess() {
	courses := []course.Course{
		{ID: "1", Code: "COURSE-001", Name: "Go Basics"},
		{ID: "2", Code: "COURSE-002", Name: "Advanced Go"},
	}

	ts.mockIService.On("FindAll", mock.Anything).Return(courses, nil).Once()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/courses", nil)
	rec := httptest.NewRecorder()
	c := ts.e.NewContext(req, rec)

	err := ts.handler.FindAll(c)

	ts.NoError(err)
	ts.Equal(http.StatusOK, rec.Code)
	ts.Contains(rec.Body.String(), "COURSE-001")
	ts.mockIService.AssertExpectations(ts.T())
}

func (ts *HandlerTestSuite) TestFindAllHttpError() {
	httpErr := &xexception.Http{
		Code:    http.StatusNotFound,
		Message: "no courses available",
	}

	ts.mockIService.On("FindAll", mock.Anything).Return(nil, httpErr).Once()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/courses", nil)
	rec := httptest.NewRecorder()
	c := ts.e.NewContext(req, rec)

	err := ts.handler.FindAll(c)

	ts.NoError(err)
	ts.Equal(http.StatusNotFound, rec.Code)
	expectedBody := `{"error":"no courses available"}`
	ts.Equal(expectedBody, strings.TrimSpace(rec.Body.String()))
	ts.mockIService.AssertExpectations(ts.T())
}

func (ts *HandlerTestSuite) TestFindAllFail() {
	ts.mockIService.On("FindAll", mock.Anything).Return(nil, errors.New("db connection failed")).Once()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/courses", nil)
	rec := httptest.NewRecorder()
	c := ts.e.NewContext(req, rec)

	err := ts.handler.FindAll(c)

	ts.NoError(err)
	ts.Equal(http.StatusInternalServerError, rec.Code)
	expectedBody := `{"error":"failed to fetch courses"}`
	ts.Equal(expectedBody, strings.TrimSpace(rec.Body.String()))
	ts.mockIService.AssertExpectations(ts.T())
}

// --- Find ---

func (ts *HandlerTestSuite) TestFindParamCodeSuccess() {
	courseData := &course.Course{ID: "1", Code: "COURSE-A", Name: "Course A"}

	ts.mockIService.On("FindByCode", mock.Anything, "COURSE-A").Return(courseData, nil).Once()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/courses/find?code=COURSE-A", nil)
	rec := httptest.NewRecorder()
	c := ts.e.NewContext(req, rec)

	err := ts.handler.Find(c)

	ts.NoError(err)
	ts.Equal(http.StatusOK, rec.Code)
	ts.Contains(rec.Body.String(), "COURSE-A")
	ts.mockIService.AssertExpectations(ts.T())
}

func (ts *HandlerTestSuite) TestFindMissingCodeParamFail() {
	req := httptest.NewRequest(http.MethodGet, "/api/v1/courses/find", nil)
	rec := httptest.NewRecorder()
	c := ts.e.NewContext(req, rec)

	err := ts.handler.Find(c)

	ts.NoError(err)
	ts.Equal(http.StatusBadRequest, rec.Code)
	expectedBody := `{"error":"missing code parameter"}`
	ts.Equal(expectedBody, strings.TrimSpace(rec.Body.String()))
	ts.mockIService.AssertNotCalled(ts.T(), "FindByCode", mock.Anything, mock.Anything)
}

func (ts *HandlerTestSuite) TestFindHttpErrorByCode() {
	httpErr := &xexception.Http{
		Code:    http.StatusNotFound,
		Message: "course not found by code",
	}

	ts.mockIService.On("FindByCode", mock.Anything, "COURSE-404").Return(nil, httpErr).Once()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/courses/find?code=COURSE-404", nil)
	rec := httptest.NewRecorder()
	c := ts.e.NewContext(req, rec)

	err := ts.handler.Find(c)

	ts.NoError(err)
	ts.Equal(http.StatusNotFound, rec.Code)
	expectedBody := `{"error":"course not found by code"}`
	ts.Equal(expectedBody, strings.TrimSpace(rec.Body.String()))
	ts.mockIService.AssertExpectations(ts.T())
}

func (ts *HandlerTestSuite) TestFindInternalFailByCode() {
	ts.mockIService.On("FindByCode", mock.Anything, "COURSE-ERR").Return(nil, errors.New("db error")).Once()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/courses/find?code=COURSE-ERR", nil)
	rec := httptest.NewRecorder()
	c := ts.e.NewContext(req, rec)

	err := ts.handler.Find(c)

	ts.NoError(err)
	ts.Equal(http.StatusInternalServerError, rec.Code)
	expectedBody := `{"error":"failed to fetch course"}`
	ts.Equal(expectedBody, strings.TrimSpace(rec.Body.String()))
	ts.mockIService.AssertExpectations(ts.T())
}

// --- ReserveCourse ---

func (ts *HandlerTestSuite) TestReserveCourseSuccess() {
	courseCode := "COURSE-RES-OK"
	jsonBody := `{"code":"` + courseCode + `"}`

	ts.mockIService.On("ReserveByCode", mock.Anything, courseCode).Return(nil).Once()

	req := httptest.NewRequest(http.MethodPost, "/api/v1/courses/reserve", strings.NewReader(jsonBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := ts.e.NewContext(req, rec)

	ts.e.POST("/api/v1/courses/reserve", ts.handler.ReserveCourse)

	err := ts.handler.ReserveCourse(c)

	ts.NoError(err)
	ts.Equal(http.StatusOK, rec.Code)
	ts.Contains(rec.Body.String(), "seat reserved successfully")
	ts.mockIService.AssertExpectations(ts.T())
}

func (ts *HandlerTestSuite) TestReserveCourseInvalidPayloadFail() {
	jsonBody := `{"code":123}`

	req := httptest.NewRequest(http.MethodPost, "/api/v1/courses/reserve", strings.NewReader(jsonBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := ts.e.NewContext(req, rec)

	err := ts.handler.ReserveCourse(c)

	ts.NoError(err)
	ts.Equal(http.StatusBadRequest, rec.Code)
	ts.Contains(rec.Body.String(), "invalid request payload")
	ts.mockIService.AssertNotCalled(ts.T(), "ReserveByCode", mock.Anything, mock.Anything)
}

func (ts *HandlerTestSuite) TestReserveCourseHttpError() {
	courseCode := "COURSE-RES-BAD"
	jsonBody := `{"code":"` + courseCode + `"}`
	httpErr := &xexception.Http{
		Code:    http.StatusBadRequest,
		Message: "no available seats to reserve",
	}

	ts.mockIService.On("ReserveByCode", mock.Anything, courseCode).Return(httpErr).Once()

	req := httptest.NewRequest(http.MethodPost, "/api/v1/courses/reserve", strings.NewReader(jsonBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := ts.e.NewContext(req, rec)

	err := ts.handler.ReserveCourse(c)

	ts.NoError(err)
	ts.Equal(http.StatusBadRequest, rec.Code)
	expectedBody := `{"error":"no available seats to reserve"}`
	ts.Equal(expectedBody, strings.TrimSpace(rec.Body.String()))
	ts.mockIService.AssertExpectations(ts.T())
}

func (ts *HandlerTestSuite) TestReserveCourseInternalFail() {
	courseCode := "COURSE-RES-ERR"
	jsonBody := `{"code":"` + courseCode + `"}`

	ts.mockIService.On("ReserveByCode", mock.Anything, courseCode).Return(errors.New("db failed")).Once()

	req := httptest.NewRequest(http.MethodPost, "/api/v1/courses/reserve", strings.NewReader(jsonBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := ts.e.NewContext(req, rec)

	err := ts.handler.ReserveCourse(c)

	ts.NoError(err)
	ts.Equal(http.StatusInternalServerError, rec.Code)
	expectedBody := `{"error":"failed to reserve course"}`
	ts.Equal(expectedBody, strings.TrimSpace(rec.Body.String()))
	ts.mockIService.AssertExpectations(ts.T())
}

// --- ReleaseCourse ---

func (ts *HandlerTestSuite) TestReleaseCourseSuccess() {
	courseCode := "COURSE-REL-OK"
	jsonBody := `{"code":"` + courseCode + `"}`

	ts.mockIService.On("ReleaseByCode", mock.Anything, courseCode).Return(nil).Once()

	req := httptest.NewRequest(http.MethodPost, "/api/v1/courses/release", strings.NewReader(jsonBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := ts.e.NewContext(req, rec)

	ts.e.POST("/api/v1/courses/release", ts.handler.ReleaseCourse)

	err := ts.handler.ReleaseCourse(c)

	ts.NoError(err)
	ts.Equal(http.StatusOK, rec.Code)
	ts.Contains(rec.Body.String(), "seat released successfully")
	ts.mockIService.AssertExpectations(ts.T())
}

func (ts *HandlerTestSuite) TestReleaseCourseInvalidPayloadFail() {
	jsonBody := `{"code":""`

	req := httptest.NewRequest(http.MethodPost, "/api/v1/courses/release", strings.NewReader(jsonBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := ts.e.NewContext(req, rec)

	err := ts.handler.ReleaseCourse(c)

	ts.NoError(err)
	ts.Equal(http.StatusBadRequest, rec.Code)
	ts.Contains(rec.Body.String(), "invalid request payload")
	ts.mockIService.AssertNotCalled(ts.T(), "ReleaseByCode", mock.Anything, mock.Anything)
}

func (ts *HandlerTestSuite) TestReleaseCourseHttpError() {
	courseCode := "COURSE-REL-BAD"
	jsonBody := `{"code":"` + courseCode + `"}`
	httpErr := &xexception.Http{
		Code:    http.StatusBadRequest,
		Message: "course has already ended",
	}

	ts.mockIService.On("ReleaseByCode", mock.Anything, courseCode).Return(httpErr).Once()

	req := httptest.NewRequest(http.MethodPost, "/api/v1/courses/release", strings.NewReader(jsonBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := ts.e.NewContext(req, rec)

	err := ts.handler.ReleaseCourse(c)

	ts.NoError(err)
	ts.Equal(http.StatusBadRequest, rec.Code)
	expectedBody := `{"error":"course has already ended"}`
	ts.Equal(expectedBody, strings.TrimSpace(rec.Body.String()))
	ts.mockIService.AssertExpectations(ts.T())
}

func (ts *HandlerTestSuite) TestReleaseCourseInternalFail() {
	courseCode := "COURSE-REL-ERR"
	jsonBody := `{"code":"` + courseCode + `"}`

	ts.mockIService.On("ReleaseByCode", mock.Anything, courseCode).Return(errors.New("db failed")).Once()

	req := httptest.NewRequest(http.MethodPost, "/api/v1/courses/release", strings.NewReader(jsonBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := ts.e.NewContext(req, rec)

	err := ts.handler.ReleaseCourse(c)

	ts.NoError(err)
	ts.Equal(http.StatusInternalServerError, rec.Code)
	expectedBody := `{"error":"failed to release course"}`
	ts.Equal(expectedBody, strings.TrimSpace(rec.Body.String()))
	ts.mockIService.AssertExpectations(ts.T())
}

// --- InitDummy ---

func (ts *HandlerTestSuite) TestInitDummySuccess() {
	ts.mockIService.On("Save", mock.Anything, mock.AnythingOfType("*course.Course")).
		Return(nil, nil).Times(5)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/courses/init-dummy", nil)
	rec := httptest.NewRecorder()
	c := ts.e.NewContext(req, rec)

	ts.e.POST("/api/v1/courses/init-dummy", ts.handler.InitDummy)
	err := ts.handler.InitDummy(c)

	ts.NoError(err)
	ts.Equal(http.StatusOK, rec.Code)
	ts.Contains(rec.Body.String(), "COURSE-001")
	ts.Contains(rec.Body.String(), "COURSE-005")
	ts.mockIService.AssertExpectations(ts.T())
}

func (ts *HandlerTestSuite) TestInitDummyHttpError() {
	httpErr := &xexception.Http{
		Code:    http.StatusConflict,
		Message: "course code already exists",
	}

	ts.mockIService.On("Save", mock.Anything, mock.AnythingOfType("*course.Course")).
		Return(nil, nil).Once().
		On("Save", mock.Anything, mock.AnythingOfType("*course.Course")).
		Return(nil, httpErr).Once()

	req := httptest.NewRequest(http.MethodPost, "/api/v1/courses/init-dummy", nil)
	rec := httptest.NewRecorder()
	c := ts.e.NewContext(req, rec)

	ts.e.POST("/api/v1/courses/init-dummy", ts.handler.InitDummy)
	err := ts.handler.InitDummy(c)

	ts.NoError(err)
	ts.Equal(http.StatusConflict, rec.Code)
	expectedBody := `{"error":"course code already exists"}`
	ts.Equal(expectedBody, strings.TrimSpace(rec.Body.String()))
	ts.mockIService.AssertExpectations(ts.T())
}

func (ts *HandlerTestSuite) TestInitDummyInternalFail() {
	ts.mockIService.On("Save", mock.Anything, mock.AnythingOfType("*course.Course")).
		Return(nil, errors.New("db error")).Once()

	req := httptest.NewRequest(http.MethodPost, "/api/v1/courses/init-dummy", nil)
	rec := httptest.NewRecorder()
	c := ts.e.NewContext(req, rec)

	ts.e.POST("/api/v1/courses/init-dummy", ts.handler.InitDummy)
	err := ts.handler.InitDummy(c)

	ts.NoError(err)
	ts.Equal(http.StatusInternalServerError, rec.Code)
	expectedBody := `{"error":"failed to insert dummy course"}`
	ts.Equal(expectedBody, strings.TrimSpace(rec.Body.String()))
	ts.mockIService.AssertExpectations(ts.T())
}

// --- CleanDummy ---

func (ts *HandlerTestSuite) TestCleanDummySuccess() {
	courses := []course.Course{
		{ID: "100", Code: "COURSE-001"},
		{ID: "101", Code: "COURSE-002"},
	}

	ts.mockIService.On("FindAll", mock.Anything).Return(courses, nil).Once()
	ts.mockIService.On("DeleteByID", mock.Anything, "100").Return(nil).Once()
	ts.mockIService.On("DeleteByID", mock.Anything, "101").Return(nil).Once()

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/courses/clean-dummy", nil)
	rec := httptest.NewRecorder()
	c := ts.e.NewContext(req, rec)

	ts.e.DELETE("/api/v1/courses/clean-dummy", ts.handler.CleanDummy)
	err := ts.handler.CleanDummy(c)

	ts.NoError(err)
	ts.Equal(http.StatusOK, rec.Code)
	ts.Contains(rec.Body.String(), "dummy courses cleaned")
	ts.mockIService.AssertExpectations(ts.T())
}

func (ts *HandlerTestSuite) TestCleanDummyFindAllFail() {
	ts.mockIService.On("FindAll", mock.Anything).Return(nil, errors.New("db error")).Once()

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/courses/clean-dummy", nil)
	rec := httptest.NewRecorder()
	c := ts.e.NewContext(req, rec)

	err := ts.handler.CleanDummy(c)

	ts.NoError(err)
	ts.Equal(http.StatusInternalServerError, rec.Code)
	expectedBody := `{"error":"failed to fetch courses"}`
	ts.Equal(expectedBody, strings.TrimSpace(rec.Body.String()))
	ts.mockIService.AssertNotCalled(ts.T(), "DeleteByID", mock.Anything, mock.Anything)
	ts.mockIService.AssertExpectations(ts.T())
}

func (ts *HandlerTestSuite) TestCleanDummyDeleteInternalFail() {
	courses := []course.Course{
		{ID: "100", Code: "COURSE-001"},
	}

	ts.mockIService.On("FindAll", mock.Anything).Return(courses, nil).Once()
	ts.mockIService.On("DeleteByID", mock.Anything, "100").Return(errors.New("delete failed")).Once()

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/courses/clean-dummy", nil)
	rec := httptest.NewRecorder()
	c := ts.e.NewContext(req, rec)

	err := ts.handler.CleanDummy(c)

	ts.NoError(err)
	ts.Equal(http.StatusInternalServerError, rec.Code)
	expectedBody := `{"error":"failed to delete dummy course"}`
	ts.Equal(expectedBody, strings.TrimSpace(rec.Body.String()))
	ts.mockIService.AssertExpectations(ts.T())
}

func (ts *HandlerTestSuite) TestCleanDummyDeleteHttpError() {
	courses := []course.Course{
		{ID: "100", Code: "COURSE-001"},
	}
	httpErr := &xexception.Http{
		Code:    http.StatusNotFound,
		Message: "course id not found",
	}

	ts.mockIService.On("FindAll", mock.Anything).Return(courses, nil).Once()
	ts.mockIService.On("DeleteByID", mock.Anything, "100").Return(httpErr).Once()

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/courses/clean-dummy", nil)
	rec := httptest.NewRecorder()
	c := ts.e.NewContext(req, rec)

	err := ts.handler.CleanDummy(c)

	ts.NoError(err)
	ts.Equal(http.StatusNotFound, rec.Code)
	expectedBody := `{"error":"course id not found"}`
	ts.Equal(expectedBody, strings.TrimSpace(rec.Body.String()))
	ts.mockIService.AssertExpectations(ts.T())
}
