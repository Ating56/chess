package chess

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	board    *Board // 棋局
	text     string // 文字展示
	gameOver bool   // 游戏结束标志
}

func NewGame() (*Game, error) {
	g := &Game{}
	var err error
	g.board, err = NewBoard()
	if err != nil {
		return nil, err
	}
	return g, nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		xPos := (x - Edge) / SquareSize / 2
		yPos := (y - Edge) / SquareSize / 2
		fmt.Printf("定位到坐标%d, %d\n", xPos, yPos)
		g.ClickSquare(xPos, yPos)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.board.Draw(screen)
}

// ClickSquare 点击棋盘格子
func (g *Game) ClickSquare(xPos, yPos int) {
	p := g.board.piecePos[yPos][xPos]
	if p == g.board.player {
		// 轮到一方走并点击己方的棋子
		g.board.selectedPointX = xPos
		g.board.selectedPointY = yPos
	} else if g.board.selectedPointX != -1 && g.board.selectedPointY != -1 && !g.gameOver {
		// 轮到一方走，已经选中过棋子，点到另一个位置
		if g.board.LegalMove(g.board.selectedPointX, g.board.selectedPointY, xPos, yPos) {
			g.board.Move(g.board.selectedPointX, g.board.selectedPointY, xPos, yPos)
			g.board.Eat(xPos, yPos, g.board.player)
			// 清除选中的格子
			g.board.selectedPointX = -1
			g.board.selectedPointY = -1
			if g.board.IsOver() {
				// 是否游戏结束
				if g.board.player == 1 {
					g.text = "黑子胜"
				} else {
					g.text = "白子胜"
				}
				g.gameOver = true
			}
		}
	}
}
