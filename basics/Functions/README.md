### Functions
```go
func add(x int, y int) int {
  return x + y
}
```
위 예시에서 `add`함수는 두 개의 `int` 매개변수를 받는다.<br>
`Go`에서는 type이 변수명 뒤에 온다.<br>

그 이유는 type이 복잡해지더라도 가독성이 좋다고 생각하기 때문이다.<br>
(개인적으로, Go가 좀 더 가독성이 좋은 것 같다...)<br>
- **C에서 복잡한 함수 포인터**
  ```c++
  int (*(*fp)(int (*)(int, int), int))(int, int)
  ```
- **Go에서 복잡한 함수 포인터**
  ```go
  func fp(func(int, int) int, int) func (int, int) int
  ```

### go run
![image](https://github.com/user-attachments/assets/147bbb1c-7491-456a-b683-5592779745fe)

### Reference
[Basics/Functions](https://go.dev/tour/basics/4)<br>
[article on Go's declaration syntax](https://go.dev/blog/declaration-syntax)<br>
