# go-jsonstrconv
## Whats's this?
This library converts all values of json to string type.

### Original Json
```json
{"id":12345,"name":"John","tel":12345}
```

### Converted
```json
{"id":"12345","name":"John","tel":"12345"}
```

## How to use
```go
	import "github.com/rung/go-jsonstrconv"
```
```go
	msg := "{\"id\":12345,\"name\":\"John\",\"tel\":12345}"
	converted, err := jsonstrconv.ToString([]byte(msg))
```
Please see [example](example/convert/main.go)
