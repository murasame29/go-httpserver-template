# go-server-template
Golang (echo)のHttpServerテンプレート  
echoでrouterが現状実装されているがhttp.Handlerを満たしていればなんでもOK  
NewRelicの実装をいくつかしてあるけど使わなかったら削除して  


## Makefile
幾つかのヘルパーがあるよ

サーバを起動する
```bash
make run
# envファイルを指定して起動 例:.env
make ENV_FILE=.env run
```

全てのテストの実行とカバレッジ表示
```bash
make test
```

## ディレクトリ構成
```
|─ cmd  　　　　　　 <- cmd関連のDir
|   └─ app 
|       └─ main.go  <- こいつがサーバ
└─ internal         <- 内部的な実装
    |─ config       <- 環境変数とかサーバの設定など
    |─ pkg  　      <- pkg
    |─ router       <- ルーティングの設定など、フレームワークに依存するところ
    └─ server       <- serverの起動とかシャットダウンするとこ。基本フレームワークに依存しない
```


