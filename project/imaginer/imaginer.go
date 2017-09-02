package imaginer

import (
	"os"
	"image"
	"strings"
	"image/png"
)

// Применить некоторый фильтр
/*
Параметры - путь до исходного файла, путь до результата, применяемый фильтр
 */
func SetFilter(srcFilename string, dstFilename string, filter Filter) error {

	srcFile, err := os.Open(srcFilename)
	if err != nil {
		return err
	}

	srcImage, _, err := image.Decode(srcFile)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// Создается папка, если еще не
	os.MkdirAll(dstFilename[:strings.LastIndex(dstFilename,"/")], os.ModePerm)

	dstFile, err := os.Create(dstFilename)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	// Применяем кастомный фильтр
	dstImage, err := filter.ApplyFilter(srcImage)
	if err != nil {
		return err
	}

	png.Encode(dstFile, dstImage)
	return nil
}