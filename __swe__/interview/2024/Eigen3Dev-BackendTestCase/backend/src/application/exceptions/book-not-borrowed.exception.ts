import { HttpException, HttpStatus } from '@nestjs/common';

export class BookNotBorrowedException extends HttpException {
  constructor() {
    super('Book is not borrowed by this member', HttpStatus.BAD_REQUEST);
  }
}
