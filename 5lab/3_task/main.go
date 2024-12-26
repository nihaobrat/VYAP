package task3

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"sync"
	"time"
)

func RunTask3() {
	file, err := os.Open("task3/image.png")
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

	var wg sync.WaitGroup
	start := time.Now()

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		wg.Add(1)
		go func(y int) {
			defer wg.Done()
			for x := bounds.Min.X; x < bounds.Max.X; x++ {
				originalColor := rgba.At(x, y).(color.RGBA)
				gray := uint8((int(originalColor.R) + int(originalColor.G) + int(originalColor.B)) / 3)
				newColor := color.RGBA{R: gray, G: gray, B: gray, A: originalColor.A}
				rgba.Set(x, y, newColor)
			}
		}(y)
	}

	wg.Wait()
	duration := time.Since(start)
	fmt.Println("Время параллельной обработки:", duration)

	outputFile, err := os.Create("task3/output_parallel.png")
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

	fmt.Println("Обработанное изображение сохранено как task3/output_parallel.png")
}
