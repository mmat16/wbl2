package patterns

import "fmt"

/*
Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры
использования данного примера на практике.
*/

// Комманда это поведенческий паттерн проектирования, инкапсулирующий запрос в
// виде объекта. Такой объект комманды содержит в себе всю информацию о вызове
// и таким образом может выполнять его независимо.

// Плюсы: 1) убирает прямую зависимость между объектами, вызывающими операции,
// и объектами, которые их непосредственно выполняют; 2) позволяет реализовать
// простую отмену и повтор операций; 3) позволяет реализовать отложенный запуск
// операций; 4) позволяет предоставить унифицированный интерфейс для выполнения
// различных операций, приводящих к одному и тому же результату.

// Минусы: 1) усложняет код программы из-за введения множества дополнительных
// классов; 2) клиент должен знать о конкретном классе команды, который он
// использует; 3) может привести к утечкам памяти, если в объекте команды
// хранятся ссылки на объекты получателей.

// Примеры использования: разделение бизнес-логики и пользовательского
// интерфейса. Создание общего хендлера для обработки событий из разных
// источников.

// Owner - владелец собаки, который может давать команду которой он научил
// собаку
type Owner struct {
	Command Command
}

// GiveCommand - владелец дает команду собаке
func (o *Owner) GiveCommand() {
	o.Command.Execute()
}

// тип комманды - интерфейс с методом Execute
type Command interface {
	Execute()
}

// SitCommand - комманда собаке сесть
type SitCommand struct {
	Dog Dog
}

// Execute - собака садится
func (c *SitCommand) Execute() {
	c.Dog.Sit()
}

// DownCommand - комманда собаке лечь
type DownCommand struct {
	Dog Dog
}

// Execute - собака ложится
func (c *DownCommand) Execute() {
	c.Dog.Down()
}

// Dog - собака, которая может выполнять комманды
type Dog interface {
	Sit()
	Down()
}

// Labrador - порода собаки
type Labrador struct {
	Name string
}

// Sit - собака садится
func (l *Labrador) Sit() {
	fmt.Printf("%s: \"сижу.... мячик! птичка! котик! я побежал!\"\n", l.Name)
}

// Down - собака ложится
func (l *Labrador) Down() {
	fmt.Printf("%s: \"лежу... сплю zzz...\"\n", l.Name)
}

type BorderCollie struct {
	Name string
}

// Sit - собака садится
func (b *BorderCollie) Sit() {
	fmt.Printf("%s: \"СИЖУ\"\n", b.Name)
}

// Down - собака ложится
func (b *BorderCollie) Down() {
	fmt.Printf("%s: \"ЛЕЖУ\"\n", b.Name)
}

// пример использования
// func main() {
// 	lab := patterns.Labrador{Name: "Локи"}
// 	collie := patterns.BorderCollie{Name: "Мики"}
//
// 	labSit := patterns.SitCommand{Dog: &lab}
// 	labDown := patterns.DownCommand{Dog: &lab}
//
// 	collieSit := patterns.SitCommand{Dog: &collie}
// 	collieDown := patterns.DownCommand{Dog: &collie}
//
// 	labOwner := patterns.Owner{Command: &labSit}
// 	collieOwner := patterns.Owner{Command: &collieSit}
//
// 	labOwner.GiveCommand()
// 	collieOwner.GiveCommand()
//
// 	labOwner.Command = &labDown
//	collieOwner.Command = &collieDown
//
//	labOwner.GiveCommand()
//	collieOwner.GiveCommand()
//}
// >> ❯ go run main.go
// Локи: "сижу.... мячик! птичка! котик! я побежал!"
// Мики: "СИЖУ"
// Локи: "лежу... сплю zzz..."
// Мики: "ЛЕЖУ"
