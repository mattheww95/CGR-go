
package main

import (
	"os"
	"bufio"
	"strings"
	"fmt"
)

// TODO need to make the sequence a byte array
type Fasta struct {
	Header string
	Sequence string
}


/// Read Fasta file
func ReadFasta(FileIn string) []*Fasta {
	f, err := os.Open(FileIn)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	sequences := make([]*Fasta, 0, 1000)
	s := bufio.NewScanner(f)
	headerFound := false
	sequence := make([]string, 0, 100)
	var header string
	for s.Scan() {
		line := s.Text()
		switch {
		case line == "":
			continue
		case line[0] != '>':
			if !headerFound {
				panic("missing header")
			}
			sequence = append(sequence, strings.TrimSuffix(line, "\n"))
		case headerFound:
			fallthrough
		case line[0] == '>':
			header = strings.TrimSuffix(line, "\n");
			if headerFound {
				record := &Fasta{Header: header, Sequence: strings.Join(sequence[:], "")}
				sequences = append(sequences, record)
				sequence = sequence[:0]
			}
			headerFound = true
		}
	}
	fmt.Println(sequences[0])

	if err := s.Err(); err != nil {
		panic(err)
	}
	return sequences
}
