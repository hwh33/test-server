package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"math"
	"os"
	"strings"
)

/* Define the image type and implement the
 * Image interface.
 */

type Image struct {
	color_model color.Model
	bounds      image.Rectangle
	color_grid  [][]color.RGBA
}

func (img Image) ColorModel() color.Model {
	return img.color_model
}

func (img Image) Bounds() image.Rectangle {
	return img.bounds
}

func (img Image) At(x, y int) color.Color {
	return img.color_grid[x][y]
}

/* Maps a value in the input range to one in the output range. */
func RangeMap(value, in_range_min, in_range_max, out_range_min, out_range_max float64) float64 {
	fractional_loc := (value - in_range_min) / (in_range_max - in_range_min)
	return fractional_loc*(out_range_max-out_range_min) + out_range_min
}

/* f(x, y) is the function plotted. */
func f(x, y float64) float64 {
	// return -1 * x
	// return (450-x)*(450-x) + (450-y)*(450-y)
	return -1 * ((450-x)*(450-x) + (450-y)*(450-y))
}

/* Pic(max_x, max_y) generates a color plot of f. */
func Pic(bounds image.Rectangle, x_increment, y_increment int) Image {
	// First we find all values f(x, y) over the range.
	num_x_increments := bounds.Max.X - bounds.Min.X + 1
	num_y_increments := bounds.Max.Y - bounds.Min.Y + 1
	num_f_vals := num_x_increments * num_y_increments
	f_vals := make([][]float64, num_f_vals)
	min, max := math.Inf(1), math.Inf(-1)
	for x := bounds.Min.X; x <= bounds.Max.X; x += x_increment {
		f_vals[x] = make([]float64, num_y_increments)
		for y := bounds.Min.Y; y <= bounds.Max.Y; y += y_increment {
			f_vals[x][y] = f(float64(x), float64(y))
			if f_vals[x][y] > max {
				max = f_vals[x][y]
			}
			if f_vals[x][y] < min {
				min = f_vals[x][y]
			}
		}
	}

	// Now we translate these values to RGB colors.
	max_RGBA := 255
	alpha := max_RGBA
	color_grid := make([][]color.RGBA, num_f_vals)
	// debugging
	fmt.Printf("(min, max) of f(x,y): (%f, %f)\n", min, max)
	for x := bounds.Min.X; x <= bounds.Max.X; x += x_increment {
		color_grid[x] = make([]color.RGBA, num_y_increments)
		for y := bounds.Min.Y; y <= bounds.Max.Y; y += y_increment {
			curr_val := f_vals[x][y]
			var blue, green, red float64
			// Assign blue values. These are largest at the minimum of f.
			if curr_val > max {
				blue = 0
			} else {
				blue = RangeMap(curr_val, min, max, float64(max_RGBA), 0)
			}
			// Assign green values. These are largest in the middle of the range.
			if curr_val <= max/2 {
				green = RangeMap(curr_val, min, max/2, 0, float64(max_RGBA))
			} else {
				green = RangeMap(curr_val, max/2, max, float64(max_RGBA), 0)
			}
			// Assign red values. These are largest at the maximum of f.
			if curr_val < min {
				red = 0
			} else {
				red = RangeMap(curr_val, min, max, 0, float64(max_RGBA))
			}
			color_grid[x][y] = color.RGBA{uint8(red), uint8(green), uint8(blue), uint8(alpha)}
			// debugging
			if y == 450 && (x == 150 || x == 200) {
				fmt.Printf("RGB = [%f, %f, %f] at (%d, %d)\n", red, green, blue, x, y)
			}
		}
	}

	// Now we create and return the Image.
	return Image{color.RGBAModel, bounds, color_grid}
}

func main() {
	filepath := "/Users/Harry/Desktop/test_plot.jpeg"

	// Get the directory and file names.
	separator := strings.LastIndex(filepath, "/")
	dir_name, file_name := filepath[:separator], filepath[separator+1:]
	cd_err := os.Chdir(dir_name)
	image_file, create_err := os.Create(file_name)
	if cd_err != nil || create_err != nil {
		fmt.Println("Error creating file:")
		if cd_err != nil {
			fmt.Println(cd_err)
		} else if create_err != nil {
			fmt.Println(create_err)
		}
		os.Exit(1)
	}

	bounds := image.Rectangle{image.Point{0, 0}, image.Point{900, 900}}
	m := Pic(bounds, 1, 1)
	jpeg.Encode(image_file, m, &jpeg.Options{jpeg.DefaultQuality})
}
