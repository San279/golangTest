package repositories

import "github.com/stretchr/testify/mock"

type promotionRepositoryMock struct {
	mock.Mock
}

func NewPromotionRepositoryMock() *promotionRepositoryMock {
	return &promotionRepositoryMock{}
}

func (p *promotionRepositoryMock) GetPromotion() (Promotion, error) {
	args := p.Called()
	return args.Get(0).(Promotion), args.Error(1)

}
