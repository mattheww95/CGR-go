package main

import (
	"os"
	"fmt"
	"math/rand"
	"log/slog"
	"github.com/golang/protobuf/proto"

)

func main(){
	fastas := ReadFasta(os.Args[1])
	var cgr_size uint64 = 1024
	rand.Seed(42)
	x := rand.Uint64() % cgr_size
	y := rand.Uint64() % cgr_size
	slog.Debug("Points", slog.Uint64("x", uint64(x)), slog.Uint64("y", uint64(y)))
	for _, f := range fastas {
		fmt.Println(f)
		i := *f
		cgr_map := CreateCGRMap(cgr_size)
		cgr_map.AddPoint(x, y)
		for _, nuc := range i.Sequence {
			slog.Debug("Points",slog.Any("Nuc", nuc), slog.Uint64("x", uint64(x)), slog.Uint64("y", uint64(y)))
			x, y = cgr_map.NextPoint(x, y, nuc)
			cgr_map.AddPoint(x, y)
			slog.Debug("Points", slog.Uint64("x", uint64(x)), slog.Uint64("y", uint64(y)))
		}
		output_obj := &CGR{
			Name: i.Header,
			Cgr: cgr_map.Data,
			Size: cgr_map.Size,
		}
		data, err := proto.Marshal(output_obj)
		if err != nil {
			slog.Error("Could not serialize fasta.")
		}

		new_data := &CGR{}
		err = proto.Unmarshal(data, new_data)
		if err != nil {
			slog.Error("Could not unmarshal data.")
		}
		fmt.Println("Data:", output_obj)
		fmt.Println("New Data:", new_data)

	}
}
