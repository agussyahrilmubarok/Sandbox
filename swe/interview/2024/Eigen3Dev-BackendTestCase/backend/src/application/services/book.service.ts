import { Injectable, Logger } from '@nestjs/common';
import { EventEmitter2 } from '@nestjs/event-emitter';
import { BookRepository } from '../../infrastructure/repositories/book.repository';
import { MemberRepository } from '../../infrastructure/repositories/member.repository';
import { BookBorrowedEvent } from '../../domain/events/book-borrowed.event';
import { BookReturnedEvent } from '../../domain/events/book-returned.event';
import { Book } from '../../domain/entities/book.entity';
import { Member } from '../../domain/entities/member.entity';
import { BorrowedBookDTO } from '../../domain/dto/borrowed-book.dto';
import {
  MemberNotFoundException,
  BookNotFoundException,
  BookNotAvailableException,
  CannotBorrowMoreBooksException,
  BookNotBorrowedException,
  MemberIsPenalizedException,
} from '../../application/exceptions';

@Injectable()
export class BookService {
  private readonly logger = new Logger(BookService.name);

  constructor(
    private readonly bookRepository: BookRepository,
    private readonly memberRepository: MemberRepository,
    private readonly eventEmitter: EventEmitter2,
  ) {}

  /**
   * Borrow Book
   * @param memberCode string
   * @param bookCode string
   */
  async borrowBook(memberCode: string, bookCode: string): Promise<void> {
    const { member, book } = await this.findMemberAndBook(memberCode, bookCode);

    if (book.stock < 1) {
      this.logger.error(`Book not available: ${bookCode}`);
      throw new BookNotAvailableException();
    }

    const currentDate = new Date();
    if (member.penaltyUntil && member.penaltyUntil > currentDate) {
      this.logger.error(`Member: ${memberCode} cannot borrow cause penalized`);
      throw new MemberIsPenalizedException();
    }

    if (member.penaltyUntil < currentDate) {
      member.penaltyUntil = null;
    }

    const borrowedBooksArrays: BorrowedBookDTO[] = JSON.parse(
      member.borrowedBooks,
    );
    if (borrowedBooksArrays.length >= 2) {
      this.logger.error(`Member: ${memberCode} cannot borrow more books`);
      throw new CannotBorrowMoreBooksException();
    }

    const newBorrowedBook = new BorrowedBookDTO(bookCode, new Date());
    borrowedBooksArrays.push(newBorrowedBook);
    book.stock -= 1;
    member.borrowedBooks = JSON.stringify(borrowedBooksArrays);

    await this.saveEntities(member, book);

    this.eventEmitter.emit(
      'book.borrowed',
      new BookBorrowedEvent(memberCode, bookCode),
    );
    this.logger.log(`Book borrowed: ${bookCode} by member: ${memberCode}`);
  }

  /**
   * Return Book
   * @param memberCode string
   * @param bookCode string
   */
  async returnBook(memberCode: string, bookCode: string): Promise<void> {
    const { member, book } = await this.findMemberAndBook(memberCode, bookCode);

    const borrowedBooksArrays: BorrowedBookDTO[] = JSON.parse(
      member.borrowedBooks,
    );
    const borrowedBook = borrowedBooksArrays.find(
      (book) => book.bookCode === bookCode,
    );
    if (!borrowedBook) {
      this.logger.error(
        `Book: ${bookCode} was not borrowed by member: ${memberCode}`,
      );
      throw new BookNotBorrowedException();
    }

    const returnedLate =
      (new Date().getTime() - new Date(borrowedBook.borrowedDate).getTime()) /
        (1000 * 3600 * 24) >
      7;
    const index = borrowedBooksArrays.findIndex(
      (book) => book.bookCode === bookCode,
    );
    if (index > -1) {
      borrowedBooksArrays.splice(index, 1);
    }

    book.stock += 1;
    if (returnedLate) {
      member.penaltyUntil = new Date(
        new Date().getTime() + 3 * 24 * 3600 * 1000,
      );
    }
    member.borrowedBooks = JSON.stringify(borrowedBooksArrays);

    await this.saveEntities(member, book);

    this.eventEmitter.emit(
      'book.returned',
      new BookReturnedEvent(memberCode, bookCode, returnedLate),
    );
    this.logger.log(`Book returned: ${bookCode} by member: ${memberCode}`);
  }

  /**
   * Get All Books
   * @returns Books[]
   */
  async getAllBooks(): Promise<Book[]> {
    this.logger.log('Fetching all books');
    const books = await this.bookRepository.findAll();
    return books.filter((book) => book.stock > 0);
  }

  private async findMemberAndBook(memberCode: string, bookCode: string) {
    const member = await this.memberRepository.findByCode(memberCode);
    const book = await this.bookRepository.findByCode(bookCode);

    if (!member) {
      this.logger.error(`Member not found: ${memberCode}`);
      throw new MemberNotFoundException();
    }

    if (!book) {
      this.logger.error(`Book not found: ${bookCode}`);
      throw new BookNotFoundException();
    }

    return { member, book };
  }

  private async saveEntities(member: Member, book: Book) {
    await this.memberRepository.save(member);
    await this.bookRepository.save(book);
  }
}
