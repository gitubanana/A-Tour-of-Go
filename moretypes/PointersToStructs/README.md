### Pointers to structs
```go
type Vertex struct {
  X int
  Y int
}

func main() {
  v := Vertex{1, 2}
  p := &v
  p.X = 1e9
  fmt.Println(v)
}
```
`struct`의 필드는 `struct pointer`를 통해서 접근될 수 있다.<br>
`struct pointer`를 통해 접근할 때 `(*p).X`로 접근할 수 있겠지만, 이는 길고 복잡하다.<br>

따라서, `Go`에서는 그냥 `p.X`로 접근을 허용해줬다!<br>
(`C`에서는 `p->X`였는데, 흠...)

### go run
![image](https://github.com/user-attachments/assets/fb2a3765-076a-4d60-82a0-3dec04922d9b)


### Reference
[Moretypes/Pointers to structs](https://go.dev/tour/moretypes/4)<br>
