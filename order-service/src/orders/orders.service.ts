import {Injectable, NotFoundException} from '@nestjs/common';
import {CreateOrderDto, UpdateOrderDto} from './dto';

@Injectable()
export class OrderService {
    private orders = [];

    create(createOrderDto: CreateOrderDto) {
        const newOrder = {id: this.orders.length + 1, ...createOrderDto};
        this.orders.push(newOrder);
        return String(newOrder.id);
    }

    findAll() {
        return this.orders;
    }

    findOne(id: string) {
        console.log('busca', id, this.orders)

        const order = this.orders.find((order) => order.id == id);
        if (!order) {
            throw new NotFoundException(`Order with id ${id} not found`);
        }
        return order;
    }

    update(id: string, updateOrderDto: UpdateOrderDto) {
        const order = this.findOne(id);
        // Actualizar solo las propiedades proporcionadas
        Object.assign(order, updateOrderDto);
        return order;
    }

    remove(id: string) {
        const index = this.orders.findIndex((order) => order.id == id);
        if (index === -1) {
            throw new NotFoundException(`Order with id ${id} not found`);
        }
        this.orders.splice(index, 1);
        return `Order with id ${id} has been deleted`;
    }
}
