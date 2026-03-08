import { HttpException, HttpStatus } from '@nestjs/common';

export class BookNotFoundException extends HttpException {
  constructor() {
    super('Book is not available', HttpStatus.NOT_FOUND);
  }
}
