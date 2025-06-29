package chess

import (
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Board struct {
	images         map[int]*ebiten.Image // 图片
	piecePos       [4][4]int
	selectedPointX int // 选中格子的X坐标
	selectedPointY int // 选中格子的Y坐标
	player         int // 当前玩家，白子1，黑子2
}

func NewBoard() (*Board, error) {
	b := &Board{
		images:         make(map[int]*ebiten.Image),
		piecePos:       PieceStartPos,
		selectedPointX: -1,
		selectedPointY: -1,
		player:         1,
	}
	b.LoadPieceImg()
	return b, nil
}

// Draw 绘制棋盘和棋子
func (b *Board) Draw(screen *ebiten.Image) {
	// 棋盘
	img, _, err := ebitenutil.NewImageFromFile("image/board.png")
	if err != nil {
		log.Fatal(err)
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	screen.DrawImage(img, op)

	// 棋子
	for y := range RowPos {
		for x := range ColumnPos {
			xPos, yPos := 0, 0
			if IsBoardVertical {
				xPos = Edge + x*(SquareSize+Gap)
				yPos = Edge + y*(SquareSize+Gap)
			}
			p := b.piecePos[y][x]
			if p != 0 {
				b.DrawPiece(xPos, yPos, screen, b.images[p])
			}
			if x == b.selectedPointX && y == b.selectedPointY {
				const scale = 1.1
				b.DrawPieceScale(float64(xPos)/scale, float64(yPos)/scale, scale, scale, screen, b.images[p])
			}
		}
	}
}

// LoadPieceImg 加载图片
func (b *Board) LoadPieceImg() {
	img1, _, err := ebitenutil.NewImageFromFile("image/whitePiece.png")
	if err != nil {
		log.Fatal(err)
	}
	b.images[1] = img1

	img2, _, err := ebitenutil.NewImageFromFile("image/blackPiece.png")
	if err != nil {
		log.Fatal(err)
	}
	b.images[2] = img2
}

// DrawPiece 绘制棋子
func (b *Board) DrawPiece(x, y int, screen, img *ebiten.Image) {
	if img == nil {
		return
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(img, op)
}

// DrawPieceScale 点击棋子放大
func (b *Board) DrawPieceScale(x, y, scaleX, scaleY float64, screen, img *ebiten.Image) {
	if img == nil {
		return
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)
	op.GeoM.Scale(scaleX, scaleY)
	screen.DrawImage(img, op)
}

// LegalMove 判断当前走法是否合理
func (b *Board) LegalMove(startX, startY, endX, endY int) bool {
	// end没子
	if b.GetPiece(endX, endY) == 0 {
		if startX == endX && math.Abs(float64(startY-endY)) == 1 {
			return true
		} else if startY == endY && math.Abs(float64(startX-endX)) == 1 {
			return true
		}
	}
	return false
}

// Move 执行走法
func (b *Board) Move(startX, startY, endX, endY int) {
	startPieceV := b.GetPiece(startX, startY)
	b.DelPiece(startX, startY)
	b.AddPiece(endX, endY, startPieceV)
}

// Eat 吃子
func (b *Board) Eat(xPos, yPos int) {
	b.CheckRowCanEat(xPos, yPos, b.player)
	b.CheckColumnCanEat(xPos, yPos, b.player)
}

// CheckRowCanEat 检查行是否能吃子
func (b *Board) CheckRowCanEat(xPos, yPos int, player int) {
	xSlice := b.piecePos[yPos][0:]
	playerPieceNum, otherPlayerPieceNum := 0, 0
	x, otherX := 0, 0 // 记录被吃的点的x
LOOP:
	for k, v := range xSlice {
		// k是坐标x，v是值
		switch v {
		case player:
			x = k
			playerPieceNum++
		case player ^ ReversalNum:
			otherX = k
			otherPlayerPieceNum++
		default:
			if k == 1 || k == 2 {
				// 中间有空位，不会有吃子的可能清除计数
				playerPieceNum, otherPlayerPieceNum = 0, 0
				break LOOP
			}
			continue LOOP
		}
	}
	if playerPieceNum >= 2 && otherPlayerPieceNum == 1 {
		b.DelPiece(otherX, yPos)
	} else if otherPlayerPieceNum >= 2 && playerPieceNum == 1 {
		b.DelPiece(x, yPos)
	}
}

// CheckColumnCanEat 检查列是否能吃子
func (b *Board) CheckColumnCanEat(xPos, yPos int, player int) {
	ySlice := make([]int, 4)
	for yk, yv := range b.piecePos {
		ySlice[yk] = yv[xPos]
	}
	playerPieceNum, otherPlayerPieceNum := 0, 0
	y, otherY := 0, 0 // 记录被吃的点的y
LOOP:
	for k, v := range ySlice {
		switch v {
		case player:
			y = k
			playerPieceNum++
		case player ^ ReversalNum:
			otherY = k
			otherPlayerPieceNum++
		default:
			if k == 1 || k == 2 {
				// 中间有空位，不会有吃子的可能清除计数
				playerPieceNum, otherPlayerPieceNum = 0, 0
				break LOOP
			}
			continue LOOP
		}
	}
	if playerPieceNum >= 2 && otherPlayerPieceNum == 1 {
		b.DelPiece(xPos, otherY)
	} else if otherPlayerPieceNum >= 2 && playerPieceNum == 1 {
		b.DelPiece(xPos, y)
	}
}

// IsOver 游戏是否结束
func (b *Board) IsOver() bool {
	whiteNum, blackNum := 0, 0
	for y := range b.piecePos {
		for x := range b.piecePos[y] {
			if b.GetPiece(x, y) == 1 {
				whiteNum++
			} else if b.GetPiece(x, y) == 2 {
				blackNum++
			}
		}
	}
	if whiteNum == 1 || blackNum == 1 {
		return true
	}
	return false
}

// ChangePlayer 交换玩家
func (b *Board) ChangePlayer() {
	b.player ^= ReversalNum
}

// GetPiece 获取棋盘上点位的值
func (b *Board) GetPiece(x, y int) int {
	return b.piecePos[y][x]
}

// AddPiece 在棋盘上放置一颗棋子
func (b *Board) AddPiece(x, y, v int) {
	b.piecePos[y][x] = v
}

// DelPiece 在棋盘上移除一颗棋子
func (b *Board) DelPiece(x, y int) {
	b.piecePos[y][x] = 0
}
