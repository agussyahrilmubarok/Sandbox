import { Injectable } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { Repository } from 'typeorm';
import { Member } from '../../domain/entities/member.entity';

@Injectable()
export class MemberRepository {
  constructor(
    @InjectRepository(Member)
    private readonly memberRepository: Repository<Member>,
  ) {}

  findAll(): Promise<Member[]> {
    return this.memberRepository.find();
  }

  findByCode(code: string): Promise<Member | undefined> {
    return this.memberRepository.findOne({ where: { code } });
  }

  save(member: Member): Promise<Member> {
    return this.memberRepository.save(member);
  }
}
