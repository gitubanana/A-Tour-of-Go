### Pointers
```go
var p *int
```
`Go`는 포인터가 있다.<br>
`*T`는 타입 `T`의 포인터이다. 포인터의 `zero value`는 `nil`이다.<br>
```go
i := 42
p = &i
```
`&` 연산자는 피연산자의 메모리 주솟값을 가져온다.<br>
```go
fmt.Println(*p)
*p = 21
```
`*` 연산자는 포인터가 가리키고 있는 값을 가져온다.<br>
(`dereferencing` 혹은 `indirecting`으로 알려져 있다.)<br>
> [!NOTE]
> 여기까지만 보면, `C`와 똑같이 보이지만<br>
> `C`와 다르게, `Go`에는 포인터 연산이 없다. (ㅠㅠ)<br>
> 즉, `p++`을 할 수 없다.<br>

### go run
![image](https://github.com/user-attachments/assets/94b70068-f8d6-42e8-98f7-bfa2f2c78e97)

### Reference
[Moretypes/Pointers](https://go.dev/tour/moretypes/1)<br>
