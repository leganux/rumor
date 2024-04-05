// products.controller.ts

import {Controller} from '@nestjs/common';
import {GrpcMethod} from '@nestjs/microservices';
import {ProductsService} from './products.service';
import {
    CreateProductRequest,
    CreateProductResponse,
    FindAllProductsRequest,
    FindAllProductsResponse,
    FindProductByIdRequest,
    FindProductByIdResponse,
    UpdateProductRequest,
    UpdateProductResponse,
    DeleteProductRequest,
    DeleteProductResponse,
} from './../generated/product';

@Controller()
export class ProductController {
    constructor(private readonly productService: ProductsService) {
    }

    @GrpcMethod('ProductService', 'CreateProduct')
    async createProduct(
        data: CreateProductRequest,
    ): Promise<{ productId: number }> {
        const product = await this.productService.create(data);
        return {productId: Number(product.id)};
    }

    @GrpcMethod('ProductService', 'FindAllProducts')
    async findAllProducts(
        data: FindAllProductsRequest,
    ): Promise<FindAllProductsResponse> {
        const products = await this.productService.findAll();
        return {products};
    }

    @GrpcMethod('ProductService', 'FindProductById')
    async findProductById(
        data: FindProductByIdRequest,
    ): Promise<FindProductByIdResponse> {
        const product = await this.productService.findOne(data.id);
        return {product};
    }

    @GrpcMethod('ProductService', 'UpdateProduct')
    async updateProduct(
        data: UpdateProductRequest,
    ): Promise<UpdateProductResponse> {
        const product = await this.productService.update(data.id, data);
        return {product};
    }

    @GrpcMethod('ProductService', 'DeleteProduct')
    async deleteProduct(
        data: DeleteProductRequest,
    ): Promise<DeleteProductResponse> {
        const message = this.productService.delete(data.id);
        return {message};
    }
}
