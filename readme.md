### go ebiten
1. ```go mod init <项目名称>```
2. ```go get github.com/hajimehoshi/ebiten/v2```
3. ```go run main.go```

- run examples: ```go run -tags=example github.com/hajimehoshi/ebiten/v2/examples/...```
- 打包为exe: ```go build -ldflags="-H windowsgui" -o game.exe```
