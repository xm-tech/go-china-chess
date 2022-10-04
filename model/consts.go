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
