import { Test, TestingModule } from '@nestjs/testing';
import { INestApplication } from '@nestjs/common';
import * as request from 'supertest';
import { AppModule } from './../src/app.module';
import { DataSource } from 'typeorm';
import { cleanDatabase, seedDatabase } from './seed';

describe('BookController (e2e)', () => {
  let app: INestApplication;
  let dataSource: DataSource;

  beforeEach(async () => {
    const moduleFixture: TestingModule = await Test.createTestingModule({
      imports: [AppModule],
    }).compile();
    app = moduleFixture.createNestApplication();
    await app.init();

    dataSource = app.get(DataSource);
    await cleanDatabase(dataSource);
    await seedDatabase(dataSource);
  });

  afterAll(async () => {
    await app.close();
  });

  it('should borrow a book', async () => {
    const borrowBookDto = { memberCode: 'M001', bookCode: 'JK-45' };
    return request(app.getHttpServer())
      .post('/api/books/borrow')
      .send(borrowBookDto)
      .expect(200)
      .expect({ message: `Book JK-45 borrowed by member M001` });
  });

  it('should not borrow more than 2 books', async () => {
    await request(app.getHttpServer())
      .post('/api/books/borrow')
      .send({ memberCode: 'M001', bookCode: 'JK-45' })
      .expect(200);

    await request(app.getHttpServer())
      .post('/api/books/borrow')
      .send({ memberCode: 'M001', bookCode: 'SHR-1' })
      .expect(200);

    const response = await request(app.getHttpServer())
      .post('/api/books/borrow')
      .send({ memberCode: 'M001', bookCode: 'TW-11' })
      .expect(400);

    expect(response.body).toEqual({
      statusCode: 400,
      timestamp: expect.any(String),
      path: '/api/books/borrow',
      message: 'Cannot borrow more than 2 books',
    });
  });

  it('should not borrow a book that is not available', async () => {
    await request(app.getHttpServer())
      .post('/api/books/borrow')
      .send({ memberCode: 'M001', bookCode: 'JK-45' })
      .expect(200);

    const response = await request(app.getHttpServer())
      .post('/api/books/borrow')
      .send({ memberCode: 'M002', bookCode: 'JK-45' })
      .expect(400);

    expect(response.body).toEqual({
      statusCode: 400,
      timestamp: expect.any(String),
      path: '/api/books/borrow',
      message: 'Book is not available',
    });
  });

  it('should not borrow a book if member is penalized', async () => {
    const response = await request(app.getHttpServer())
      .post('/api/books/borrow')
      .send({ memberCode: 'M002', bookCode: 'JK-45' })
      .expect(400);

    expect(response.body).toEqual({
      statusCode: 400,
      timestamp: expect.any(String),
      path: '/api/books/borrow',
      message: 'Cannot borrow, member is penalized',
    });
  });

  it('should return a book', async () => {
    const returnBookDto = { memberCode: 'M004', bookCode: 'CB-01' };
    return request(app.getHttpServer())
      .post('/api/books/return')
      .send(returnBookDto)
      .expect(200)
      .expect({ message: `Book CB-01 returned by member M004` });
  });

  it('should not return a book that is not borrowed', async () => {
    const response = await request(app.getHttpServer())
      .post('/api/books/return')
      .send({ memberCode: 'M004', bookCode: 'JK-45' })
      .expect(400);

    expect(response.body).toEqual({
      statusCode: 400,
      timestamp: expect.any(String),
      path: '/api/books/return',
      message: 'Book is not borrowed by this member',
    });
  });

  it('check all existing books and quantities', async () => {
    const response = await request(app.getHttpServer())
      .get('/api/books')
      .expect(200);

    expect(response.body).toEqual([
      { author: 'J.K Rowling', code: 'JK-45', stock: 1, title: 'Harry Potter' },
      {
        author: 'Arthur Conan Doyle',
        code: 'SHR-1',
        stock: 1,
        title: 'A Study in Scarlet',
      },
      { author: 'Stephenie Meyer', code: 'TW-11', stock: 1, title: 'Twilight' },
      {
        author: 'J.R.R. Tolkien',
        code: 'HOB-83',
        stock: 1,
        title: 'The Hobbit, or There and Back Again',
      },
      {
        author: 'C.S. Lewis',
        code: 'NRN-7',
        stock: 1,
        title: 'The Lion, the Witch and the Wardrobe',
      },
      { author: 'C.B. Lewis', code: 'CB-01', stock: 1, title: 'The Coba Coba' },
    ]);
  });

  it('check all members', async () => {
    const response = await request(app.getHttpServer())
      .get('/api/members')
      .expect(200);

    expect(response.body).toEqual([
      { borrowedBooksCount: 0, code: 'M001', name: 'Angga' },
      { borrowedBooksCount: 0, code: 'M002', name: 'Ferry' },
      { borrowedBooksCount: 0, code: 'M003', name: 'Putri' },
      { borrowedBooksCount: 1, code: 'M004', name: 'Agus' },
    ]);
  });
});
