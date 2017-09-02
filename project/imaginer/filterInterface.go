package imaginer

import (
	"image"
)

// Интерфейс для реализации фильтров
type Filter interface {
	ApplyFilter(src image.Image) (image.Image, error)
}