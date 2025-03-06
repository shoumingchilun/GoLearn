package weightconv

import "fmt"

type Kilogram float64

func (k Kilogram) String() string {
	return fmt.Sprintf("%gkg", k)
}

type Pound float64

func (p Pound) String() string {
	return fmt.Sprintf("%glb", p)
}
