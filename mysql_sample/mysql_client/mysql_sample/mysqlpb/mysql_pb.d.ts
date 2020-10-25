import * as jspb from 'google-protobuf'



export class Articles extends jspb.Message {
  getCategory(): string;
  setCategory(value: string): Articles;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Articles.AsObject;
  static toObject(includeInstance: boolean, msg: Articles): Articles.AsObject;
  static serializeBinaryToWriter(message: Articles, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Articles;
  static deserializeBinaryFromReader(message: Articles, reader: jspb.BinaryReader): Articles;
}

export namespace Articles {
  export type AsObject = {
    category: string,
  }
}

export class ArticlesRequest extends jspb.Message {
  getCategory(): Articles | undefined;
  setCategory(value?: Articles): ArticlesRequest;
  hasCategory(): boolean;
  clearCategory(): ArticlesRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ArticlesRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ArticlesRequest): ArticlesRequest.AsObject;
  static serializeBinaryToWriter(message: ArticlesRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ArticlesRequest;
  static deserializeBinaryFromReader(message: ArticlesRequest, reader: jspb.BinaryReader): ArticlesRequest;
}

export namespace ArticlesRequest {
  export type AsObject = {
    category?: Articles.AsObject,
  }
}

export class ArticlesResponse extends jspb.Message {
  getResult(): string;
  setResult(value: string): ArticlesResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ArticlesResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ArticlesResponse): ArticlesResponse.AsObject;
  static serializeBinaryToWriter(message: ArticlesResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ArticlesResponse;
  static deserializeBinaryFromReader(message: ArticlesResponse, reader: jspb.BinaryReader): ArticlesResponse;
}

export namespace ArticlesResponse {
  export type AsObject = {
    result: string,
  }
}

