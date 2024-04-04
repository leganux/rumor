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

export interface FindAllOrdersRequest {
}

export interface FindAllOrdersResponse {
  orders: Order[];
}

export interface UpdateOrderRequest {
  id: string;
  customerId: string;
  products: OrderProduct[];
}

export interface UpdateOrderResponse {
  order: Order | undefined;
}

export interface DeleteOrderRequest {
  id: string;
}

export interface DeleteOrderResponse {
  message: string;
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

function createBaseFindAllOrdersRequest(): FindAllOrdersRequest {
  return {};
}

export const FindAllOrdersRequest = {
  encode(_: FindAllOrdersRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): FindAllOrdersRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseFindAllOrdersRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(_: any): FindAllOrdersRequest {
    return {};
  },

  toJSON(_: FindAllOrdersRequest): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<FindAllOrdersRequest>, I>>(base?: I): FindAllOrdersRequest {
    return FindAllOrdersRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<FindAllOrdersRequest>, I>>(_: I): FindAllOrdersRequest {
    const message = createBaseFindAllOrdersRequest();
    return message;
  },
};

function createBaseFindAllOrdersResponse(): FindAllOrdersResponse {
  return { orders: [] };
}

export const FindAllOrdersResponse = {
  encode(message: FindAllOrdersResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.orders) {
      Order.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): FindAllOrdersResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseFindAllOrdersResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.orders.push(Order.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): FindAllOrdersResponse {
    return { orders: globalThis.Array.isArray(object?.orders) ? object.orders.map((e: any) => Order.fromJSON(e)) : [] };
  },

  toJSON(message: FindAllOrdersResponse): unknown {
    const obj: any = {};
    if (message.orders?.length) {
      obj.orders = message.orders.map((e) => Order.toJSON(e));
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<FindAllOrdersResponse>, I>>(base?: I): FindAllOrdersResponse {
    return FindAllOrdersResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<FindAllOrdersResponse>, I>>(object: I): FindAllOrdersResponse {
    const message = createBaseFindAllOrdersResponse();
    message.orders = object.orders?.map((e) => Order.fromPartial(e)) || [];
    return message;
  },
};

function createBaseUpdateOrderRequest(): UpdateOrderRequest {
  return { id: "", customerId: "", products: [] };
}

export const UpdateOrderRequest = {
  encode(message: UpdateOrderRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdateOrderRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateOrderRequest();
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

  fromJSON(object: any): UpdateOrderRequest {
    return {
      id: isSet(object.id) ? globalThis.String(object.id) : "",
      customerId: isSet(object.customerId) ? globalThis.String(object.customerId) : "",
      products: globalThis.Array.isArray(object?.products)
        ? object.products.map((e: any) => OrderProduct.fromJSON(e))
        : [],
    };
  },

  toJSON(message: UpdateOrderRequest): unknown {
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

  create<I extends Exact<DeepPartial<UpdateOrderRequest>, I>>(base?: I): UpdateOrderRequest {
    return UpdateOrderRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<UpdateOrderRequest>, I>>(object: I): UpdateOrderRequest {
    const message = createBaseUpdateOrderRequest();
    message.id = object.id ?? "";
    message.customerId = object.customerId ?? "";
    message.products = object.products?.map((e) => OrderProduct.fromPartial(e)) || [];
    return message;
  },
};

function createBaseUpdateOrderResponse(): UpdateOrderResponse {
  return { order: undefined };
}

export const UpdateOrderResponse = {
  encode(message: UpdateOrderResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.order !== undefined) {
      Order.encode(message.order, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdateOrderResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateOrderResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.order = Order.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): UpdateOrderResponse {
    return { order: isSet(object.order) ? Order.fromJSON(object.order) : undefined };
  },

  toJSON(message: UpdateOrderResponse): unknown {
    const obj: any = {};
    if (message.order !== undefined) {
      obj.order = Order.toJSON(message.order);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<UpdateOrderResponse>, I>>(base?: I): UpdateOrderResponse {
    return UpdateOrderResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<UpdateOrderResponse>, I>>(object: I): UpdateOrderResponse {
    const message = createBaseUpdateOrderResponse();
    message.order = (object.order !== undefined && object.order !== null) ? Order.fromPartial(object.order) : undefined;
    return message;
  },
};

function createBaseDeleteOrderRequest(): DeleteOrderRequest {
  return { id: "" };
}

export const DeleteOrderRequest = {
  encode(message: DeleteOrderRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteOrderRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteOrderRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.id = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): DeleteOrderRequest {
    return { id: isSet(object.id) ? globalThis.String(object.id) : "" };
  },

  toJSON(message: DeleteOrderRequest): unknown {
    const obj: any = {};
    if (message.id !== "") {
      obj.id = message.id;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<DeleteOrderRequest>, I>>(base?: I): DeleteOrderRequest {
    return DeleteOrderRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<DeleteOrderRequest>, I>>(object: I): DeleteOrderRequest {
    const message = createBaseDeleteOrderRequest();
    message.id = object.id ?? "";
    return message;
  },
};

function createBaseDeleteOrderResponse(): DeleteOrderResponse {
  return { message: "" };
}

export const DeleteOrderResponse = {
  encode(message: DeleteOrderResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.message !== "") {
      writer.uint32(10).string(message.message);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteOrderResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteOrderResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.message = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): DeleteOrderResponse {
    return { message: isSet(object.message) ? globalThis.String(object.message) : "" };
  },

  toJSON(message: DeleteOrderResponse): unknown {
    const obj: any = {};
    if (message.message !== "") {
      obj.message = message.message;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<DeleteOrderResponse>, I>>(base?: I): DeleteOrderResponse {
    return DeleteOrderResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<DeleteOrderResponse>, I>>(object: I): DeleteOrderResponse {
    const message = createBaseDeleteOrderResponse();
    message.message = object.message ?? "";
    return message;
  },
};

export interface OrderService {
  CreateOrder(request: CreateOrderRequest): Promise<CreateOrderResponse>;
  FindAllOrders(request: FindAllOrdersRequest): Promise<FindAllOrdersResponse>;
  UpdateOrder(request: UpdateOrderRequest): Promise<UpdateOrderResponse>;
  DeleteOrder(request: DeleteOrderRequest): Promise<DeleteOrderResponse>;
}

export const OrderServiceServiceName = "order.OrderService";
export class OrderServiceClientImpl implements OrderService {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || OrderServiceServiceName;
    this.rpc = rpc;
    this.CreateOrder = this.CreateOrder.bind(this);
    this.FindAllOrders = this.FindAllOrders.bind(this);
    this.UpdateOrder = this.UpdateOrder.bind(this);
    this.DeleteOrder = this.DeleteOrder.bind(this);
  }
  CreateOrder(request: CreateOrderRequest): Promise<CreateOrderResponse> {
    const data = CreateOrderRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "CreateOrder", data);
    return promise.then((data) => CreateOrderResponse.decode(_m0.Reader.create(data)));
  }

  FindAllOrders(request: FindAllOrdersRequest): Promise<FindAllOrdersResponse> {
    const data = FindAllOrdersRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "FindAllOrders", data);
    return promise.then((data) => FindAllOrdersResponse.decode(_m0.Reader.create(data)));
  }

  UpdateOrder(request: UpdateOrderRequest): Promise<UpdateOrderResponse> {
    const data = UpdateOrderRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "UpdateOrder", data);
    return promise.then((data) => UpdateOrderResponse.decode(_m0.Reader.create(data)));
  }

  DeleteOrder(request: DeleteOrderRequest): Promise<DeleteOrderResponse> {
    const data = DeleteOrderRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "DeleteOrder", data);
    return promise.then((data) => DeleteOrderResponse.decode(_m0.Reader.create(data)));
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
