package  main
/// Create a string of random nucleotides for testing and playing
import (
	"math/rand"
	"fmt"
)



const nucleotides uint64 = 4

func CreateRandomSequence(sequence_length uint64){
	rand.Seed(42)
	fmt.Println(">RandomFasta")
	for i := uint64(0); i < sequence_length; i++ {
		num := rand.Uint64() % nucleotides
		switch num {
		case 0:
			fmt.Print("a")
		case 1:
			fmt.Print("t")
		case 2:
			fmt.Print("c")
		case 3:
			fmt.Print("g")
		}
		if i % 80 == 0 {
			fmt.Println()
		}
	}
	fmt.Println()

}
