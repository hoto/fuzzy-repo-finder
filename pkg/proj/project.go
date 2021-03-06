package proj

import (
	"fmt"
)

type Project struct {
	FullPath string
	Group    string
	Name     string
}

func (p Project) String() string {
	return fmt.Sprintf("Name=[%s], Group=[%s], FullPath=[%s]", p.Name, p.Group, p.FullPath)
}
