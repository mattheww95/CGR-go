package main

import (
	"image"
	"image/color"
	"strings"
	"image/png"
	"path"
	"fmt"
	"os"
	"golang.org/x/image/font"
    "golang.org/x/image/font/basicfont"
    "golang.org/x/image/math/fixed"
)


func addLabel(img *image.RGBA, x, y int, label string){
	// Lovingly taken from: https://stackoverflow.com/a/38300583
	col := color.Gray{Y: 200}
	pt := fixed.Point26_6{fixed.I(x), fixed.I(y)}

	d := &font.Drawer{
		Dst: img,
		Src: image.NewUniform(col),
		Face: basicfont.Face7x13,
		Dot: pt,
	}
	d.DrawString(label)
}

func EncodeCGRToPoints(values *CGR, r *image.Rectangle, imgx, imgy int, cgr_size int) image.Image {

	new_image := image.NewRGBA((*r).Bounds())

	points := (*values).Cgr
	size := int((*values).Size)
	gray_value := color.Gray{Y: 200}
	off_color := color.Gray{Y: 0}

	for i := 0; i < imgx + (2 * cgr_size); i++ {
		for y := 0; y < imgy + (2 * cgr_size); y++ {
			new_image.Set(i, y, off_color)
		}
	}

	for idx, value := range points {
		// values idx % size gives column, and values / size gives row
		column := idx % size
		row := idx / size
		if value != 0 {
			new_image.Set(row + imgx, column + imgy, gray_value)
		}else{
			new_image.Set(row + imgx, column + imgy, off_color)
		}
	}
	bounds := new_image.Bounds()
	label_offset_x := imgx - 10
	label_offset_y := imgy - 10
	// TODO need to verify the coordinates match up
	//addLabel(new_image, 0, cgr_size, "A")
	addLabel(new_image, bounds.Min.X + label_offset_x, 0 + label_offset_y, "A")
	//addLabel(new_image, cgr_size, 0, "T")
	addLabel(new_image, bounds.Max.X - label_offset_x, bounds.Max.Y - label_offset_y, "T")
	//addLabel(new_image, cgr_size, cgr_size, "C")
	addLabel(new_image, bounds.Max.X - label_offset_x, bounds.Min.Y + label_offset_y, "C")
	//addLabel(new_image, 0, 0, "G")
	addLabel(new_image, 0 + label_offset_x, bounds.Max.Y - label_offset_y, "G")
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
	const offset_img_x, offset_img_y int = 50, 50
	new_rectangle := image.Rect(0, 0, int(cgr_size) + (2 * offset_img_x), int(cgr_size) + (2 * offset_img_y))
	new_image := EncodeCGRToPoints(cgr, &new_rectangle, offset_img_x, offset_img_y, int(cgr_size))
	EncodeImage(&new_image, strings.Replace((*cgr).Name, ">", "", -1), output_dir)
}
