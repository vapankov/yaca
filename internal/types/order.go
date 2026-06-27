package types

type Order uint8

const (
	OrderNA Order = iota
	OrderAsc
	OrderDesc
)
