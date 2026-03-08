package course_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"example.com/course/internal/course"
	"github.com/google/uuid"
	mock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type ServiceTestSuite struct {
	suite.Suite
	mockIStore *MockIStore
	service    course.IService
}

func (ts *ServiceTestSuite) SetupTest() {
	ts.mockIStore = NewMockIStore(ts.T())
	ts.service = course.NewService(ts.mockIStore)
}

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}

// -- FindAll --

func (ts *ServiceTestSuite) TestFindAllSuccess() {
	expected := []course.Course{
		{
			ID:            uuid.New().String(),
			Code:          "COURSE-001",
			Name:          "Go Programming Basics",
			Price:         199.99,
			StartDate:     time.Now(),
			EndDate:       time.Now().Add(7 * 24 * time.Hour),
			Seat:          10,
			SeatAvailable: 10,
		},
		{
			ID:            uuid.New().String(),
			Code:          "COURSE-002",
			Name:          "Advanced Golang",
			Price:         249.99,
			StartDate:     time.Now().Add(-7 * 24 * time.Hour),
			EndDate:       time.Now().Add(-1 * 24 * time.Hour),
			Seat:          0,
			SeatAvailable: 0,
		},
	}
	ts.mockIStore.On("FindAll", mock.Anything).Return(expected, nil).Once()

	result, err := ts.service.FindAll(context.Background())

	ts.NoError(err)
	ts.Equal(2, len(result))
	ts.Equal("Go Programming Basics", result[0].Name)
	ts.mockIStore.AssertExpectations(ts.T())
}

func (ts *ServiceTestSuite) TestFindAllFail() {
	ts.mockIStore.On("FindAll", mock.Anything).Return(nil, errors.New("db error")).Once()

	result, err := ts.service.FindAll(context.Background())

	ts.Error(err)
	ts.Nil(result)
	ts.mockIStore.AssertExpectations(ts.T())
}

// -- FindByID --

func (ts *ServiceTestSuite) TestFindByIDSuccess() {
	expected := &course.Course{
		ID:            uuid.New().String(),
		Code:          "COURSE-001",
		Name:          "Go Programming Basics",
		Price:         199.99,
		StartDate:     time.Now(),
		EndDate:       time.Now().Add(7 * 24 * time.Hour),
		Seat:          10,
		SeatAvailable: 10,
	}
	ts.mockIStore.On("FindByID", mock.Anything, expected.ID).Return(expected, nil).Once()

	result, err := ts.service.FindByID(context.Background(), expected.ID)

	ts.NoError(err)
	ts.Equal(expected, result)
	ts.mockIStore.AssertExpectations(ts.T())
}

func (ts *ServiceTestSuite) TestFindByIDFailNotFoundID() {
	ts.mockIStore.On("FindByID", mock.Anything, "999").Return(nil, errors.New("not found")).Once()

	result, err := ts.service.FindByID(context.Background(), "999")

	ts.Error(err)
	ts.Nil(result)
	ts.Contains(err.Error(), "not found")
	ts.mockIStore.AssertExpectations(ts.T())
}

func (ts *ServiceTestSuite) TestFindByIDFailNotFoundCourse() {
	ts.mockIStore.On("FindByID", mock.Anything, "888").Return(nil, nil).Once()

	result, err := ts.service.FindByID(context.Background(), "888")

	ts.Error(err)
	ts.Nil(result)
	ts.Contains(err.Error(), "not found")
	ts.mockIStore.AssertExpectations(ts.T())
}

// -- FindByCode --

func (ts *ServiceTestSuite) TestFindByCodeSuccess() {
	expected := &course.Course{
		ID:            uuid.New().String(),
		Code:          "COURSE-001",
		Name:          "Go Programming Basics",
		Price:         199.99,
		StartDate:     time.Now(),
		EndDate:       time.Now().Add(7 * 24 * time.Hour),
		Seat:          10,
		SeatAvailable: 10,
	}
	ts.mockIStore.On("FindByCode", mock.Anything, "COURSE-001").Return(expected, nil).Once()

	result, err := ts.service.FindByCode(context.Background(), "COURSE-001")

	ts.NoError(err)
	ts.Equal(expected, result)
	ts.mockIStore.AssertExpectations(ts.T())
}

func (ts *ServiceTestSuite) TestFindByCodeFailNotFoundCode() {
	ts.mockIStore.On("FindByCode", mock.Anything, "ZZZ").Return(nil, errors.New("not found")).Once()

	result, err := ts.service.FindByCode(context.Background(), "ZZZ")

	ts.Error(err)
	ts.Nil(result)
	ts.Contains(err.Error(), "not found")
	ts.mockIStore.AssertExpectations(ts.T())
}

func (ts *ServiceTestSuite) TestFindByCodeFailNotFoundCourse() {
	ts.mockIStore.On("FindByCode", mock.Anything, "YYY").Return(nil, nil).Once()

	result, err := ts.service.FindByCode(context.Background(), "YYY")

	ts.Error(err)
	ts.Nil(result)
	ts.Contains(err.Error(), "not found")
	ts.mockIStore.AssertExpectations(ts.T())
}

// -- Save --

func (ts *ServiceTestSuite) TestSaveSuccess() {
	m := &course.Course{
		ID:            uuid.New().String(),
		Code:          "COURSE-001",
		Name:          "Go Programming Basics",
		Price:         199.99,
		StartDate:     time.Now(),
		EndDate:       time.Now().Add(7 * 24 * time.Hour),
		Seat:          10,
		SeatAvailable: 10,
	}
	ts.mockIStore.On("Save", mock.Anything, m).Return(nil).Once()

	result, err := ts.service.Save(context.Background(), m)

	ts.NoError(err)
	ts.Equal(m, result)
	ts.mockIStore.AssertExpectations(ts.T())
}

func (ts *ServiceTestSuite) TestSaveFail() {
	m := &course.Course{
		ID:            uuid.New().String(),
		Code:          "COURSE-001",
		Name:          "Go Programming Basics",
		Price:         199.99,
		StartDate:     time.Now(),
		EndDate:       time.Now().Add(7 * 24 * time.Hour),
		Seat:          10,
		SeatAvailable: 10,
	}
	ts.mockIStore.On("Save", mock.Anything, m).Return(errors.New("db fail")).Once()

	result, err := ts.service.Save(context.Background(), m)

	ts.Error(err)
	ts.Nil(result)
	ts.mockIStore.AssertExpectations(ts.T())
}

// -- DeleteByID --

func (ts *ServiceTestSuite) TestDeleteByIDSuccess() {
	ts.mockIStore.On("DeleteByID", mock.Anything, "1").Return(nil).Once()

	err := ts.service.DeleteByID(context.Background(), "1")

	ts.NoError(err)
	ts.mockIStore.AssertExpectations(ts.T())
}

func (ts *ServiceTestSuite) TestDeleteByIDFail() {
	ts.mockIStore.On("DeleteByID", mock.Anything, "404").Return(errors.New("delete fail")).Once()

	err := ts.service.DeleteByID(context.Background(), "404")

	ts.Error(err)
	ts.Contains(err.Error(), "delete fail")
	ts.mockIStore.AssertExpectations(ts.T())
}

// -- ReserveByCode --

func (ts *ServiceTestSuite) TestReserveByCodeSuccess() {
	initialSeatAvailable := 5
	m := &course.Course{
		ID:            uuid.New().String(),
		Code:          "COURSE-RES-001",
		Name:          "Reserve Success",
		Price:         300.00,
		StartDate:     time.Now().Add(-1 * time.Hour),
		EndDate:       time.Now().Add(7 * 24 * time.Hour),
		Seat:          10,
		SeatAvailable: initialSeatAvailable,
	}

	ts.mockIStore.On("FindByCode", mock.Anything, "COURSE-RES-001").Return(m, nil).Once()

	expectedAfterReserve := initialSeatAvailable - 1 // 4
	ts.mockIStore.On("Save", mock.Anything, mock.MatchedBy(func(c *course.Course) bool {
		return c.Code == "COURSE-RES-001" && c.SeatAvailable == expectedAfterReserve
	})).Return(nil).Once()

	err := ts.service.ReserveByCode(context.Background(), "COURSE-RES-001")

	ts.NoError(err)
	ts.mockIStore.AssertExpectations(ts.T())
	ts.Equal(expectedAfterReserve, m.SeatAvailable)
}

func (ts *ServiceTestSuite) TestReserveByCodeFailNotFoundCode() {
	ts.mockIStore.On("FindByCode", mock.Anything, "ZZZ-RES").Return(nil, errors.New("db error")).Once()

	err := ts.service.ReserveByCode(context.Background(), "ZZZ-RES")

	ts.Error(err)
	ts.Contains(err.Error(), "course not found by code")
	ts.mockIStore.AssertExpectations(ts.T())
}

func (ts *ServiceTestSuite) TestReserveByCodeFailNotFoundCourse() {
	ts.mockIStore.On("FindByCode", mock.Anything, "YYY-RES").Return(nil, nil).Once()

	err := ts.service.ReserveByCode(context.Background(), "YYY-RES")

	ts.Error(err)
	ts.Contains(err.Error(), "course not found")
	ts.mockIStore.AssertExpectations(ts.T())
}

func (ts *ServiceTestSuite) TestReserveByCodeFailCourseEnded() {
	m := &course.Course{
		Code:          "COURSE-RES-END",
		EndDate:       time.Now().Add(-1 * time.Hour),
		SeatAvailable: 5,
	}

	ts.mockIStore.On("FindByCode", mock.Anything, "COURSE-RES-END").Return(m, nil).Once()

	err := ts.service.ReserveByCode(context.Background(), "COURSE-RES-END")

	ts.Error(err)
	ts.Contains(err.Error(), "course has already ended")
	ts.mockIStore.AssertNotCalled(ts.T(), "Save", mock.Anything, mock.Anything)
	ts.mockIStore.AssertExpectations(ts.T())
}

func (ts *ServiceTestSuite) TestReserveByCodeFailNoSeatAvailable() {
	m := &course.Course{
		Code:          "COURSE-RES-FULL",
		EndDate:       time.Now().Add(7 * 24 * time.Hour),
		SeatAvailable: 0,
	}

	ts.mockIStore.On("FindByCode", mock.Anything, "COURSE-RES-FULL").Return(m, nil).Once()

	err := ts.service.ReserveByCode(context.Background(), "COURSE-RES-FULL")

	ts.Error(err)
	ts.Contains(err.Error(), "no available seats to reserve")
	ts.mockIStore.AssertNotCalled(ts.T(), "Save", mock.Anything, mock.Anything)
	ts.mockIStore.AssertExpectations(ts.T())
}

func (ts *ServiceTestSuite) TestReserveByCodeFailSave() {
	initialSeatAvailable := 1
	m := &course.Course{
		Code:          "COURSE-RES-SAVEFAIL",
		EndDate:       time.Now().Add(7 * 24 * time.Hour),
		SeatAvailable: initialSeatAvailable,
	}

	ts.mockIStore.On("FindByCode", mock.Anything, "COURSE-RES-SAVEFAIL").Return(m, nil).Once()
	ts.mockIStore.On("Save", mock.Anything, mock.Anything).Return(errors.New("db save failed")).Once()

	err := ts.service.ReserveByCode(context.Background(), "COURSE-RES-SAVEFAIL")

	ts.Error(err)
	ts.Contains(err.Error(), "db save failed")
	ts.mockIStore.AssertExpectations(ts.T())
	ts.Equal(initialSeatAvailable-1, m.SeatAvailable)
}

// -- ReleaseByCode --

func (ts *ServiceTestSuite) TestReleaseByCodeSuccess() {
	initialSeatAvailable := 5
	m := &course.Course{
		ID:            uuid.New().String(),
		Code:          "COURSE-REL-001",
		Name:          "Release Success",
		Price:         300.00,
		StartDate:     time.Now().Add(-1 * time.Hour),
		EndDate:       time.Now().Add(7 * 24 * time.Hour),
		Seat:          10,
		SeatAvailable: initialSeatAvailable,
	}

	ts.mockIStore.On("FindByCode", mock.Anything, "COURSE-REL-001").Return(m, nil).Once()

	expectedAfterRelease := initialSeatAvailable + 1 // 6
	ts.mockIStore.On("Save", mock.Anything, mock.MatchedBy(func(c *course.Course) bool {
		return c.Code == "COURSE-REL-001" &&
			c.SeatAvailable == expectedAfterRelease
	})).Return(nil).Once()

	err := ts.service.ReleaseByCode(context.Background(), "COURSE-REL-001")

	ts.NoError(err)
	ts.Equal(expectedAfterRelease, m.SeatAvailable)
	ts.mockIStore.AssertExpectations(ts.T())
}

func (ts *ServiceTestSuite) TestReleaseByCodeFailNotFoundCode() {

	ts.mockIStore.On("FindByCode", mock.Anything, "ZZZ-REL").
		Return(nil, errors.New("db error")).Once()

	err := ts.service.ReleaseByCode(context.Background(), "ZZZ-REL")

	ts.Error(err)
	ts.Contains(err.Error(), "course not found by code")
	ts.mockIStore.AssertExpectations(ts.T())
}

func (ts *ServiceTestSuite) TestReleaseByCodeFailNotFoundCourse() {

	ts.mockIStore.On("FindByCode", mock.Anything, "YYY-REL").
		Return(nil, nil).Once()

	err := ts.service.ReleaseByCode(context.Background(), "YYY-REL")

	ts.Error(err)
	ts.Contains(err.Error(), "course not found")
	ts.mockIStore.AssertExpectations(ts.T())
}

func (ts *ServiceTestSuite) TestReleaseByCodeFailCourseEnded() {

	m := &course.Course{
		Code:          "COURSE-REL-END",
		EndDate:       time.Now().Add(-1 * time.Hour), // already ended
		Seat:          10,
		SeatAvailable: 5,
	}

	ts.mockIStore.On("FindByCode", mock.Anything, "COURSE-REL-END").
		Return(m, nil).Once()

	err := ts.service.ReleaseByCode(context.Background(), "COURSE-REL-END")

	ts.Error(err)
	ts.Contains(err.Error(), "course has already ended")
	ts.mockIStore.AssertNotCalled(ts.T(), "Save", mock.Anything, mock.Anything)
	ts.mockIStore.AssertExpectations(ts.T())
}

func (ts *ServiceTestSuite) TestReleaseByCodeFailSeatAlreadyFull() {
	m := &course.Course{
		Code:          "COURSE-REL-FULL",
		Seat:          10,
		SeatAvailable: 10, // FULL — cannot release more
		EndDate:       time.Now().Add(2 * time.Hour),
	}

	ts.mockIStore.On("FindByCode", mock.Anything, "COURSE-REL-FULL").
		Return(m, nil).Once()

	err := ts.service.ReleaseByCode(context.Background(), "COURSE-REL-FULL")

	ts.Error(err)
	ts.Contains(err.Error(), "all seats are already available")

	ts.mockIStore.AssertNotCalled(ts.T(), "Save", mock.Anything, mock.Anything)
	ts.mockIStore.AssertExpectations(ts.T())
}

func (ts *ServiceTestSuite) TestReleaseByCodeFailSave() {

	initialSeatAvailable := 5
	m := &course.Course{
		Code:          "COURSE-REL-SAVEFAIL",
		Seat:          10,
		EndDate:       time.Now().Add(7 * 24 * time.Hour), // valid course
		SeatAvailable: initialSeatAvailable,
	}

	ts.mockIStore.On("FindByCode", mock.Anything, "COURSE-REL-SAVEFAIL").
		Return(m, nil).Once()

	ts.mockIStore.On("Save", mock.Anything, mock.Anything).
		Return(errors.New("db save failed")).Once()

	err := ts.service.ReleaseByCode(context.Background(), "COURSE-REL-SAVEFAIL")

	ts.Error(err)
	ts.Contains(err.Error(), "db save failed")

	// Seat must still increment even if save fails
	ts.Equal(initialSeatAvailable+1, m.SeatAvailable)

	ts.mockIStore.AssertExpectations(ts.T())
}
