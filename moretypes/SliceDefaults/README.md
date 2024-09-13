### Slice defaults
```go
var a [10]int

// 다음 slice는 모두 똑같다.
a[0:10]
a[:10]
a[0:]
a[:]
```
`slicing`을 할 때, `high` 혹은 `low`를 생략해도 된다.<br>
생략하게 되면, `high`는 `slice의 크기`가 되고, `low`는 `0`이 된다.<br>

### go run
![image](https://github.com/user-attachments/assets/4db7c0cc-f988-4ebf-afcc-7995b2213433)


### Reference
[Moretypes/Slice defaults](https://go.dev/tour/moretypes/10)<br>
