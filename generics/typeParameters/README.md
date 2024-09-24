### Type parameters
`Go`함수에서는 `type parameter`를 이용해 여러 타입에 쓰일 수 있다.<br>
`type parameter`는 `[]` 사이에 쓰고, 함수 인자 전에 위치한다.<br>
```go
func Index[T comparable](s []T, x T) int
```

`s`는 `built-in constraint` `comparable`에 해당하는 타입 `T`로 이루어진 `slice`이다.<br>
`x`도 타입 `T` 값이다.<br>

`comparable`은 `==`과 `!=` 연산자를 이용해 값을 비교할 수 있는 유용한 `constraint`이다. 
```go
// Index returns the index of x in s, or -1 if not found.
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v and x are type T, which has the comparable
		// constraint, so we can use == here.
		if v == x {
			return i
		}
	}
	return -1
}

func main() {
	// Index works on a slice of ints
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	// Index also works on a slice of strings
	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello"))
}
```

### go run
![image](https://github.com/user-attachments/assets/58b0ee42-4722-4742-bc56-1ca497c1b975)


### Reference
[Generics/Type parameters](https://go.dev/tour/generics/1)<br>
