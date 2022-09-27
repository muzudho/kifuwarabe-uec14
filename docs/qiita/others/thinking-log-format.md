チェスや将棋のAIで、
局面を与えると、次の一手を示すようなプログラムを Engine と呼ぶ。  
日本語では 思考エンジン と呼ぶ  

**通信プロトコル** や、 **記録のフォーマット** には、いろいろな仕様がある  

チェスには 📖 [UCI](https://www.chessprogramming.org/UCI), 📖 [FEN](https://www.chess.com/terms/fen-chess) 、
ライフゲームには 📖 [RLE](https://conwaylife.com/wiki/Run_Length_Encoded)  
ドラフツ（チェッカー）には 📖 [PDN](https://en.wikipedia.org/wiki/Portable_Draughts_Notation)  
碁には 📖 [GTP](http://www.lysator.liu.se/~gunnar/gtp/)  
将棋には 📖 [USI, SFEN](http://shogidokoro.starfree.jp/usi.html), 📖 [CSA](http://www2.computer-shogi.org/protocol/)  

それに比べ、  

**思考ログのフォーマット** は 好きにしたらええやん、という状況だ。  
そこで自分用に仕様を決める。  
仕様の名前は `Thinking log format` 、略して `TLF` とでもしておく

`TLF` が参考にしたのは `CSA` と `GTP` だ  

## 使用例

📖 [Tic Tac Toe Engine Test](http://tic.warabenture.com:8000/tic-tac-toe/vol2.0/engine-manual/ver1.0/)  

# 思考エンジン内コマンドラインからコマンド入力

開発者がコマンドラインにコマンドを入力したら、  
それが単一行か、複数行かに関わらず、  
行頭に `# ` を付ける。  


## 例: 単一行 board

Input:  

```
board
```

Log:  

```
# board
```

## 例: 複数行 board

Input:  

```
board:'''
..x
.x.
o..
'''
```

Log:  

```
# board:'''
# ..x
# .x.
# o..
# '''
```

# 思考エンジン内コマンドラインからコメント入力

開発者がコメントを入力したかったら、行頭に常に `#` を付ける  

コメントが単一行か複数行かに関わらず、  
ログの行頭には `# ` を付ける  

## 例: 単一行コメント

Input:  

```
#飯が食いたい
```

Log:  

```
# #飯が食いたい
```

## 例: 複数行コメント

Input:  

```
#飯が食いたい
#何食お
```

Log:  

```
# #飯が食いたい
# #何食お
```

# 思考エンジンからの返事

思考エンジンが何か返事をしたかったら、ログの行頭に `= ` を付ける
複数行なら、２行目以降のログの行頭に `. ` を付ける  

## 例: 空返事（からへんじ）

空返事は Ok の意味を表す  

Output & Log:  

```
=
```

## 例: 返事

座標 A1 という返事なら  

Output & Log:  

```
= A1
```

* この `A1` はメッセージと呼ぶ

Output & Log:  

```
= coord:'A1'
```

* `文字列:` で始まるなら、プロパティと呼ぶ

## 例: 複数プロパティの返事

Output & Log:  

```
= profile name:'nihon taro' weight:92.6
```

* 先頭から読んで `文字列:` が出てきたとき、プロパティのリストの始まりとする
* プロパティのリストより前の部分をメッセージと呼ぶ
  * メッセージは省略してもよい
* プロパティ名と値の間に `:` を置く。スペースは開けない
  * プロパティ名に `msg` を使うのは避ける。 `msg` はメッセージを指す
* 複数のプロパティは半角空白区切りで置くことができる
* 文字列型の値のクォーテーションの `"` と `'` の違いや、エスケープシーケンスは TOML の考え方に倣う

## 例: 複数行の返事

Output & Log:  

```
= board:'''
. ..x
. .x.
. o.o
. '''
```

* ```'''``` から ```'''``` まで
* エスケープシーケンスは使えない
* 文字列型の値のクォーテーションの `"` と `'` の違いや、エスケープシーケンスは TOML の考え方に倣う

# 思考エンジンからのコメント

思考エンジンはコメントを出力しないが、ログには書く。  
常に行頭に `# ` を付ける。  

開発者のコメントは、打鍵の手間を考えて `#` の右隣にスペースを入れる必要はないが、  
思考エンジンのコメントは `#` の右隣にスペースを入れる必要がある  

基本的に、開発者のコメントか、思考エンジンのコメントかを、出力から区別しようとしていない  

## 例: 思考エンジンの単一行コメント

座標 J10 というコメントなら  

Log:  

```
# # J10
```

## 例: 思考エンジンの複数行コメント

Log:  

```
# # .xx
# # .x.
# # o.o
```

# 思考エンジンのパニックの返事

思考エンジンが　返答に困っているときは 出力の行頭に `? ` を付ける  

## 例: エラーメッセージ

Output & Log:  

```
? unknown_command
```

* スネークケースでもよい

Output & Log:  

```
? internal error
```

* 単語間にスペースを開けても良い
* コロン `:` はプロパティを書くときに使うので、メッセージ内に使うのは避ける

## 例: プロパティ付き

Output & Log:  

```
? unexpected_profile name:'nihon taro' weight:92.6
```

* 解説は `思考エンジンからの返事` を参照

## 例: 複数行文字列付き

Output & Log:  

```
? unexpected board:'''
. xxx
. xox
. xxx
. '''
```

* 解説は `思考エンジンからの返事` を参照

# 設計思想

**思考エンジンの返事** は、大会サーバーに送る必要がある。  
それ以外は 前後の文脈 や 計算途中の記録 などとして付属しているだけで、  
大会サーバーに送ると失格になることがあるだろう。  

そこで、大会サーバーに送らないものは 行頭が `# ` で始まり、  
大会サーバーに送ったものは行頭が `= ` か、 `. ` で始まり、  
エラーなら `? ` で始まるようになっている  