package providers

import "fmt"

type Container struct {
	services    map[string]func() (any, error)
	initialised map[string]any
}

type ContainerError struct {
	message string
}

func (ce *ContainerError) Error() string {
	return ce.message
}

func (c *Container) add(key string, callable func() (any, error)) {
	c.services[key] = callable
}

func (c *Container) get(key string) (any, error) {
	service := c.initialised[key]

	if service != nil {
		return service, nil
	}

	serviceProvider := c.services[key]

	if serviceProvider == nil {
		return nil, &ContainerError{message: fmt.Sprintf("provider %s not found", key)}
	}

	result, err := serviceProvider()

	if result != nil {
		c.initialised[key] = result
	}

	return result, err
}
