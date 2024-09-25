### Range and close
`sender`는 더 이상 값을 보내지 않는다는 것을 알려주기 위해 `channel`을 `close`할 수 있다.<br>
`receiver`는 `<-`연산자의 두 번째 인자를 통해 `channel`이 닫혔는지 확인할 수 있다.<br>
아래 코드에서, `ok`는 더 이상 `channel`에서 받을 값이 없고, `channel`이 닫혔으면 `false`가 된다.
```go
v, ok := <-ch
```

`for i := range c`는 `channel` `c`가 닫힐 때까지 계속 값을 받아온다.<br>
> [!IMPORTANT]
> `sender`만 `channel`를 닫아야 한다. (`receiver`는 닫으면 안 된다.)<br>
> 닫힌 `channel`에 값을 보내면 `panic`이 호출된다.<br>

> [!NOTE]
> `channel`은 파일과 다르기 때문에, 주로 닫을 필요가 없다.<br>
> `close`는 `range`루프를 종료시키는 것과 같이 `receiver`에게 더 이상 값이 없는 것을 알려줄 때만 필요하다.<br>

### go run
![image](https://github.com/user-attachments/assets/908679bf-4f95-4bec-a8f2-bf0093e3edb1)


### Reference
[Concurrency/Range and close](https://go.dev/tour/concurrency/4)<br>
