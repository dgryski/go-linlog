package linlog

import "testing"

func TestBinOf(t *testing.T) {

	tests := []struct {
		sz  uint64
		l   uint64
		b   uint64
		wb  uint64
		wsz uint64
	}{
		{
			0, 4, 2,
			0, 0,
		},
		{
			1, 4, 2,
			1, 4,
		},
		{
			4, 4, 2,
			1, 4,
		},
		{
			5, 4, 2,
			2, 8,
		},
		{
			9, 4, 2,
			3, 12,
		},
		{
			15, 4, 2,
			4, 16,
		},
		{
			17, 4, 2,
			5, 20,
		},
		{
			34, 4, 2,
			9, 40,
		},
	}

	for _, tt := range tests {
		if r, b := BinOf(tt.sz, tt.l, tt.b); r != tt.wsz || b != tt.wb {
			t.Errorf("BinOf(%d,%d,%d)=(%d,%d), want (%d,%d)", tt.sz, tt.l, tt.b, r, b, tt.wsz, tt.wb)
		}
	}
}

func TestBinDownOf(t *testing.T) {

	tests := []struct {
		sz  uint64
		l   uint64
		b   uint64
		wb  uint64
		wsz uint64
	}{
		{
			0, 4, 2,
			0, 0,
		},
		{
			1, 4, 2,
			0, 0,
		},
		{
			3, 4, 2,
			0, 0,
		},
		{
			4, 4, 2,
			1, 4,
		},
		{
			7, 4, 2,
			1, 4,
		},
		{
			15, 4, 2,
			3, 12,
		},
		{
			16, 4, 2,
			4, 16,
		},
		{
			17, 4, 2,
			4, 16,
		},
		{
			34, 4, 2,
			8, 32,
		},
	}

	for _, tt := range tests {
		if r, b := BinDownOf(tt.sz, tt.l, tt.b); r != tt.wsz || b != tt.wb {
			t.Errorf("BinDownOf(%d,%d,%d)=(%d,%d), want (%d,%d)", tt.sz, tt.l, tt.b, r, b, tt.wsz, tt.wb)
		}
	}
}
