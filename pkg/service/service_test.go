package service

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	ozon_fintech "ozon-fintech"
	mock_repository "ozon-fintech/pkg/repository/mocks"
	"testing"
)

func TestGetBaseURL(t *testing.T) {
	type mockBehavior func(r *mock_repository.MockRepository, input *ozon_fintech.Link)

	testTable := []struct {
		name               string
		input              *ozon_fintech.Link
		want               string
		mockBehavior       mockBehavior
		expectedStatusCode int
	}{
		{
			name:  "OK",
			input: &ozon_fintech.Link{Token: "abc_012_yz"},
			want:  "https://yandex.ru",
			mockBehavior: func(r *mock_repository.MockRepository, input *ozon_fintech.Link) {
				r.EXPECT().GetBaseURL(gomock.Any(), input).Return("https://yandex.ru", nil)
			},
		},
		{
			name:  "ERROR",
			input: &ozon_fintech.Link{Token: "abc_012_yz"},
			mockBehavior: func(r *mock_repository.MockRepository, input *ozon_fintech.Link) {
				r.EXPECT().GetBaseURL(gomock.Any(), input).Return("", fmt.Errorf("some error"))
			},
		},
		{
			name:  "ERROR_NOT_FOUND",
			input: &ozon_fintech.Link{Token: "abc_012_yz"},
			mockBehavior: func(r *mock_repository.MockRepository, input *ozon_fintech.Link) {
				r.EXPECT().GetBaseURL(gomock.Any(), input).Return("", sql.ErrNoRows)
			},
		},
	}
	c := gomock.NewController(t)
	defer c.Finish()

	repos := mock_repository.NewMockRepository(c)
	service := NewService(repos)

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockBehavior(repos, tc.input)

			got, err := service.GetBaseURL(context.Background(), tc.input)
			if err != nil {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tc.want, got)
			}
		})
	}
}