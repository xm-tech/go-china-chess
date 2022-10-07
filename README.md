# 先做出核心的东西，快速原型

## 棋盘摆放

```go
// 棋子
type Chessman struct {
  x int
  y int
  img string
  camp int
}

var chesses []*Chessman

```

