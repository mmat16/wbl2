package patterns

import (
	"fmt"
)

/*
Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры
использования данного примера на практике.
*/

// Паттерн «посетитель» является поведенческим паттерном проектирования,
// предназначен для добавления новых операций в уже существующие структуры и
// интерфейсы. Позволяет добавлять новые операции без изменения изначальной
// структуры. Допустим, у нас есть интерфейс, содержащий методы для работы с
// различными типами данных. При этом, мы хотим добавить новый метод, который
// будет работать только с одним типом данных, или несколько новых методов,
// специализированных под определённые типы данных. Чтобы не изменять исходный
// интерфейс, добавляя в него новые методы (что может нарушить работу типов,
// которые уже реализуют этот интерфейс) и ухудшая абстракцию (чем меньше
// методов в интерфейсе - тем лучше абстракция), можно добавить лишь один
// новый метод accept, который будет принимать в качестве аргумента объект
// интерфейса "посетитель", который в свою очередь будет содержать методы для
// обработки различных типов данных. Таким образом, мы сможем добавлять новые
// методы для работы с различными типами данных, не изменяя исходный интерфейс.

// Плюсы: добавление новых операций в уже существующие структуры и интерфейсы,
// не изменяя их и не нарушая работу уже существующих методов.

// Минусы: усложнение кода, необходимость создания новых структур и интерфейсов.

// Пример: реализация интерфейса "посетитель" для работы с различными типами
// данных.

// shop интерфейс магазина с методами accept для работы с различными типами и
// Buy для покупки товара
type Shop interface {
	Accept(visitor)
	Buy()
}

// CofeeShop реализация интерфейса shop для кофейни
type CofeeShop struct {
	name  string
	price int
}

// NewCofeeShop создаёт новый экземпляр кофейни с указанным названием и ценой
func NewCofeeShop(name string, price int) *CofeeShop {
	return &CofeeShop{name: name, price: price}
}

// accept принимает в качестве аргумента объект интерфейса "посетитель" и
// вызывает его метод для работы с кофейней
func (c *CofeeShop) Accept(v visitor) {
	v.visitCofeeShop(c)
}

// Buy покупает кофе
func (c *CofeeShop) Buy() {
	fmt.Printf("Buying coffee at %s for %d\n", c.name, c.price)
}

// NewsPaperShop реализация интерфейса shop для газетного киоска
type NewsPaperShop struct {
	name  string
	price int
}

// NewNewsPaperShop создаёт новый экземпляр газетного киоска с указанным
// названием
func NewNewsPaperShop(name string, price int) *NewsPaperShop {
	return &NewsPaperShop{name: name, price: price}
}

// accept принимает в качестве аргумента объект интерфейса "посетитель" и
// вызывает его метод для работы с газетным киоском
func (n *NewsPaperShop) Accept(v visitor) {
	v.visitNewsPaperShop(n)
}

// Buy покупает газету
func (n *NewsPaperShop) Buy() {
	fmt.Printf("Buying newspaper at %s for %d\n", n.name, n.price)
}

// CigaretteShop реализация интерфейса shop для табачного киоска
type CigaretteShop struct {
	name  string
	price int
}

// NewCigaretteShop создаёт новый экземпляр табачного киоска с указанным
// названием и ценой
func NewCigaretteShop(name string, price int) *CigaretteShop {
	return &CigaretteShop{name: name, price: price}
}

// accept принимает в качестве аргумента объект интерфейса "посетитель" и
// вызывает его метод для работы с табачным киоском
func (c *CigaretteShop) Accept(v visitor) {
	v.visitCigaretteShop(c)
}

// Buy покупает сигареты
func (c *CigaretteShop) Buy() {
	fmt.Printf("Buying cigarettes at %s for %d\n", c.name, c.price)
}

// visitor интерфейс "посетитель" с методами для работы с различными типами
// данных
type visitor interface {
	visitCofeeShop(*CofeeShop)
	visitNewsPaperShop(*NewsPaperShop)
	visitCigaretteShop(*CigaretteShop)
}

// ShopVisitor реализация интерфейса "посетитель" для работы с различными
// типами данных
type ShopVisitor struct{}

// NewShopVisitor создаёт новый экземпляр "посетителя"
func NewShopVisitor() *ShopVisitor {
	return &ShopVisitor{}
}

// visitCofeeShop реализация метода visitCofeeShop интерфейса "посетитель"
func (s *ShopVisitor) visitCofeeShop(c *CofeeShop) {
	fmt.Printf("Visiting %s\n", c.name)
	c.Buy()
}

// visitNewsPaperShop реализация метода visitNewsPaperShop интерфейса "посетитель"
func (s *ShopVisitor) visitNewsPaperShop(n *NewsPaperShop) {
	fmt.Printf("Visiting %s\n", n.name)
	n.Buy()
}

// visitCigaretteShop реализация метода visitCigaretteShop интерфейса "посетитель"
func (s *ShopVisitor) visitCigaretteShop(c *CigaretteShop) {
	fmt.Printf("Visiting %s\n", c.name)
	c.Buy()
}

// ComplainVisitor реализация интерфейса "посетитель" для работы с различными
// типами данных
type ComplainVisitor struct{}

// NewComplainVisitor создаёт новый экземпляр "посетителя"
func NewComplainVisitor() *ComplainVisitor {
	return &ComplainVisitor{}
}

// visitCofeeShop реализация метода visitCofeeShop интерфейса "посетитель"
func (c *ComplainVisitor) visitCofeeShop(cs *CofeeShop) {
	fmt.Printf("Visiting %s\n", cs.name)
	fmt.Printf("Coffee is bad and expensive in %s\n", cs.name)
}

// visitNewsPaperShop реализация метода visitNewsPaperShop интерфейса
// "посетитель"
func (c *ComplainVisitor) visitNewsPaperShop(n *NewsPaperShop) {
	fmt.Printf("Visiting %s\n", n.name)
	fmt.Printf("All news are fake in %s\n", n.name)
}

// visitCigaretteShop реализация метода visitCigaretteShop интерфейса
// "посетитель"
func (c *ComplainVisitor) visitCigaretteShop(cs *CigaretteShop) {
	fmt.Printf("Visiting %s\n", cs.name)
	fmt.Println("Smoking kills")
}

// Пример использования
// func main() {
// 	shops := []Shop{
// 		NewCofeeShop("Starbucks", 100),
// 		NewNewsPaperShop("Kiosk", 50),
// 		NewCigaretteShop("Tobacco", 200),
// 	}
// 	visitor := NewShopVisitor()
// 	for _, s := range shops {
// 		s.Accept(visitor)
// 	}
// 	fmt.Println()
// 	complainVisitor := NewComplainVisitor()
// 	for _, s := range shops {
// 		s.Accept(complainVisitor)
// 	}
// }
