package assets

import "testing"

func TestVanillaSize(t *testing.T) {
	if size := len(vanillaRaw); size != 164064 {
		t.Fatalf("expected 164064 bytes, got %d", size)
	}
}
