module bloom

go 1.20

require (
	github.com/bits-and-blooms/bloom/v3 v3.3.1
	github.com/strider2038/otus-algo v0.0.0-00010101000000-000000000000
)

require github.com/bits-and-blooms/bitset v1.6.0 // indirect

replace github.com/strider2038/otus-algo => ./..
