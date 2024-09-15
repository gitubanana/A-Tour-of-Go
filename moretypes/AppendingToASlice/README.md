### Appending to a slice
```go
func append(s []T, vs ...T) []T
```
built-in 함수인 `append`를 통해 `slice`에 새로운 원소를 끝에 추가할 수 있다.<br>
`s`는 원소를 추가할 `slice`이고, 나머지는 `s`에 추가될 원소들이다.<br>
`append`의 리턴값은 기존의 원소 + 추가 원소를 가리키고 있는 `slice`이다.<br>

`s`가 가리키고 있는 배열의 크기가 원소를 추가하기에 작다면, 더 큰 배열을 새로 할당하고 원소를 추가한 후, 그걸 가리키는 `silce`를 리턴한다.<br>

### go run
![image](https://github.com/user-attachments/assets/0d2105f7-5727-4d27-8bb9-7c1d19939c6b)


### Reference
[Moretypes/Appending to a slice](https://go.dev/tour/moretypes/15)<br>
[Slices: usage and interals](https://go.dev/blog/slices-intro)<br>
