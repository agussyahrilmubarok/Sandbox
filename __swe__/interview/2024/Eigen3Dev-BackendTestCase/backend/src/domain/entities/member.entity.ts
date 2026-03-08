import { Entity, PrimaryColumn, Column } from 'typeorm';

@Entity()
export class Member {
  @PrimaryColumn()
  code: string;

  @Column()
  name: string;

  @Column('text', { default: '[]' })
  borrowedBooks: string;

  @Column({ nullable: true })
  penaltyUntil?: Date;

  constructor(
    code: string,
    name: string,
    borrowedBooks: string = '[]',
    penaltyUntil?: Date,
  ) {
    this.code = code;
    this.name = name;
    this.borrowedBooks = borrowedBooks;
    this.penaltyUntil = penaltyUntil;
  }
}
