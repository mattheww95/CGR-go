package main

import (
	"os"
	"fmt"
	"math/rand"
	"strings"
	"log/slog"
	"github.com/golang/protobuf/proto"
	"path"
	"github.com/integrii/flaggy"
)


func MakeDir(name string) string {

	output_path := path.Join(name, "cgr");
	stat, err := os.Stat(output_path)
	if err == nil && stat.IsDir() {
		return output_path
	}
	slog.Info("Creating directory", slog.String("Path", output_path))
	err = os.MkdirAll(output_path, os.ModePerm)
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

var INPUT_FASTAS string = ""
var OUTPUT_DIRECTORY string = ""
var CGR_SIZE uint64 = 1024
var create_cgr *flaggy.Subcommand

const version string = "0.0.1"

func cli(){
	flaggy.SetName("Chaos Game Representation with go")
	flaggy.SetDescription("A tool kit for creating and manipulating chaos game representations of genomes")
	flaggy.ShowHelpOnUnexpectedEnable()

	create_cgr = flaggy.NewSubcommand("create")
	create_cgr.Description = "Compute a CGR for all sequences in a multi-fasta"
	create_cgr.String(&INPUT_FASTAS, "i", "input", "A multifasta file to use for CGR computation")
	create_cgr.String(&OUTPUT_DIRECTORY, "o", "output-directory", "The output directory for computed CGR files.")
	cgr_size_help := fmt.Sprintf("The cgr map size to create: %d", CGR_SIZE)
	create_cgr.UInt64(&CGR_SIZE, "c", "cgr-size", cgr_size_help)


	flaggy.AttachSubcommand(create_cgr, 1)
	flaggy.Parse()

}

func main(){
	cli()
	if len(os.Args) <= 1 {
		flaggy.ShowHelpAndExit("No inputs passed")
	}

	if create_cgr.Used {
		if INPUT_FASTAS == "" {
			flaggy.ShowHelpAndExit("input fastas is a mandatory argument")
		}

		if OUTPUT_DIRECTORY == "" {
			var err error
			OUTPUT_DIRECTORY, err = os.Getwd()
			if err != nil {
				slog.Error("Could not get current directory", slog.Any("err", err))
			}
		}

		fastas := ReadFasta(INPUT_FASTAS)
		var cgr_size uint64 = CGR_SIZE
		rand.Seed(42)
		x := rand.Uint64() % cgr_size
		y := rand.Uint64() % cgr_size
		slog.Debug("Points", slog.Uint64("x", uint64(x)), slog.Uint64("y", uint64(y)))
		output_path := MakeDir(OUTPUT_DIRECTORY)
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
}
