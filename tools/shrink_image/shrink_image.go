package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"io/ioutil"
	"os"
	"strings"
)

const inDir = "static/images/full-originals/"
const outDir = "static/images/fulls/"

func shrinkImage(args []string) error {
	for _, name := range args {
		fin, err := os.Open(name)
		if err != nil {
			return err
		}
		defer fin.Close()
		img, _, err := image.Decode(fin)
		if err != nil {
			return err
		}
		splitName := strings.Split(fin.Name(), "/")
		fout, err := os.Create(outDir + splitName[len(splitName)-1])
		if err != nil {
			return err
		}
		err = jpeg.Encode(fout, img, &jpeg.Options{Quality: 30})
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	args := os.Args
	var err error

	if len(args) == 1 {
		fmt.Println("need image name arguments or directory name for all images")
	} else if len(args) == 2 {
		file, err := os.Open(args[1])
		if err != nil {
			panic(err)
		}
		fs, err := file.Stat()
		if err != nil {
			panic(err)
		}
		if fs.IsDir() {
			files, err := ioutil.ReadDir(args[1])
			if err != nil {
				panic(err)
			}
			args := func() []string {
				fileNames := make([]string, len(files))
				for i, file := range files {
					if strings.HasPrefix(file.Name(), ".") {
						continue
					}
					fileNames[i] = file.Name()
				}
				return fileNames
			}()
			err = shrinkImage(args)
		} else {
			err = shrinkImage(args[1:])
		}
	} else {
		err = shrinkImage(args)
	}

	if err != nil {
		panic(err)
	}
}
