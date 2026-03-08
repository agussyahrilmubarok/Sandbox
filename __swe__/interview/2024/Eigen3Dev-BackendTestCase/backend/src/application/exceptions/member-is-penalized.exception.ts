import { HttpException, HttpStatus } from '@nestjs/common';

export class MemberIsPenalizedException extends HttpException {
  constructor() {
    super('Cannot borrow, member is penalized', HttpStatus.BAD_REQUEST);
  }
}
