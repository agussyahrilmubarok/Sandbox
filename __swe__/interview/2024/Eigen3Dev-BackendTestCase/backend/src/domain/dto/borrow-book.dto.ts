import { IsNotEmpty, IsString } from 'class-validator';

export class BorrowBookDto {
  @IsNotEmpty()
  @IsString()
  memberCode: string;

  @IsNotEmpty()
  @IsString()
  bookCode: string;
}
