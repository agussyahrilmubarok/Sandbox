import { Test, TestingModule } from '@nestjs/testing';
import { BookController } from './book.controller';
import { BookService } from '../../application/services/book.service';
import {
  MemberNotFoundException,
  BookNotFoundException,
  BookNotAvailableException,
  CannotBorrowMoreBooksException,
  BookNotBorrowedException,
} from '../../application/exceptions';

describe('BookController', () => {
  let bookController: BookController;
  let bookService: BookService;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      controllers: [BookController],
      providers: [
        {
          provide: BookService,
          useValue: {
            borrowBook: jest.fn(),
            returnBook: jest.fn(),
            getAllBooks: jest.fn(),
          },
        },
      ],
    }).compile();

    bookController = module.get<BookController>(BookController);
    bookService = module.get<BookService>(BookService);
  });

  /**
   * Test Suites for BorrowBook
   */
  describe('borrowBook', () => {
    it('should borrow a book successfully', async () => {
      const borrowBookDto = { memberCode: 'M001', bookCode: 'BK-01' };

      await bookController.borrowBook(borrowBookDto);

      expect(bookService.borrowBook).toHaveBeenCalledWith(
        borrowBookDto.memberCode,
        borrowBookDto.bookCode,
      );
    });

    it('should throw MemberNotFoundException', async () => {
      const borrowBookDto = { memberCode: 'M001', bookCode: 'BK-01' };
      jest
        .spyOn(bookService, 'borrowBook')
        .mockRejectedValue(new MemberNotFoundException());

      await expect(bookController.borrowBook(borrowBookDto)).rejects.toThrow(
        MemberNotFoundException,
      );
    });

    it('should throw BookNotFoundException', async () => {
      const borrowBookDto = { memberCode: 'M001', bookCode: 'BK-01' };
      jest
        .spyOn(bookService, 'borrowBook')
        .mockRejectedValue(new BookNotFoundException());

      await expect(bookController.borrowBook(borrowBookDto)).rejects.toThrow(
        BookNotFoundException,
      );
    });

    it('should throw BookNotAvailableException', async () => {
      const borrowBookDto = { memberCode: 'M001', bookCode: 'BK-01' };
      jest
        .spyOn(bookService, 'borrowBook')
        .mockRejectedValue(new BookNotAvailableException());

      await expect(bookController.borrowBook(borrowBookDto)).rejects.toThrow(
        BookNotAvailableException,
      );
    });

    it('should throw CannotBorrowMoreBooksException', async () => {
      const borrowBookDto = { memberCode: 'M001', bookCode: 'BK-01' };
      jest
        .spyOn(bookService, 'borrowBook')
        .mockRejectedValue(new CannotBorrowMoreBooksException());

      await expect(bookController.borrowBook(borrowBookDto)).rejects.toThrow(
        CannotBorrowMoreBooksException,
      );
    });
  });

  /**
   * Test Suites for ReturnBook
   */
  describe('returnBook', () => {
    it('should return a book successfully', async () => {
      const returnBookDto = { memberCode: 'M001', bookCode: 'BK-01' };

      await bookController.returnBook(returnBookDto);

      expect(bookService.returnBook).toHaveBeenCalledWith(
        returnBookDto.memberCode,
        returnBookDto.bookCode,
      );
    });

    it('should throw MemberNotFoundException', async () => {
      const returnBookDto = { memberCode: 'M001', bookCode: 'BK-01' };
      jest
        .spyOn(bookService, 'returnBook')
        .mockRejectedValue(new MemberNotFoundException());

      await expect(bookController.returnBook(returnBookDto)).rejects.toThrow(
        MemberNotFoundException,
      );
    });

    it('should throw BookNotFoundException', async () => {
      const returnBookDto = { memberCode: 'M001', bookCode: 'BK-01' };
      jest
        .spyOn(bookService, 'returnBook')
        .mockRejectedValue(new BookNotFoundException());

      await expect(bookController.returnBook(returnBookDto)).rejects.toThrow(
        BookNotFoundException,
      );
    });

    it('should throw BookNotBorrowedException', async () => {
      const returnBookDto = { memberCode: 'M001', bookCode: 'BK-01' };
      jest
        .spyOn(bookService, 'returnBook')
        .mockRejectedValue(new BookNotBorrowedException());

      await expect(bookController.returnBook(returnBookDto)).rejects.toThrow(
        BookNotBorrowedException,
      );
    });
  });

  /**
   * Test Suites for GetAllBooks
   */
  describe('getAllBooks', () => {
    it('should return all books with stock greater than 0', async () => {
      const books = [
        { code: 'BK-01', title: 'Book 1', author: 'Author 1', stock: 3 },
        { code: 'BK-02', title: 'Book 2', author: 'Author 2', stock: 0 },
        { code: 'BK-03', title: 'Book 3', author: 'Author 3', stock: 5 },
      ];
      jest
        .spyOn(bookService, 'getAllBooks')
        .mockResolvedValue(books.filter((book) => book.stock > 0));

      const result = await bookController.getAllBooks();

      expect(result).toEqual([
        { code: 'BK-01', title: 'Book 1', author: 'Author 1', stock: 3 },
        { code: 'BK-03', title: 'Book 3', author: 'Author 3', stock: 5 },
      ]);
      expect(bookService.getAllBooks).toHaveBeenCalled();
    });
  });
});
