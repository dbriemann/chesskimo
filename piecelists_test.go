package chesskimo

import (
	"testing"
)

// TestPieceList tests the whole functionality of a PieceList
func TestPieceList(t *testing.T) {
	// TODO - currently we iterate and test all items. When OTB is removed from list
	// -> only test 'Size' items.

	result0 := [16]Square{OTB, OTB, OTB, OTB, OTB, OTB, OTB, OTB, OTB, OTB, OTB, OTB, OTB, OTB, OTB, OTB}
	result1 := [16]Square{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, OTB, OTB, OTB, OTB, OTB, OTB}
	result2 := [16]Square{9, 1, 2, 3, 4, 5, 6, 7, 8, OTB, OTB, OTB, OTB, OTB, OTB, OTB}
	result3 := [16]Square{9, 1, 2, 3, 8, 5, 6, 7, OTB, OTB, OTB, OTB, OTB, OTB, OTB, OTB}

	pl := NewPieceList()
	// Test basic list.
	for i, p := range pl.Pieces {
		if result0[i] != p {
			t.Fatalf("PieceList content should be %d, but is %d", result0[i], p)
		}
	}

	// Add 10 elements to the piece list
	for i := Square(0); i < 10; i++ {
		pl.Add(i)
	}

	// Test if all are there.
	for i := uint8(0); i < pl.Size; i++ {
		if result1[i] != pl.Pieces[i] {
			t.Fatalf("PieceList content should be %d, but is %d", result1[i], pl.Pieces[i])
		}
	}

	// Test the size.
	if pl.Size != 10 {
		t.Fatalf("PieceList should have size 10 but has size %d", pl.Size)
	}

	// Remove the first element by index.
	pl.RemoveIdx(0)

	// Test the size.
	if pl.Size != 9 {
		t.Fatalf("PieceList should have size 9 but has size %d", pl.Size)
	}

	// Test if all remaining are there and correctly ordered.
	for i := uint8(0); i < pl.Size; i++ {
		if result2[i] != pl.Pieces[i] {
			t.Fatalf("PieceList content should be %d, but is %d", result2[i], pl.Pieces[i])
		}
	}

	// Remove '4' by square.
	pl.Remove(4)

	// Test the size.
	if pl.Size != 8 {
		t.Fatalf("PieceList should have size 8 but has size %d", pl.Size)
	}

	// Test if all remaining are there and correctly ordered.
	for i := uint8(0); i < pl.Size; i++ {
		if result3[i] != pl.Pieces[i] {
			t.Fatalf("PieceList content should be %d, but is %d", result3[i], pl.Pieces[i])
		}
	}

	// Clear list
	pl.Clear()

	// Test the size.
	if pl.Size != 0 {
		t.Fatalf("PieceList should have size 0 but has size %d", pl.Size)
	}

	// Test if all remaining are there and correctly ordered.
	for i := uint8(0); i < pl.Size; i++ {
		if result0[i] != pl.Pieces[i] {
			t.Fatalf("PieceList content should be %d, but is %d", result0[i], pl.Pieces[i])
		}
	}

}
