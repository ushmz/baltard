# Koolhas(search-system) Cloned

簡易な検索システムのバックエンドサーバー

![OMA-CCTV-building-Beijing_dezeen](https://user-images.githubusercontent.com/38400669/75150699-51365b00-5748-11ea-8995-0997223fb6c5.jpg)

## ✅ Requirements

- Go 1.15+
- Docker

## Setup

1. git clone する
2. docker-compose にて build & daemon で MySQL コンテナを起動する
3. 立てるコンテナの PORT は，`3366`としている

```bash
cd backend
make build
make up
```

- seeder 以下には，ihoriya が利用している初期データ構築用クエリがある
- 現在は，Docker 上の MySQL へクエリをそのまま投げているため非常に面倒
- TODO: コンテナ起動 -> マイグレーション実行 -> シードデータ投入を自動で行う用の
  ビルドファイルを書く

3. マイグレーション実行

- マイグレーションツールには
  ，[sql-migrate](https://github.com/rubenv/sql-migrate)を用いている
- sql-migrate の利用方法は，上記リンクの README を参照(非常に簡単)

```bash
cd backend/db/migrations
make up
```

4. Go 環境があるとことを仮定して，`go run main.go`で API サーバーを起動

- なければ，goenvを使ったり自由に環境作成(brewでもok)
- Ref: [Go言語の開発環境をMacとVScodeで作りコードを動かしてみる](https://sagantaf.hatenablog.com/entry/2020/02/08/221720)

TODO: Go も docker 化させる

## DBの確認方法
1. `make up`を実行し，MySQLコンテナを起動
2. 起動されたら，`docker exec -it MySQLコンテナのID`でコンテナ内に入る(ターミナルのユーザーが`root@5621abe683d4:/#`のように変更されたらOK)
3. `mysql -u koolhaas -p`を実行する．passwordが求められるので，`koolhaas`と入力．
4. MySQL内に入り，`use koolhaas;`でDBを切り替え，操作可能となる

## NOTE

- 現状では小規模な API なので，ある程度の行数になるまで，レイヤーを分けない方針
  でいく
- 基本的に DB は，実験を行う者がそれぞれ持つことにする
  - 初期のマイグレーションファイルには，汎用的なテーブル設計(& ihoriya が使うもの)を
    追加している
  - 利用しないテーブルが作られるのが嫌な場合は，マイグレーションを実行す
    る前に，実行したくないマイグレーションファイルを削除する
  - また必要に応じて，migration ファイルを追加(e.g. テーブルやカラムを増やす目的)すればよ
    い
- DB の正規化は分析の際に面倒を減らすため崩している
