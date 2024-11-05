package main

import (
	"os"
	"fmt"
	"math/rand"
	"strings"
	"log/slog"
	"github.com/golang/protobuf/proto"
	"path"

)


func MakeDir(name string) string {

	cwd, err := os.Getwd()
	if err != nil {
		slog.Error("Could not get current directory", slog.Any("err", err))
	}

	output_path := path.Join(cwd, name);
	stat, err := os.Stat(output_path)
	if err == nil && stat.IsDir() {
		return output_path
	}
	slog.Info("Creating directory", slog.String("Path", output_path))
	err = os.Mkdir(output_path, os.ModePerm)
	if err != nil {
		slog.Error("Could not make directory", slog.Any("err", err))
	}
	return output_path
}


func WriteProtoBuff(output_obj *CGR, output_dir string){

	data, err := proto.Marshal(output_obj)
	if err != nil {
		slog.Error("Could not serialize fasta.")
	}
	output_name := fmt.Sprintf("%s.pb", strings.Replace((*output_obj).Name, ">", "", -1))
	output_fi := path.Join(output_dir, output_name)
	fi, err := os.OpenFile(output_fi, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()

	fi.Write(data)

}

func main(){
	fastas := ReadFasta(os.Args[1])
	var cgr_size uint64 = 1024
	rand.Seed(42)
	x := rand.Uint64() % cgr_size
	y := rand.Uint64() % cgr_size
	slog.Debug("Points", slog.Uint64("x", uint64(x)), slog.Uint64("y", uint64(y)))
	output_path := MakeDir("cgr")
	for _, f := range fastas {
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

		WriteProtoBuff(output_obj, output_path)
	}
}
