package chess

const (
	ScreenWidth     = 440  // 窗口宽度, 棋盘像素宽度
	ScreenHeight    = 440  // 窗口高度, 棋盘像素高度
	Edge            = 10   // 边缘宽度/高度
	IsBoardVertical = true // 竖放棋盘true; 横放棋盘false
)

const (
	SquareSize  = 60 // 棋子边长
	Gap         = 60 // 棋子间距
	ColumnPos   = 4  // 横向点位 x
	RowPos      = 4  // 纵向点位 y
	ReversalNum = 3  // player反转所需异或的值
	WhiteValue  = 1  // 白子的值
	BlackValue  = 2  // 黑子的值
	EmptyValue  = 0  // 空位的值
)

// 棋盘位置初始值y, x
var PieceStartPos = [4][4]int{
	{2, 2, 2, 2},
	{0, 0, 0, 0},
	{0, 0, 0, 0},
	{1, 1, 1, 1},
}
