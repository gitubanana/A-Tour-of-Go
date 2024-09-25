### Select
`select`문은 여러 통신에 대해서 `goroutine`이 기다릴 수 있게 한다.<br>
`select`문은 `case`중 하나가 실행할 수 있을 때까지 `block`되고, 그리고 해당 `case`를 실행한다.<br>
여러 개가 준비되었다면, 하나를 랜덤하게 골라 실행한다.<br>

### go run
![image](https://github.com/user-attachments/assets/a454a164-2b59-4fbe-9003-d565bb85da3f)


### Reference
[Concurrency/Select](https://go.dev/tour/concurrency/5)<br>
