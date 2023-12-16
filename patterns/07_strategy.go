package patterns

import "sort"

// реализовать паттерн Стратегия
// объяснить применимость паттерна, его плюсы и минусы, а так же
// примеры использования данного паттерна на практике.

// Стратегия - поведенческий паттерн проектирования, выделяющий набор
// алгоритмов, решающих конкретную задачу. Позволяет выбирать алгоритм
// в процессе выполнения программы.

// Плюсы: позволяет изменять поведение объекта во время выполнения программы, не
// изменяя его структуры. Уменьшает зависимость между клиентом и объектами,
// предоставляя клиентскому коду в качестве зависимости абстракцию. Позволяет
// добавлять новые алгоритмы без изменения существующих и выбирать наиболее
// подходящий алгоритм из набора во время выполнения.

// Минусы: может усложнить код программы за счёт дополнительных классов и
// интерфейсов стратегий. В случае если стратегии схожи, код может стать
// избыточным и повторяющимся. Клиент должен знать о существовании стратегий и
// уметь выбирать их.

// примеры использования: платёжные системы (выбор различных способов оплаты),
// сортировка (выбор различных алгоритмов сортировки), компиляторы (выбор
// различных алгоритмов оптимизации кода).

// пример реализации патерна Стратегия в программе сортировки массивов.
type SortStrategy interface {
	Sort([]int) []int
}

// Конкретная стратегия - сортировка в возрастающем порядке.
type AscendingSortStrategy struct{}

func (a *AscendingSortStrategy) Sort(arr []int) []int {
	sort.Ints(arr)
	return arr
}

// Конкретная стратегия - сортировка в убывающем порядке.
type DescendingSortStrategy struct{}

func (d *DescendingSortStrategy) Sort(arr []int) []int {
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	return arr
}

// Контекст - объект, который использует стратегии.
type SortContext struct {
	Strategy SortStrategy
}

// Метод для установки стратегии.
func (s *SortContext) SetStrategy(strategy SortStrategy) {
	s.Strategy = strategy
}

// Метод для выполнения стратегии.
func (s *SortContext) Execute(arr []int) []int {
	return s.Strategy.Sort(arr)
}

// Пример использования.
// func main() {
// 	data := []int{5, 124, 2598, 3, 1, 0, 9, 7, 6, 4, 2, 8}
// 	ascStategy := &AscendingSortStrategy{}
// 	context := &SortContext{Strategy: ascStategy}
// 	result := context.Execute(data)
// 	fmt.Println(result)
//
// 	descStrategy := &DescendingSortStrategy{}
// 	context.SetStrategy(descStrategy)
// 	result = context.Execute(data)
// 	fmt.Println(result)
// }
