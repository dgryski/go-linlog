// Package linlog implements linear-log bucketing
/*

http://pvk.ca/Blog/2015/06/27/linear-log-bucketing-fast-versatile-simple/

*/
package linlog

func BinOf(size uint64, linear, subbin uint64) (uint64, uint64) {
	nBits := lb(size | (1 << uint(linear)))
	shift := nBits - subbin
	mask := uint64(1<<shift - 1)
	rounded := size + mask /* XXX: overflow. */
	subIndex := rounded >> shift
	xrange := nBits - linear

	return (rounded & ^mask), (xrange << subbin) + subIndex
}

func BinDownOf(size uint64, linear, subbin uint64) (uint64, uint64) {
	nBits := lb(size | (1 << linear))
	shift := nBits - subbin
	subIndex := size >> shift
	xrange := nBits - linear

	return (subIndex << shift), (xrange << subbin) + subIndex
}

func lb(x uint64) uint64 { return 64 - nlz(x) - 1 }

func nlz(x uint64) uint64 {
	var n uint64

	n = 1

	if (x >> 32) == 0 {
		n = n + 32
		x = x << 32
	}
	if (x >> (32 + 16)) == 0 {
		n = n + 16
		x = x << 16
	}

	if (x >> (32 + 16 + 8)) == 0 {
		n = n + 8
		x = x << 8
	}

	if (x >> (32 + 16 + 8 + 4)) == 0 {
		n = n + 4
		x = x << 4
	}

	if (x >> (32 + 16 + 8 + 4 + 2)) == 0 {
		n = n + 2
		x = x << 2
	}

	n = n - (x >> 63)
	return uint64(n)
}
