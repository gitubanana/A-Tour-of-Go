### Creating a slice with make
`slice`는 built-in 함수인 `make`로 만들어질 수 있다.<br>
(이 방법을 통해서 크기가 바뀌는 배열을 만들 수 있다.)<br>

`make`함수는 `zero value`로 된 배열을 할당하고, 그 배열을 묘사하는 `slice`를 리턴한다.<br>
```go
a := make([]int, 5) // len(a)=5

// capacity를 표시하려면, 3번째 인자를 주면 된다.
b := make([]int, 0, 5) // len(b)=0, cap(b)=5

b = b[:cap(b)] // len(b)=5, cap(b)=5
b = b[1:] // len(b)=4, cap(b)=4
```

### go run
![image](https://github.com/user-attachments/assets/9c90f4e0-c125-4830-86c1-af04c4c05fc1)


### Reference
[Moretypes/Creating a slice with make](https://go.dev/tour/moretypes/13)<br>
