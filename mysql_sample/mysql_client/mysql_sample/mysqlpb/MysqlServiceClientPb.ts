/**
 * @fileoverview gRPC-Web generated client stub for mysql_sample
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as mysql_sample_mysqlpb_mysql_pb from '../../mysql_sample/mysqlpb/mysql_pb';


export class AritcleServiceClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: any; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodInfoAritcle = new grpcWeb.AbstractClientBase.MethodInfo(
    mysql_sample_mysqlpb_mysql_pb.ArticlesResponse,
    (request: mysql_sample_mysqlpb_mysql_pb.ArticlesRequest) => {
      return request.serializeBinary();
    },
    mysql_sample_mysqlpb_mysql_pb.ArticlesResponse.deserializeBinary
  );

  aritcle(
    request: mysql_sample_mysqlpb_mysql_pb.ArticlesRequest,
    metadata: grpcWeb.Metadata | null): Promise<mysql_sample_mysqlpb_mysql_pb.ArticlesResponse>;

  aritcle(
    request: mysql_sample_mysqlpb_mysql_pb.ArticlesRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: mysql_sample_mysqlpb_mysql_pb.ArticlesResponse) => void): grpcWeb.ClientReadableStream<mysql_sample_mysqlpb_mysql_pb.ArticlesResponse>;

  aritcle(
    request: mysql_sample_mysqlpb_mysql_pb.ArticlesRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: mysql_sample_mysqlpb_mysql_pb.ArticlesResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/mysql_sample.AritcleService/Aritcle',
        request,
        metadata || {},
        this.methodInfoAritcle,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/mysql_sample.AritcleService/Aritcle',
    request,
    metadata || {},
    this.methodInfoAritcle);
  }

}

