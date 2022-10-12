目指せ！第１４回ＵＥＣ杯コンピューター囲碁大会＜本編＞ Step [O11o_5o0] 石定義

# 連載の目次

📖 [目指せ！第１４回ＵＥＣ杯コンピューター囲碁大会＜本編＞](https://qiita.com/muzudho1/items/744f6051954525878b84)  

# Step [O11o_5o0] 石定義

## Step [O11o0] ファイル作成 - stone.go ファイル

👇 以下のファイルを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 kernel
	│	├── 📄 go.mod
 	│	├── 📄 kernel.go
 	│	├── 📄 logger.go
👉 	│	└── 📄 stone.go
    ├── 📄 .gitignore
 	├── 📄 engine_config.go
  	├── 📄 engine.toml
    ├── 📄 go.mod
  	├── 📄 go.work
 	└── 📄 main.go
```

```go
// BOF [O11o0]

package kernel

import "fmt"

type Stone uint

const (
	Space Stone = iota
	Black
	White
	Wall
)

// GetStoneFromName - 文字列の名前を与えると、Stone値を返します
//
// Returns
// -------
// isOk : bool
// stone : Stone
func GetStoneFromName(stoneName string, getDefaultStone func() (bool, Stone)) (bool, Stone) {
	switch stoneName {
	case "space":
		return true, Space
	case "black":
		return true, Black
	case "white":
		return true, White
	case "wall":
		return true, Wall
	default:
		return getDefaultStone()
	}
}

// GetStoneOrDefaultFromTurn - black または white を与えると、Stone値を返します
//
// Returns
// -------
// stone : Stone
func GetStoneOrDefaultFromTurn(stoneName string, getDefaultStone func() Stone) Stone {
	switch stoneName {
	case "black":
		return Black
	case "white":
		return White
	default:
		return getDefaultStone()
	}
}

// GetStoneFromChar - １文字与えると、Stone値を返します
//
// Returns
// -------
// isOk : bool
// stone : Stone
func GetStoneFromChar(stoneChar string, getDefaultStone func() (bool, Stone)) (bool, Stone) {
	switch stoneChar {
	case ".":
		return true, Space
	case "x":
		return true, Black
	case "o":
		return true, White
	case "+":
		return true, Wall
	default:
		return getDefaultStone()
	}
}

// String - 文字列化
func (s Stone) String() string {
	switch s {
	case Space:
		return "."
	case Black:
		return "x"
	case White:
		return "o"
	case Wall:
		return "+"
	default:
		panic(fmt.Sprintf("%d", int(s)))
	}
}

// GetColor - 色の取得
func (s Stone) GetColor() Color {
	switch s {
	case Space:
		return Color_None
	case Black:
		return Color_Black
	case White:
		return Color_White
	case Wall:
		return Color_None
	default:
		panic(fmt.Sprintf("%d", int(s)))
	}
}

// EOF [O11o0]
```

## Step [O11o1o0] リモートリポジトリにプッシュ

がんばって git などを使い、 `github.com/muzudho/kifuwarabe-uec14/kernel` モジュールの各パッケージのソースを  
リモートリポジトリにプッシュしてほしい  

## Step [O11o2o0] カレントディレクトリーを移動

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

```shell
cd kernel
```

## Step [O11o3o0] tidy

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
go get -u
go mod tidy
```

## Step [O11o4o0] カレントディレクトリーを戻す

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

```shell
cd ..
```

## Step [O11o5o0] tidy

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
go mod tidy
```
