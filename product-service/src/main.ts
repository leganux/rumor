import {NestFactory} from '@nestjs/core';
import {AppModule} from './app.module';

import {ProductsModule} from './products/products.module';
import {MicroserviceOptions, Transport} from '@nestjs/microservices';
import * as path from 'path';

async function bootstrap() {
    const app = await NestFactory.create(AppModule);
    const app_ = await NestFactory.create(ProductsModule);

    const protofile = path.join(__dirname, '..', 'public', 'product.proto')


    const grpcOptions: MicroserviceOptions = {
        transport: Transport.GRPC,
        options: {
            url: 'product_service:50057',
            package: 'product', // Nombre de tu paquete gRPC
            protoPath: protofile, // Ruta a tu archivo proto
        },
    };

    app_.connectMicroservice(grpcOptions);
    await app_.startAllMicroservices();


    await app.listen(3007);


}

bootstrap();
