# デバッグ方法を調べる

Windows を使っているとする  

👇 とりあえずググって適当に記事を読む  

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

このとき、 `dlv` もインストールしたものとします  

# Step [O2o0] デバッガーを使用する
