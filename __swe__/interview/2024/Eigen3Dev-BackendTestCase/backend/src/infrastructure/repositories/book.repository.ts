import { Injectable } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { Repository } from 'typeorm';
import { Book } from '../../domain/entities/book.entity';

@Injectable()
export class BookRepository {
  constructor(
    @InjectRepository(Book)
    private readonly bookRepository: Repository<Book>,
  ) {}

  findAll(): Promise<Book[]> {
    return this.bookRepository.find();
  }

  findByCode(code: string): Promise<Book | undefined> {
    return this.bookRepository.findOne({ where: { code } });
  }

  save(book: Book): Promise<Book> {
    return this.bookRepository.save(book);
  }
}
