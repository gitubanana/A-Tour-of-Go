### Range continued
```go
for i, _ := range pow // skip the value
```
```go
for _, value := range pow // skip the index
```
인덱스 혹은 값은 `_`를 통해, 사용 안 할 수 있다.<br>

```go
for i := range pow
```
인덱스만 원한다면, 두 번째 값은 안 써도 된다.<br>

### go run
![image](https://github.com/user-attachments/assets/69c8b7d3-face-4363-8066-422dc05bdaba)


### Reference
[Moretypes/Range continued](https://go.dev/tour/moretypes/17)<br>
