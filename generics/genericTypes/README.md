### Generic types
`generic function`에 더해서, `Go`는 `generic type`도 제공한다.<br>
`type`에 `type parameter`를 인자로 넣어 `generic data structure`를 만들 때 유용하게 사용할 수 있다.<br>
```go
type List[T any] struct {
	next *List[T]
	val  T
}
```

### go run
![image](https://github.com/user-attachments/assets/91184dd2-ccdc-4ca8-adb5-7772da70fa73)


### Reference
[Generics/Generic types](https://go.dev/tour/generics/2)<br>
