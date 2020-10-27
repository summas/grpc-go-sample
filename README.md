# grpc-go-sample
goでgRPCの試作を作りました。

以下のgrpcurl似てmysqlのDBから値が帰ってくることまで動作確認済み。

grpcurl -d '{ "category": "32" }' -plaintext localhost:50051 mysql_sample.AritcleService/Aritcle

{
  "result": "category is test1234567"
}
