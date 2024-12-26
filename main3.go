package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Задание 1: Работа с query-параметрами
	r.GET("/greet", greetHandler)

	// Задание 2: Маршруты для арифметических операций
	r.GET("/add", func(c *gin.Context) { arithmeticHandler(c, "+") })
	r.GET("/sub", func(c *gin.Context) { arithmeticHandler(c, "-") })
	r.GET("/mul", func(c *gin.Context) { arithmeticHandler(c, "*") })
	r.GET("/div", func(c *gin.Context) { arithmeticHandler(c, "/") })

	// Задание 3: Работа с JSON
	r.POST("/charcount", charCountHandler)

	// Запуск сервера
	r.Run(":8080")
}

// Обработчик для приветствия
func greetHandler(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")

	if name == "" || age == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name and age are required"})
		return
	}

	c.String(http.StatusOK, "Меня зовут %s, мне %s лет", name, age)
}

// Обработчик арифметических операций
func arithmeticHandler(c *gin.Context, operation string) {
	a, err1 := strconv.ParseFloat(c.Query("a"), 64)
	b, err2 := strconv.ParseFloat(c.Query("b"), 64)

	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Both parameters must be numbers"})
		return
	}

	var result float64
	switch operation {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot divide by zero"})
			return
		}
		result = a / b
	}

	c.JSON(http.StatusOK, gin.H{"result": result})
}

// Обработчик подсчета символов в строке
func charCountHandler(c *gin.Context) {
	var input struct {
		Text string `json:"text"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// Создаем карту для хранения количества символов
	charCount := make(map[string]int)
	for _, char := range input.Text {
		charCount[string(char)]++
	}

	c.JSON(http.StatusOK, charCount)
}

// Пример запроса для charCountHandler:
// POST http://localhost:8080/charcount

// Примеры запросов для arithmeticHandler:
// GET http://localhost:8080/add?a=10&b=5       // Сложение
// GET http://localhost:8080/sub?a=10&b=5       // Вычитание
// GET http://localhost:8080/mul?a=10&b=5       // Умножение
// GET http://localhost:8080/div?a=10&b=5       // Деление

// Пример запроса для greetHandler:
// GET http://localhost:8080/greet?name=Alice&age=19
