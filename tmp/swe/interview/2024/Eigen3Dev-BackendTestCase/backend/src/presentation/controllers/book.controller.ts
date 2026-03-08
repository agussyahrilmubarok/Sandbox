import {
  Controller,
  Get,
  Post,
  Body,
  UseFilters,
  HttpCode,
} from '@nestjs/common';
import { ApiBody, ApiOperation, ApiResponse, ApiTags } from '@nestjs/swagger';
import { BookService } from '../../application/services/book.service';
import { HttpExceptionFilter } from '../filters/http-exception.filter';

@ApiTags('api/books')
@Controller('api/books')
@UseFilters(new HttpExceptionFilter())
export class BookController {
  constructor(private readonly bookService: BookService) {}

  @Post('borrow')
  @HttpCode(200)
  @ApiOperation({ summary: 'Borrow a book' })
  @ApiBody({
    schema: {
      type: 'object',
      properties: {
        memberCode: { type: 'string' },
        bookCode: { type: 'string' },
      },
    },
  })
  @ApiResponse({ status: 200, description: 'Member borrow a book.' })
  async borrowBook(
    @Body() borrowBookDto: { memberCode: string; bookCode: string },
  ) {
    const { memberCode, bookCode } = borrowBookDto;
    await this.bookService.borrowBook(memberCode, bookCode);
    return { message: `Book ${bookCode} borrowed by member ${memberCode}` };
  }

  @Post('return')
  @HttpCode(200)
  @ApiOperation({ summary: 'Return a book' })
  @ApiBody({
    schema: {
      type: 'object',
      properties: {
        memberCode: { type: 'string' },
        bookCode: { type: 'string' },
      },
    },
  })
  @ApiResponse({ status: 200, description: 'Member return a book.' })
  async returnBook(
    @Body() returnBookDto: { memberCode: string; bookCode: string },
  ) {
    const { memberCode, bookCode } = returnBookDto;
    await this.bookService.returnBook(memberCode, bookCode);
    return { message: `Book ${bookCode} returned by member ${memberCode}` };
  }

  @ApiOperation({ summary: 'Get all books' })
  @ApiResponse({ status: 200, description: 'Return all books.' })
  @Get()
  async getAllBooks() {
    const books = await this.bookService.getAllBooks();
    return books;
  }
}
