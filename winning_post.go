package assets

import (
	_ "embed"

	"github.com/filecoin-project/go-state-types/abi"
)

//go:embed winning-post-vanilla/vanilla.dat
var vanillaRaw []byte

var WinningPoStWarmUp = struct {
	CommR      [32]byte
	SectorID   abi.SectorID
	Vanilla    []byte
	Randomness abi.PoStRandomness
	SealProof  abi.RegisteredSealProof
	PoStProof  abi.RegisteredPoStProof
}{
	CommR: [32]byte{187, 43, 211, 193, 248, 228, 156, 255, 247, 51, 129, 20, 122, 131, 181, 10, 97, 58, 42, 17, 234, 107, 231, 84, 79, 154, 199, 112, 109, 238, 50, 96},
	SectorID: abi.SectorID{
		Miner:  10938,
		Number: 99,
	},
	Vanilla:    vanillaRaw,
	Randomness: make(abi.PoStRandomness, abi.RandomnessLength),
	SealProof:  abi.RegisteredSealProof_StackedDrg32GiBV1_1,
	PoStProof:  abi.RegisteredPoStProof_StackedDrgWinning32GiBV1,
}
