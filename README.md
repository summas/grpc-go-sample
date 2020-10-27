# grpc-go-sample
goによるgRPCの試作。

以下のgrpcurlコマンドによりmysqlのDBから所得した値が帰ってくることを動作確認済み。

grpcurl -d '{ "category": "32" }' -plaintext localhost:50051 mysql_sample.AritcleService/Aritcle

{
  "result": "category is test1234567"
}
