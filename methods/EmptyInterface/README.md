### The empty interface
어떠한 `method`도 없는 `interface type`을 `empty interface`라고 부른다.<br>
```go
interface{}
```

`empty interface`는 어떠한 타입의 값이든 가지고 있을 수 있다.<br>
(모든 타입은 최소 0개의 `method`를 구현하기 때문이다.)<br>

`empty interface`는 `type`이 알려지지 않는 코드에서 쓰인다.<br>
예를 들면, `fmt.Print`은 `interface{}` 타입의 인자를 받아 출력한다.<br>
즉, `interface{}`로 받기 때문에, 어떠한 타입의 값이든 출력할 수 있다.<br>

### go run
![image](https://github.com/user-attachments/assets/f88553f9-fd81-46c4-ab88-acd3432bb2a1)


### Reference
[Methods/The empty interface](https://go.dev/tour/methods/14)<br>
