package main
/*
 * BMP array generation for fasta sequences
 */

import (
	"log/slog"
)

type Boundry struct {
	X, Y uint64
}


type CGRMap struct {
	Data []uint8
	Size uint64
	A Boundry
	T Boundry
	C Boundry
	G Boundry
}


/// Create a new CGRMap object of the specified size
func CreateCGRMap(size uint64) CGRMap {
	map_size := size * size
	cgr_data := make([]uint8, map_size, map_size)
	return CGRMap{
		Data: cgr_data,
		Size: size,
		A: Boundry{0, size},
		T: Boundry{size, 0},
		C: Boundry{size, size},
		G: Boundry{0, 0},
	}
}

/// Data points arte stored in a  contiguous array
/// therefore a special method is needed to add points
func (m *CGRMap) AddPoint(x uint64, y uint64){
	slog.Debug("Adding Point:", slog.Uint64("x", x), slog.Uint64("y", y))
	point := (x * m.Size) + y
	m.Data[point] = 1
}

/// Func get next point
func (m *CGRMap) NextPoint(previous_x uint64, previous_y uint64, nucleotide rune) (uint64, uint64) {
	var new_x uint64
	var new_y uint64
	slog.Debug("Adding nucleotied", slog.Any("new nuc", nucleotide))
	switch(nucleotide){
		case 'a', 'A':
		new_x = previous_x - (previous_x >> 1)
		new_y = previous_y + ((m.A.Y - previous_y) >> 1)
		case 't', 'T':
		new_x = previous_x + ((m.T.X - previous_x) >> 1)
		new_y = previous_y - (previous_y >> 1)
		case 'g', 'G':
		new_x = previous_x - (previous_x >> 1)
		new_y = previous_y - (previous_y >> 1)
		case 'c', 'C':
		new_x = previous_x + ((m.C.X - previous_x) >> 1)
		new_y = previous_y + ((m.C.Y - previous_y) >> 1)
		default:
		new_x = previous_x
		new_y = previous_y
	}
	return new_x, new_y
}
