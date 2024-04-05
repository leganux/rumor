import {NestFactory} from '@nestjs/core';
import {AppModule} from './app.module';

import {OrderModule} from './orders/orders.module';
import {MicroserviceOptions, Transport} from '@nestjs/microservices';
import {join} from "path";

async function bootstrap() {
    const app = await NestFactory.create(AppModule);
    const app_ = await NestFactory.create(OrderModule);

    const protofile = join(__dirname, '..', 'public', 'order.proto')

    // Configurar el servidor gRPC
    const grpcOptions: MicroserviceOptions = {
        transport: Transport.GRPC,
        options: {
            url: 'order_service:50052',
            package: 'order', // Nombre de tu paquete gRPC
            protoPath: protofile
        },
    };

    app_.connectMicroservice(grpcOptions);

    await app_.startAllMicroservices();
    //await app_.listen(3008);


    await app.listen(3008);
}

bootstrap();
