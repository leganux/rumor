import {NestFactory} from '@nestjs/core';
import {AppModule} from './app.module';

import {ProductsModule} from './products/products.module';
import {MicroserviceOptions, Transport} from '@nestjs/microservices';
import * as path from 'path';
import * as dotenv from 'dotenv';


async function bootstrap() {

    try {
        dotenv.config();
    } catch (e) {

    }
    const app = await NestFactory.create(AppModule);
    const app_ = await NestFactory.create(ProductsModule);

    const protofile = path.join(__dirname, '..', 'public', 'product.proto')


    let uri_grpc = 'product_service:'
    if (process.env.ENVIRONMENT == 'develop') {
        uri_grpc = 'localhost:'
    }
    uri_grpc = uri_grpc + (process.env.PORT_GRCP || 50057)

    console.log(uri_grpc)


    const grpcOptions: MicroserviceOptions = {
        transport: Transport.GRPC,
        options: {
            url: uri_grpc,
            package: 'product', // Nombre de tu paquete gRPC
            protoPath: protofile, // Ruta a tu archivo proto
        },
    };

    app_.connectMicroservice(grpcOptions);
    await app_.startAllMicroservices();


    await app.listen(process.env.PORT_API || 3007);


}

bootstrap();
