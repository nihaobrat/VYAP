package task2

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"time"
)

func RunTask2() {
	file, err := os.Open("task2/image.png")
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("Ошибка при декодировании изображения:", err)
		return
	}

	bounds := img.Bounds()
	rgba := image.NewRGBA(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			rgba.Set(x, y, img.At(x, y))
		}
	}

	start := time.Now()
	filter(rgba)
	duration := time.Since(start)

	fmt.Println("Время обработки:", duration)

	outputFile, err := os.Create("task2/output.png")
	if err != nil {
		fmt.Println("Ошибка при создании файла:", err)
		return
	}
	defer outputFile.Close()

	err = png.Encode(outputFile, rgba)
	if err != nil {
		fmt.Println("Ошибка при сохранении изображения:", err)
		return
	}

	fmt.Println("Обработанное изображение сохранено как task2/output.png")
}

func filter(img *image.RGBA) {
	bounds := img.Bounds()

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			originalColor := img.At(x, y).(color.RGBA)
			gray := uint8((int(originalColor.R) + int(originalColor.G) + int(originalColor.B)) / 3)
			newColor := color.RGBA{R: gray, G: gray, B: gray, A: originalColor.A}
			img.Set(x, y, newColor)
		}
	}
}
