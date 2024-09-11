### Switch with no condition
```go
t := time.Now()
switch {
case t.Hour() < 12:
  fmt.Println("Good morning!")
case t.Hour() < 17:
  fmt.Println("Good afternoon.")
default:
  fmt.Println("Good evening.")
}
```
비교할 값이 없는 `switch`문은 `switch true`랑 똑같다.<br>
이 덕분에, 긴 `if-else if-else`보다 깔끔하게 코드를 짤 수 있다.<br>

### go run
![image](https://github.com/user-attachments/assets/63258205-c6f1-4a21-930c-1e310583c132)


### Reference
[Flowcontrol/Switch with no condition](https://go.dev/tour/flowcontrol/11)<br>
