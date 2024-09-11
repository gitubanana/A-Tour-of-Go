### Switch evaluation order
```go
switch i {
case 0:
case f():
}
```
`switch`문은 위부터 아래로 비교를 한다.<br>
따라서, 위 예시에서 `i == 0`일 경우에는 `f()`가 안 불린다.<br>

### go run
![image](https://github.com/user-attachments/assets/df1af7fa-c951-4d00-bae0-bb11ab8695fe)


### Reference
[Flowcontrol/Switch evaluation order](https://go.dev/tour/flowcontrol/10)<br>
