import { IsNotEmpty, IsString } from 'class-validator';

export class ReturnBookDto {
  @IsNotEmpty()
  @IsString()
  memberCode: string;

  @IsNotEmpty()
  @IsString()
  bookCode: string;
}
