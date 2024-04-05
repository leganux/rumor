// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var order_pb = require('./order_pb.js');

function serialize_order_CreateOrderRequest(arg) {
  if (!(arg instanceof order_pb.CreateOrderRequest)) {
    throw new Error('Expected argument of type order.CreateOrderRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_order_CreateOrderRequest(buffer_arg) {
  return order_pb.CreateOrderRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_order_CreateOrderResponse(arg) {
  if (!(arg instanceof order_pb.CreateOrderResponse)) {
    throw new Error('Expected argument of type order.CreateOrderResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_order_CreateOrderResponse(buffer_arg) {
  return order_pb.CreateOrderResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var OrderServiceService = exports.OrderServiceService = {
  createOrder: {
    path: '/order.OrderService/CreateOrder',
    requestStream: false,
    responseStream: false,
    requestType: order_pb.CreateOrderRequest,
    responseType: order_pb.CreateOrderResponse,
    requestSerialize: serialize_order_CreateOrderRequest,
    requestDeserialize: deserialize_order_CreateOrderRequest,
    responseSerialize: serialize_order_CreateOrderResponse,
    responseDeserialize: deserialize_order_CreateOrderResponse,
  },
};

exports.OrderServiceClient = grpc.makeGenericClientConstructor(OrderServiceService);
