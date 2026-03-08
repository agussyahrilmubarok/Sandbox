import { Entity, PrimaryColumn, Column } from 'typeorm';

@Entity()
export class Book {
  @PrimaryColumn()
  code: string;

  @Column()
  title: string;

  @Column()
  author: string;

  @Column()
  stock: number;

  constructor(code: string, title: string, author: string, stock: number) {
    this.code = code;
    this.title = title;
    this.author = author;
    this.stock = stock;
  }
}
