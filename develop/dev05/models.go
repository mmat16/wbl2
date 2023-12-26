package grep

import (
	"regexp"
	"strings"
)

// matcher интерфейс который содержит метод проверки совпадения паттерна со
// строкой
type matcher interface {
	match() bool
	addLine(line string)
	getLine() string
}

// reMatcher проверяет совпадение строки по регулярному выражению
type reMatcher struct {
	pattern *regexp.Regexp
	line    string
}

// match - метод проверки совпадения
func (rm *reMatcher) match() bool {
	if ignore {
		return rm.pattern.MatchString(strings.ToLower(rm.line))
	}
	return rm.pattern.MatchString(rm.line)
}

// addLine - метод добавления строки в matcher
func (rm *reMatcher) addLine(line string) {
	rm.line = line
}

// getLine - метод возвращает строку из matcher
func (rm *reMatcher) getLine() string {
	return rm.line
}

// stringMatcher проверяет точное совпадение паттерна со строкой
type stringMatcher struct {
	pattern string
	line    string
}

// match - метод проверки совпадения
func (sm *stringMatcher) match() bool {
	if ignore {
		return strings.Contains(strings.ToLower(sm.line), sm.pattern)
	}
	return strings.Contains(sm.pattern, sm.line)
}

// addLine - метод добавления строки в matcher
func (sm *stringMatcher) addLine(line string) {
	sm.line = line
}

// getLine - метод возвращает строку из matcher
func (sm *stringMatcher) getLine() string {
	return sm.line
}

// queue - структура данных очередь, для хранения строк, предшествующих
// строке, в которой найдено совпадение
type queue struct {
	lines []string
	limit int
}

// newQueue - конструктор очереди
func newQueue(limit int) *queue {
	return &queue{limit: limit + 1}
}

// push - метод добавления строки в очередь
func (q *queue) push(line string) {
	if len(q.lines) == q.limit {
		q.lines = q.lines[1:]
	}
	q.lines = append(q.lines, line)
}

// popAll - метод возвращает все строки из очереди
func (q *queue) popAll() []string {
	return q.lines
}
