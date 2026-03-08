import { HttpException, HttpStatus } from '@nestjs/common';

export class CannotBorrowMoreBooksException extends HttpException {
  constructor() {
    super('Cannot borrow more than 2 books', HttpStatus.BAD_REQUEST);
  }
}
