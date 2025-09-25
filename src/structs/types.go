package structs

type (
	Uint8  uint8
	IUint8 interface {
		Uint8() Uint8
	}
	Uint16  uint16
	IUint16 interface {
		Uint16() Uint16
	}
	Uint32  uint32
	IUint32 interface {
		Uint32() Uint32
	}
	Uint64  uint64
	IUint64 interface {
		Uint64() Uint64
	}

	Int8  int8
	IInt8 interface {
		Int8() Int8
	}
	Int16  int16
	IInt16 interface {
		Int16() Int16
	}
	Int32  int32
	IInt32 interface {
		Int32() Int32
	}
	Int64  int64
	IInt64 interface {
		Int64() Int64
	}
	Float32  float32
	IFloat32 interface {
		Float32() Float32
	}
	Float64  float64
	IFloat64 interface {
		Float64() Float64
	}
	Complex64  complex64
	IComplex64 interface {
		Complex64() Complex64
	}
	Complex128  complex128
	IComplex128 interface {
		Complex128() Complex128
	}
	Byte  byte
	IByte interface {
		Byte() Byte
	}

	IBool interface {
		Bool() Bool
	}

	Uint  uint
	IUint interface {
		Uint() Uint
	}
)
