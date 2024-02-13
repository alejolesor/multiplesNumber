package usecases

import (
	"TRAFILEA/app/test/mocks"
	"TRAFILEA/usecases"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var serviceMock mocks.ServicesMultipliers
var typeMock mocks.Type

func TestSaveMultiples(t *testing.T) {

	tests := []struct {
		name    string
		mock    func()
		wantErr bool
	}{
		{
			//Given
			name: "Success Method save multiplier",
			mock: func() {
				typeMock.On("Match", 3).Return(true)
				typeMock.On("Name").Return("1")
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			serviceMock = mocks.ServicesMultipliers{}
			typeMock = mocks.Type{}
			tt.mock()

			multipleOperations := usecases.NewType()

			usecase := usecases.NewServiceMultipliers(multipleOperations)

			//when
			result := usecase.SaveMultiplier(3)

			//then
			assert.True(t, result, "True is true!")

		})
	}

}

func TestGetMultiplier(t *testing.T) {

	tests := []struct {
		name    string
		mock    func()
		wantErr bool
	}{
		{
			//Given
			name: "Success Method Get multiplier",
			mock: func() {

			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			serviceMock = mocks.ServicesMultipliers{}
			typeMock = mocks.Type{}

			tt.mock()

			multipleOperations := usecases.NewType()

			usecase := usecases.NewServiceMultipliers(multipleOperations)

			//when
			result := usecase.GetMultipliers(10)

			//then
			assert.NotNil(t, result, "not nil")

		})
	}

}

func TestGetCollections(t *testing.T) {

	serviceMock = mocks.ServicesMultipliers{}

	//Given
	multiplesByType := []usecases.MultiplesByType{{Num: 3, Type: "3"}}

	serviceMock.On("GetMultiplierCollection").Return(multiplesByType, nil)

	//when
	result := serviceMock.GetMultiplierCollection()

	//then
	assert.Equal(t, multiplesByType, result)

}

func TestGetValueByNumber(t *testing.T) {

	serviceMock = mocks.ServicesMultipliers{}

	//Given
	multiplesByType := usecases.MultiplesByType{Num: 3, Type: "1"}

	serviceMock.On("GetValueByNumber", 3).Return(multiplesByType, nil)

	//when
	result, _ := serviceMock.GetValueByNumber(3)

	//then
	assert.Equal(t, multiplesByType, result, "not nil")

}

func TestGetValueByNumberError(t *testing.T) {

	serviceMock = mocks.ServicesMultipliers{}

	//Given
	multiplesByType := usecases.MultiplesByType{}

	serviceMock.On("GetValueByNumber", 3).Return(multiplesByType, errors.New("not found"))

	//when
	_, err := serviceMock.GetValueByNumber(3)

	//then
	assert.Error(t, err)

}
