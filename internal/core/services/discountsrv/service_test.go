package discountsrv

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mendezdev/mytheresa-api/internal/core/domain"
	"github.com/mendezdev/mytheresa-api/internal/core/ports"
	"github.com/mendezdev/mytheresa-api/pkg/apierrors"
	"github.com/stretchr/testify/assert"
)

func TestDiscountOk(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := ports.NewMockDiscountRepository(mockCtrl)
	srv := New(mockRepo)

	d1 := domain.Discount{}
	d2 := domain.Discount{}

	discounts := make([]domain.Discount, 0)
	discounts = append(discounts, d1, d2)

	mockRepo.
		EXPECT().
		GetDiscounts().
		Return(discounts, nil).
		Times(1)

	results, err := srv.GetDiscounts()
	assert.Nil(t, err)
	assert.Len(t, results, 2)
}

func TestDiscountErr(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := ports.NewMockDiscountRepository(mockCtrl)
	srv := New(mockRepo)

	mockRepo.
		EXPECT().
		GetDiscounts().
		Return(nil, apierrors.NewInternalServerError("error")).
		Times(1)

	_, err := srv.GetDiscounts()
	assert.NotNil(t, err)
}
