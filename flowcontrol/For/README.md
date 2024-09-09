### For
```go
sum := 0
for i := 0; i < 10; i++ {
  sum += i
}
```
`Go`에서는 반복문이 `for` 하나밖에 없다. (while이 없어서 놀랐다.)<br>
`C`의 반복문처럼 3개의 구문이 있다. 초기화; 반복 조건; 증감문 <br>
초기화로 생겨진 변수는 `for`문 안에서만 사용할 수 있다.<br>
> [!Note]
> 다른 언어와 다르게, `Go`의 `for`문에는 `()`이 없다. (처음 보는 for문의 생김새에 놀랐다.) <br>
> 또한, `{}`는 항상 있어야 한다. (C/C++에서는 for문 안이 한 줄일 때는 생략 가능했는데...)<br>

### go run
![image](https://github.com/user-attachments/assets/f9b5cb6a-fa73-4cca-98ee-9b9e78337584)

### Reference
[Flowcontrol/For](https://go.dev/tour/flowcontrol/1)<br>
