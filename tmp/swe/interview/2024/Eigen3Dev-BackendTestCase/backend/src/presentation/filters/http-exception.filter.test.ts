import { ArgumentsHost, HttpException } from '@nestjs/common';
import { Request, Response } from 'express';
import { HttpExceptionFilter } from './http-exception.filter';

describe('HttpExceptionFilter', () => {
  let filter: HttpExceptionFilter;
  let mockArgumentsHost: ArgumentsHost;

  beforeEach(() => {
    filter = new HttpExceptionFilter();
    mockArgumentsHost = {
      switchToHttp: jest.fn().mockReturnValue({
        getResponse: jest.fn().mockReturnValue({
          status: jest.fn().mockReturnThis(),
          json: jest.fn().mockReturnThis(),
        }),
        getRequest: jest.fn().mockReturnValue({
          url: '/test-url',
        }),
      }),
    } as unknown as ArgumentsHost;
  });

  it('should return a formatted response', () => {
    const exception = new HttpException('Test error', 400);
    const response = mockArgumentsHost.switchToHttp().getResponse<Response>();
    const request = mockArgumentsHost.switchToHttp().getRequest<Request>();

    filter.catch(exception, mockArgumentsHost);

    expect(response.status).toHaveBeenCalledWith(400);
    expect(response.json).toHaveBeenCalledWith({
      statusCode: 400,
      timestamp: expect.any(String),
      path: request.url,
      message: 'Test error',
    });
  });
});
