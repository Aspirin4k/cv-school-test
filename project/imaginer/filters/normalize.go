package filters

import (
	"image"
	"errors"
	"image/color"

	"github.com/Aspirin4k/cv_school_test/project/imaginer"
	"fmt"
)

var (
	GREYSCALED = 0
	COLORED = 1

	MAX = float64(65535)
	MIN = float64(0)
)

type Normalize struct {
	imaginer.Filter
	// Тип нормализации
	// Цветное или серое изображение
	NormalizeType int
}

func (f *Normalize)  ApplyFilter(src image.Image) (image.Image, error) {
	switch {
	case f.NormalizeType == GREYSCALED:
		dstBounds := src.Bounds()
		dstImage := image.NewRGBA(dstBounds)

		// Максимальная и минимальная яркость
		min := MAX
		max := MIN

		// Для нормализации необходимо знать max и min, а значит пройти по всем пикселям,
		// что занимает n операций (очевидно, меньше невозможно)
		// в дополнение к этому необходимо второй раз пройти по сетке, изменив
		// каждый отдельный пиксель, что займет еще n времени
		// по хорошему - необходимо распараллелить...
		for i := src.Bounds().Min.X ; i < src.Bounds().Max.X; i++ {
			for j := src.Bounds().Min.Y ; j < src.Bounds().Max.Y; j++ {
				// К сожалению, файл енкодится в формат, в котором каждый пиксель
				// закодирован в rgba, из-за чего по новой надо вычислять оттенок серого
				r, g, b, _ := src.At(i, j).RGBA()
				Y := R * float64(r) + G * float64(g) + B * float64(b)
				if Y > max {
					max = Y
				}
				if Y < min {
					min = Y
				}
			}
		}

		for i := src.Bounds().Min.X ; i < src.Bounds().Max.X; i++ {
			for j := src.Bounds().Min.Y ; j < src.Bounds().Max.Y; j++ {
				r, g, b, _ := src.At(i, j).RGBA()
				Y := R * float64(r) + G * float64(g) + B * float64(b)
				// Если использовать эту формулу (из википедии),
				// то при Y = max (самое белое место на фото) произойдет переполнение переменной
				// Из-за чего белый конвертнеться в черный...
				newY := (Y - min) * (MAX - MIN) / (max - min) + min
				pixel := color.Gray{uint8(newY / 256)}
				dstImage.Set(i, j, pixel)
			}
		}

		return dstImage, nil
	case f.NormalizeType == COLORED:
		dstBounds := src.Bounds()
		dstImage := image.NewRGBA(dstBounds)

		for i := src.Bounds().Min.X ; i < src.Bounds().Max.X; i++ {
			for j := src.Bounds().Min.Y; j < src.Bounds().Max.Y; j++ {
				r, g, b, a := src.At(i, j).RGBA()
				sum := r + g + b
				fmt.Printf("%f %f %f\n", float64(r) / float64(sum), g/sum, b/sum)
				pixel := color.RGBA{
					uint8(float64(r) / float64(sum) * 255),
					uint8(float64(g) / float64(sum) * 255),
					uint8(float64(b) / float64(sum) * 255),
					uint8(a)}
				dstImage.Set(i, j, pixel)
			}
		}

		return dstImage, nil
	default:
		return nil, errors.New("unexpected normalization type!!!")
	}
}