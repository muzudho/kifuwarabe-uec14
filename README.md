# kifuwarabe-uec14

コンピュータ囲碁の思考エンジン。  


# Run

Input:  

```shell
go run .
```


# 環境設定

ブログにも書いてある：  
📖 [ブログ](https://qiita.com/muzudho1/items/cea62be01f7418bbf150)  


## Go 言語の環境設定

(1) PC に Go言語をインストールする：  
📖 [GO](https://go.dev/)  

Go 言語のインストール先は `C:\Program Files\Go\` にした。  


(2) Visual Studio Code の Extensions から、 `Go` （Go Team at Google）をインストールしておく。  


## ソースコードのデプロイ先

Input:  

```shell
echo %HOMEPATH%
```

Output:  

```
\Users\muzud
```

ユーザー・ディレクトリーが `\Users\muzud` と分かった。  
Go ではユーザー・ディレクトリーの下に `go/src` というフォルダーを作るものなので、そのようにフォルダーを作る。  

さらに `go/src` フォルダーの下に `github.com\muzudho\kifuwarabe-uec14` というフォルダーを作った。  

つまり、📁 `C:\Users\muzud\go\src\github.com\muzudho\kifuwarabe-uec14` というフォルダーを作った。  

この中にソースコードを置くことにする。  
