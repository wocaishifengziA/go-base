// 选项模式
package option

import "time"

type Connection struct {
	addr    string
	cache   bool
	timeout time.Duration
}

const (
	defaultTimeout = 10
	defaultCaching = false
)

type options struct {
	Cache   bool
	timeOut time.Duration
}

type Option interface {
	apply(*options)
}

type FuncOption func(*options)

func (f FuncOption) apply(o *options) {
	f(o)
}

func WithTimeout(t time.Duration) Option {
	return FuncOption(func(o *options) {
		o.timeOut = t
	})
}

func WithCaching(cache bool) Option {
	return FuncOption(func(o *options) {
		o.Cache = cache
	})
}

func NewConnection(addr string, op ...Option) *Connection {
	options := options{
		Cache:   defaultCaching,
		timeOut: defaultTimeout,
	}
	for _, o := range op {
		o.apply(&options)
	}
	return &Connection{
		addr:    addr,
		cache:   options.Cache,
		timeout: options.timeOut,
	}
}
