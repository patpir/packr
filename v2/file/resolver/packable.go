package resolver

import "github.com/patpir/packr/v2/file"

type Packable interface {
	Pack(name string, f file.File) error
}
