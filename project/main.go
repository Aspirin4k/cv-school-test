package main

import (
	"io/ioutil"
	"flag"
	"fmt"
	"strings"

	"github.com/Aspirin4k/cv_school_test/project/fragmenter"
	"github.com/Aspirin4k/cv_school_test/project/imaginer/filters"
	"github.com/Aspirin4k/cv_school_test/project/imaginer"
)

func main() {
	// Возможные входные параметры
	annotationsLocation := flag.String("annotations", "./../annotations/", "location of annotations")
	imagesLocation := flag.String("images", "./../images/", "location of images")
	fragmentsLocation := flag.String("fragments", "./../fragments/", "location for fragments")
	greyscaleLocation := flag.String("greyscale", "./../fragments_grey/","location for greyscaled fragments")
	flipLocation := flag.String("flip", "./../fragments_flip/", "location for flipped fragments")
	flag.Parse()

	annotations, err := ioutil.ReadDir(*annotationsLocation)
	if err != nil {
		fmt.Println(err)
	}

	for _, f := range annotations {
		// Считываем файл с аннотациями и передаем обработчику
		var filesData []byte
		filesData, err = ioutil.ReadFile(*annotationsLocation + f.Name())
		if err != nil {
			fmt.Println(err)
		}
		// Убираем расширение
		fragmenter.ProcceedImage(
			f.Name()[:strings.LastIndex(f.Name(),".")],
			filesData,
			*imagesLocation,
			*fragmentsLocation)
		if err != nil {
			fmt.Println(err)
		}
	}

	// Далее идет дополнительная часть задания
	fragments, err := ioutil.ReadDir(*fragmentsLocation)
	if err != nil {
		fmt.Println(err)
	}

	// 1. greyscale
	for _, f :=  range fragments {
		provider := &filters.GreyScale{}
		fileName := f.Name()[:strings.LastIndex(f.Name(),".")]
		err = imaginer.SetFilter(*fragmentsLocation + f.Name(), *greyscaleLocation + fileName + "_grey.png", provider)
		if err != nil {
			fmt.Println(err)
		}
	}

	// 2. flip
	for _, f :=  range fragments {
		provider := &filters.Flip{}
		fileName := f.Name()[:strings.LastIndex(f.Name(),".")]
		err = imaginer.SetFilter(*fragmentsLocation + f.Name(), *flipLocation + fileName + "_flip.png", provider)
		if err != nil {
			fmt.Println(err)
		}
	}

	// Возможно задание построенно некорректно (Предлагается сделать нормализацию изображений, полученных в пункте 2,
	// предоставляя ссылку на метод нормализации черно-белых)
	// Поэтому было реализовано 2 способа: для черно-белых и цветных
	// 3 normalization - серые
	fragmentsGrey, err := ioutil.ReadDir(*greyscaleLocation)
	if err != nil {
		fmt.Println(err)
	}

	for _, f :=  range fragmentsGrey {
		provider := &filters.Normalize{}
		provider.NormalizeType = filters.GREYSCALED
		fileName := f.Name()[:strings.LastIndex(f.Name(),".")]
		err = imaginer.SetFilter(*greyscaleLocation + f.Name(), *greyscaleLocation +  fileName + "_normalized.png", provider)
		if err != nil {
			fmt.Println(err)
		}
	}

	// 3.5 normalization - цветные
	fragmentsFlip, err := ioutil.ReadDir(*flipLocation)
	if err != nil {
		fmt.Println(err)
	}

	for _, f :=  range fragmentsFlip {
		provider := &filters.Normalize{}
		provider.NormalizeType = filters.COLORED
		fileName := f.Name()[:strings.LastIndex(f.Name(),".")]
		err = imaginer.SetFilter(*flipLocation + f.Name(), *flipLocation + fileName + "_normalized.png", provider)
		if err != nil {
			fmt.Println(err)
		}
	}
}