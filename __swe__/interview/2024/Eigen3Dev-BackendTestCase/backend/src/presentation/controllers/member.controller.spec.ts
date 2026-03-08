import { Test, TestingModule } from '@nestjs/testing';
import { MemberController } from './member.controller';
import { MemberService } from '../../application/services/member.service';
import { MemberDetailDto } from 'src/domain/dto/member-detail.dto';

describe('MemberController', () => {
  let memberController: MemberController;
  let memberService: MemberService;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      controllers: [MemberController],
      providers: [
        {
          provide: MemberService,
          useValue: {
            getAllMembers: jest.fn(),
          },
        },
      ],
    }).compile();

    memberController = module.get<MemberController>(MemberController);
    memberService = module.get<MemberService>(MemberService);
  });

  describe('getAllMembers', () => {
    it('should return all members', async () => {
      const members: MemberDetailDto[] = [
        {
          code: 'M001',
          name: 'John Doe',

          borrowedBooksCount: 2,
        },
        {
          code: 'M002',
          name: 'Jane Doe',
          borrowedBooksCount: 1,
        },
        {
          code: 'M003',
          name: 'Jim Doe',
          borrowedBooksCount: 0,
        },
      ];
      jest.spyOn(memberService, 'getAllMembers').mockResolvedValue(members);

      const result = await memberController.getAllMembers();

      expect(result).toEqual(members);
      expect(memberService.getAllMembers).toHaveBeenCalled();
    });
  });
});
