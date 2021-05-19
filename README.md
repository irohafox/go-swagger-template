# golang-rest-api-template
このリポジトリはgo/swagger-goを用いたWebAPI実装のベースとなる試験的なリポジトリです。

## TODO
- [ ] swagger-ui
- [ ] DB/ORM
- [ ] token認証

## Dockerコンテナ環境自動構築
```
docker/bootstrap
```

## Dockerコンテナ起動
```
docker-compose up

# デーモン起動の場合には-dオプション付きで
docker-compose up -d
```


## Swaggerコマンド集
swagger.yamlのバリデーション
```
docker-compose run --rm web swagger validate swagger.yaml
```
swagger.yamlをベースとしたサーバサイドコード生成
```
docker-compose run --rm web swagger generate server -f swagger.yaml
```
