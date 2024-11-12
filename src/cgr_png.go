package main

import (
	"image"
	"image/color"
	"strings"
	"image/png"
	"path"
	"fmt"
	"os"
)



func EncodeCGRToPoints(values *CGR, r *image.Rectangle) image.Image {

	new_image := image.NewGray(*r)

	points := (*values).Cgr
	size := int((*values).Size)
	gray_value := color.Gray{Y: 200}
	for idx, value := range points {
		// values idx % size gives column, and values / size gives row
		if value != 0 {
			column := idx % size
			row := idx / size
			new_image.Set(row, column, gray_value)
		}
	}
	return new_image
}

func EncodeImage(new_image *image.Image, output_name string, output_dir string){
	output_fn := fmt.Sprintf("%s.png", output_name)
	output_file := path.Join(output_dir, output_fn)

	f, err := os.Create(output_file)
	if err != nil {
		panic(err)
	}

	if err := png.Encode(f, *new_image); err != nil {
		f.Close()
		panic(err)
	}

	if err := f.Close(); err != nil {
		panic(err)
	}
}

func WriteImage(cgr *CGR, output_dir string) {
	cgr_size := (*cgr).Size
	new_rectangle := image.Rect(0, 0, int(cgr_size), int(cgr_size))
	new_image := EncodeCGRToPoints(cgr, &new_rectangle)
	EncodeImage(&new_image, strings.Replace((*cgr).Name, ">", "", -1), output_dir)
}
