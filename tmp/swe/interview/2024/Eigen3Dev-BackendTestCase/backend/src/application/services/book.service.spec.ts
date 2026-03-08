import { Test, TestingModule } from '@nestjs/testing';
import { EventEmitter2 } from '@nestjs/event-emitter';
import { BookService } from './book.service';
import { BookRepository } from '../../infrastructure/repositories/book.repository';
import { MemberRepository } from '../../infrastructure/repositories/member.repository';
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

describe('BookService', () => {
  let service: BookService;
  let bookRepository: BookRepository;
  let memberRepository: MemberRepository;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      providers: [
        BookService,
        {
          provide: BookRepository,
          useValue: {
            findByCode: jest.fn(),
            save: jest.fn(),
            findAll: jest.fn(),
          },
        },
        {
          provide: MemberRepository,
          useValue: {
            findByCode: jest.fn(),
            save: jest.fn(),
          },
        },
        EventEmitter2,
      ],
    }).compile();

    service = module.get<BookService>(BookService);
    bookRepository = module.get<BookRepository>(BookRepository);
    memberRepository = module.get<MemberRepository>(MemberRepository);
  });

  it('should be defined', () => {
    expect(service).toBeDefined();
  });

  /**
   * Test Suite for BorrowBook
   */
  describe('borrowBook', () => {
    it('should borrow a book successfully', async () => {
      const member = new Member('M001', 'John Doe');
      const book = new Book('JK-45', 'Harry Potter', 'J.K. Rowling', 1);

      jest.spyOn(memberRepository, 'findByCode').mockResolvedValue(member);
      jest.spyOn(bookRepository, 'findByCode').mockResolvedValue(book);
      jest.spyOn(memberRepository, 'save').mockResolvedValue(member);
      jest.spyOn(bookRepository, 'save').mockResolvedValue(book);

      await service.borrowBook('M001', 'JK-45');

      expect(memberRepository.save).toHaveBeenCalledWith(expect.any(Member));
      expect(bookRepository.save).toHaveBeenCalledWith(expect.any(Book));
    });

    it('should throw an error if member is not found', async () => {
      jest.spyOn(memberRepository, 'findByCode').mockResolvedValue(null);

      await expect(service.borrowBook('M001', 'BK-01')).rejects.toThrow(
        MemberNotFoundException,
      );
    });

    it('should throw an error if member is not found', async () => {
      jest.spyOn(memberRepository, 'findByCode').mockResolvedValue(null);

      await expect(service.borrowBook('M001', 'BK-01')).rejects.toThrow(
        MemberNotFoundException,
      );
    });

    it('should throw an error if book is not found', async () => {
      jest
        .spyOn(memberRepository, 'findByCode')
        .mockResolvedValue(new Member('M001', 'John Doe'));
      jest.spyOn(bookRepository, 'findByCode').mockResolvedValue(null);

      await expect(service.borrowBook('M001', 'BK-01')).rejects.toThrow(
        BookNotFoundException,
      );
    });

    it('should throw an error if book is not available', async () => {
      jest
        .spyOn(memberRepository, 'findByCode')
        .mockResolvedValue(new Member('M001', 'John Doe'));
      jest
        .spyOn(bookRepository, 'findByCode')
        .mockResolvedValue(new Book('BK-01', 'Book Title', 'Author', 0));

      await expect(service.borrowBook('M001', 'BK-01')).rejects.toThrow(
        BookNotAvailableException,
      );
    });

    it('should throw an error if member is penalized', async () => {
      const member = new Member('M001', 'John Doe');
      member.penaltyUntil = new Date(Date.now() + 24 * 60 * 60 * 1000); // Penalized until tomorrow

      jest.spyOn(memberRepository, 'findByCode').mockResolvedValue(member);
      jest
        .spyOn(bookRepository, 'findByCode')
        .mockResolvedValue(new Book('BK-01', 'Book Title', 'Author', 1));

      await expect(service.borrowBook('M001', 'BK-01')).rejects.toThrow(
        MemberIsPenalizedException,
      );
    });

    it('should throw an error if member has already borrowed 2 books', async () => {
      const member = new Member('M001', 'John Doe');
      member.borrowedBooks = JSON.stringify([
        new BorrowedBookDTO('BK-01', new Date()),
        new BorrowedBookDTO('BK-02', new Date()),
      ]);

      jest.spyOn(memberRepository, 'findByCode').mockResolvedValue(member);
      jest
        .spyOn(bookRepository, 'findByCode')
        .mockResolvedValue(new Book('BK-03', 'Book Title', 'Author', 1));

      await expect(service.borrowBook('M001', 'BK-03')).rejects.toThrow(
        CannotBorrowMoreBooksException,
      );
    });
  });

  /**
   * Test Suite for ReturnBook
   */
  describe('returnBook', () => {
    it('should return a book successfully', async () => {
      const member = new Member(
        'M001',
        'John Doe',
        JSON.stringify([new BorrowedBookDTO('JK-45', new Date())]),
      );
      const book = new Book('JK-45', 'Harry Potter', 'J.K. Rowling', 0);

      jest.spyOn(memberRepository, 'findByCode').mockResolvedValue(member);
      jest.spyOn(bookRepository, 'findByCode').mockResolvedValue(book);
      jest.spyOn(memberRepository, 'save').mockResolvedValue(member);
      jest.spyOn(bookRepository, 'save').mockResolvedValue(book);

      await service.returnBook('M001', 'JK-45');

      expect(memberRepository.save).toHaveBeenCalledWith(expect.any(Member));
      expect(bookRepository.save).toHaveBeenCalledWith(expect.any(Book));
    });

    it('should throw an error if member is not found', async () => {
      jest.spyOn(memberRepository, 'findByCode').mockResolvedValue(null);

      await expect(service.returnBook('M001', 'BK-01')).rejects.toThrow(
        MemberNotFoundException,
      );
    });

    it('should throw an error if book is not found', async () => {
      jest
        .spyOn(memberRepository, 'findByCode')
        .mockResolvedValue(new Member('M001', 'John Doe'));
      jest.spyOn(bookRepository, 'findByCode').mockResolvedValue(null);

      await expect(service.returnBook('M001', 'BK-01')).rejects.toThrow(
        BookNotFoundException,
      );
    });

    it('should throw an error if book was not borrowed by member', async () => {
      const member = new Member('M001', 'John Doe');
      const book = new Book('JK-45', 'Harry Potter', 'J.K. Rowling', 0);

      jest.spyOn(memberRepository, 'findByCode').mockResolvedValue(member);
      jest.spyOn(bookRepository, 'findByCode').mockResolvedValue(book);

      await expect(service.returnBook('M001', 'JK-45')).rejects.toThrow(
        BookNotBorrowedException,
      );
    });

    it('should apply a penalty if the book is returned late', async () => {
      jest.useFakeTimers();
      const member = new Member(
        'M001',
        'John Doe',
        JSON.stringify([
          new BorrowedBookDTO(
            'JK-45',
            new Date(Date.now() - 8 * 24 * 3600 * 1000),
          ),
        ]),
      );
      const book = new Book('JK-45', 'Harry Potter', 'J.K. Rowling', 0);

      jest.spyOn(memberRepository, 'findByCode').mockResolvedValue(member);
      jest.spyOn(bookRepository, 'findByCode').mockResolvedValue(book);
      jest.spyOn(memberRepository, 'save').mockResolvedValue(member);
      jest.spyOn(bookRepository, 'save').mockResolvedValue(book);

      const currentDate = new Date();
      jest.setSystemTime(currentDate);

      await service.returnBook('M001', 'JK-45');

      expect(member.penaltyUntil).toBeDefined();
      expect(member.penaltyUntil!.getTime()).toBeGreaterThan(
        currentDate.getTime(),
      );
      expect(memberRepository.save).toHaveBeenCalledWith(member);
      expect(bookRepository.save).toHaveBeenCalledWith(book);

      jest.useRealTimers();
    });
  });

  /**
   * Test Suite for GetAllBooks
   */
  describe('getAllBooks', () => {
    it('should return all books with their quantities, excluding borrowed books', async () => {
      const books = [
        { code: 'BK-01', title: 'Book 1', author: 'Author 1', stock: 3 },
        { code: 'BK-02', title: 'Book 2', author: 'Author 2', stock: 0 },
        { code: 'BK-03', title: 'Book 3', author: 'Author 3', stock: 5 },
      ];

      jest.spyOn(bookRepository, 'findAll').mockResolvedValue(books);

      const result = await service.getAllBooks();

      expect(result).toEqual([
        { code: 'BK-01', title: 'Book 1', author: 'Author 1', stock: 3 },
        { code: 'BK-03', title: 'Book 3', author: 'Author 3', stock: 5 },
      ]);
      expect(bookRepository.findAll).toHaveBeenCalled();
    });
  });
});
