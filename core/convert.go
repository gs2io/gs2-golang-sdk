package core

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
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
	{
		v, ok := value.(map[string]interface{})
		if ok {
			v2, _ := json.Marshal(v)
			v3 := string(v2)
			return &v3
		}
	}
	return nil
}

func CastInt32(value interface{}) *int32 {
	if v, ok := value.(*int32); ok {
		return CastInt32(*v)
	}
	if v, ok := value.(int32); ok {
		return &v
	}
	if v, ok := value.(*int64); ok {
		return CastInt32(*v)
	}
	if v, ok := value.(int64); ok {
		v2 := int32(v)
		return &v2
	}
	if v, ok := value.(*float32); ok {
		return CastInt32(*v)
	}
	if v, ok := value.(float32); ok {
		v2 := int32(v)
		return &v2
	}
	if v, ok := value.(*float64); ok {
		return CastInt32(*v)
	}
	if v, ok := value.(float64); ok {
		v2 := int32(v)
		return &v2
	}
	if v, ok := value.(*string); ok {
		return CastInt32(*v)
	}
	if v, ok := value.(string); ok {
		v2, err := strconv.ParseInt(v, 10, 32)
		if err != nil {
			return nil
		}
		return CastInt32(v2)
	}
	v, err := strconv.ParseInt(fmt.Sprintf("%v", value), 10, 32)
	if err != nil {
		return nil
	}
	v2 := int32(v)
	return &v2
}

func CastInt64(value interface{}) *int64 {
	if v, ok := value.(*int32); ok {
		return CastInt64(*v)
	}
	if v, ok := value.(int32); ok {
		v2 := int64(v)
		return &v2
	}
	if v, ok := value.(*int64); ok {
		return CastInt64(*v)
	}
	if v, ok := value.(int64); ok {
		return &v
	}
	if v, ok := value.(*float32); ok {
		return CastInt64(*v)
	}
	if v, ok := value.(float32); ok {
		v2 := int64(v)
		return &v2
	}
	if v, ok := value.(*float64); ok {
		return CastInt64(*v)
	}
	if v, ok := value.(float64); ok {
		v2 := int64(v)
		return &v2
	}
	if v, ok := value.(*string); ok {
		if len(*v) == 25 && (*v)[10] == uint8('T') {
			if t, err := time.Parse(time.RFC3339Nano, *v); err == nil {
				v2 := t.UTC().UnixNano() / int64(time.Millisecond)
				return &v2
			}
		}
		return CastInt64(*v)
	}
	if v, ok := value.(string); ok {
		if len(v) == 25 && v[10] == uint8('T') {
			if t, err := time.Parse(time.RFC3339Nano, v); err == nil {
				v2 := t.UTC().UnixNano() / int64(time.Millisecond)
				return &v2
			}
		}
		v2, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil
		}
		return CastInt64(v2)
	}
	v, err := strconv.ParseInt(fmt.Sprintf("%v", value), 10, 64)
	if err != nil {
		return nil
	}
	return &v
}

func CastFloat32(value interface{}) *float32 {
	if v, ok := value.(*int32); ok {
		return CastFloat32(*v)
	}
	if v, ok := value.(int32); ok {
		v2 := float32(v)
		return &v2
	}
	if v, ok := value.(*int64); ok {
		return CastFloat32(*v)
	}
	if v, ok := value.(int64); ok {
		v2 := float32(v)
		return &v2
	}
	if v, ok := value.(*float32); ok {
		return CastFloat32(*v)
	}
	if v, ok := value.(float32); ok {
		return &v
	}
	if v, ok := value.(*float64); ok {
		return CastFloat32(*v)
	}
	if v, ok := value.(float64); ok {
		v2 := float32(v)
		return &v2
	}
	if v, ok := value.(*string); ok {
		return CastFloat32(*v)
	}
	if v, ok := value.(string); ok {
		v2, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil
		}
		return CastFloat32(v2)
	}
	v, err := strconv.ParseFloat(fmt.Sprintf("%v", value), 32)
	if err != nil {
		return nil
	}
	v2 := float32(v)
	return &v2
}

func CastFloat64(value interface{}) *float64 {
	if v, ok := value.(*int32); ok {
		return CastFloat64(*v)
	}
	if v, ok := value.(int32); ok {
		v2 := float64(v)
		return &v2
	}
	if v, ok := value.(*int64); ok {
		return CastFloat64(*v)
	}
	if v, ok := value.(int64); ok {
		v2 := float64(v)
		return &v2
	}
	if v, ok := value.(*float32); ok {
		return CastFloat64(*v)
	}
	if v, ok := value.(float32); ok {
		v2 := float64(v)
		return &v2
	}
	if v, ok := value.(*float64); ok {
		return CastFloat64(*v)
	}
	if v, ok := value.(float64); ok {
		return &v
	}
	if v, ok := value.(*string); ok {
		return CastFloat64(*v)
	}
	if v, ok := value.(string); ok {
		v2, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil
		}
		return CastFloat64(v2)
	}
	v, err := strconv.ParseFloat(fmt.Sprintf("%v", value), 64)
	if err != nil {
		return nil
	}
	return &v
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
