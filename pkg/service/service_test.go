package service

import (
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
		name         string
		input        *ozon_fintech.Link
		want         string
		mockBehavior mockBehavior
	}{
		{
			name:  "OK",
			input: &ozon_fintech.Link{Token: "abc_012_yz"},
			want:  "https://yandex.ru",
			mockBehavior: func(r *mock_repository.MockRepository, input *ozon_fintech.Link) {
				r.EXPECT().GetBaseURL(input).Return("https://yandex.ru", nil)
			},
		},
		{
			name:  "ERROR",
			input: &ozon_fintech.Link{Token: "abc_012_yz"},
			mockBehavior: func(r *mock_repository.MockRepository, input *ozon_fintech.Link) {
				r.EXPECT().GetBaseURL(input).Return("", fmt.Errorf("some error"))
			},
		},
		{
			name:  "ERROR_NOT_FOUND",
			input: &ozon_fintech.Link{Token: "abc_012_yz"},
			mockBehavior: func(r *mock_repository.MockRepository, input *ozon_fintech.Link) {
				r.EXPECT().GetBaseURL(input).Return("", sql.ErrNoRows)
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

			got, err := service.GetBaseURL(tc.input)
			if err != nil {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tc.want, got)
			}
		})
	}
}

func TestCreateShortURL(t *testing.T) {
	type mockBehavior func(r *mock_repository.MockRepository, input *ozon_fintech.Link)

	testTable := []struct {
		name         string
		input        *ozon_fintech.Link
		want         string
		mockBehavior mockBehavior
	}{
		{
			name:  "OK",
			input: &ozon_fintech.Link{BaseURL: "https://yandex.ru"},
			want:  "TOKEN_1234",
			mockBehavior: func(r *mock_repository.MockRepository, input *ozon_fintech.Link) {
				r.EXPECT().CreateShortURL(input).Return("TOKEN_1234", nil)
			},
		},
		{
			name:  "ERROR",
			input: &ozon_fintech.Link{BaseURL: "https://yandex.ru"},
			mockBehavior: func(r *mock_repository.MockRepository, input *ozon_fintech.Link) {
				r.EXPECT().CreateShortURL(input).Return("", fmt.Errorf("some error"))
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

			got, err := service.CreateShortURL(tc.input)
			if err != nil {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tc.want, got)
			}
		})
	}
}
