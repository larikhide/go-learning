package quenec

//cоздать очередь
var q []string

//AddQ добавить новый элемент в очередь
func AddQ(str string) {
  q = append(q, str)
}

//DiscardQ удалить из очереди
func DiscardQ() string {
  if len(q) > 0 {
    elem := q[0]
    q = q[1:]
    return elem
  }
  return "Очереди нет"
}
