package filters

import (
	"image"
	"image/color"

	"github.com/Aspirin4k/cv_school_test/project/imaginer"
)

var (
	R = 0.2126
	G = 0.7152
	B = 0.0722
)

type GreyScale struct {
	imaginer.Filter
}

func (f *GreyScale)  ApplyFilter(src image.Image) (image.Image, error) {
	dstBounds := src.Bounds()
	dstImage := image.NewRGBA(dstBounds)

	// По сетке пикселей
	// ... дополнить комментарий
	for i := src.Bounds().Min.X ; i < src.Bounds().Max.X; i++ {
		for j := src.Bounds().Min.Y ; j < src.Bounds().Max.Y; j++ {
			r, g, b, _ := src.At(i, j).RGBA()
			Y := R * float64(r) + G * float64(g) + B * float64(b)
			pixel := color.Gray{uint8(Y / 256)}
			dstImage.Set(i, j, pixel)
		}
	}

	return dstImage, nil
}