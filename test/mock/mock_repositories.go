package mock

import gomock "github.com/golang/mock/gomock"

type MockRepositories struct {
	MockITokenRepository *MockITokenRepository
}

func NewMockRepository(mockCtrl *gomock.Controller) MockRepositories {
	return MockRepositories{
		MockITokenRepository: NewMockITokenRepository(mockCtrl),
	}
}
