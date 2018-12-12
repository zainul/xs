package redis

import "github.com/zainul/xs/internal/pkg/counter"

type counterImpl struct {
}

// NewCounter ...
func NewCounter(config string) counter.Counter {
	return &counterImpl{}
}

func (c *counterImpl) Increament(key string) (int64, error) {
	return 0, nil
}
func (c *counterImpl) Decreament(key string) (int64, error) {
	return 0, nil
}
func (c *counterImpl) Get(key string) (int64, error) {
	return 0, nil
}
