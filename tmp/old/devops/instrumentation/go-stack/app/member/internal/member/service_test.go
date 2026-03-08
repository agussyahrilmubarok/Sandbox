package member_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"example.com/member/internal/member"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type ServiceTestSuite struct {
	suite.Suite
	mockIStore *MockIStore
	service    member.IService
}

func (ts *ServiceTestSuite) SetupTest() {
	ts.mockIStore = NewMockIStore(ts.T())
	ts.service = member.NewService(ts.mockIStore)
}

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}

// -- FindAll --

func (ts *ServiceTestSuite) TestFindAllSuccess() {
	expected := []member.Member{
		{ID: "1", Name: "John"},
		{ID: "2", Name: "Doe"},
	}
	ts.mockIStore.On("FindAll", mock.Anything).Return(expected, nil).Once()

	result, err := ts.service.FindAll(context.Background())

	ts.NoError(err)
	ts.Equal(2, len(result))
	ts.Equal("John", result[0].Name)
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
	expected := &member.Member{ID: "1", Name: "John"}
	ts.mockIStore.On("FindByID", mock.Anything, "1").Return(expected, nil).Once()

	result, err := ts.service.FindByID(context.Background(), "1")

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

func (ts *ServiceTestSuite) TestFindByIDFailNotFoundMember() {
	ts.mockIStore.On("FindByID", mock.Anything, "888").Return(nil, nil).Once()

	result, err := ts.service.FindByID(context.Background(), "888")

	ts.Error(err)
	ts.Nil(result)
	ts.Contains(err.Error(), "not found")
	ts.mockIStore.AssertExpectations(ts.T())
}

// -- FindByCode --

func (ts *ServiceTestSuite) TestFindByCodeSuccess() {
	expected := &member.Member{ID: "1", Code: "ABC123"}
	ts.mockIStore.On("FindByCode", mock.Anything, "ABC123").Return(expected, nil).Once()

	result, err := ts.service.FindByCode(context.Background(), "ABC123")

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

func (ts *ServiceTestSuite) TestFindByCodeFailNotFoundMember() {
	ts.mockIStore.On("FindByCode", mock.Anything, "YYY").Return(nil, nil).Once()

	result, err := ts.service.FindByCode(context.Background(), "YYY")

	ts.Error(err)
	ts.Nil(result)
	ts.Contains(err.Error(), "not found")
	ts.mockIStore.AssertExpectations(ts.T())
}

// -- FindByEmail --

func (ts *ServiceTestSuite) TestFindByEmailSuccess() {
	expected := &member.Member{ID: "1", Email: "a@b.com"}
	ts.mockIStore.On("FindByEmail", mock.Anything, "a@b.com").Return(expected, nil).Once()

	result, err := ts.service.FindByEmail(context.Background(), "a@b.com")

	ts.NoError(err)
	ts.Equal(expected, result)
	ts.mockIStore.AssertExpectations(ts.T())
}

func (ts *ServiceTestSuite) TestFindByEmailFailNotFoundEmail() {
	ts.mockIStore.On("FindByEmail", mock.Anything, "x@y.com").Return(nil, errors.New("not found")).Once()

	result, err := ts.service.FindByEmail(context.Background(), "x@y.com")

	ts.Error(err)
	ts.Nil(result)
	ts.Contains(err.Error(), "not found")
	ts.mockIStore.AssertExpectations(ts.T())
}

func (ts *ServiceTestSuite) TestFindByEmailFailNotFoundMember() {
	ts.mockIStore.On("FindByEmail", mock.Anything, "none@b.com").Return(nil, nil).Once()

	result, err := ts.service.FindByEmail(context.Background(), "none@b.com")

	ts.Error(err)
	ts.Nil(result)
	ts.Contains(err.Error(), "not found")
	ts.mockIStore.AssertExpectations(ts.T())
}

// -- Save --

func (ts *ServiceTestSuite) TestSaveSuccess() {
	m := &member.Member{ID: "1", Code: "A", Name: "B", CreatedAt: time.Now()}
	ts.mockIStore.On("Save", mock.Anything, m).Return(nil).Once()

	result, err := ts.service.Save(context.Background(), m)

	ts.NoError(err)
	ts.Equal(m, result)
	ts.mockIStore.AssertExpectations(ts.T())
}

func (ts *ServiceTestSuite) TestSaveFail() {
	m := &member.Member{ID: "1", Code: "A", Name: "B"}
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
