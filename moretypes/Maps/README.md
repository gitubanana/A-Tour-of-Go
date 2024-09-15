### Maps
```go
type Vertex struct {
  Lat, Long float64
}

var m map[string]Vertex

func main() {
  m = make(map[string]Vertex)
  m["Bell Labs"] = Vertex{
    40.68433, -74.39967,
  }
  fmt.Println(m["Bell Labs"])
}
```

`map`은 `key`를 `value`와 연결해 저장한다.<br>

`map`의 `zero value`는 `nil`이다.<br>
`nil map`은 `key`가 없고, `key`를 추가할 수도 없다.<br>

`make`함수는 주어진 타입의 `map`을 반환한다.<br>

### go run
![image](https://github.com/user-attachments/assets/e76d6b85-7c95-422a-b46b-0d5db33e8a68)


### Reference
[Moretypes/Maps](https://go.dev/tour/moretypes/19)<br>
