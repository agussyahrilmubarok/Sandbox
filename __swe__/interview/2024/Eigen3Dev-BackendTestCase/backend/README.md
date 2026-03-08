# Web Services

## How it runs

!Note: please use .env is not for production for testing.

```bash
# npm install
npm install

# unit test
npm run test

# e2e test
npm run test:e2e
```

## Task

Entities

- [x] Member
- [x] Book

Use Case

- [x] Members can borrow books with conditions
  - [x] Members may not borrow more than 2 books
  - [x] Borrowed books are not borrowed by other members
  - [x] Member is currently not being penalized
- [x] Member returns the book with conditions
  - [x] The returned book is a book that the member has borrowed
  - [x] If the book is returned after more than 7 days, the member will be subject to a penalty.
  - [x] Member with penalty cannot able to borrow the book for 3 days
- [x] Check the book
  - [x] Shows all existing books and quantities
  - [x] Books that are being borrowed are not counted
- [x] Member check

  - [x] Shows all existing members
  - [x] The number of books being borrowed by each member

Requirements

- [x] it should be use any framework, but prefered NestJS Framework Or ExpressJS
- [x] it should be use Swagger as API Documentation
- [x] it should be use Database (SQL/NoSQL)
- [x] it should be open sourced on your github repo

Extras

- [ ] Implement DDD Patern
- [x] Implement Unit Testing
