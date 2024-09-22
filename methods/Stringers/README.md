### Stringers
아주 흔한 `interface` 중 하나는 `fmt`패키지에 정의되어 있는 `Stringer`이다.<br>
```go
type Stringer interface {
  String() string
}
```
`Stringer`는 그 자신을 문자열로 표현할 수 있는 타입이다.<br>
`fmt`패키지는 값을 출력하기 위해 `Stringer interface`를 기대한다.<br>
(`source code`를 보면 이해가 될 것이다.)<br>

### go run
![image](https://github.com/user-attachments/assets/433f794b-c3aa-4740-bf09-301ae6353511)



### Reference
[Methods/Stringers](https://go.dev/tour/methods/17)<br>
