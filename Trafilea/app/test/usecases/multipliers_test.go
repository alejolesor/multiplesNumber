package usecases

import (
	"TRAFILEA/app/test/mocks"
	"TRAFILEA/usecases"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {

	assert.True(t, true, "True is true!")

}

var serviceMock mocks.ServicesMultipliers
var typeMock mocks.Type

func TestMultiples(t *testing.T) {

	tests := []struct {
		name    string
		mock    func()
		wantErr bool
	}{
		{
			//Given
			name: "Success Method Deposit",
			mock: func() {
				serviceMock.On("SaveMultiplier", 3).Return(nil)
				//mockRepositoryAccount.On("Get", "francisco").Return(&accountMap, nil)
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
