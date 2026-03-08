import { DataSource } from 'typeorm';
import { Book } from '../src/domain/entities/book.entity';
import { Member } from '../src/domain/entities/member.entity';

export async function seedDatabase(dataSource: DataSource) {
  const bookRepository = dataSource.getRepository(Book);
  const memberRepository = dataSource.getRepository(Member);

  const books = [
    { code: 'JK-45', title: 'Harry Potter', author: 'J.K Rowling', stock: 1 },
    {
      code: 'SHR-1',
      title: 'A Study in Scarlet',
      author: 'Arthur Conan Doyle',
      stock: 1,
    },
    { code: 'TW-11', title: 'Twilight', author: 'Stephenie Meyer', stock: 1 },
    {
      code: 'HOB-83',
      title: 'The Hobbit, or There and Back Again',
      author: 'J.R.R. Tolkien',
      stock: 1,
    },
    {
      code: 'NRN-7',
      title: 'The Lion, the Witch and the Wardrobe',
      author: 'C.S. Lewis',
      stock: 1,
    },
    {
      code: 'CB-01',
      title: 'The Coba Coba',
      author: 'C.B. Lewis',
      stock: 1,
    },
  ];

  const members = [
    { code: 'M001', name: 'Angga' },
    {
      code: 'M002',
      name: 'Ferry',
      borrowedBooks: '[]',
      penaltyUntil: new Date(Date.now() + 2 * 24 * 60 * 60 * 1000),
    },
    { code: 'M003', name: 'Putri', borrowedBooks: '[]', penaltyUntil: null },
    {
      code: 'M004',
      name: 'Agus',
      borrowedBooks: JSON.stringify([
        {
          bookCode: 'CB-01',
          borrowDate: new Date(Date.now() - 8 * 24 * 60 * 60 * 1000),
        },
      ]),
      penaltyUntil: null,
    },
  ];

  await bookRepository.save(books);
  await memberRepository.save(members);
}

export async function cleanDatabase(dataSource: DataSource) {
  const bookRepository = dataSource.getRepository(Book);
  const memberRepository = dataSource.getRepository(Member);

  await bookRepository.clear();
  await memberRepository.clear();
}
