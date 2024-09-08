### Type inference
변수의 타입이 명시적으로 정해지지 않은 경우, 오른쪽 값에서 변수의 타입을 추론한다.<br>
```go
var i int
j := i // j is an int

i := 42           // int
f := 3.142        // float64
g := 0.867 + 0.5i // complex128
```

### go run
![image](https://github.com/user-attachments/assets/af7e19c5-ec81-44b3-8ab8-7e2a95fa4d4c)

### Reference
[Basics/Type inference](https://go.dev/tour/basics/14)<br>
