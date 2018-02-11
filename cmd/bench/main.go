package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
	"testing"

	"github.com/dbriemann/chesskimo"
)

const (
	NANOS_PER_MILLI = 1000000
	NANOS_PER_SEC   = 1000000000
)

var version = "undefined"

type set struct {
	Fen    string
	Depth  int
	Result uint64
}

var profile = flag.String("profile", "", "specify a file to write profile info")

var (
	board    chesskimo.Board
	active   int   = 0
	testsets []set = []set{
		set{Fen: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1", Depth: 6, Result: 119060324},
		set{Fen: "r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1", Depth: 5, Result: 193690690},
		set{Fen: "n1n5/PPPk4/8/8/8/8/4Kppp/5N1N b - - 0 1", Depth: 6, Result: 71179139},
		set{Fen: "r3k2r/Pppp1ppp/1b3nbN/nP6/BBP1P3/q4N2/Pp1P2PP/R2Q1RK1 w kq - 0 1", Depth: 5, Result: 15833292},
		set{Fen: "8/2p5/3p4/KP5r/1R3p1k/8/4P1P1/8 w - - 0 0", Depth: 7, Result: 178633661},
	}
)

func main() {
	fmt.Println("Version", version)

	flag.Parse()
	if *profile != "" {
		f, err := os.Create(*profile)
		if err != nil {
			panic(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	avgNPS := float64(0)

	fmt.Println("Starting Chesskimo move generator benchmarks..")
	for i := 0; i < len(testsets); i++ {
		set := testsets[i]
		res := testing.Benchmark(benchActive)
		nps := float64(testsets[i].Result) / (float64(res.NsPerOp()) / NANOS_PER_SEC)
		avgNPS += nps
		fmt.Printf("FEN: %s, depth: %d, nodes: %d, duration: %f sec, NPS: %f\n", set.Fen, set.Depth, set.Result, float64(res.NsPerOp())/NANOS_PER_SEC, nps)
	}

	avgNPS /= float64(len(testsets))

	fmt.Printf("Average NPS: %f\n", avgNPS)
}

func benchActive(b *testing.B) {
	board.SetFEN(testsets[active].Fen)
	result := board.Perft(testsets[active].Depth)
	if testsets[active].Result != result {
		b.Fatalf("Wrong Perft result for FEN %s: %d but should be %d", testsets[active].Fen, result, testsets[active].Result)
	}
}
