### Named return values
`Go`에서는 리턴값의 이름이 정해질 수 있다. 이 값을 `named return values`라고 부른다.<br>
`named return values`는 함수의 맨 위에 선언된 변수처럼 취급된다.<br>
(이름은 가독성을 위해, 리턴값의 의미로 지정돼야 한다.)<br>

리턴값이 없는 `return`은 `named return values`을 리턴한다. 이러한 `return`을 `naked return`이라고 부른다.<br>
(`naked return`은 긴 함수에서는 가독성을 해칠 수 있기 때문에, 짧은 함수에서만 써야한다.)<br>
```go
func split(sum int) (x, y int) {
  x = sum * 4 / 9
  y = sum - x
  return // 실험해본 결과, `named return values`가 있다면 `return`이 있어야 build가 된다.
}
```

### go run
![image](https://github.com/user-attachments/assets/8da920e7-138d-4899-aa4a-fc534b0a23e9)


### Reference
[Basics/Named return values](https://go.dev/tour/basics/7)<br>
[The benefit of using named return values in Golang](https://medium.com/@xcoulon/the-benefit-of-using-named-return-values-in-golang-7a946bdc0894)<br>
