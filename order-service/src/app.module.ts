import {Module} from '@nestjs/common';
import {AppController} from './app.controller';
import {AppService} from './app.service';
import {OrderModule} from './orders/orders.module';
import {join} from "path";
import {ServeStaticModule} from '@nestjs/serve-static';
@Module({
    imports: [OrderModule, ServeStaticModule.forRoot({
        rootPath: join(__dirname, '..', 'public'),
    }),],
    controllers: [AppController],
    providers: [AppService],
})
export class AppModule {
}
