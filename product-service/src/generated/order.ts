/* eslint-disable */
import * as _m0 from "protobufjs/minimal";

export const protobufPackage = "order";

export interface OrderProduct {
  productId: string;
  quantity: number;
}

export interface Order {
  id: string;
  customerId: string;
  products: OrderProduct[];
}

export interface CreateOrderRequest {
  customerId: string;
  products: OrderProduct[];
}

export interface CreateOrderResponse {
  orderId: string;
}

function createBaseOrderProduct(): OrderProduct {
  return { productId: "", quantity: 0 };
}

export const OrderProduct = {
  encode(message: OrderProduct, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.productId !== "") {
      writer.uint32(10).string(message.productId);
    }
    if (message.quantity !== 0) {
      writer.uint32(16).int32(message.quantity);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): OrderProduct {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseOrderProduct();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.productId = reader.string();
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.quantity = reader.int32();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): OrderProduct {
    return {
      productId: isSet(object.productId) ? globalThis.String(object.productId) : "",
      quantity: isSet(object.quantity) ? globalThis.Number(object.quantity) : 0,
    };
  },

  toJSON(message: OrderProduct): unknown {
    const obj: any = {};
    if (message.productId !== "") {
      obj.productId = message.productId;
    }
    if (message.quantity !== 0) {
      obj.quantity = Math.round(message.quantity);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<OrderProduct>, I>>(base?: I): OrderProduct {
    return OrderProduct.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<OrderProduct>, I>>(object: I): OrderProduct {
    const message = createBaseOrderProduct();
    message.productId = object.productId ?? "";
    message.quantity = object.quantity ?? 0;
    return message;
  },
};

function createBaseOrder(): Order {
  return { id: "", customerId: "", products: [] };
}

export const Order = {
  encode(message: Order, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.customerId !== "") {
      writer.uint32(18).string(message.customerId);
    }
    for (const v of message.products) {
      OrderProduct.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Order {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseOrder();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.id = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.customerId = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.products.push(OrderProduct.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Order {
    return {
      id: isSet(object.id) ? globalThis.String(object.id) : "",
      customerId: isSet(object.customerId) ? globalThis.String(object.customerId) : "",
      products: globalThis.Array.isArray(object?.products)
        ? object.products.map((e: any) => OrderProduct.fromJSON(e))
        : [],
    };
  },

  toJSON(message: Order): unknown {
    const obj: any = {};
    if (message.id !== "") {
      obj.id = message.id;
    }
    if (message.customerId !== "") {
      obj.customerId = message.customerId;
    }
    if (message.products?.length) {
      obj.products = message.products.map((e) => OrderProduct.toJSON(e));
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<Order>, I>>(base?: I): Order {
    return Order.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<Order>, I>>(object: I): Order {
    const message = createBaseOrder();
    message.id = object.id ?? "";
    message.customerId = object.customerId ?? "";
    message.products = object.products?.map((e) => OrderProduct.fromPartial(e)) || [];
    return message;
  },
};

function createBaseCreateOrderRequest(): CreateOrderRequest {
  return { customerId: "", products: [] };
}

export const CreateOrderRequest = {
  encode(message: CreateOrderRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.customerId !== "") {
      writer.uint32(10).string(message.customerId);
    }
    for (const v of message.products) {
      OrderProduct.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateOrderRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateOrderRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.customerId = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.products.push(OrderProduct.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CreateOrderRequest {
    return {
      customerId: isSet(object.customerId) ? globalThis.String(object.customerId) : "",
      products: globalThis.Array.isArray(object?.products)
        ? object.products.map((e: any) => OrderProduct.fromJSON(e))
        : [],
    };
  },

  toJSON(message: CreateOrderRequest): unknown {
    const obj: any = {};
    if (message.customerId !== "") {
      obj.customerId = message.customerId;
    }
    if (message.products?.length) {
      obj.products = message.products.map((e) => OrderProduct.toJSON(e));
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<CreateOrderRequest>, I>>(base?: I): CreateOrderRequest {
    return CreateOrderRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<CreateOrderRequest>, I>>(object: I): CreateOrderRequest {
    const message = createBaseCreateOrderRequest();
    message.customerId = object.customerId ?? "";
    message.products = object.products?.map((e) => OrderProduct.fromPartial(e)) || [];
    return message;
  },
};

function createBaseCreateOrderResponse(): CreateOrderResponse {
  return { orderId: "" };
}

export const CreateOrderResponse = {
  encode(message: CreateOrderResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.orderId !== "") {
      writer.uint32(10).string(message.orderId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateOrderResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateOrderResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.orderId = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CreateOrderResponse {
    return { orderId: isSet(object.orderId) ? globalThis.String(object.orderId) : "" };
  },

  toJSON(message: CreateOrderResponse): unknown {
    const obj: any = {};
    if (message.orderId !== "") {
      obj.orderId = message.orderId;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<CreateOrderResponse>, I>>(base?: I): CreateOrderResponse {
    return CreateOrderResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<CreateOrderResponse>, I>>(object: I): CreateOrderResponse {
    const message = createBaseCreateOrderResponse();
    message.orderId = object.orderId ?? "";
    return message;
  },
};

export interface OrderService {
  CreateOrder(request: CreateOrderRequest): Promise<CreateOrderResponse>;
}

export const OrderServiceServiceName = "order.OrderService";
export class OrderServiceClientImpl implements OrderService {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || OrderServiceServiceName;
    this.rpc = rpc;
    this.CreateOrder = this.CreateOrder.bind(this);
  }
  CreateOrder(request: CreateOrderRequest): Promise<CreateOrderResponse> {
    const data = CreateOrderRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "CreateOrder", data);
    return promise.then((data) => CreateOrderResponse.decode(_m0.Reader.create(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends globalThis.Array<infer U> ? globalThis.Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
