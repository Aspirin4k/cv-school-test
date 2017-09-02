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
	// 1. greyscale
	fragments, err := ioutil.ReadDir(*fragmentsLocation)
	if err != nil {
		fmt.Println(err)
	}

	for _, f :=  range fragments {
		provider := &filters.GreyScale{}
		fileName := f.Name()[:strings.LastIndex(f.Name(),".")]
		imaginer.SetFilter(*fragmentsLocation + f.Name(), *greyscaleLocation + fileName + "_grey.png", provider)
	}

	// 2. flip
	for _, f :=  range fragments {
		provider := &filters.Flip{}
		fileName := f.Name()[:strings.LastIndex(f.Name(),".")]
		imaginer.SetFilter(*fragmentsLocation + f.Name(), *flipLocation + fileName + "_flip.png", provider)
	}
}