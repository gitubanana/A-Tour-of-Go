### Slices are like references to arrays
```go
names := [4]string{
  "John",
  "Paul",
  "George",
  "Ringo",
}

a := names[0:2]
b := names[1:3]

b[0] = "XXX"
fmt.Println(a, b) // [John XXX] [XXX George]
fmt.Println(names) // [John XXX George Ringo]
```
`slice`는 어떠한 데이터도 저장하지 않는다.<br>
그냥 배열의 부분을 묘사하고 있는 것이다.<br>

따라서, `slice`의 원소를 변경하는 것은 원래 배열의 원소를 변경하는 것과 똑같다.<br>
그 배열의 묘사하고 있는 다른 `slice`도 당연히 그 변화를 보게 될 것이다.<br>

### go run
![image](https://github.com/user-attachments/assets/dacaa63e-6dcc-493e-9c2b-5904c37827f5)


### Reference
[Moretypes/Slices are like references to arrays](https://go.dev/tour/moretypes/8)<br>
