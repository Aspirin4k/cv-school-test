package fragmenter

import (
	"strings"
	"strconv"
	"os"
	"image"
	"image/draw"
	"image/png"
	"path/filepath"
)

func ProcceedImage(fileName string, data []byte, imagesLocation string, fragmentsLocation string) error {
	iterator := 0
	srcFile, err := os.Open(imagesLocation + fileName + ".png")
	if err != nil {
		return err
	}
	defer srcFile.Close()

	srcImage, _, err := image.Decode(srcFile)
	if err != nil {
		return err
	}

	// Сплитим байтовый массив как строку по строкам
	lines := strings.Split(string(data[:]), "\n")
	for _, line := range lines {
		// Виндовые спецсимволы
		line = strings.Replace(line, "\r", "", -1)
		coords := strings.Split(line, ",")

		// Неэлегантно
		lx, err := strconv.Atoi(coords[0])
		if err != nil {
			return err
		}
		ly, err := strconv.Atoi(coords[1])
		if err != nil {
			return err
		}
		rx, err := strconv.Atoi(coords[2])
		if err != nil {
			return err
		}
		ry, err := strconv.Atoi(coords[3])
		if err != nil {
			return err
		}

		// Создается изображение размера фрагмента
		dstImage := image.NewRGBA(image.Rect(0, 0, rx - lx, ry - ly))
		// В целевое изображение рисуется фрагмент исходного
		draw.Draw(dstImage, dstImage.Bounds(), srcImage, image.Point{lx,ly}, draw.Src)

		// Создается папка, если еще не
		dstPath, err := filepath.Abs(fragmentsLocation)
		if err != nil {
			return err
		}
		os.MkdirAll(dstPath, os.ModePerm)
		// Создается файл в системе
		dstFile, err := os.Create(fragmentsLocation + fileName + "_" + strconv.Itoa(iterator) + ".png")
		if err != nil {
			return err
		}

		// В созданный файл заносится полученное изображение
		png.Encode(dstFile, dstImage)
		// Т.к. в цикле, то лучше сразу закрыть файл без откладывания
		dstFile.Close()

		iterator++
	}

	return nil
}