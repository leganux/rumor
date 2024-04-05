// orders.controller.ts

import { Controller, NotFoundException } from '@nestjs/common';
import { GrpcMethod } from '@nestjs/microservices';
import { OrderService } from './orders.service';
import {
  CreateOrderRequest,
  CreateOrderResponse,
  FindAllOrdersRequest,
  FindAllOrdersResponse,
  UpdateOrderRequest,
  UpdateOrderResponse,
  DeleteOrderRequest,
  DeleteOrderResponse,
} from './../generated/order';

@Controller()
export class OrderController {
  constructor(private readonly orderService: OrderService) {}

  @GrpcMethod('OrderService', 'CreateOrder')
  async createOrder(data: CreateOrderRequest): Promise<CreateOrderResponse> {
    const orderId = await this.orderService.create(data);

    const response: CreateOrderResponse = {
      orderId: orderId,
    };

    return response;
  }

  @GrpcMethod('OrderService', 'FindAllOrders')
  async findAllOrders(
    data: FindAllOrdersRequest,
  ): Promise<FindAllOrdersResponse> {
    const orders = await this.orderService.findAll();
    console.log(data);
    const response: FindAllOrdersResponse = {
      orders: orders,
    };

    return response;
  }

  @GrpcMethod('OrderService', 'UpdateOrder')
  async updateOrder(data: UpdateOrderRequest): Promise<UpdateOrderResponse> {
    const { id, ...updateData } = data;
    try {
      const updatedOrder = await this.orderService.update(id, updateData);
      return { order: updatedOrder };
    } catch (error) {
      if (error instanceof NotFoundException) {
        throw new NotFoundException(error.message);
      }
      throw error;
    }
  }

  @GrpcMethod('OrderService', 'DeleteOrder')
  async deleteOrder(data: DeleteOrderRequest): Promise<DeleteOrderResponse> {
    const { id } = data;
    try {
      const result = this.orderService.remove(id);
      return { message: result };
    } catch (error) {
      if (error instanceof NotFoundException) {
        throw new NotFoundException(error.message);
      }
      throw error;
    }
  }
}
