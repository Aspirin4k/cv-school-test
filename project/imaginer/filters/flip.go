package filters

import (
	"image"
	"image/color"

	"github.com/Aspirin4k/cv_school_test/project/imaginer"
)

type Flip struct {
	imaginer.Filter
}

func (f *Flip)  ApplyFilter(src image.Image) (image.Image, error) {
	dstBounds := src.Bounds()
	dstImage := image.NewRGBA(dstBounds)

	// По сетке пикселей
	// ... дополнить комментарий
	for i := src.Bounds().Min.X ; i < src.Bounds().Max.X; i++ {
		for j := src.Bounds().Min.Y ; j < src.Bounds().Max.Y; j++ {
			r, g, b, a := src.At(i, j).RGBA()
			pixel := color.RGBA{uint8(r),uint8(g),uint8(b),uint8(a)}
			dstImage.Set(src.Bounds().Max.X - i - 1, j, pixel)
		}
	}

	return dstImage, nil
}
