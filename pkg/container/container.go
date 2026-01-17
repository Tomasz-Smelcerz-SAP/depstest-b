package container

import (
	"github.com/Tomasz-Smelcerz-SAP/depstest-a/pkg/component"
)

type Container struct {
	elements []component.Component
}

type ContainerSet struct {
	containers []Container
}
