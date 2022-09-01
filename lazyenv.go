package lazyenv

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

type cache struct {
	sync.Mutex
	values map[string]string
}

type GetDefaultValueParams struct {
	Key string
}

type GetDefaultValue[T any] func(params GetDefaultValueParams) (T, error)
type Mapper[T any] func(value string) (T, error)

var cacheInstance *cache

func castAs[T any](v any) (T, error) {
	cast, ok := v.(T)
	if !ok {
		return cast, fmt.Errorf("failed to cast %v, did you forget to add a mapper?", v)
	}
	return cast, nil
}

func (c *cache) get(key string) (string, bool) {
	c.Lock()
	defer c.Unlock()
	value, exists := c.values[key]
	return value, exists
}

func (c *cache) set(key string, value string) {
	c.Lock()
	defer c.Unlock()
	c.values[key] = value
}

// Reset clears the cache so that each call to Get will fetch the value from the environment
func Reset () {
	cacheInstance.Lock()
	defer cacheInstance.Unlock()
	cacheInstance.values = make(map[string]string)
}

func getEnv(key string) (string, bool) {
	if cacheInstance == nil {
		cacheInstance = &cache{values: make(map[string]string)}
	}
	value, exists := cacheInstance.get(key)
	if !exists {
		value, exists = os.LookupEnv(key)
		if !exists {
			return "", false
		}
		cacheInstance.set(key, value)
	}
	return value, true
}

// Get returns the value of the environment variable with the given key
// if the variable is not set, it returns the value returned by getDefaultValue parameter
// optionalMapper is a variadic parameter that can be used to map the value to a different type
// if the first element of the variadic parameter is provided, it will be considered, otherwise the value
// will be returned as a string
func Get[T any](key string, getDefaultValue GetDefaultValue[T], optionalMapper ...Mapper[T]) (T, error) {
	value, exists := getEnv(key)
	if !exists {
		return getDefaultValue(GetDefaultValueParams{
			key,
		})
	}
	if len(optionalMapper) > 0 {
		val, err := optionalMapper[0](value)
		if err != nil {
			return getDefaultValue(GetDefaultValueParams{
				key,
			})
		}
		return val, nil
	}
	return castAs[T](value)
}

// MustGet returns the value or panics if it's not available
func MustGet[T any](key string, optionalMapper ...Mapper[T]) T {
	// it's ok to ignore err, OrPanic stops Get from returning it
	value, _ := Get(key, OrPanic[T], optionalMapper...)
	return value
}

// Required is a default value getter that returns the value of the environment variable if it is set, otherwise it returns an error
func Required[T any](params GetDefaultValueParams) (T, error) {
	var v T
	return v, errors.New("required variable not found: " + params.Key)
}

// Optional is a default value getter that returns the value of the environment variable if it is set, otherwise it returns the value the type initialises to and no error
func Optional[T any](key GetDefaultValueParams) (T, error) {
	var v T
	return v, nil
}

// OrReturn is a default value getter that returns the value of the environment variable if it is set, otherwise it returns the value provided
func OrReturn[T any](defaultValue T) func (key GetDefaultValueParams) (T, error) {
	return func (key GetDefaultValueParams) (T, error) {
		return defaultValue, nil
	}
}

func OrPanic[T any](params GetDefaultValueParams) (T, error) {
	panic(errors.New("required variable not found: " + params.Key))
}

// String is a mapper that returns the value of the variable as a string
func String(value string) (string, error) {
	return value, nil
}

// Int is a mapper that returns the value of the variable as a int
func Int(value string) (int, error) {
	// convert value to int
	return strconv.Atoi(value)
}

// Int8 is a mapper that returns the value of the variable as a int8
func Int8(value string) (int8, error) {
	i64, err := strconv.ParseInt(value, 10, 8)
	if err != nil {
		return 0, err
	}
	return int8(i64), nil
}

// Int16 is a mapper that returns the value of the variable as a int16
func Int16(value string) (int16, error) {
	i64, err := strconv.ParseInt(value, 10, 16)
	if err != nil {
		return 0, err
	}
	return int16(i64), nil
}

// Int32 is a mapper that returns the value of the variable as a int32
func Int32(value string) (int32, error) {
	i64, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(i64), nil
}

// Int64 is a mapper that returns the value of the variable as a int64
func Int64(value string) (int64, error) {
	return strconv.ParseInt(value, 10, 64)
}

// Uint is a mapper that returns the value of the variable as a uint
func Uint(value string) (uint, error) {
	u64, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(u64), nil
}

// Uint8 is a mapper that returns the value of the variable as a uint8
func Uint8(value string) (uint8, error) {
	u64, err := strconv.ParseUint(value, 10, 8)
	if err != nil {
		return 0, err
	}
	return uint8(u64), nil
}

// Uint16 is a mapper that returns the value of the variable as a Uint16
func Uint16(value string) (uint16, error) {
	u64, err := strconv.ParseUint(value, 10, 16)
	if err != nil {
		return 0, err
	}
	return uint16(u64), nil
}

// Uint32 is a mapper that returns the value of the variable as a uint32
func Uint32(value string) (uint32, error) {
	u64, err := strconv.ParseUint(value, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint32(u64), nil
}

// Uint64 is a mapper that returns the value of the variable as a uint64
func Uint64(value string) (uint64, error) {
	return strconv.ParseUint(value, 10, 64)
}

// Float32 is a mapper that returns the value of the variable as a float32
func Float32(value string) (float32, error) {
	f64, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return 0, err
	}
	return float32(f64), nil
}

// Float64 is a mapper that returns the value of the variable as a float64
func Float64(value string) (float64, error) {
	return strconv.ParseFloat(value, 64)
}

// Bool is a mapper that converts 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False. Any other value returns an error.
func Bool(value string) (bool, error) {
	return strconv.ParseBool(value)
}

// SliceOf returns a mapper that will convert the value to a slice of the given type
func SliceOf[T any](separator string, mapper Mapper[T]) Mapper[[]T] {
	return func(value string) ([]T, error) {
		values := strings.Split(value, separator)
		results := make([]T, len(values))
		for i, v := range values {
			result, err := mapper(v)
			if err != nil {
				return nil, err
			}
			results[i] = result
		}
		return results, nil
	}
}

// JSONOf is a mapper that converts the value of the variable to a JSON object of a given type
func JSONOf[T any](value string)  (T, error) {
		var results T
		err := json.Unmarshal([]byte(value), &results)
		return results, err
}
