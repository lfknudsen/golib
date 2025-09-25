package structs

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
}
