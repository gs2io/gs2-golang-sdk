package core

import (
	"fmt"
)

func ToString(data interface{}) string {
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
