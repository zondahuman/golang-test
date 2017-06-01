package enums

import "fmt"

type Type int

//const (
//	One Type = 1 << iota // 1
//	Two Type
//	Three Type
//)
const (
       One = 1 << iota // 1 (i.e. 1 << 0)
       Two // 2 (i.e. 1 << 1)
       Three // 4 (i.e 1 << 2)
    )

func (t Type) String() string {
	s := ""
	if t & One == One {
		s += "One"
	}

	return s
}

func main() {
	fmt.Println(One)
}