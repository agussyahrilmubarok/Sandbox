export class BorrowedBookDTO {
  constructor(
    public bookCode: string,
    public borrowedDate: Date,
  ) {}
}
