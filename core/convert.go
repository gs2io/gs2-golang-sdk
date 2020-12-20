package core

import (
	"fmt"
)

func ToString(data interface{}) string {
	{
		v, ok := data.(String)
		if ok {
			return string(v)
		}
	}
	{
		v, ok := data.(string)
		if ok {
			return v
		}
	}
	{
		v, ok := data.(int32)
		if ok {
			return fmt.Sprintf("%d", v)
		}
	}
	{
		v, ok := data.(int64)
		if ok {
			return fmt.Sprintf("%d", v)
		}
	}
	{
		v, ok := data.(float32)
		if ok {
			return fmt.Sprintf("%f", v)
		}
	}
	{
		v, ok := data.(float64)
		if ok {
			return fmt.Sprintf("%f", v)
		}
	}
	{
		v, ok := data.(bool)
		if ok {
			if v {
				return "true"
			} else {
				return "false"
			}
		}
	}
	return "null"
}

func ToStringPointer(data string) *string {
	return &data
}

func ToInt32Pointer(data int32) *int32 {
	return &data
}

func ToInt64Pointer(data int64) *int64 {
	return &data
}

func ToFloat32Pointer(data float32) *float32 {
	return &data
}

func ToFloat64Pointer(data float64) *float64 {
	return &data
}

func ToBoolPointer(data bool) *bool {
	return &data
}
