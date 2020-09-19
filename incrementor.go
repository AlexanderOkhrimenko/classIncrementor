package classIncrementor

// Incrementor базовая структура для счетчика
type Incrementor struct {
	Value int
	Max   int
}

// GetNumber получение текущего значения счетчика
func (num Incrementor) GetNumber() int {
	return num.Value
}

// IncrementNumber метод увеличения счетчика
func (num *Incrementor) IncrementNumber() {
	num.Value++
	if num.Value > num.Max {
		num.Value = 0
	}
}

// SetMaximumValue задание масксимального значения счетчика
func (num *Incrementor) SetMaximumValue(maximumValue int) {
	num.Max = maximumValue
}

// CreateTestClass тестовая функция проверки класса где value - начальное значение счетчика,
// maxValue - максимальное значение счетчика, numberIterationsint - количестов итераций увеличения счетчика.
func CreateTestClass(value, maxValue, numberIterationsint int) int {
	var a = Incrementor{value, maxValue}
	for i := 0; i <= numberIterationsint; i++ {
		a.IncrementNumber()
	}
	return a.GetNumber()
}
