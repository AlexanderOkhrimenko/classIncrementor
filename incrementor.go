package classIncrementor

// Incrementor базовая структура для счетчика
type Incrementor struct {
	Value int
	Max   int
}

// getNumber получение текущего значения счетчика
func (num Incrementor) getNumber() int {
	return num.Value
}

// incrementNumber метод увеличения счетчика
func (num *Incrementor) incrementNumber() {
	num.Value++
	if num.Value > num.Max {
		num.Value = 0
	}
}

// setMaximumValue задание масксимального значения счетчика
func (num *Incrementor) setMaximumValue(maximumValue int) {
	num.Max = maximumValue
}

// CreateTestClass тестовая функция проверки класса где value - начальное значение счетчика,
// maxValue - максимальное значение счетчика, numberIterationsint - количестов итераций увеличения счетчика.
func CreateTestClass(value, maxValue, numberIterationsint int) int {
	var a = Incrementor{value, maxValue}
	for i := 0; i <= numberIterationsint; i++ {
		a.incrementNumber()
	}
	return a.getNumber()
}
