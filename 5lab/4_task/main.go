package task4

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"sync"
	"time"
)

var kernel = [3][3]float64{
	{1.0 / 16, 2.0 / 16, 1.0 / 16},
	{2.0 / 16, 4.0 / 16, 2.0 / 16},
	{1.0 / 16, 2.0 / 16, 1.0 / 16},
}

func RunTask4() {
	file, err := os.Open("task4/image.png")
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

	result := image.NewRGBA(bounds)
	var wg sync.WaitGroup
	start := time.Now()

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		wg.Add(1)
		go func(y int) {
			defer wg.Done()
			for x := bounds.Min.X; x < bounds.Max.X; x++ {
				newColor := applyKernel(rgba, x, y)
				result.Set(x, y, newColor)
			}
		}(y)
	}

	wg.Wait()
	duration := time.Since(start)
	fmt.Println("Время параллельной обработки с фильтром:", duration)

	outputFile, err := os.Create("task4/output_blur.png")
	if err != nil {
		fmt.Println("Ошибка при создании файла:", err)
		return
	}
	defer outputFile.Close()

	err = png.Encode(outputFile, result)
	if err != nil {
		fmt.Println("Ошибка при сохранении изображения:", err)
		return
	}

	fmt.Println("Обработанное изображение сохранено как task4/output_blur.png")
}

func applyKernel(img *image.RGBA, x, y int) color.RGBA {
	bounds := img.Bounds()
	var r, g, b float64

	for ky := -1; ky <= 1; ky++ {
		for kx := -1; kx <= 1; kx++ {
			px := x + kx
			py := y + ky

			if px >= bounds.Min.X && px < bounds.Max.X && py >= bounds.Min.Y && py < bounds.Max.Y {
				neighborColor := img.At(px, py).(color.RGBA)
				weight := kernel[ky+1][kx+1]

				r += float64(neighborColor.R) * weight
				g += float64(neighborColor.G) * weight
				b += float64(neighborColor.B) * weight
			}
		}
	}

	return color.RGBA{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
		A: img.At(x, y).(color.RGBA).A,
	}
}
