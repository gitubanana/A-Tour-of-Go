### Type assertions
`type assertion`을 통해 `interface`값이 묘사하고 있는 실제 값을 접근할 수 있다.<br>
아래 구문은 `interface i`의  `concrete type`이 `T`인지 확인하고, 타입 `T` 값을 `t`에 대입한다.<br>
만약, `i`의 `concrete type`이 `T`가 아니라면, `panic`을 일으킨다.<br>
```go
t := i.(T)
```

`interface`의 값이 어떤 타입을 가지고 있는지 테스트해보고 싶다면, `type assertion`으로부터 두 개의 값을 받아오면 된다.<br>
실제 값과 성공 여부를 리턴한다.<br>
`i`의 `concrete type`이 `T`라면, `실제 값`과 `true`를 반환한다.<br>
아니라면, 타입 `T`의 `zero value`와 `false`를 반환한다. 그리고, `panic`은 일어나지 않는다.<br>
```go
t, ok := i.(T) //  map사용할 때와 똑같은 구조
```

### go run
![image](https://github.com/user-attachments/assets/2fc3913a-3261-492f-afb1-166005321591)



### Reference
[Methods/Type assertions](https://go.dev/tour/methods/15)<br>
