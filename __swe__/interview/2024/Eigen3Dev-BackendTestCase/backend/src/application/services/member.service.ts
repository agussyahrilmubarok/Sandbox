import { Injectable, Logger } from '@nestjs/common';
import { MemberRepository } from '../../infrastructure/repositories/member.repository';
import { MemberDetailDto } from '../../domain/dto/member-detail.dto';

@Injectable()
export class MemberService {
  private readonly logger = new Logger(MemberService.name);

  constructor(private readonly memberRepository: MemberRepository) {}

  /**
   * Get All Member
   * @returns AllMembersDto
   */
  async getAllMembers(): Promise<MemberDetailDto[]> {
    this.logger.log('Fetching all members');
    const members = await this.memberRepository.findAll();
    return members.map(
      (member) =>
        new MemberDetailDto(
          member.code,
          member.name,
          JSON.parse(member.borrowedBooks).length,
        ),
    );
  }
}
