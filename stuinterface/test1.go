package stuinterface

import "fmt"

type I struct {
	Name string
}

func (i1 *I) String() string {
	return fmt.Sprintf("%s", (*i1).Name)
}


