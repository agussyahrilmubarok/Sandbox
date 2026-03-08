import { TypeOrmModuleOptions } from '@nestjs/typeorm';
import { Book } from '../domain/entities/book.entity';
import { Member } from '../domain/entities/member.entity';

export const typeOrmConfig: TypeOrmModuleOptions = {
  type: process.env.NODE_ENV === 'production' ? 'postgres' : 'sqlite',
  database:
    process.env.NODE_ENV === 'production'
      ? process.env.DATABASE_URL
      : 'data/database.sqlite',
  entities: [Book, Member],
  synchronize: true,
  logging: process.env.NODE_ENV === 'production' ? false : true,
};
