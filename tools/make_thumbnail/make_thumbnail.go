package main

import (
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/nfnt/resize"
)

const inDir = "static/images/fulls/"
const outDir = "static/images/thumbs/"

func makeThumbnail(full os.FileInfo, wg *sync.WaitGroup) {
	if strings.HasPrefix(full.Name(), ".") {
		wg.Done()
		return
	}
	fin, err := os.Open(inDir + full.Name())
	if err != nil {
		panic(err)
	}
	defer fin.Close()

	img, _, err := image.Decode(fin)
	if err != nil {
		log.Fatal("Cannot decode image:", err)
	}
	resizedImg := resize.Resize(370, 217, img, resize.Lanczos3)
	fout, err := os.Create(outDir + full.Name())
	if err != nil {
		panic(err)
	}
	defer fout.Close()
	if strings.HasSuffix(fin.Name(), "png") {
		err = png.Encode(fout, resizedImg)
		if err != nil {
			panic(err)
		}
	} else {
		err = jpeg.Encode(fout, resizedImg, nil)
		if err != nil {
			panic(err)
		}
	}
	wg.Done()
}

func main() {
	fulls, err := ioutil.ReadDir(inDir)
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	wg.Add(len(fulls))
	for _, full := range fulls {
		go makeThumbnail(full, &wg)
	}
	wg.Wait()
}
