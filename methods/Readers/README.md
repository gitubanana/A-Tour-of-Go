### Readers
`io`패키지에는 데이터 스트림을 읽는 `io.Reader` `interface`가 있다.<br>
`Go standard library`에는 이 `interface`를 구현한 많은 것들이 있다.<br>
(파일, 네트워크, 압축, 암호 등...)<br>

`io.Reader` `interface`는 `Read` 메소드가 있다.<br>
`Read` 메소드는 주어진 슬라이스에 데이터를 덧붙이고, 덧붙인 바이트 수와 에러값을 리턴한다.<br>
스트림이 끝났을 때는 `io.EOF`에러를 리턴한다.<br>
```go
func (T) Read(b []byte) (n int, err error)
```

### go run
![image](https://github.com/user-attachments/assets/da2b212f-43b1-4df8-b494-ffbfe95430e6)


### Reference
[Methods/Readers](https://go.dev/tour/methods/21)<br>
