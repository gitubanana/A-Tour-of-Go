### Interface values with nil underlying values
`interface`의 값이 `nil`이라면, `method`는 `nil receiver`와 함께 불린다.<br>

다른 언어에서는 `null pointer exception`으로 처리하겠지만, `Go`에서는 `nil receiver`와 `method`가 불리는 건 흔한 일이다.<br>

### go run
![image](https://github.com/user-attachments/assets/879cfb23-937e-4dff-a43a-7cd50b978c0b)


### Reference
[Methods/Interface values with nil underlying values](https://go.dev/tour/methods/12)<br>
