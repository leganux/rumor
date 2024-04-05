import {NestFactory} from '@nestjs/core';
import {AppModule} from './app.module';

import {OrderModule} from './orders/orders.module';
import {MicroserviceOptions, Transport} from '@nestjs/microservices';
import {join} from "path";
import * as dotenv from 'dotenv';

async function bootstrap() {
    try {
        dotenv.config();
    } catch (e) {

    }
    const app = await NestFactory.create(AppModule);
    const app_ = await NestFactory.create(OrderModule);

    const protofile = join(__dirname, '..', 'public', 'order.proto')

    let uri_grpc = 'order_service:'
    if (process.env.ENVIRONMENT == 'develop') {
        uri_grpc = 'localhost:'
    }
    uri_grpc = uri_grpc + (process.env.PORT_GRCP || 50052)

    console.log(uri_grpc)

    // Configurar el servidor gRPC
    const grpcOptions: MicroserviceOptions = {
        transport: Transport.GRPC,
        options: {
            url: uri_grpc,
            package: 'order', // Nombre de tu paquete gRPC
            protoPath: protofile
        },
    };

    app_.connectMicroservice(grpcOptions);

    await app_.startAllMicroservices();
    await app.listen(process.env.PORT_API || 3008);
}

bootstrap();
