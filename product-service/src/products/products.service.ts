// products.service.ts

import {Injectable, NotFoundException} from '@nestjs/common';
import {CreateProductDto, UpdateProductDto} from './dto';

interface Product {
    id: string;
    name: string;
    description: string;
    price: number;
    quantity: number;
}

@Injectable()
export class ProductsService {
    private products: Product[] = [];

    create(productData: CreateProductDto): Product {
        const id = (this.products.length + 1).toString(); // Asignar un ID incremental
        const product: Product = {
            id,
            ...productData,
        };
        this.products.push(product);
        return product;
    }

    findAll() {
        return this.products;
    }

    findOne(id: string) {
        const product = this.products.find(product => product.id === id);
        if (!product) {
            throw new NotFoundException('Product not found');
        }
        return product;
    }

    update(id: string, updateOrderDto: UpdateProductDto) {
        const product = this.findOne(id);
        // Actualizar solo las propiedades proporcionadas
        Object.assign(product, updateOrderDto);
        return product;
    }

    delete(id: string): string {
        const index = this.products.findIndex((product) => product.id === id);
        if (index === -1) {
            throw new NotFoundException('Product not found');
        }
        this.products.splice(index, 1);
        return 'Product has been deleted';
    }
}
