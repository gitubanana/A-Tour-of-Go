### Switch
```go
linux := "linux"
switch os := runtime.GOOS; os {
case "darwin":
  fmt.Println("OS X.")
case linux:
  fmt.Println("Linux.")
default:
  // freebsd, openbsd
  // plan9, windowns......
  fmt.Printf("%s.\n", os)
}
```
다른 언어들과 다르게, `Go`의 `switch`문은 오로지 비교값과 같은 부분만 실행한다.<br>
(다른 언어들은 같으면, 아랫부분까지 함께 실행되어 `break`를 써줘야 한다. 실은, `Go`에서는 자동으로 `break`가 들어간다.)<br>
또한, `Go`에서는 상수 외에도 다양한 변수와 비교할 수 있다.<br>

### go run
![image](https://github.com/user-attachments/assets/2cb93aea-0ba9-4fb4-b1e3-6a9cec5c8b98)


### Reference
[Flowcontrol/Switch](https://go.dev/tour/flowcontrol/9)<br>
