# Go `lazyenv`

## Installation

```bash
go get -u github.com/danielkov/lazyenv
```

## Usage

```go
package main

import (
  "fmt"
  "github.com/danielkov/lazyenv"
)

func main() {
  fmt.Println(lazyenv.Get("FOO"), lazyenv.Required[string])
}
```

Resetting the cache:

```go
lazyenv.Reset()
```

Implementing a custom mapper:

```go
func MapOf[K comparable, V any](keyMapper lazyenv.Mapper[K], valueSeparator string, valueMapper lazyenv.Mapper[V], entrySeparator string) lazyenv.Mapper[map[K]V] {
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

func main() {
	value, err := lazyenv.Get("TEST_MAP", lazyenv.Required[map[string]string], MapOf(lazyenv.String, "=", lazyenv.String, ","))
}
```

## Explanation

If you want to read my journal of how and why I've created this library, [here's a link to my blog post on Dev.to](https://dev.to/danielkov/taking-go-generics-for-a-spin-29l4).

## Contributing

PR's are welcome! Please see the [contributing guide](/CONTRIBUTING.md) for more information.

## License

MIT
