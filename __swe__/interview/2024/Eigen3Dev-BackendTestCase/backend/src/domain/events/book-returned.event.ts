export class BookReturnedEvent {
  constructor(
    public readonly memberCode: string,
    public readonly bookCode: string,
    public readonly returnedLate: boolean,
  ) {}
}
