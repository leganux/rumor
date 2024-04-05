import {Module} from '@nestjs/common';
import {AppController} from './app.controller';
import {AppService} from './app.service';
import {ProductsModule} from './products/products.module';


import {ServeStaticModule} from '@nestjs/serve-static';
import { join } from 'path';

@Module({
    imports: [
        ProductsModule,
        ServeStaticModule.forRoot({
            rootPath: join(__dirname, '..', 'public'),
        }),
    ],
    controllers: [AppController],
    providers: [AppService],
})
export class AppModule {}
