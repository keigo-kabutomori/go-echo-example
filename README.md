# Go勉強会用資料

## 動作手順

1. `.env.sample`を`.env`に名前を変更してください。
2. `go run server.go`を実行して下さい。
3. Paw等のクライアントソフトでAPIにアクセスしてください。

## 環境変数

| 変数名  | 説明                     |
| ------- | ------------------------ |
| DB_TYPE | データベースのタイプ     |
| DB_NAME | DB名を入れて下さい       |
| PORT    | ポート番号を入れて下さい |

### ローカルで確認する場合
`.env.sample`を`.env`に名前を変更してください。

### Heroku

Herokuで動かす場合

| 変数名  | 説明                                                                                   |
| ------- | -------------------------------------------------------------------------------------- |
| DB_TYPE | `postgres`                                                                             |
| DB_NAME | PostgreSQLをHerokuにインストールすると自動的に設定されるのでユーザー側で設定不要です。 |
| PORT    | Heroku側で自動的に設定されるのでユーザー側で設定不要です。                             |

## API一覧

### `GET /logs`

ログ一覧を取得する

#### Example Response

```json
[
  {
    "ID":1,
    "CreatedAt":"2020-04-18T23:36:27.706047+09:00",
    "UpdatedAt":"2020-04-18T23:38:38.225574+09:00",
    "DeletedAt":null,
    "text":"samplelelelelele"
  },{
    "ID":2,
    "CreatedAt":"2020-04-18T23:36:47.637507+09:00",
    "UpdatedAt":"2020-04-18T23:36:47.637507+09:00",
    "DeletedAt":null,
    "text":"sample"
  },{
    "ID":3,
    "CreatedAt":"2020-04-18T23:36:48.058618+09:00",
    "UpdatedAt":"2020-04-18T23:36:48.058618+09:00",
    "DeletedAt":null,"text":"sample"
  }
]
```

### `GET /logs/:id`

ログを取得する

#### Example Response

```json
{
  "ID":1,
  "CreatedAt":"2020-04-18T23:36:27.706047+09:00",
  "UpdatedAt":"2020-04-18T23:36:27.706047+09:00",
  "DeletedAt":null,
  "text":"sample"
}
```

### `POST /logs`

ログを作成する

#### Exapmle Request

```json
{
  "text":"sample"
}
```

#### Example Response

```json
{
  "ID":1,
  "CreatedAt":"2020-04-18T23:36:27.706047+09:00",
  "UpdatedAt":"2020-04-18T23:36:27.706047+09:00",
  "DeletedAt":null,
  "text":"sample"
}
```

### `PUT /logs/:id`

ログを更新する

#### Exapmle Request

```json
{
  "text":"samplelelelele"
}
```

#### Example Response

```json
{
  "ID":1,
  "CreatedAt":"2020-04-18T23:36:27.706047+09:00",
  "UpdatedAt":"2020-04-18T23:36:27.706047+09:00",
  "DeletedAt":null,
  "text":"samplelelelele"
}
```

### `DELETE /logs/:id`

ログを削除する

#### Example Response

```json
{}