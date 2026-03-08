import { HttpException, HttpStatus } from '@nestjs/common';

export class MemberNotFoundException extends HttpException {
  constructor() {
    super('Member not found', HttpStatus.NOT_FOUND);
  }
}
