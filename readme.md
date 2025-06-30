### A chess mini game

A 3x3 grid with four pieces from both sides, with two pieces from our side placed side by side, can eat the opponent's piece placed side by side with these two pieces.
Each chess piece can only take one step up, down, left, and right.
If one side only has one chess piece left, they lose the game.

For example, in this case, white characters can eat black spots:
![image](/image/example1.png)
But this situation cannot:
![image](/image/example2.png)

Some commands:
- run examples: ```go run -tags=example github.com/hajimehoshi/ebiten/v2/examples/...```
- exe: ```go build -ldflags="-H windowsgui" -o game.exe```
