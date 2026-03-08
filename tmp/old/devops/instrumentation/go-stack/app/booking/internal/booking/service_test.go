package booking_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"example.com/booking/internal/booking"
	mock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type ServiceTestSuite struct {
	suite.Suite
	mockIStore  *MockIStore
	mockIClient *MockIClient
	service     booking.IService
}

func (ts *ServiceTestSuite) SetupTest() {
	ts.mockIStore = new(MockIStore)
	ts.mockIClient = new(MockIClient)
	ts.service = booking.NewService(ts.mockIStore, ts.mockIClient)
}

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}

// --- Booking (Synchronous) ---

func (ts *ServiceTestSuite) TestBookingSuccess() {
	req := booking.BookingRequest{
		MemberCode: "MEM-001",
		CourseCode: "COURSE-A",
	}
	memberResp := &booking.MemberServiceMemberResponse{Code: "MEM-001", Name: "John Doe"}
	courseResp := &booking.CourseServiceMessageResponse{Message: "Reserved successfully"}

	ts.mockIClient.On("FindMemberByCode", mock.Anything, "MEM-001").Return(memberResp, nil).Once()
	ts.mockIClient.On("ReserveCourseByCode", mock.Anything, "COURSE-A").Return(courseResp, nil).Once()
	ts.mockIStore.On("Save", mock.Anything, mock.AnythingOfType("*booking.Booking")).Return(nil).Once()

	result, err := ts.service.Booking(context.Background(), req)

	ts.NoError(err)
	ts.NotNil(result)
	ts.Equal("MEM-001", result.MemberCode)
	ts.Equal("COURSE-A", result.CourseCode)
	ts.mockIClient.AssertExpectations(ts.T())
	ts.mockIStore.AssertExpectations(ts.T())
}

func (ts *ServiceTestSuite) TestBookingFailFindMember() {
	req := booking.BookingRequest{MemberCode: "MEM-404", CourseCode: "COURSE-A"}
	memberErr := errors.New("member not found in client")

	ts.mockIClient.On("FindMemberByCode", mock.Anything, "MEM-404").Return(nil, memberErr).Once()

	result, err := ts.service.Booking(context.Background(), req)

	ts.Error(err)
	ts.Nil(result)
	ts.Contains(err.Error(), "member not found")
	ts.mockIClient.AssertExpectations(ts.T())
	ts.mockIClient.AssertNotCalled(ts.T(), "ReserveCourseByCode", mock.Anything, mock.Anything)
	ts.mockIStore.AssertNotCalled(ts.T(), "Save", mock.Anything, mock.Anything)
}

func (ts *ServiceTestSuite) TestBookingFailReserveCourse() {
	req := booking.BookingRequest{MemberCode: "MEM-001", CourseCode: "COURSE-B"}
	memberResp := &booking.MemberServiceMemberResponse{Code: "MEM-001"}
	courseErr := errors.New("course fully booked")

	ts.mockIClient.On("FindMemberByCode", mock.Anything, "MEM-001").Return(memberResp, nil).Once()
	ts.mockIClient.On("ReserveCourseByCode", mock.Anything, "COURSE-B").Return(nil, courseErr).Once()

	result, err := ts.service.Booking(context.Background(), req)

	ts.Error(err)
	ts.Nil(result)
	ts.Contains(err.Error(), "failed to reserve course")
	ts.mockIClient.AssertExpectations(ts.T())
	ts.mockIStore.AssertNotCalled(ts.T(), "Save", mock.Anything, mock.Anything)
}

func (ts *ServiceTestSuite) TestBookingFailSaveAndReleaseSuccess() {
	req := booking.BookingRequest{MemberCode: "MEM-001", CourseCode: "COURSE-C"}
	memberResp := &booking.MemberServiceMemberResponse{Code: "MEM-001"}
	courseResp := &booking.CourseServiceMessageResponse{Message: "Reserved"}
	saveErr := errors.New("db save failed")
	releaseResp := &booking.CourseServiceMessageResponse{Message: "Released"}

	ts.mockIClient.On("FindMemberByCode", mock.Anything, "MEM-001").Return(memberResp, nil).Once()
	ts.mockIClient.On("ReserveCourseByCode", mock.Anything, "COURSE-C").Return(courseResp, nil).Once()
	ts.mockIStore.On("Save", mock.Anything, mock.AnythingOfType("*booking.Booking")).Return(saveErr).Once()
	ts.mockIClient.On("ReleaseCourseByCode", mock.Anything, "COURSE-C").Return(releaseResp, nil).Once() // Rollback Success

	result, err := ts.service.Booking(context.Background(), req)

	ts.Error(err)
	ts.Nil(result)
	ts.Contains(err.Error(), "failed to save booking")
	ts.mockIClient.AssertExpectations(ts.T())
	ts.mockIStore.AssertExpectations(ts.T())
}

func (ts *ServiceTestSuite) TestBookingFailSaveAndReleaseFail() {
	req := booking.BookingRequest{MemberCode: "MEM-001", CourseCode: "COURSE-D"}
	memberResp := &booking.MemberServiceMemberResponse{Code: "MEM-001"}
	courseResp := &booking.CourseServiceMessageResponse{Message: "Reserved"}
	saveErr := errors.New("db save failed")
	releaseErr := errors.New("release failed too")

	ts.mockIClient.On("FindMemberByCode", mock.Anything, "MEM-001").Return(memberResp, nil).Once()
	ts.mockIClient.On("ReserveCourseByCode", mock.Anything, "COURSE-D").Return(courseResp, nil).Once()
	ts.mockIStore.On("Save", mock.Anything, mock.AnythingOfType("*booking.Booking")).Return(saveErr).Once()
	ts.mockIClient.On("ReleaseCourseByCode", mock.Anything, "COURSE-D").Return(nil, releaseErr).Once() // Rollback Fail

	result, err := ts.service.Booking(context.Background(), req)

	ts.Error(err)
	ts.Nil(result)
	ts.Contains(err.Error(), "failed to save booking")
	ts.mockIClient.AssertExpectations(ts.T())
	ts.mockIStore.AssertExpectations(ts.T())
}

// --- BookingV2 (Concurrent) ---

func (ts *ServiceTestSuite) TestBookingV2Success() {
	req := booking.BookingRequest{
		MemberCode: "MEM-002",
		CourseCode: "COURSE-X",
	}
	memberResp := &booking.MemberServiceMemberResponse{Code: "MEM-002", Name: "Jane Smith"}
	courseResp := &booking.CourseServiceMessageResponse{Message: "Reserved V2"}

	ts.mockIClient.On("FindMemberByCode", mock.Anything, "MEM-002").Return(memberResp, nil).Once()
	ts.mockIClient.On("ReserveCourseByCode", mock.Anything, "COURSE-X").Return(courseResp, nil).Once()
	ts.mockIStore.On("Save", mock.Anything, mock.AnythingOfType("*booking.Booking")).Return(nil).Once()

	result, err := ts.service.BookingV2(context.Background(), req)

	ts.NoError(err)
	ts.NotNil(result)
	ts.Equal("MEM-002", result.MemberCode)
	ts.Equal("COURSE-X", result.CourseCode)
	ts.mockIClient.AssertExpectations(ts.T())
	ts.mockIStore.AssertExpectations(ts.T())
}

func (ts *ServiceTestSuite) TestBookingV2FailMemberNotFound() {
	req := booking.BookingRequest{MemberCode: "MEM-FAIL", CourseCode: "COURSE-Y"}
	memberErr := errors.New("member not found client error")
	courseResp := &booking.CourseServiceMessageResponse{Message: "Reserved V2"}

	ts.mockIClient.On("FindMemberByCode", mock.Anything, "MEM-FAIL").Return(nil, memberErr).Once()
	ts.mockIClient.On("ReserveCourseByCode", mock.Anything, "COURSE-Y").Return(courseResp, nil).Once()

	result, err := ts.service.BookingV2(context.Background(), req)

	ts.Error(err)
	ts.Nil(result)
	ts.Contains(err.Error(), "member not found")
	ts.mockIClient.AssertExpectations(ts.T())
	ts.mockIStore.AssertNotCalled(ts.T(), "Save", mock.Anything, mock.Anything)
}

func (ts *ServiceTestSuite) TestBookingV2FailCourseReserve() {
	req := booking.BookingRequest{MemberCode: "MEM-002", CourseCode: "COURSE-Z-FAIL"}
	memberResp := &booking.MemberServiceMemberResponse{Code: "MEM-002"}
	courseErr := errors.New("course full client error")

	ts.mockIClient.On("FindMemberByCode", mock.Anything, "MEM-002").Return(memberResp, nil).Once()
	ts.mockIClient.On("ReserveCourseByCode", mock.Anything, "COURSE-Z-FAIL").Return(nil, courseErr).Once()

	result, err := ts.service.BookingV2(context.Background(), req)

	ts.Error(err)
	ts.Nil(result)
	ts.Contains(err.Error(), "failed to reserve course")
	ts.mockIClient.AssertExpectations(ts.T())
	ts.mockIStore.AssertNotCalled(ts.T(), "Save", mock.Anything, mock.Anything)
}

func (ts *ServiceTestSuite) TestBookingV2FailSaveAndRollback() {
	req := booking.BookingRequest{
		MemberCode: "MEM-002",
		CourseCode: "COURSE-ROLLBACK",
	}
	memberResp := &booking.MemberServiceMemberResponse{Code: "MEM-002"}
	courseResp := &booking.CourseServiceMessageResponse{Message: "Reserved V2"}
	saveErr := errors.New("db save failed V2")
	releaseErr := errors.New("release failed too V2")

	ts.mockIClient.On("FindMemberByCode", mock.Anything, "MEM-002").Return(memberResp, nil).Once()
	ts.mockIClient.On("ReserveCourseByCode", mock.Anything, "COURSE-ROLLBACK").Return(courseResp, nil).Once()
	ts.mockIStore.On("Save", mock.Anything, mock.AnythingOfType("*booking.Booking")).Return(saveErr).Once()
	ts.mockIClient.On("ReleaseCourseByCode", mock.Anything, "COURSE-ROLLBACK").Return(nil, releaseErr).Once()

	result, err := ts.service.BookingV2(context.Background(), req)

	ts.Error(err)
	ts.Nil(result)
	ts.Contains(err.Error(), "failed to save booking")
	ts.mockIClient.AssertCalled(ts.T(), "ReleaseCourseByCode", mock.Anything, "COURSE-ROLLBACK")
	ts.mockIClient.AssertExpectations(ts.T())
	ts.mockIStore.AssertExpectations(ts.T())
}

func (ts *ServiceTestSuite) TestBookingV2FailTimeout() {
	req := booking.BookingRequest{MemberCode: "MEM-TIMEOUT", CourseCode: "COURSE-TIMEOUT"}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
	defer cancel()

	ts.mockIClient.On("FindMemberByCode", mock.Anything, "MEM-TIMEOUT").Return(nil, errors.New("simulated block")).WaitUntil(time.After(50 * time.Millisecond)).Once()
	ts.mockIClient.On("ReserveCourseByCode", mock.Anything, "COURSE-TIMEOUT").Return(nil, errors.New("simulated block")).WaitUntil(time.After(50 * time.Millisecond)).Once()

	time.Sleep(10 * time.Millisecond)

	result, err := ts.service.BookingV2(ctx, req)

	ts.Error(err)
	ts.Nil(result)
	ts.Contains(err.Error(), "context cancelled or timed out")
	ts.mockIStore.AssertNotCalled(ts.T(), "Save", mock.Anything, mock.Anything)
}
