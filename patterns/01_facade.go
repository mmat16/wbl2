package patterns

import (
	"log"
	"math/rand"
)

// реализовать паттер фасад. объяснить применимость паттерна, его плюсы и минусы
// а так же реальные примеры использования паттерна на практике

// Паттерн фасад является структурным паттерном проектирования, предназначен
// для упрощения внешнего интерфейса сложной системы. Фасад позволяет скрыть
// сложность системы и предоставляет набор простых методов для её использования.
// Грубо говоря - добавляет ещё один уровень абстракции над сложной системой.

// Применимость: конкретный пример - использование онлайн кошельков: достаточно
// привязать карту к кошельку и пользоваться им, не забивая голову тем,
// какие денежные и информационные потоки происходят внутри системы. Более общий
// пример - использование библиотеки для работы с базой данных (ORM), или API
// крупных сервисов.

// Плюсы: защита и инкапсуляция логики сложной системы, упрощение внешнего
// интерфейса, доступного стороннему пользователю.

// Минусы: возможное сокрытие важной информации о работе системы за неважной
// информацией. увелечиние сложности системы и внешнего API, требующих
// поддержки и документации.

// реализация "магического шара" с ответами на вопросы. Внутри системы
// реализована сложная логика получения случайного ответа из списка, валидация
// вопроса и обработка возможной паники. Внешний интерфейс представляет собой
// простой метод Shake, который принимает вопрос и возвращает ответ.
type Magic8Ball struct {
	answers answers
}

// NewMagic8Ball создаёт новый экземпляр "магического шара" с ответами
func NewMagic8Ball() *Magic8Ball {
	return &Magic8Ball{answers: *newAnswers()}
}

// Shake возвращает случайный ответ на вопрос и является фасадом для "сложной"
// системы "магического шара" - получения случайного ответа из списка, валидация
// вопроса и обработка возможной паники
func (m *Magic8Ball) Shake(question string) string {
	defer handlePanic()
	validateQuestion(question)
	return m.answers.getAnswer()
}

// answers список ответов шара, реализованный в виде map
type answers map[int]string

// newAnswers создаёт список ответов
func newAnswers() *answers {
	return &answers{
		1:  "It is certain",
		2:  "It is decidedly so",
		3:  "Without a doubt",
		4:  "Yes — definitely",
		5:  "You may rely on it",
		6:  "As I see it, yes",
		7:  "Most likely",
		8:  "Outlook good",
		9:  "Signs point to yes",
		10: "Yes",
		11: "Reply hazy, try again",
		12: "Ask again later",
		13: "Better not tell you now",
		14: "Cannot predict now",
		15: "Concentrate and ask again",
		16: "Don’t count on it",
		17: "My reply is no",
		18: "My sources say no",
		19: "Outlook not so good",
		20: "Very doubtful",
	}
}

// getAnswer возвращает случайный ответ из списка
func (a *answers) getAnswer() string {
	return (*a)[getRandomInt(1, 20)]
}

// getRandomInt возвращает случайное число в диапазоне от min до max
func getRandomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

// valideaQuestion проверяет входящий вопрос на корректность - не пустой ли он,
// и оканчивается на знак вопроса
func validateQuestion(question string) {
	if question == "" {
		panic("question is empty")
	}
	if question[len(question)-1:] != "?" {
		panic("question has no question mark")
	}
}

// errHandler обрабатывает возможнную панику, вызванную некорректным вопросом c
// помощью recover
func handlePanic() {
	if err := recover(); err != nil {
		log.Fatal(err)
	}
}

// Пример использования
// func main() {
// 	magic8Ball := NewMagic8Ball()
// 	fmt.Println(magic8Ball.Shake("Will I win the lottery?"))
// }
