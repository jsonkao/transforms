// Package parser contains useful functions to parse a scripts file.
package parser

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"../display"
	"../draw"
	"../matrix"
)


/* ParseFile goes through the file named filename and performs all of the
   actions listed in that file.

The file follows the following format:
  Every command is a single character that takes up a line
  Any command that requires arguments must have those arguments in the 2nd line.
  The commands are as follows:
    line: add a line to the edge matrix -
	    takes 6 arguemnts (x0, y0, z0, x1, y1, z1)
	  ident: set the transform matrix to the identity matrix -
	  scale: create a scale matrix, then multiply the transform matrix by the
      scale matrix -
	    takes 3 arguments (sx, sy, sz)
    translate: create a translation matrix, then multiply the transform matrix
      by the translation matrix -
	    takes 3 arguments (tx, ty, tz)
    rotate: create an rotation matrix, then  multiply the transform matrix by
      the rotation matrix -
	    takes 2 arguments (axis, theta) axis should be x y or z
    apply: apply the current transformation matrix to the edge  matrix
	  display: draw the lines of the edge matrix to the screen display  the screen
	  save: draw the lines of the edge matrix to the screen save the screen to a
       file -
	    takes 1 argument (file name)
	  quit: end parsing
*/
func ParseFile(filename string,
	transform [][]float64,
	edges [][]float64,
	screen [][][]int) {

	file, err := os.Open(filename)
	if (err != nil) {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Immediate operations (no parameters)
		if line == "ident" {
			matrix.MakeIdentity(transform)
		} else if line == "display" {
			display.DisplayScreen(screen)
		} else if line == "apply" {
			matrix.MultiplyMatrices(&transform, &edges)
		} else if line == "save" {
			display.WriteScreenToPPM(screen)
		}

		scanner.Scan()

		// Non-immediate operations (has parameters)
		if line == "line" {
			draw.DrawLineFromParams(screen, FloatParams(scanner.Text)...)
		} else if line == "move" {
			// translate ...
		} else if line == "scale" {
			// dilation ...
		} else if line == "rotate" {
			// rotate ...
		}
}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func FloatParams(text string) (args []float64) {
	args = []float64{}
	for _, v := range strings.Fields(text) {
		floated, err := strconv.ParseFloat(v, 64)
		if (err != nil) {
			panic(err)
		}
		args = append(args, floated)
	}
	return
}
