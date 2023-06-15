package service

import (
	"test/internal/models"
	"test/internal/repository/mock"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestGetDataService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	GetDataService := mock.NewMockGetData(ctrl)

	link := models.Data{
		Id:           1,
		Active_link:  "internal/repoitory/mock/mock.go",
		History_link: "mock/mock.go/repository/internal",
	}

	testcases := []struct {
		mockFuncs func()
		wantErr   bool
	}{
		{
			mockFuncs: func() {
				GetDataService.EXPECT().GetLinkByID(link.Id).Return(link, nil)
			},
			wantErr: false,
		},
	}

	for _, test := range testcases {
		test.mockFuncs()
		_, err := GetDataService.GetLinkByID(link.Id)
		if test.wantErr && err == nil {
			t.Error("expected error")
		} else if !test.wantErr && err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	}
}
