// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var product_pb = require('./product_pb.js');

function serialize_product_GetProductsRequest(arg) {
  if (!(arg instanceof product_pb.GetProductsRequest)) {
    throw new Error('Expected argument of type product.GetProductsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_product_GetProductsRequest(buffer_arg) {
  return product_pb.GetProductsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_product_GetProductsResponse(arg) {
  if (!(arg instanceof product_pb.GetProductsResponse)) {
    throw new Error('Expected argument of type product.GetProductsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_product_GetProductsResponse(buffer_arg) {
  return product_pb.GetProductsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var ProductServiceService = exports.ProductServiceService = {
  getProducts: {
    path: '/product.ProductService/GetProducts',
    requestStream: false,
    responseStream: false,
    requestType: product_pb.GetProductsRequest,
    responseType: product_pb.GetProductsResponse,
    requestSerialize: serialize_product_GetProductsRequest,
    requestDeserialize: deserialize_product_GetProductsRequest,
    responseSerialize: serialize_product_GetProductsResponse,
    responseDeserialize: deserialize_product_GetProductsResponse,
  },
};

exports.ProductServiceClient = grpc.makeGenericClientConstructor(ProductServiceService);
