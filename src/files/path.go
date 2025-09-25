package files

import (
	"os"
	"strings"
)

const Separator = string(os.PathSeparator)

type Path struct {
	DriveLetter string
	Root        string
	Name        string
	Extension   string
	Separator   string
}

func (p Path) String() string {
	return p.Root + p.Separator + p.Root + p.Separator + p.Name + p.Extension
}

type PathParts []string

func (p PathParts) String() string {
	return strings.Join(p, Separator)
}

func PathPartsFromString(s string) PathParts {
	array := strings.Split(s, Separator)
	newLength := 0
	for _, part := range array {
		if part != "" {
			newLength++
		}
	}
	out := make([]string, newLength)
	i := 0
	for _, part := range array {
		if part != "" {
			out[i] = part
			i++
		}
	}
	return out
}

func (p PathParts) Root() string {
	return p[0]
}

func (p PathParts) stem() PathParts {
	if len(p) == 0 {
		return PathParts{}
	}
	return p[1 : len(p)-1]
}
