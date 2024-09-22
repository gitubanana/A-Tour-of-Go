### Errors
`Go`는 에러 상태를 `error`값으로 표현한다.<br>
`error`타입은 `fmt.Stringer`와 비슷한 `built-in interface`다.<br>
(`fmt.Stringer`와 비슷하게, `fmt`패키지는 `error`를 출력할 때 `error` `interface`를 찾는다.)<br>
```go
type error interface {
  Error() string
}
```

함수는 주로 `error`값을 리턴한다.<br>
그리고, 그 값이 `nil`과 다르면 에러를 처리해야 한다.<br>
`nil error`는 함수가 성공을, `non-nil error`는 실패를 의미한다.<br>
```go
i, err := strconv.Atoi("42")
if err != nil {
  fmt.Printf("couldn't convert number: %v\n", err)
  return
}
fmt.Println("Converted integer:", i)
```

### go run
![image](https://github.com/user-attachments/assets/b28b299c-b8bf-46e1-9ce4-d8bfa7a642dc)


### Reference
[Methods/Errors](https://go.dev/tour/methods/19)<br>
