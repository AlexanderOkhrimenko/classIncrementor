package classIncrementor

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestCreateTestClass(t *testing.T) {
	// 0 значение счетчика
	// 10 максимальное значение счетчика
	// 10 количество итераций увеличения счетчика
	x := CreateTestClass(0, 10, 13)
	// Ожидаемое значение при "переполнении" счетчика 3
	assert.Equal(t, 3, x)
}
