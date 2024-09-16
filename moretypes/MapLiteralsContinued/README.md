### Map literals continued
```go
type Vertex struct {
  Lat, Long float64
}

var m = map[string]Vertex{
  "Bell Labs": {40.68433, -74.39967},
  "Google":    {37.42202, -122.08408},
}

/*
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
*/
```
`Vertex{...}`에서 `Vertex`를 생략해도 된다.<br>

### go run
![image](https://github.com/user-attachments/assets/357f07e1-09ea-4ede-b00c-7948a7f780dc)


### Reference
[Moretypes/Map literals continued](https://go.dev/tour/moretypes/21)<br>
