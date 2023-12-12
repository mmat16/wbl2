package patterns

/*
Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры
использования данного примера на практике.
*/

// цепочка вызовов - это поведенческий паттерн проектирования, который позволяет
// передавать запросы последовательно по цепочке обработчиков. Каждый обработчик
// решает, может ли он обработать запрос сам и стоит ли передавать запрос дальше
// по цепи, либо же дальнейшая обработка не имеет смысла и цепочку можно
// прервать.

// Применимость: можно применять, когда есть более одного объекта, который может
// обработать запрос, и точно неизвестно, какой именно объект должен это сделать.
// сокрытие конкретных классов обработчиков от клиента, которые формирует
// цепочку обработчиков.

// Плюсы: уменьшает зависимость между клиентом и обработчиками, реализует
// принцип единственной обязанности, реализует принцип открытости/закрытости.

// Минусы: запрос может остаться никем не обработанным. возможно увеличение
// времени обработки запроса из-за того, что запрос может быть обработан
// несколькими обработчиками. усложнение логики программы.

// Примеры: обработка запросов веб-сервером, обработка исключений в языках
// программирования, обработка событий в системах GUI.

// Request - запрос, который передается по цепочке обработчиков.
type Request struct {
	RID  int
	Data string
}

// Handler - обработчик запроса.
type Handler interface {
	HandleRequest(*Request)
	SetNext(Handler)
}

// ConcreteHandler - конкретный обработчик запроса.
type CacheHandler struct {
	Cache map[int]string
	Next  Handler
}

// HandleRequest - обработка запроса.
func (ch *CacheHandler) HandleRequest(r *Request) {
	if data, ok := ch.Cache[r.RID]; ok {
		r.Data = data
		return
	}
	if ch.Next != nil {
		ch.Next.HandleRequest(r)
	} else {
		r.Data = "not found"
	}
}

// SetNext - установка следующего обработчика.
func (ch *CacheHandler) SetNext(h Handler) {
	ch.Next = h
}

// ORMHandler - конкретный обработчик запроса.
type ORMHandler struct {
	Orm  map[int]string
	Next Handler
}

// HandleRequest - обработка запроса.
func (oh *ORMHandler) HandleRequest(r *Request) {
	if data, ok := oh.Orm[r.RID]; ok {
		r.Data = data
		return
	}
	if oh.Next != nil {
		oh.Next.HandleRequest(r)
	} else {
		r.Data = "not found"
	}
}

// SetNext - установка следующего обработчика.
func (oh *ORMHandler) SetNext(h Handler) {
	oh.Next = h
}

// Пример использования:
// func main() {
// 	cacheHandler := &patterns.CacheHandler{Cache: make(map[int]string)}
// 	ormHandler := &patterns.ORMHandler{Orm: make(map[int]string)}
// 	cacheHandler.SetNext(ormHandler)
//
// 	req := &patterns.Request{RID: 1}
// 	cacheHandler.HandleRequest(req)
// 	println(req.Data)
// }
// >> not found

// func main() {
// 	cacheHandler := &patterns.CacheHandler{Cache: make(map[int]string)}
// 	ormHandler := &patterns.ORMHandler{Orm: make(map[int]string)}
// 	cacheHandler.SetNext(ormHandler)
//
// 	ormHandler.Orm[1] = "test"
//
// 	req := &patterns.Request{RID: 1}
// 	cacheHandler.HandleRequest(req)
// 	println(req.Data)
// }
// >> test

// func main() {
// 	cacheHandler := &patterns.CacheHandler{Cache: make(map[int]string)}
// 	ormHandler := &patterns.ORMHandler{Orm: make(map[int]string)}
// 	cacheHandler.SetNext(ormHandler)
//
// 	ormHandler.Orm[1] = "test"
// 	cacheHandler.Cache[1] = "new test"
//
// 	req := &patterns.Request{RID: 1}
// 	cacheHandler.HandleRequest(req)
// 	println(req.Data)
// }
// >> new tes
