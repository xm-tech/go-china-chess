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

// 棋子ID
const (
	RedJu0 = iota
	RedMa0
	RedXiang0
	RedShi0
	RedShuai
	RedShi1
	RedXiang1
	RedMa1
	RedJu1
	RedPao0
	RedPao1
	RedBing0
	RedBing1
	RedBing2
	RedBing3
	RedBing4

	BlackJu0
	BlackMa0
	BlackXiang0
	BlackShi0
	BlackJiang
	BlackShi1
	BlackXiang1
	BlackMa1
	BlackJu1
	BlackPao0
	BlackPao1
	BlackZu0
	BlackZu1
	BlackZu2
	BlackZu3
	BlackZu4
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
