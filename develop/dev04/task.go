package anagrams

import (
	"slices"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// FindAnagrams ищет все множества анаграмм в словаре words
func FindAnagrams(words *[]string) *map[string]*[]string {
	anagrams := make(map[string][]string)
	for _, word := range *words {
		word = strings.ToLower(word)
		key := order(word)
		anagrams[key] = append(anagrams[key], word)
	}

	result := make(map[string]*[]string)
	for _, value := range anagrams {
		value := value
		value = removeDuplicates(value)
		if len(value) > 1 {
			slices.Sort(value)
			result[value[0]] = &value
		}
	}

	return &result
}

// order сортирует буквы в слове, создавая ключ для мапы анаграмм
func order(word string) string {
	runes := []rune(word)
	slices.Sort(runes)
	// sort.Sort(runes)
	return string(runes)
}

// removeDuplicates удаляет дубликаты из слайса строк
func removeDuplicates(words []string) []string {
	unique := make(map[string]bool)
	for _, word := range words {
		unique[word] = true
	}

	result := make([]string, 0, len(unique))
	for word := range unique {
		result = append(result, word)
	}

	return result
}
