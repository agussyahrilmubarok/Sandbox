import { Module } from '@nestjs/common';
import { BookController } from '../../presentation/controllers/book.controller';
import { EventEmitterModule } from '@nestjs/event-emitter';
import { BookService } from '../../application/services/book.service';
import { BookRepository } from '../repositories/book.repository';
import { MemberModule } from './member.module';
import { TypeOrmModule } from '@nestjs/typeorm';
import { Book } from '../../domain/entities/book.entity';

@Module({
  imports: [
    TypeOrmModule.forFeature([Book]),
    EventEmitterModule.forRoot(),
    MemberModule,
  ],
  controllers: [BookController],
  providers: [BookService, BookRepository],
})
export class BookModule {}
