export class BookBorrowedEvent {
  constructor(
    public readonly memberCode: string,
    public readonly bookCode: string,
  ) {}
}
