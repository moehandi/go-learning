package main

import "fmt"

// the billionaire have house
type House struct {
	name string
	address string
	size float32 // house size
	price float32 // house price
}

func (h *House) Value() float32 {
	return h.price * 100
}

// the billionaire have Golds
type Gold struct {
	goldType string
	carat int
	weight float32
}

func (g *Gold) Value() float32 {
	return g.weight * 5
}

type AssetValuable interface {
	Value() float32
}

func ShowAssetsValue(asset AssetValuable) {
	fmt.Println("Asset value is:", asset.Value())
}
func main()  {
	var o AssetValuable = &House{"Smart House","Bali", 230, 120}
	ShowAssetsValue(o)
	o = &Gold{"PUre Gold", 24, 120}
	ShowAssetsValue(o)
}