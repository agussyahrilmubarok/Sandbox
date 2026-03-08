import { IsNotEmpty } from 'class-validator';

export class MemberDetailDto {
  @IsNotEmpty()
  code: string;

  @IsNotEmpty()
  name: string;

  @IsNotEmpty()
  borrowedBooksCount: number;

  constructor(code: string, name: string, borrowedBooksCount: number) {
    this.code = code;
    this.name = name;
    this.borrowedBooksCount = borrowedBooksCount;
  }
}
