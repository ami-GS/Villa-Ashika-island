package main

import (
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/nfnt/resize"
)

const inDir = "static/images/fulls/"
const outDir = "static/images/thumbs/"

func makeThumbnail(full os.FileInfo, finChan chan bool) {
	if strings.HasPrefix(full.Name(), ".") {
		finChan <- true
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
	/*
		croppedImg, err := cutter.Crop(img, cutter.Config{
			Width:   370,
			Height:  217,
			Options: cutter.Copy,
			Mode:    cutter.Centered,
		})
	*/
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
		jpeg.Encode(fout, resizedImg, nil)
	}
	finChan <- true
}

func main() {
	fulls, err := ioutil.ReadDir(inDir)
	if err != nil {
		panic(err)
	}

	finChan := make(chan bool)
	for _, full := range fulls {
		go makeThumbnail(full, finChan)
	}
	count := 0
	for count != len(fulls) {
		select {
		case _ = <-finChan:
			count++
		default:
		}
	}

}
