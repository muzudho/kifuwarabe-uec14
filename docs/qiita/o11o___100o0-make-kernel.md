目指せ！第１４回ＵＥＣ杯コンピューター囲碁大会＜本編＞ Step [O11o___100o0] カーネル作成

# 連載の目次

📖 [目指せ！第１４回ＵＥＣ杯コンピューター囲碁大会＜本編＞](https://qiita.com/muzudho1/items/744f6051954525878b84)  

# 本文

## Step [O11o___100o0] カーネル作成

👇 ここでは、カーネルは以下の意味を指す  

* 思考エンジンのプログラムのうち、おおまかに言って **ゲームの知識（ドメイン）以外の部分**
* １つのカーネルは、１つの対局に対応する

## Step [O11o___100o1p0] フォルダー作成

👇 以下のフォルダーを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14
👉	├── 📂 kernel
    ├── 📄 .gitignore
    ├── 📄 go.mod
  	├── 📄 go.work
 	└── 📄 main.go
```

## Step [O11o___100o2p0] カレントディレクトリーを移動

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

```shell
cd kernel
```

## Step [O11o___100o3o_1o0] ワークスペースズへ登録

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
go work use .
```

👇 以下のファイルが自動で編集されている  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 kernel
    ├── 📄 .gitignore
    ├── 📄 go.mod
👉 	├── 📄 go.work
 	└── 📄 main.go
```

```go
// ...略...


// * 以下の行が自動削除
// use .
// * 以下が自動追加
use (
	.
	./kernel
)
```

## Step [O11o___100o3o0] Goモジュールの作成

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
go mod init github.com/muzudho/kifuwarabe-uec14/kernel
```

Output:  

```shell
go: creating new go.mod: module github.com/muzudho/kifuwarabe-uec14/kernel
```

👇 以下のファイルが自動生成された  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 kernel
👉	│	└── 📄 go.mod
    ├── 📄 .gitignore
    ├── 📄 go.mod
  	├── 📄 go.work
 	└── 📄 main.go
```

```go
module github.com/muzudho/kifuwarabe-uec14/kernel

go 1.19
```

## Step [O11o___100o4p0] カレントディレクトリーを戻す

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

```shell
cd ..
```
