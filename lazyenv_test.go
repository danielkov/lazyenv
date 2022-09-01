package lazyenv_test

import (
	"errors"
	"os"
	"strings"
	"testing"

	"github.com/danielkov/lazyenv"
)

func TestLazyGet_String(t *testing.T) {
	os.Setenv("TEST_STRING", "test")
	defer os.Unsetenv("TEST_STRING")
	value, err := lazyenv.Get("TEST_STRING", lazyenv.Required[string])
	if err != nil {
		t.Error(err)
	}
	if value != "test" {
		t.Errorf("expected test, got %s", value)
	}
}


func TestLazyGet_Int(t *testing.T) {
	os.Setenv("TEST_INT", "1")
	defer os.Unsetenv("TEST_INT")
	value, err := lazyenv.Get("TEST_INT", lazyenv.Required[int], lazyenv.Int)
	if err != nil {
		t.Error(err)
	}
	if value != 1 {
		t.Errorf("expected 1, got %d", value)
	}
}

func TestLazyGet_Bool(t *testing.T) {
	os.Setenv("TEST_BOOL", "true")
	defer os.Unsetenv("TEST_BOOL")
	value, err := lazyenv.Get("TEST_BOOL", lazyenv.Required[bool], lazyenv.Bool)
	if err != nil {
		t.Error(err)
	}
	if value != true {
		t.Errorf("expected true, got %t", value)
	}
}

func TestLazyGet_Float32(t *testing.T) {
	os.Setenv("TEST_FLOAT", "1.1")
	defer os.Unsetenv("TEST_FLOAT")
	value, err := lazyenv.Get("TEST_FLOAT", lazyenv.Required[float32], lazyenv.Float32)
	if err != nil {
		t.Error(err)
	}
	if value != 1.1 {
		t.Errorf("expected 1.1, got %f", value)
	}
}

func TestLazyGet_Float64(t *testing.T) {
	os.Setenv("TEST_FLOAT", "1.1")
	defer os.Unsetenv("TEST_FLOAT")
	value, err := lazyenv.Get("TEST_FLOAT", lazyenv.Required[float64], lazyenv.Float64)
	if err != nil {
		t.Error(err)
	}
	if value != 1.1 {
		t.Errorf("expected 1.1, got %f", value)
	}
}

func TestLazyGet_Int8(t *testing.T) {
	os.Setenv("TEST_INT8", "1")
	defer os.Unsetenv("TEST_INT8")
	value, err := lazyenv.Get("TEST_INT8", lazyenv.Required[int8], lazyenv.Int8)
	if err != nil {
		t.Error(err)
	}
	if value != 1 {
		t.Errorf("expected 1, got %d", value)
	}
}

func TestLazyGet_Int8_Invalid(t *testing.T) {
	os.Setenv("TEST_INT8_INVALID", "zzz")
	defer os.Unsetenv("TEST_INT8_INVALID")
	value, err := lazyenv.Get("TEST_INT8_INVALID", lazyenv.Required[int8], lazyenv.Int8)
	if err == nil {
		t.Error("expected error, got nil")
	}
	if err.Error() != "required variable not found: TEST_INT8_INVALID" {
		t.Errorf("expected error, got %s", err)
	}
	if value != 0 {
		t.Errorf("expected 0, got %d", value)
	}
}

func TestLazyGet_Int16(t *testing.T) {
	os.Setenv("TEST_INT16", "1")
	defer os.Unsetenv("TEST_INT16")
	value, err := lazyenv.Get("TEST_INT16", lazyenv.Required[int16], lazyenv.Int16)
	if err != nil {
		t.Error(err)
	}
	if value != 1 {
		t.Errorf("expected 1, got %d", value)
	}
}

func TestLazyGet_Int16_Invalid(t *testing.T) {
	os.Setenv("TEST_INT16_INVALID", "zzz")
	defer os.Unsetenv("TEST_INT16_INVALID")
	value, err := lazyenv.Get("TEST_INT16_INVALID", lazyenv.Required[int16], lazyenv.Int16)
	if err == nil {
		t.Error("expected error, got nil")
	}
	if err.Error() != "required variable not found: TEST_INT16_INVALID" {
		t.Errorf("expected error, got %s", err)
	}
	if value != 0 {
		t.Errorf("expected 0, got %d", value)
	}
}

func TestLazyGet_Int32(t *testing.T) {
	os.Setenv("TEST_INT32", "1")
	defer os.Unsetenv("TEST_INT32")
	value, err := lazyenv.Get("TEST_INT32", lazyenv.Required[int32], lazyenv.Int32)
	if err != nil {
		t.Error(err)
	}
	if value != 1 {
		t.Errorf("expected 1, got %d", value)
	}
}

func TestLazyGet_Int32_Invalid(t *testing.T) {
	os.Setenv("TEST_INT32_INVALID", "zzz")
	defer os.Unsetenv("TEST_INT32_INVALID")
	value, err := lazyenv.Get("TEST_INT32_INVALID", lazyenv.Required[int32], lazyenv.Int32)
	if err == nil {
		t.Error("expected error, got nil")
	}
	if err.Error() != "required variable not found: TEST_INT32_INVALID" {
		t.Errorf("expected error, got %s", err)
	}
	if value != 0 {
		t.Errorf("expected 0, got %d", value)
	}
}

func TestLazyGet_Int64(t *testing.T) {
	os.Setenv("TEST_INT64", "1")
	defer os.Unsetenv("TEST_INT64")
	value, err := lazyenv.Get("TEST_INT64", lazyenv.Required[int64], lazyenv.Int64)
	if err != nil {
		t.Error(err)
	}
	if value != 1 {
		t.Errorf("expected 1, got %d", value)
	}
}

func TestLazyGet_Int64_Invalid(t *testing.T) {
	os.Setenv("TEST_INT64_INVALID", "zzz")
	defer os.Unsetenv("TEST_INT64_INVALID")
	value, err := lazyenv.Get("TEST_INT64_INVALID", lazyenv.Required[int64], lazyenv.Int64)
	if err == nil {
		t.Error("expected error, got nil")
	}
	if err.Error() != "required variable not found: TEST_INT64_INVALID" {
		t.Errorf("expected error, got %s", err)
	}
	if value != 0 {
		t.Errorf("expected 0, got %d", value)
	}
}

func TestLazyGet_Uint8(t *testing.T) {
	os.Setenv("TEST_UINT8", "1")
	defer os.Unsetenv("TEST_UINT8")
	value, err := lazyenv.Get("TEST_UINT8", lazyenv.Required[uint8], lazyenv.Uint8)
	if err != nil {
		t.Error(err)
	}
	if value != 1 {
		t.Errorf("expected 1, got %d", value)
	}
}

func TestLazyGet_Uint8_Invalid(t *testing.T) {
	os.Setenv("TEST_UINT8_INVALID", "zzz")
	defer os.Unsetenv("TEST_UINT8_INVALID")
	value, err := lazyenv.Get("TEST_UINT8_INVALID", lazyenv.Required[uint8], lazyenv.Uint8)
	if err == nil {
		t.Error("expected error, got nil")
	}
	if err.Error() != "required variable not found: TEST_UINT8_INVALID" {
		t.Errorf("expected error, got %s", err)
	}
	if value != 0 {
		t.Errorf("expected 0, got %d", value)
	}
}

func TestLazyGet_Uint16(t *testing.T) {
	os.Setenv("TEST_UINT16", "1")
	defer os.Unsetenv("TEST_UINT16")
	value, err := lazyenv.Get("TEST_UINT16", lazyenv.Required[uint16], lazyenv.Uint16)
	if err != nil {
		t.Error(err)
	}
	if value != 1 {
		t.Errorf("expected 1, got %d", value)
	}
}

func TestLazyGet_Uint16_Invalid(t *testing.T) {
	os.Setenv("TEST_UINT16_INVALID", "zzz")
	defer os.Unsetenv("TEST_UINT16_INVALID")
	value, err := lazyenv.Get("TEST_UINT16_INVALID", lazyenv.Required[uint16], lazyenv.Uint16)
	if err == nil {
		t.Error("expected error, got nil")
	}
	if err.Error() != "required variable not found: TEST_UINT16_INVALID" {
		t.Errorf("expected error, got %s", err)
	}
	if value != 0 {
		t.Errorf("expected 0, got %d", value)
	}
}

func TestLazyGet_Uint32(t *testing.T) {
	os.Setenv("TEST_UINT32", "1")
	defer os.Unsetenv("TEST_UINT32")
	value, err := lazyenv.Get("TEST_UINT32", lazyenv.Required[uint32], lazyenv.Uint32)
	if err != nil {
		t.Error(err)
	}
	if value != 1 {
		t.Errorf("expected 1, got %d", value)
	}
}

func TestLazyGet_Uint32_Invalid(t *testing.T) {
	os.Setenv("TEST_UINT32_INVALID", "zzz")
	defer os.Unsetenv("TEST_UINT32_INVALID")
	value, err := lazyenv.Get("TEST_UINT32_INVALID", lazyenv.Required[uint32], lazyenv.Uint32)
	if err == nil {
		t.Error("expected error, got nil")
	}
	if err.Error() != "required variable not found: TEST_UINT32_INVALID" {
		t.Errorf("expected error, got %s", err)
	}
	if value != 0 {
		t.Errorf("expected 0, got %d", value)
	}
}

func TestLazyGet_Uint64(t *testing.T) {
	os.Setenv("TEST_UINT64", "1")
	defer os.Unsetenv("TEST_UINT64")
	value, err := lazyenv.Get("TEST_UINT64", lazyenv.Required[uint64], lazyenv.Uint64)
	if err != nil {
		t.Error(err)
	}
	if value != 1 {
		t.Errorf("expected 1, got %d", value)
	}
}

func TestLazyGet_Uint64_Invalid(t *testing.T) {
	os.Setenv("TEST_UINT64_INVALID", "zzz")
	defer os.Unsetenv("TEST_UINT64_INVALID")
	value, err := lazyenv.Get("TEST_UINT64_INVALID", lazyenv.Required[uint64], lazyenv.Uint64)
	if err == nil {
		t.Error("expected error, got nil")
	}
	if err.Error() != "required variable not found: TEST_UINT64_INVALID" {
		t.Errorf("expected error, got %s", err)
	}
	if value != 0 {
		t.Errorf("expected 0, got %d", value)
	}
}
func TestLazyGet_Float32_Invalid(t *testing.T) {
	os.Setenv("TEST_FLOAT32_INVALID", "zzz")
	defer os.Unsetenv("TEST_FLOAT32_INVALID")
	value, err := lazyenv.Get("TEST_FLOAT32_INVALID", lazyenv.Required[float32], lazyenv.Float32)
	if err == nil {
		t.Error("expected error, got nil")
	}
	if err.Error() != "required variable not found: TEST_FLOAT32_INVALID" {
		t.Errorf("expected error, got %s", err)
	}
	if value != 0 {
		t.Errorf("expected 0, got %f", value)
	}
}

func TestLazyGet_Uintptr(t *testing.T) {
	os.Setenv("TEST_UINT", "1")
	defer os.Unsetenv("TEST_UINT")
	value, err := lazyenv.Get("TEST_UINT", lazyenv.Required[uint], lazyenv.Uint)
	if err != nil {
		t.Error(err)
	}
	if value != 1 {
		t.Errorf("expected 1, got %d", value)
	}
}

func TestLazyGet_Uintptr_Invalid(t *testing.T) {
	os.Setenv("TEST_UINT_INVALID", "zzz")
	defer os.Unsetenv("TEST_UINT_INVALID")
	value, err := lazyenv.Get("TEST_UINT_INVALID", lazyenv.Required[uint], lazyenv.Uint)
	if err == nil {
		t.Error("expected error, got nil")
	}
	if err.Error() != "required variable not found: TEST_UINT_INVALID" {
		t.Errorf("expected error, got %s", err)
	}
	if value != 0 {
		t.Errorf("expected 0, got %d", value)
	}
}

func TestLazyGet_InvalidCast(t *testing.T) {
	os.Setenv("TEST_INVALID", "invalid")
	defer os.Unsetenv("TEST_INVALID")
	_, err := lazyenv.Get("TEST_INVALID", lazyenv.Required[int], lazyenv.Int)
	if err == nil {
		t.Error("expected error, got nil")
	}
}

func TestLazyGet_InvalidCast_Optional(t *testing.T) {
	os.Setenv("TEST_INVALID", "invalid")
	defer os.Unsetenv("TEST_INVALID")
	_, err := lazyenv.Get("TEST_INVALID", lazyenv.Optional[int], lazyenv.Int)
	if err != nil {
		t.Error(err)
	}
}

func TestLazyGet_StringSlice(t *testing.T) {
	os.Setenv("TEST_STRING_SLICE", "test1,test2")
	defer os.Unsetenv("TEST_STRING_SLICE")
	value, err := lazyenv.Get("TEST_STRING_SLICE", lazyenv.Required[[]string], lazyenv.SliceOf(",", lazyenv.String))
	if err != nil {
		t.Error(err)
	}
	if value[0] != "test1" || value[1] != "test2" {
		t.Errorf("expected test1,test2, got %s", value)
	}
}

func TestLazyGet_IntSlice(t *testing.T) {
	os.Setenv("TEST_INT_SLICE", "1,2")
	defer os.Unsetenv("TEST_INT_SLICE")
	value, err := lazyenv.Get("TEST_INT_SLICE", lazyenv.Required[[]int], lazyenv.SliceOf(",", lazyenv.Int))
	if err != nil {
		t.Error(err)
	}
	if value[0] != 1 || value[1] != 2 {
		t.Errorf("expected 1,2, got %v", value)
	}
}

func TestLazyGet_BoolSlice(t *testing.T) {
	os.Setenv("TEST_BOOL_SLICE", "true,false")
	defer os.Unsetenv("TEST_BOOL_SLICE")
	value, err := lazyenv.Get("TEST_BOOL_SLICE", lazyenv.Required[[]bool], lazyenv.SliceOf(",", lazyenv.Bool))
	if err != nil {
		t.Error(err)
	}
	if value[0] != true || value[1] != false {
		t.Errorf("expected true,false, got %v", value)
	}
}

func TestLazyGet_JSONOf(t *testing.T) {
	type test struct {
		Test string `json:"test"`
	}
	os.Setenv("TEST_JSON", `{"test": "test"}`)
	defer os.Unsetenv("TEST_JSON")
	value, err := lazyenv.Get("TEST_JSON", lazyenv.Required[test], lazyenv.JSONOf[test])
	if err != nil {
		t.Error(err)
	}
	if value.Test != "test" {
		t.Errorf("expected {Test=test}, got %v", value)
	}
}

func TestLazyGet_JSONOf_OrDefault(t *testing.T) {
	type test struct {
		Test string `json:"test"`
	}

	value, err := lazyenv.Get("TestLazyGet_JSONOf_OrDefault", lazyenv.OrReturn(test{Test: "default"}), lazyenv.JSONOf[test])
	if err != nil {
		t.Error(err)
	}
	if value.Test != "default" {
		t.Errorf("expected {Test=default}, got %v", value)
	}
}

func TestLazyGet_Float32Slice_FromCache(t *testing.T) {
	os.Setenv("TEST_FLOAT_SLICE", "1.1,2.2")
	defer os.Unsetenv("TEST_FLOAT_SLICE")
	value, err := lazyenv.Get("TEST_FLOAT_SLICE", lazyenv.Required[[]float32], lazyenv.SliceOf(",", lazyenv.Float32))
	if err != nil {
		t.Error(err)
	}
	if value[0] != 1.1 || value[1] != 2.2 {
		t.Errorf("expected 1.1,2.2, got %v", value)
	}
	os.Setenv("TEST_FLOAT_SLICE", "3.3,4.4")
	value, err = lazyenv.Get("TEST_FLOAT_SLICE", lazyenv.Required[[]float32], lazyenv.SliceOf(",", lazyenv.Float32))
	if err != nil {
		t.Error(err)
	}
	if value[0] != 1.1 || value[1] != 2.2 {
		t.Errorf("expected 1.1,2.2 from cache, got %v", value)
	}

	lazyenv.Reset()

	value, err = lazyenv.Get("TEST_FLOAT_SLICE", lazyenv.Required[[]float32], lazyenv.SliceOf(",", lazyenv.Float32))
	if err != nil {
		t.Error(err)
	}
	if value[0] != 3.3 || value[1] != 4.4 {
		t.Errorf("expected 3.3,4.4, got %v", value)
	}
}

func TestLazyGet_StringSlice_FromCache(t *testing.T) {
	os.Setenv("TEST_STRING_SLICE_FROM_CACHE", "test1,test2")
	defer os.Unsetenv("TEST_STRING_SLICE_FROM_CACHE")
	value, err := lazyenv.Get("TEST_STRING_SLICE_FROM_CACHE", lazyenv.Required[[]string], lazyenv.SliceOf(",", lazyenv.String))
	if err != nil {
		t.Error(err)
	}
	if value[0] != "test1" || value[1] != "test2" {
		t.Errorf("expected test1,test2, got %s", value)
	}
	os.Setenv("TEST_STRING_SLICE_FROM_CACHE", "test3,test4")
	value, err = lazyenv.Get("TEST_STRING_SLICE_FROM_CACHE", lazyenv.Required[[]string], lazyenv.SliceOf(",", lazyenv.String))
	if err != nil {
		t.Error(err)
	}
	if value[0] != "test1" || value[1] != "test2" {
		t.Errorf("expected test1,test2 from cache, got %s", value)
	}

	lazyenv.Reset()

	value, err = lazyenv.Get("TEST_STRING_SLICE_FROM_CACHE", lazyenv.Required[[]string], lazyenv.SliceOf(",", lazyenv.String))
	if err != nil {
		t.Error(err)
	}
	if value[0] != "test3" || value[1] != "test4" {
		t.Errorf("expected test3,test4, got %s", value)
	}
}

func TestLazyGet_IntSlice_Invalid(t *testing.T) {
	os.Setenv("TEST_INT_SLICE_INVALID", "1,zzz,3")
	defer os.Unsetenv("TEST_INT_SLICE_INVALID")
	_, err := lazyenv.Get("TEST_INT_SLICE_INVALID", lazyenv.Required[[]int], lazyenv.SliceOf(",", lazyenv.Int))
	if err == nil {
		t.Error("expected error, got nil")
	}
	if err.Error() != "required variable not found: TEST_INT_SLICE_INVALID" {
		t.Errorf("expected error, got %v", err)
	}
}

func TestLazyGet_IntSlice_WithoutMapper(t *testing.T) {
	os.Setenv("TEST_INT_SLICE_WITHOUT_MAPPER", "1,2,3")
	defer os.Unsetenv("TEST_INT_SLICE_WITHOUT_MAPPER")
	_, err := lazyenv.Get("TEST_INT_SLICE_WITHOUT_MAPPER", lazyenv.Required[[]int])
	if err == nil {
		t.Error("expected error, got nil")
	}
	if err.Error() != "failed to cast 1,2,3, did you forget to add a mapper?" {
		t.Errorf("expected error, got %v", err)
	}
}

func _MapOf[K comparable, V any](entrySeparator, valueSeparator string, keyMapper lazyenv.Mapper[K], valueMapper lazyenv.Mapper[V]) lazyenv.Mapper[map[K]V] {
	return func(value string) (map[K]V, error) {
		values := strings.Split(value, entrySeparator)
		results := make(map[K]V, len(values))
		for _, v := range values {
			kv := strings.Split(v, valueSeparator)
			if len(kv) != 2 {
				return nil, errors.New("invalid map value: " + v)
			}
			k, err := keyMapper(kv[0])
			if err != nil {
				return nil, err
			}
			v, err := valueMapper(kv[1])
			if err != nil {
				return nil, err
			}
			results[k] = v
		}
		return results, nil
	}
}

func TestLazyGet_CustomMapper(t *testing.T) {
	os.Setenv("TEST_MAP", "key1=value1,key2=value2")
	defer os.Unsetenv("TEST_MAP")
	value, err := lazyenv.Get("TEST_MAP", lazyenv.Required[map[string]string], _MapOf(",", "=", lazyenv.String, lazyenv.String))
	if err != nil {
		t.Error(err)
	}
	if value["key1"] != "value1" || value["key2"] != "value2" {
		t.Errorf("expected key1=value1,key2=value2, got %v", value)
	}
}
