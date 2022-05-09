package mock

import gomock "github.com/golang/mock/gomock"

type MockUsecases struct {
}

func NewMockUsecases(mockCtrl *gomock.Controller) MockUsecases {
	return MockUsecases{}
}
