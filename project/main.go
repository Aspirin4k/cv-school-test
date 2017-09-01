package main

import (
	"io/ioutil"
	"flag"
	"fmt"

	"github.com/Aspirin4k/cv_school_test/project/fragmenter"
	"strings"
)

func main() {
	// Возможные входные параметры
	annotationsLocation := flag.String("annotations", "./../annotations/", "location of annotations")
	imagesLocation := flag.String("images", "./../images/", "location of images")
	fragmentsLocation := flag.String("fragments", "./../fragments/", "location for fragments")
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
}