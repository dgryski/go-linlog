// Package linlog implements linear-log bucketing
/*

http://pvk.ca/Blog/2015/06/27/linear-log-bucketing-fast-versatile-simple/

*/
package linlog

// BinOf rounds size as appropriate, and returns the rounded size and bucket number.
func BinOf(size uint64, linear, subbin uint64) (rounded uint64, bucket uint64) {
	nBits := lb(size | (1 << uint(linear)))
	shift := nBits - subbin
	mask := uint64(1<<shift - 1)
	round := size + mask /* XXX: overflow. */
	subIndex := round >> shift
	xrange := nBits - linear

	return (round & ^mask), (xrange << subbin) + subIndex
}

// BinDownOf rounds size down, and returns the rounded size and bucket number.
func BinDownOf(size uint64, linear, subbin uint64) (rounded uint64, bucket uint64) {
	nBits := lb(size | (1 << linear))
	shift := nBits - subbin
	subIndex := size >> shift
	xrange := nBits - linear

	return (subIndex << shift), (xrange << subbin) + subIndex
}
