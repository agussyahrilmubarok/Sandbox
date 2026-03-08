import { Module } from '@nestjs/common';
import { TypeOrmModule } from '@nestjs/typeorm';
import { BookModule } from './infrastructure/modules/book.module';
import { MemberModule } from './infrastructure/modules/member.module';
import { typeOrmConfig } from './config/typeorm.config';

@Module({
  imports: [TypeOrmModule.forRoot(typeOrmConfig), BookModule, MemberModule],
})
export class AppModule {}
