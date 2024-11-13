package main

// TODO add in GCR representation

import (
	"os"
	"fmt"
	"math/rand"
	"strings"
	"log/slog"
	"github.com/golang/protobuf/proto"
	"path"
	"github.com/integrii/flaggy"
	"path/filepath"
)


func MakeDir(name string, dir_name string) string {

	output_path := path.Join(name, dir_name);
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

// Split one fastas into multiple cgr files
func SplitFastaToCGR(x, y, cgr_size uint64, output_path string, fastas *[]*Fasta){
	for _, f := range (*fastas) {
		i := *f
		slog.Debug("Fasta sequence", slog.String("Header", i.Header), slog.String("Sequence", i.Sequence))
		slog.Info(fmt.Sprintf("Creating CGR for %s", i.Header))
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

// Convert all sequences in a fasta into its cgr representation
func MFAToCGR(x, y, cgr_size uint64, output_name string, output_path string, fastas *[]*Fasta){
	cgr_map := CreateCGRMap(cgr_size)
	for _, f := range (*fastas) {
		i := *f
		slog.Debug("Fasta sequence", slog.String("Header", i.Header), slog.String("Sequence", i.Sequence))
		slog.Info(fmt.Sprintf("Creating CGR for %s", i.Header))
		cgr_map.AddPoint(x, y)
		for _, nuc := range i.Sequence {
			slog.Debug("Points",slog.Any("Nuc", nuc), slog.Uint64("x", uint64(x)), slog.Uint64("y", uint64(y)))
			x, y = cgr_map.NextPoint(x, y, nuc)
			cgr_map.AddPoint(x, y)
			slog.Debug("Points", slog.Uint64("x", uint64(x)), slog.Uint64("y", uint64(y)))
		}

	}

	output_obj := &CGR{
		Name: output_name,
		Cgr: cgr_map.Data,
		Size: cgr_map.Size,
	}
	WriteProtoBuff(output_obj, output_path)
}

/// Convert an input multi fasta in a collection of cgr files, one per a sequence
func FastaToCGR(input string, output_dir string, cgr_size uint64, split bool){

	fastas := ReadFasta(input)
	rand.Seed(42)
	x := rand.Uint64() % cgr_size
	y := rand.Uint64() % cgr_size
	slog.Debug("Points", slog.Uint64("x", uint64(x)), slog.Uint64("y", uint64(y)))
	output_path := MakeDir(output_dir, "cgr")
	if(split){
		SplitFastaToCGR(x, y, cgr_size, output_path, &fastas)
	}else{
		output_base := filepath.Base(input)
		seperated_name := strings.Split(output_base, ".")
		MFAToCGR(x, y, cgr_size, seperated_name[0], output_path, &fastas)
	}
}



func ReadProtoBuffers(directory string) []*CGR {
	// Once better established this can be replaced with a generator
	input_expr := path.Join(directory, "*.pb")
	matches, err := filepath.Glob(input_expr)
	if err != nil {
		panic(err)
	}
	cgr_values := make([]*CGR, len(matches))
	for idx, match := range matches {
		input_bytes, err := os.ReadFile(match)
		if err != nil {
			panic(err)
		}
		new_cgr := &CGR{}
		err = proto.Unmarshal(input_bytes, new_cgr)
		cgr_values[idx] = new_cgr
	}
	return cgr_values
}

var INPUT_FASTAS string = ""
var OUTPUT_DIRECTORY string = ""
var READ_IN_CGR_DIR string = ""
var SPLIT_FASTA bool = false
var CGR_SIZE uint64 = 1024
var create_cgr *flaggy.Subcommand
var read_cgrs *flaggy.Subcommand
var random *flaggy.Subcommand

const version string = "0.0.1"

func cli(){
	flaggy.SetName("Chaos Game Representation with go")
	flaggy.SetDescription("A tool kit for creating and manipulating chaos game representations of genomes")
	flaggy.ShowHelpOnUnexpectedEnable()

	create_cgr = flaggy.NewSubcommand("create")
	create_cgr.Description = "Compute a CGR for all a fasta sequences"
	create_cgr.String(&INPUT_FASTAS, "i", "input", "A multifasta file to use for CGR computation")
	create_cgr.Bool(&SPLIT_FASTA, "s", "split", "Generate CGR for individual sequences of file input.")
	create_cgr.String(&OUTPUT_DIRECTORY, "o", "output-directory", "The output directory for computed CGR files.")
	cgr_size_help := fmt.Sprintf("The cgr map size to create: %d", CGR_SIZE)
	create_cgr.UInt64(&CGR_SIZE, "c", "cgr-size", cgr_size_help)

	read_cgrs = flaggy.NewSubcommand("image")
	read_cgrs.Description = "Create png images from pre-computed cgr objects"
	read_cgrs.String(&READ_IN_CGR_DIR, "i", "input", "A directory of pre-computed cgr representations.")
	read_cgrs.String(&OUTPUT_DIRECTORY, "o", "output-directory", "The output directory for created images.")


	random = flaggy.NewSubcommand("random")
	random.Description = "Create a sequence of a random length. Writs to stdout"
	random.UInt64(&CGR_SIZE, "c", "sequence-size", "Specify the length of a random sequence to write.")

	flaggy.AttachSubcommand(create_cgr, 1)
	flaggy.AttachSubcommand(read_cgrs, 1)
	flaggy.AttachSubcommand(random, 1)
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
		FastaToCGR(INPUT_FASTAS, OUTPUT_DIRECTORY, CGR_SIZE, SPLIT_FASTA)
	}else if read_cgrs.Used {
		if  READ_IN_CGR_DIR == "" {
			flaggy.ShowHelpAndExit("No input directory passed.")
		}

		if OUTPUT_DIRECTORY == "" {
			var err error
			OUTPUT_DIRECTORY, err = os.Getwd()
			if err != nil {
				slog.Error("Could not get current directory", slog.Any("err", err))
			}
			if stat, _ := os.Stat(OUTPUT_DIRECTORY); !stat.IsDir(){
				flaggy.ShowHelpAndExit("Output directory passed is not a directory")
			}
		}
		output_path := MakeDir(OUTPUT_DIRECTORY, "png")

		old_cgrs := ReadProtoBuffers(READ_IN_CGR_DIR)
		for _, i := range old_cgrs {
			WriteImage(i, output_path)
		}
	}else if random.Used {
		CreateRandomSequence(CGR_SIZE)
	}else{
		flaggy.ShowHelpAndExit("Command not recognized.")
	}
}
