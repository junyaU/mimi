package recipes

import (
	"github.com/golang/mock/gomock"
	mockUsecase "infall.com/meacle/command/_mock/usecase"
	"testing"
)

func TestCreateHandler_Handle(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockUsecase.NewMockEventRepository(ctrl)
	mockRepo.EXPECT().AppendToStream(gomock.Any()).Return(nil).AnyTimes()
}
