import { Test, TestingModule } from '@nestjs/testing';
import { MemberService } from './member.service';
import { MemberRepository } from '../../infrastructure/repositories/member.repository';
import { Member } from '../../domain/entities/member.entity';
import { MemberDetailDto } from '../../domain/dto/member-detail.dto';

describe('MemberService', () => {
  let service: MemberService;
  let memberRepository: MemberRepository;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      providers: [
        MemberService,
        {
          provide: MemberRepository,
          useValue: {
            findAll: jest.fn(),
            findByCode: jest.fn(),
            save: jest.fn(),
          },
        },
      ],
    }).compile();

    service = module.get<MemberService>(MemberService);
    memberRepository = module.get<MemberRepository>(MemberRepository);
  });

  it('should be defined', () => {
    expect(service).toBeDefined();
  });

  /**
   * Test Suite for GetAllMembers
   */
  describe('getAllMembers', () => {
    it('should return all members with the number of books being borrowed by each member', async () => {
      const members = [
        new Member(
          'M001',
          'John Doe',
          JSON.stringify([
            { bookCode: 'BK-01', borrowedDate: new Date('2024-07-01') },
            { bookCode: 'BK-02', borrowedDate: new Date('2024-07-02') },
          ]),
        ),
        new Member(
          'M002',
          'Jane Doe',
          JSON.stringify([
            { bookCode: 'BK-03', borrowedDate: new Date('2024-07-03') },
          ]),
        ),
        new Member('M003', 'Jim Doe', '[]'),
      ];

      jest.spyOn(memberRepository, 'findAll').mockResolvedValue(members);

      const result = await service.getAllMembers();

      const expected: MemberDetailDto[] = [
        new MemberDetailDto('M001', 'John Doe', 2),
        new MemberDetailDto('M002', 'Jane Doe', 1),
        new MemberDetailDto('M003', 'Jim Doe', 0),
      ];

      expect(result).toEqual(expected);
      expect(memberRepository.findAll).toHaveBeenCalled();
    });
  });
});
