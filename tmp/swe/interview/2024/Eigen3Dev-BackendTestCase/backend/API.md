# Library Management API

## API Endpoints

### Borrow Book

**POST** `/api/books/borrow`

**Request Body:**

```json
{
  "memberCode": "string",
  "bookCode": "string"
}
```

**Description:**
Allows a member to borrow a book. The member must not have more than 2 borrowed books, and the book must be available.

### Return Book

**POST** `/api/books/return`

**Request Body:**

```json
{
  "memberCode": "string",
  "bookCode": "string"
}
```

**Description:**
Allows a member to return a borrowed book. If the book is returned after more than 7 days, the member will be penalized and cannot borrow books for 3 days.

### Get All Books

**GET** `/api/books`

**Description:**
Retrieves a list of all existing books and their quantities. Books that are currently borrowed are not counted.

### Get All Members

**GET** `/api/members`

**Description:**
Retrieves a list of all existing members and the number of books each member is currently borrowing.