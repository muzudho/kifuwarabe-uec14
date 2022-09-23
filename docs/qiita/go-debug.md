# Go言語、Visual Studio Code でのデバッグ方法を調べる

Windows を使っているとする  

👇 とりあえずググって適当に記事を読む。大筋で、以下の記事の通りやるとできる  

📖 [Goのデバッグ環境 on VSCode](https://future-architect.github.io/articles/20201117/)  

とりあえず真似する  

# Step [O1o0] はじめに

## Step [O1o1o0] エディター

Visual studio code を使っているとする  

## Step [O1o2o0] プログラム言語

Go 言語をインストール済みとする  

## Step [O1o3o0] Go エクステンションを入れる

📖 [Go for Visual Studio Code](https://marketplace.visualstudio.com/items?itemName=golang.Go)  

## Step [O1o4o0] 関連するツールをインストールする

`[Ctrl] + [Shift] + [P]` をキーにする  

`Go: Install/Update Tools` を選ぶ  

出てくるツールを全部選んでインストールする  

Output:  

```plaintext
All tools successfully installed. You are ready to Go. :)
```

このとき、 `dlv` もインストールしたものとする  

# Step [O2o0] デバッガーを使用する

## Step [O2o1o0] 設定ファイルの作成

VSCodeのツールバーの　虫みたいなアイコンをクリック  

`create a launch.json file` をクリック  

VSCode の上の方にフォルダー名か出てくるので、とりあえず ソースを置いているフォルダーを選ぶ  

👇 以下のファイルが自動生成される

```plaintext
    📂 .vscode
    └── 📄 launch.json
```

```json
{
    // Use IntelliSense to learn about possible attributes.
    // Hover to view descriptions of existing attributes.
    // For more information, visit: https://go.microsoft.com/fwlink/?linkid=830387
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch Package",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${fileDirname}"
        }
    ]
}
```

## Step [O2o2o0] デバッグの実行

ブレークポイントを置く  

`F5` キーを押すとデバッグが開始される  
