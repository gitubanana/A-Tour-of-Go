### Map literals
```go
type Vertex struct {
  Lat, Long float64
}

var m = map[string]Vertex{
  "Bell Labs": Vertex{
    40.68433, -74.39967,
  },
  "Google": Vertex{
    37.42202, -122.08408,
  },
}
```
`map literal`은 `struct literal`과 비슷하지만, `key`는 항상 써야된다.<br>

### go run
![image](https://github.com/user-attachments/assets/e7531ab9-0598-4d13-aa98-408c32cce3ff)


### Reference
[Moretypes/Map literals](https://go.dev/tour/moretypes/20)<br>
