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

// incrementNumber метод учеличения счетчика
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

//func main() {
//
//	var a = Incrementor{ 1, 5}
//	//fmt.Println(&a.Value, a)
//
//	fmt.Println(a.getNumber())
//	a.incrementNumber()
//	a.incrementNumber()
//	a.incrementNumber()
//	a.incrementNumber()
//	a.incrementNumber()
//	a.incrementNumber()
//	a.incrementNumber()
//	fmt.Println(a.getNumber())
//
//}
