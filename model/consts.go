package model

// 窗口
const (
	WindowWidth  = 520
	WindowHeight = 576
	ScreenWidth  = 520
	ScreenHeight = 576

	GridSize       = 56
	BoardEdgeWidth = 8
	BoardWidth     = BoardEdgeWidth + GridSize*9 + BoardEdgeWidth
	BoardHeight    = BoardEdgeWidth + GridSize*10 + BoardEdgeWidth
)

// 棋盘范围
const (
	Top    = 3
	Bottom = 12
	Left   = 3
	Right  = 11
)

// 棋子动作
const (
	// 选子
	SelectChess = 100
	// 落子
	MoveChess = 101
	// 吃子
	EatChess = 102
)
