package structs

import "fmt"

type IOption interface {
	Exists() bool
	Get() any
	Set(v any)
}

type Option struct {
	val    any
	exists bool
}

func (o *Option) Exists() bool {
	return o.exists
}

func (o *Option) Get() any {
	return o.val
}

func (o *Option) Set(v any) {
	o.val = v
	o.exists = v != nil
}

func (o *Option) String() string {
	return fmt.Sprint(o.val)
}

func (o *Option) If(f func(val any)) *Option {
	if o.exists {
		f(o.val)
	}
	return o
}

func (o *Option) IfNot(f func(val any)) *Option {
	if !o.exists {
		f(o.val)
	}
	return o
}

func (o *Option) Finally(f func(val any)) *Option {
	f(o.val)
	return o
}
