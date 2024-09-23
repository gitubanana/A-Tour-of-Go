### Images
`image`패키지에는 `Image` `interface`가 정의되어 있다.<br>
```go
package image

type Image interface {
  ColorModel() color.Model
  Bounds() Rectange // image.Rectangle
  At(x, y int) color.Color
}
```

### go run
![image](https://github.com/user-attachments/assets/7ba4d5d5-fd12-4c97-b6d2-045c48feb47d)



### Reference
[Methods/Images](https://go.dev/tour/methods/24)<br>
