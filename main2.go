package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println("-------------1-------------")
	// 1.1
	var ipAddress [4]byte = [4]byte{127, 0, 0, 1}
	formattedIP := formatIP(ipAddress)
	fmt.Println("Форматированный IP-адрес:", formattedIP)

	// 1.2
	start, end := 1, 10
	evens, err := listEven(start, end)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Четные числа в диапазоне:", evens)
	}

	start, end = 10, 5
	evens, err = listEven(start, end)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Четные числа в диапазоне:", evens)
	}

	fmt.Println("-------------2-------------")
	//2
	inputString := "hello, world!"
	charCounts := countChars(inputString)

	// Выводим результат
	fmt.Println("Подсчет символов в строке:")
	for char, count := range charCounts {
		fmt.Printf("Символ '%c' встречается %d раз(а)\n", char, count)
	}

	fmt.Println("-------------3-------------")
	// 3.1 Пример использования структуры Point
	p1 := Point{X: 3.5, Y: 7.2} // Создаем точку с координатами (3.5, 7.2)
	fmt.Printf("Точка p1 имеет координаты: X = %.2f, Y = %.2f\n", p1.X, p1.Y)

	//3.2 и 3.3
	// Создаем отрезок с начальной и конечной точками
	segment := Segment{
		Start: Point{X: 1.0, Y: 2.0}, // Точка начала отрезка
		End:   Point{X: 4.0, Y: 6.0}, // Точка конца отрезка
	}

	// Выводим координаты отрезка
	fmt.Printf("Отрезок начинается в точке (%.2f, %.2f) и заканчивается в точке (%.2f, %.2f)\n",
		segment.Start.X, segment.Start.Y, segment.End.X, segment.End.Y)

	// Вычисляем и выводим длину отрезка
	length := segment.Length()
	fmt.Printf("Длина отрезка: %.2f\n", length)

	//3.4
	// Создаем треугольник с тремя вершинами
	triangle := Triangle{
		A: Point{X: 0.0, Y: 0.0}, // Вершина A
		B: Point{X: 5.0, Y: 0.0}, // Вершина B
		C: Point{X: 0.0, Y: 4.0}, // Вершина C
	}

	// Выводим координаты вершин треугольника
	fmt.Println("Координаты треугольника:")
	fmt.Printf("Вершина A: (%.2f, %.2f)\n", triangle.A.X, triangle.A.Y)
	fmt.Printf("Вершина B: (%.2f, %.2f)\n", triangle.B.X, triangle.B.Y)
	fmt.Printf("Вершина C: (%.2f, %.2f)\n", triangle.C.X, triangle.C.Y)

	//3.5
	// Создаем круг с центром в точке (0, 0) и радиусом 5
	circle := Circle{
		Center: Point{X: 0.0, Y: 0.0}, // Центр круга
		Radius: 5.0,                   // Радиус круга
	}

	// Выводим данные круга
	fmt.Println("Параметры круга:")
	fmt.Printf("Центр: (%.2f, %.2f)\n", circle.Center.X, circle.Center.Y)
	fmt.Printf("Радиус: %.2f\n", circle.Radius)

	//3.6
	// Создаем треугольник
	triangle1 := Triangle{
		A: Point{X: 0.0, Y: 0.0},
		B: Point{X: 5.0, Y: 0.0},
		C: Point{X: 0.0, Y: 4.0},
	}

	// Вычисляем и выводим площадь треугольника
	triangleArea := triangle1.Area()
	fmt.Printf("Площадь треугольника: %.2f\n", triangleArea)

	// Создаем круг
	circle1 := Circle{
		Center: Point{X: 0.0, Y: 0.0},
		Radius: 5.0,
	}

	// Вычисляем и выводим площадь круга
	circleArea := circle1.Area()
	fmt.Printf("Площадь круга: %.2f\n", circleArea)

	//3.7

	// Создаем срез интерфейса Shape и добавляем в него фигуры
	shapes := []Shape{triangle, circle}

	// Проходим по всем фигурам и выводим их площади
	for i, shape := range shapes {
		fmt.Printf("Площадь фигуры %d: %.2f\n", i+1, shape.Area())
	}

	//3.8

	// Вызываем функцию printArea для треугольника и круга
	fmt.Println("Вызов для треугольника:")
	printArea(triangle)

	fmt.Println("Вызов для круга:")
	printArea(circle)

	fmt.Println("-------------4-------------")
	//4
	// 2. Функция для возведения числа в квадрат
	square := func(x float64) float64 {
		return x * x
	}
	cubed := func(x float64) float64 {
		return x * x * x
	}
	// 3. Создаем срез и заполняем его значениями
	slice := []float64{1.0, 2.0, 3.0, 4.0, 5.0}

	// Применяем функцию Map с функцией square
	squaredSlice := Map(slice, square)
	cubedSlice := Map(slice, cubed)

	// Выводим результаты
	fmt.Println("Исходный срез:", slice)
	fmt.Println("Срез после возведения в квадрат:", squaredSlice)
	fmt.Println("Срез после возведения в куб:", cubedSlice)

}

///////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////

// 1.1
func formatIP(ip [4]byte) string {
	return strconv.Itoa(int(ip[0])) + "." +
		strconv.Itoa(int(ip[1])) + "." +
		strconv.Itoa(int(ip[2])) + "." +
		strconv.Itoa(int(ip[3]))
}

// 1.2
func listEven(start, end int) ([]int, error) {
	// Проверка на некорректный диапазон
	if start > end {
		return nil, errors.New("левая граница диапазона больше правой")
	}

	// Создаем срез для хранения четных чисел
	var evens []int

	// Находим все четные числа в диапазоне
	for i := start; i <= end; i++ {
		if i%2 == 0 {
			evens = append(evens, i)
		}
	}

	return evens, nil
}

// 2
func countChars(s string) map[rune]int {
	// Создаем карту для подсчета символов
	charCount := make(map[rune]int)

	// Проходим по строке с помощью range
	for _, char := range s {
		charCount[char]++ // Увеличиваем счетчик для символа
	}

	return charCount
}

// 3.1
// Point описывает точку на плоскости с координатами X и Y
type Point struct {
	X float64 // Координата по оси X
	Y float64 // Координата по оси Y
}

// 3.2 и 3.3

// Segment описывает отрезок с начальной и конечной точками
type Segment struct {
	Start Point // Начальная точка отрезка
	End   Point // Конечная точка отрезка
}

// Length вычисляет длину отрезка с использованием теоремы Пифагора
func (s Segment) Length() float64 {
	dx := s.End.X - s.Start.X // Разница по X
	dy := s.End.Y - s.Start.Y // Разница по Y
	return math.Sqrt(dx*dx + dy*dy)
}

// 3.4
// Triangle описывает треугольник с тремя вершинами
type Triangle struct {
	A Point // Первая вершина треугольника
	B Point // Вторая вершина треугольника
	C Point // Третья вершина треугольника
}

// 3.5
// Circle описывает круг с центром и радиусом
type Circle struct {
	Center Point   // Точка центра круга
	Radius float64 // Радиус круга
}

// 3.6
// Метод для вычисления площади треугольника с использованием формулы Герона
func (t Triangle) Area() float64 {
	// Функция для вычисления длины стороны треугольника
	distance := func(p1, p2 Point) float64 {
		return math.Sqrt(math.Pow(p2.X-p1.X, 2) + math.Pow(p2.Y-p1.Y, 2))
	}

	// Длины сторон треугольника
	a := distance(t.A, t.B)
	b := distance(t.B, t.C)
	c := distance(t.C, t.A)

	// Полупериметр треугольника
	s := (a + b + c) / 2

	// Формула Герона
	return math.Sqrt(s * (s - a) * (s - b) * (s - c))
}

// Метод для вычисления площади круга
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// 3.7
type Shape interface {
	Area() float64
}

//3.8

// Функция printArea принимает интерфейс Shape и выводит площадь фигуры
func printArea(s Shape) {
	result := s.Area()
	fmt.Printf("Площадь фигуры: %.2f\n", result)
}

// 4.1
// Map применяет функцию f к каждому элементу среза.
// Возвращает новый срез, не изменяя исходный.
func Map(slice []float64, f func(float64) float64) []float64 {
	newSlice := make([]float64, len(slice)) // Создаем новый срез
	copy(newSlice, slice)                   // Копируем исходный срез

	for i, v := range newSlice {
		newSlice[i] = f(v) // Применяем функцию к каждому элементу
	}

	return newSlice
}
