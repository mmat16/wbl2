package patterns

import "fmt"

// реализовать паттерн "Фабричный метод"
// объяснить применимость паттерна, его плюсы и минусы, а так же реальные
// примеры использования метода на практике.

// фабричный метод - это порождающий паттерн проектирования, определяющий общий
// интерфейс для создания объектов в родительском классе, позволяюший
// изменять объекты, которые будут создаваться в дочерних классах.

// Применяется, когда необходима возможность расширять функционал класса,
// библиотеки, не изменяя при этом код самого класса, библиотеки не
// нагромождая его множеством условий и проверок.

// Преимущества: разделение кода, создающего объекты, от кода, использующего
// их. Упрощение добавления новых типов объектов в программу. Собдюдение
// принципа открытости/закрытости (сущности программы должны быть открыты для
// расширения, но закрыты для изменения).

// Недостатки: усложнение параллельной иерархии классов. Необходимость создания
// подклассов для каждого нового типа продукта и новой фабрики для этого типа.

// например, в пакете database/sql есть интерфейс Driver, который определяет
// метод Open, который возвращает объект типа Conn, который представляет
// собой соединение с базой данных. Каждый драйвер реализует свой метод Open,
// который возвращает свой тип Conn. Таким образом, мы можем использовать
// один и тот же код для работы с разными базами данных, не заботясь о том,
// какой именно драйвер используется.

// Filter - интерфейс фильтра для обработки изображений с методом Apply
type Filter interface {
	Apply(image string) string
}

// BlackAndWhiteFilter - реализация фильтра для преобразования изображения в
// черно-белое
type BlackAndWhiteFilter struct{}

// Apply - применение фильтра
func (f *BlackAndWhiteFilter) Apply(image string) string {
	return fmt.Sprintf("Black and white filter applied to %s", image)
}

// SepiaFilter - реализация фильтра для преобразования изображения в сепию
type SepiaFilter struct{}

// Apply - применение фильтра
func (f *SepiaFilter) Apply(image string) string {
	return fmt.Sprintf("Sepia filter applied to %s", image)
}

// FilterFactory - интерфейс фабрики для создания фильтров
type FilterFactory interface {
	Create() Filter
}

// BlackAndWhiteFilterFactory - реализация фабрики для создания черно-белого
// фильтра
type BlackAndWhiteFilterFactory struct{}

// Create - создание черно-белого фильтра
func (f *BlackAndWhiteFilterFactory) Create() Filter {
	return &BlackAndWhiteFilter{}
}

// SepiaFilterFactory - реализация фабрики для создания сепия фильтра
type SepiaFilterFactory struct{}

// Create - создание сепия фильтра
func (f *SepiaFilterFactory) Create() Filter {
	return &SepiaFilter{}
}

// использование
func ApplyFilter(factory FilterFactory, image string) {
	filter := factory.Create()
	result := filter.Apply(image)
	fmt.Println(result)
}

// func main() {
// 	fmt.Println(ApplyFilter(&BlackAndWhiteFilterFactory{}, "image.jpg"))
// 	fmt.Println(ApplyFilter(&SepiaFilterFactory{}, "image.jpg"))
// }
