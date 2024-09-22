### Type switches
`type switch`는 여러 `type assertion`을 연속해서 할 수 있는 구문이다.<br>
그냥 `switch`문과 비슷하지만, `type switch`의 `case`는 `interface`의 `type`과 `case`의 타입을 비교해 `type`을 특정할 수 있다.<br>
```go
switch v := i.(type) { // v는 i의 concrete 타입에 해당하는 값을 가지게 된다.
case T:
  // here v has type T
case S:
  // here v has type S
default:
  // no match; here v has the same type as i
}
```

### go run
![image](https://github.com/user-attachments/assets/6948c9ac-0f24-4d33-ab02-484ed6e66826)



### Reference
[Methods/Type switches](https://go.dev/tour/methods/16)<br>
