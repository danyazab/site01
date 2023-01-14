package internal

import "go.uber.org/dig"

func Invoke(fn, cfg interface{}) error {
	c, err := container(cfg)
	if err != nil {
		return err
	}
	return c.Invoke(fn)
}

func container(cfg interface{}) (*dig.Container, error) {
	c := dig.New()
	if err := c.Provide(func() *dig.Container { return c }); err != nil {
		return nil, err
	}
	if err := c.Provide(cfg); err != nil {
		return nil, err
	}

	return c, nil
}
