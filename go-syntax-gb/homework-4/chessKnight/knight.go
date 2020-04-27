package main

import "fmt"

//Point это структура, которая далее используется как тип, который описывает точку на доске
type Point struct {
	x int
	y int
}

//мапа для соответствия цифрого обозначения поля буквенному
var coordMap = map[int]string{
	1: "a",
	2: "b",
	3: "c",
	4: "d",
	5: "e",
	6: "f",
	7: "g",
	8: "h",
}

//GetX метод для получения значения х из структуры типа Поинт
func (p *Point) GetX() int {
	return p.x
}

//GetY метод для получения значения y из структуры типа Поинт
func (p *Point) GetY() int {
	return p.y
}

//String метод для перевода цифрогово обозначения поля в буквенное
func (p Point) String() string {
	return fmt.Sprintf("%s%d", coordMap[p.x], p.y)
}

//NewPoint функция, которая создает структуру с типом Point с координатами х и у
func NewPoint(x, y int) Point {
	return Point{
		x: x,
		y: y,
	}
}

//массив из стуктур типа Point который хранит варианты всех ходов, созданных функцией NewPoint
var avMoves = []Point{
	NewPoint(1, 2),
	NewPoint(1, -2),
	NewPoint(2, 1),
	NewPoint(2, -1),
	NewPoint(-1, -2),
	NewPoint(-1, 2),
	NewPoint(-2, -1),
	NewPoint(-2, 1),
}

//Knight структура, которая имеет поле position c типом Point с координатами х и у
type Knight struct {
	position Point
}

//GetPosition метод, который возвращает координаты коня. При этом возвращает не самого коня, а именно его позицию
func (k *Knight) GetPosition() *Point {
	return &k.position
}

//CalcAvMoves возвращает массив доступных ходов
func (k *Knight) CalcAvMoves() []Point {
	res := make([]Point, 0, 4)
	for _, v := range avMoves {
		newX := k.GetPosition().GetX() + v.GetX()
		newY := k.GetPosition().GetY() + v.GetY()
		if checkCoord(newX, newY) {
			res = append(res, NewPoint(newX, newY))
		}

	}
	return res
}

//checkCoord проверяет существует ли такая клетка на доске шахматной
func checkCoord(x, y int) bool {
	return x >= 1 && x <= 8 && y >= 1 && y <= 8
}

//NewKnight функция которая принимает на вход 2 точки, записывает координаты для коня и выплевывет коня
func NewKnight(x, y int) Knight {
	return Knight{
		position: Point{
			x: x,
			y: y,
		},
	}
}

func main() {
	kn := NewKnight(4, 2)
	fmt.Println(kn.CalcAvMoves())
}
