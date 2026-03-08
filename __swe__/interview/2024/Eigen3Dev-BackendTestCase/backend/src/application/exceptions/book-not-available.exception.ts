import { HttpException, HttpStatus } from '@nestjs/common';

export class BookNotAvailableException extends HttpException {
  constructor() {
    super('Book is not available', HttpStatus.BAD_REQUEST);
  }
}
