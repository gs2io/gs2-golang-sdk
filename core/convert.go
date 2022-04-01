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

func CastString(value interface{}) *string {
	{
		v, ok := value.(string)
		if ok {
			return &v
		}
	}
	{
		v, ok := value.(*string)
		if ok {
			return v
		}
	}
	return nil
}

func CastInt32(value interface{}) *int32 {
	{
		v, ok := value.(int32)
		if ok {
			return &v
		}
	}
	{
		v, ok := value.(*int32)
		if ok {
			return v
		}
	}
	return nil
}

func CastInt64(value interface{}) *int64 {
	{
		v, ok := value.(int64)
		if ok {
			return &v
		}
	}
	{
		v, ok := value.(*int64)
		if ok {
			return v
		}
	}
	return nil
}

func CastFloat32(value interface{}) *float32 {
	{
		v, ok := value.(float32)
		if ok {
			return &v
		}
	}
	{
		v, ok := value.(*float32)
		if ok {
			return v
		}
	}
	return nil
}

func CastFloat64(value interface{}) *float64 {
	{
		v, ok := value.(float64)
		if ok {
			return &v
		}
	}
	{
		v, ok := value.(*float64)
		if ok {
			return v
		}
	}
	return nil
}

func CastBool(value interface{}) *bool {
	{
		v, ok := value.(bool)
		if ok {
			return &v
		}
	}
	{
		v, ok := value.(*bool)
		if ok {
			return v
		}
	}
	return nil
}

func CastUserId(value interface{}) *string {
	{
		v, ok := value.(string)
		if ok {
			return &v
		}
	}
	{
		v, ok := value.(*string)
		if ok {
			return v
		}
	}
	return nil
}

func CastMap(value interface{}) map[string]interface{} {
	v, ok := value.(map[string]interface{})
	if ok {
		return v
	}
	return nil
}

func CastArray(value interface{}) []interface{} {
	v, ok := value.([]interface{})
	if ok {
		return v
	}
	return nil
}

func CastStrings(value []interface{}) []string {
	v := make([]string, 0)
	for _, d := range value {
		v = append(v, *CastString(d))
	}
	return v
}

func CastStringsFromDict(value []string) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range value {
		v = append(v, d)
	}
	return v
}

func Cast2Strings(value []interface{}) [][]string {
	v := make([][]string, 0)
	for _, d := range value {
		v = append(v, CastStrings(CastArray(d)))
	}
	return v
}

func Cast2StringsFromDict(value [][]string) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range value {
		v = append(v, d)
	}
	return v
}

func CastInt32s(value []interface{}) []int32 {
	v := make([]int32, 0)
	for _, d := range value {
		v = append(v, *CastInt32(d))
	}
	return v
}

func CastInt32sFromDict(value []int32) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range value {
		v = append(v, d)
	}
	return v
}

func CastInt64s(value []interface{}) []int64 {
	v := make([]int64, 0)
	for _, d := range value {
		v = append(v, *CastInt64(d))
	}
	return v
}

func CastInt64sFromDict(value []int64) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range value {
		v = append(v, d)
	}
	return v
}

func CastFloat32s(value []interface{}) []float32 {
	v := make([]float32, 0)
	for _, d := range value {
		v = append(v, *CastFloat32(d))
	}
	return v
}

func CastFloat32sFromDict(value []float32) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range value {
		v = append(v, d)
	}
	return v
}

func CastFloat64s(value []interface{}) []float64 {
	v := make([]float64, 0)
	for _, d := range value {
		v = append(v, *CastFloat64(d))
	}
	return v
}

func CastFloat64sFromDict(value []float64) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range value {
		v = append(v, d)
	}
	return v
}

func CastBools(value []interface{}) []bool {
	v := make([]bool, 0)
	for _, d := range value {
		v = append(v, *CastBool(d))
	}
	return v
}

func CastBoolsFromDict(value []bool) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range value {
		v = append(v, d)
	}
	return v
}

func CastUserIds(value []interface{}) []string {
	v := make([]string, 0)
	for _, d := range value {
		v = append(v, *CastUserId(d))
	}
	return v
}

func CastUserIdsFromDict(value []string) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range value {
		v = append(v, d)
	}
	return v
}
