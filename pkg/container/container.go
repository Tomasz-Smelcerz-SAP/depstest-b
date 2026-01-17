package container

import (
	"github.com/Tomasz-Smelcerz-SAP/depstest-a/pkg/component"
)

type Container struct {
	Elements []component.Component
}

func (c *Container) Visit(visitor func(*component.Component) error) error {
	for i := range c.Elements {
		if err := visitor(&c.Elements[i]); err != nil {
			return err
		}
	}
	return nil
}

type ContainerSet struct {
	Containers []Container
}
