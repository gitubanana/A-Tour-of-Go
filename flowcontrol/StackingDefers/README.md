### Stacking defers
```go
func main() {
  fmt.Println("counting")

  for i := 0; i < 10; i++ {
    defer fmt.Println(i)
  }

  fmt.Println("done")
}
```
`defer`된 함수는 stack에 푸쉬된다.<br>
함수가 `return`되면, 후입선출로 `defer`된 함수가 실행된다.<br>

### go run
![image](https://github.com/user-attachments/assets/f3c4900f-a002-49ad-963b-412bcfe6d04b)

### defer
```go
func CopyFile(dstName, srcName string) (written int64, err error) {
  src, err := os.Open(srcName)
  if err != nil {
    return
  }
  
  dst, err := os.Create(dstName)
  if err != nil {
    return
  }
  
  written, err = io.Copy(dst, src)
  dst.Close()
  src.Close()
  return
}
```
위 함수는 잘 동작하는 것 같지만, 버그가 있다.<br>
`os.Create`가 실패한다면, 파일을 닫지 않기 때문이다.<br>
`src.close()` 추가해주면 되겠지만, 복잡한 함수일 때는 여러 줄을 써줘야 할 것이다.<br>
이럴 경우 `defer`가 큰 도움이 된다.<br>
```go
func CopyFile(dstName, srcName string) (written int64, err error) {
  src, err := os.Open(srcName)
  if err != nil {
    return
  }
  defer src.Close()
  
  dst, err := os.Create(dstName)
  if err != nil {
    return
  }
  defer dst.Close()
  
  return io.Copy(dst, src)
}
```
`defer`를 통해서, 우리는 `src`파일이 열린다면 항상 `src.Close()`를 함수가 끝나 후에 실행시킬 보장할 수 있다.<br>
(확실히 코드가 깔끔해졌다!)<br>

### defer의 특성
1. `deferred`된 함수의 인자는 바로 계산된다.<br>
2. `deferred`된 함수는 감싼 함수가 끝난 후에 후입선출로 실행된다.<br>
3. `deferred`된 함수는 `named return values`를 읽거나 수정할 수 있다.<br>
  ```go
  func c() (i int) {
    defer func() { i++ }()
    return 1 // i는 1일 된다.
    // return 후에, func이 불려 i++이 되어 리턴값은 2가 된다!
  }
  ```

### Reference
[Flowcontrol/Stacking defers](https://go.dev/tour/flowcontrol/13)<br>
[Defer, panic and recover](https://go.dev/blog/defer-panic-and-recover)<br>
