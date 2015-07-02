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
