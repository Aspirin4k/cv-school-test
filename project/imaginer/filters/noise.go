package filters

import (
	"image"
	"image/color"
	"math/rand"

	"github.com/Aspirin4k/cv_school_test/project/imaginer"
)

var (
	avg = float64(2)
	stddev = float64(0.05)
)

type Noise struct {
	imaginer.Filter
}

func (f *Noise)  ApplyFilter(src image.Image) (image.Image, error) {
	dstBounds := src.Bounds()
	dstImage := image.NewRGBA(dstBounds)

	for i := src.Bounds().Min.X ; i < src.Bounds().Max.X; i++ {
		for j := src.Bounds().Min.Y ; j < src.Bounds().Max.Y; j++ {
			r, g, b, _ := src.At(i, j).RGBA()
			Y := R * float64(r) + G * float64(g) + B * float64(b)

			// Нормальное распределение
			rnd := (rand.NormFloat64() * stddev + avg) * MAX

			newY := uint32(Y + rnd) % uint32(MAX)

			pixel := color.Gray{uint8(newY / 256)}
			dstImage.Set(i, j, pixel)
		}
	}

	return dstImage, nil
}