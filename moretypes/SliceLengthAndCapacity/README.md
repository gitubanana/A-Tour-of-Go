### Slice length and capacity
```go
s := []int{2, 3, 5, 7, 11, 13}

fmt.Printf("len=%d cap=%d\n", len(s), cap(s)) // len과 cap으로 구할 수 있다.

s = s[:10] // panic: runtime error: slice bounds out of range [:10] with capacity 6
```
`slice`는 `length`와 `capacity`를 모두 가지고 있다.<br>
`length`는 `slice`에 있는 원소의 개수이다.<br>
`capacity`는 `slice`가 묘사하는 배열의 원소 개수이다. (`slice`의 첫 원소부터 센다.)<br>

`slice`의 길이는 다시 `slicing`을 통해 키울 수 있다.<br>
만약, `capacity`를 넘게 길이를 키운다면 런타임 에러를 맞이하게 된다.<br>

### go run
![image](https://github.com/user-attachments/assets/0cbbe158-8a67-4d41-9789-26b065854a60)


### Reference
[Moretypes/Slice length and capacity](https://go.dev/tour/moretypes/11)<br>
