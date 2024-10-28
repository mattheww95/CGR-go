package main

import (
	"os"
	_ "fmt"
	"math/rand"
	"log/slog"
)

func main(){
	fastas := ReadFasta(os.Args[1])
	var cgr_size uint64 = 1024
	rand.Seed(42)
	x := rand.Uint64() % cgr_size
	y := rand.Uint64() % cgr_size
	cgr_map := CreateCGRMap(cgr_size)
	cgr_map.AddPoint(x, y)
	slog.Debug("Points", slog.Uint64("x", uint64(x)), slog.Uint64("y", uint64(y)))
	for _, i := range fastas {
		for _, nuc := range i.Sequence {
			slog.Debug("Points",slog.Any("Nuc", nuc), slog.Uint64("x", uint64(x)), slog.Uint64("y", uint64(y)))
			x, y = cgr_map.NextPoint(x, y, nuc)
			cgr_map.AddPoint(x, y)
			slog.Debug("Points", slog.Uint64("x", uint64(x)), slog.Uint64("y", uint64(y)))
		}
	}
}
