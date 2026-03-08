package booking_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"example.com/booking/internal/booking"
	"github.com/agussyahrilmubarok/gox/pkg/xexception"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	mock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type HandlerTestSuite struct {
	suite.Suite
	mockIService *MockIService
	handler      *booking.Handler
	e            *echo.Echo
}

func (ts *HandlerTestSuite) SetupTest() {
	ts.mockIService = NewMockIService(ts.T())
	ts.handler = booking.NewHandler(ts.mockIService)
	ts.e = echo.New()
	ts.e.POST("/api/v1/booking/course", ts.handler.Booking)
	ts.e.POST("/api/v1/booking/course", ts.handler.BookingV2)
}

func TestHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(HandlerTestSuite))
}

// --- Helper Functions ---

func createBookingBody(memberCode, courseCode string) string {
	return `{"member_code":"` + memberCode + `","course_code":"` + courseCode + `"}`
}

func createValidBooking(memberCode, courseCode string) *booking.Booking {
	return &booking.Booking{
		ID:          uuid.New().String(),
		MemberCode:  memberCode,
		CourseCode:  courseCode,
		BookingDate: time.Now(),
		Status:      booking.BookingStatusPending,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

// --- Booking (V1) ---

func (ts *HandlerTestSuite) TestBookingSuccess() {
	memberCode := "MEM-101"
	courseCode := "COURSE-GOLANG"
	bookingData := createValidBooking(memberCode, courseCode)
	jsonBody := createBookingBody(memberCode, courseCode)

	ts.mockIService.On("Booking", mock.Anything, booking.BookingRequest{
		MemberCode: memberCode,
		CourseCode: courseCode,
	}).Return(bookingData, nil).Once()

	req := httptest.NewRequest(http.MethodPost, "/api/v1/booking/course", strings.NewReader(jsonBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := ts.e.NewContext(req, rec)

	err := ts.handler.Booking(c)

	ts.NoError(err)
	ts.Equal(http.StatusOK, rec.Code)
	ts.Contains(rec.Body.String(), memberCode)
	ts.Contains(rec.Body.String(), courseCode)
	ts.mockIService.AssertExpectations(ts.T())
}

func (ts *HandlerTestSuite) TestBookingInvalidPayload() {
	jsonBody := `{"member_code":123,"course_code":"COURSE-GOLANG"}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/booking/course", strings.NewReader(jsonBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := ts.e.NewContext(req, rec)

	err := ts.handler.Booking(c)

	ts.NoError(err)
	ts.Equal(http.StatusBadRequest, rec.Code)
	expectedBody := `{"error":"invalid request payload"}`
	ts.Equal(expectedBody, strings.TrimSpace(rec.Body.String()))
	ts.mockIService.AssertNotCalled(ts.T(), "Booking", mock.Anything, mock.Anything)
}

func (ts *HandlerTestSuite) TestBookingHttpErrorFromService() {
	memberCode := "MEM-404"
	courseCode := "COURSE-FULL"
	jsonBody := createBookingBody(memberCode, courseCode)
	httpErr := xexception.NewHTTPNotFound("member not found", errors.New("client error"))

	ts.mockIService.On("Booking", mock.Anything, mock.AnythingOfType("booking.BookingRequest")).
		Return(nil, httpErr).Once()

	req := httptest.NewRequest(http.MethodPost, "/api/v1/booking/course", strings.NewReader(jsonBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := ts.e.NewContext(req, rec)

	err := ts.handler.Booking(c)

	ts.NoError(err)
	ts.Equal(http.StatusNotFound, rec.Code)
	expectedBody := `{"error":"member not found"}`
	ts.Equal(expectedBody, strings.TrimSpace(rec.Body.String()))
	ts.mockIService.AssertExpectations(ts.T())
}

func (ts *HandlerTestSuite) TestBookingInternalErrorFromService() {
	memberCode := "MEM-101"
	courseCode := "COURSE-DB-ERR"
	jsonBody := createBookingBody(memberCode, courseCode)
	internalErr := errors.New("database connection failed")

	ts.mockIService.On("Booking", mock.Anything, mock.AnythingOfType("booking.BookingRequest")).
		Return(nil, internalErr).Once()

	req := httptest.NewRequest(http.MethodPost, "/api/v1/booking/course", strings.NewReader(jsonBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := ts.e.NewContext(req, rec)

	err := ts.handler.Booking(c)

	ts.NoError(err)
	ts.Equal(http.StatusInternalServerError, rec.Code)
	expectedBody := `{"error":"failed to book course"}`
	ts.Equal(expectedBody, strings.TrimSpace(rec.Body.String()))
	ts.mockIService.AssertExpectations(ts.T())
}

// --- BookingV2 (V2) ---

func (ts *HandlerTestSuite) TestBookingV2Success() {
	memberCode := "MEM-202"
	courseCode := "COURSE-K8S"
	bookingData := createValidBooking(memberCode, courseCode)
	jsonBody := createBookingBody(memberCode, courseCode)

	ts.mockIService.On("BookingV2", mock.Anything, booking.BookingRequest{
		MemberCode: memberCode,
		CourseCode: courseCode,
	}).Return(bookingData, nil).Once()

	req := httptest.NewRequest(http.MethodPost, "/api/v2/booking/course", strings.NewReader(jsonBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := ts.e.NewContext(req, rec)

	err := ts.handler.BookingV2(c)

	ts.NoError(err)
	ts.Equal(http.StatusOK, rec.Code)
	ts.Contains(rec.Body.String(), memberCode)
	ts.mockIService.AssertExpectations(ts.T())
}

func (ts *HandlerTestSuite) TestBookingV2InvalidPayload() {
	jsonBody := `{"member_code":"MEM-101", "course_code":[]}`
	req := httptest.NewRequest(http.MethodPost, "/api/v2/booking/course", strings.NewReader(jsonBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := ts.e.NewContext(req, rec)

	err := ts.handler.BookingV2(c)

	ts.NoError(err)
	ts.Equal(http.StatusBadRequest, rec.Code)
	expectedBody := `{"error":"invalid request payload"}`
	ts.Equal(expectedBody, strings.TrimSpace(rec.Body.String()))
	ts.mockIService.AssertNotCalled(ts.T(), "Booking", mock.Anything, mock.Anything)
	ts.mockIService.AssertNotCalled(ts.T(), "BookingV2", mock.Anything, mock.Anything)
}

func (ts *HandlerTestSuite) TestBookingV2HttpErrorFromService() {
	memberCode := "MEM-400"
	courseCode := "COURSE-TIMEOUT"
	jsonBody := createBookingBody(memberCode, courseCode)
	httpErr := xexception.NewHTTPBadRequest("request timed out", errors.New("timeout client"))

	ts.mockIService.On("BookingV2", mock.Anything, mock.AnythingOfType("booking.BookingRequest")).
		Return(nil, httpErr).Once()

	req := httptest.NewRequest(http.MethodPost, "/api/v2/booking/course", strings.NewReader(jsonBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := ts.e.NewContext(req, rec)

	err := ts.handler.BookingV2(c)

	ts.NoError(err)
	ts.Equal(http.StatusBadRequest, rec.Code)
	expectedBody := `{"error":"request timed out"}`
	ts.Equal(expectedBody, strings.TrimSpace(rec.Body.String()))
	ts.mockIService.AssertExpectations(ts.T())
}

func (ts *HandlerTestSuite) TestBookingV2InternalErrorFromService() {
	memberCode := "MEM-303"
	courseCode := "COURSE-TX-ERR"
	jsonBody := createBookingBody(memberCode, courseCode)
	internalErr := errors.New("transaction failed")

	ts.mockIService.On("BookingV2", mock.Anything, mock.AnythingOfType("booking.BookingRequest")).
		Return(nil, internalErr).Once()

	req := httptest.NewRequest(http.MethodPost, "/api/v2/booking/course", strings.NewReader(jsonBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := ts.e.NewContext(req, rec)

	err := ts.handler.BookingV2(c)

	ts.NoError(err)
	ts.Equal(http.StatusInternalServerError, rec.Code)
	expectedBody := `{"error":"failed to book course"}`
	ts.Equal(expectedBody, strings.TrimSpace(rec.Body.String()))
	ts.mockIService.AssertExpectations(ts.T())
}
