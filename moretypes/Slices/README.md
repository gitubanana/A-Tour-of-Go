### Slices
`array`은 고정된 크기를 가지고 있다.<br>
반대로, `slice`는 크기가 바뀔 수 있다.<br>

`[]T`는 원소 타입이 `T`인 `slice`이다.<br>
```go
var s []int = a[1:4]
```

`slice`는 두 개의 인덱스로 만들어질 수 있다.<br>
`[low, high)` 범위의 원소를 `slice`로 만든다. (`low`부터 `high` 전까지, `high`는 제외)<br>
```go
a[low : high]
```

### go run
![image](https://github.com/user-attachments/assets/6afc6400-da88-4aad-af5c-06ebf494536e)


### Reference
[Moretypes/Slices](https://go.dev/tour/moretypes/7)<br>
