### Goroutines
`goroutine`은 `Go runtime`에 의해 관리되는 경량의 쓰레드다.<br>
```go
go f(x, y, z)
```

`f(x, y, z)`를 실행하는 새로운 `goroutine`을 시작시킨다.<br>
`f`, `x`, `y`, `z`에 대한 연산은 현재 `goroutine`에서 이뤄지고, `f`의 실행이 새로운 `goroutine`에서 이뤄진다.<br>
`goroutine`들은 같은 주소 공간에서 실행되기 때문에, 공유 메모리에 대한 동기화가 필요하다.<br>
`sync`패키지는 유용한 기본 요소를 제공한다.<br>
`Go`에는 다른 기본 요소가 있기 때문에, 그렇게 필요하진 않답니다!<br>

### go run
![image](https://github.com/user-attachments/assets/91969375-32dd-49ea-8d15-9cc07fd53b61)


### Reference
[Concurrency/Goroutines](https://go.dev/tour/concurrency/1)<br>
