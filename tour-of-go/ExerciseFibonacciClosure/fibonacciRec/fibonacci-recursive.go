package fib

var (
	cache map[int]int
)

//не понял зачем создается функция в которой создается мапа
//мапу мы создали глобальную внутри пакета fib, она заполнилась значеними по умолчанию и так
//и почему эту функцию нигде не используют?
//почему она ничего не возвращает?
func init() {
	cache = make(map[int]int)
}

//Fib with cached temp
func Fib(n int) int {
	if n < 2 {
		return n
	}

	if r, ok := cache[n]; ok {
		return r
	}
	sum := Fib(n-1) + Fib(n-2)
	cache[n] = sum //замыкание происходит здесь? вызываемая функция Fib меняет значение в мапе за своими пределами?
	return sum
}
