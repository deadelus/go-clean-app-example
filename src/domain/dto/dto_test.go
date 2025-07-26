package dto_test

import (
	"go-clean-app-project/src/domain/dto"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResult(t *testing.T) {
	t.Run("should return a success result", func(t *testing.T) {
		// Given
		data := "test"

		// When
		result := dto.Success(data)

		// Then
		assert.True(t, result.Success)
		assert.Equal(t, &data, result.Data)
		assert.Empty(t, result.Error)
	})

	t.Run("should return a failure result", func(t *testing.T) {
		// Given
		err := "error"

		// When
		result := dto.Failure[any](err)

		// Then
		assert.False(t, result.Success)
		assert.Nil(t, result.Data)
		assert.Equal(t, err, result.Error)
	})
}
