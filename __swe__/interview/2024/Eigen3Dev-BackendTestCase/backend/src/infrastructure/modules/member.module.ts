import { Module } from '@nestjs/common';
import { TypeOrmModule } from '@nestjs/typeorm';
import { MemberController } from '../../presentation/controllers/member.controller';
import { MemberService } from '../../application/services/member.service';
import { MemberRepository } from '../repositories/member.repository';
import { Member } from '../../domain/entities/member.entity';

@Module({
  imports: [TypeOrmModule.forFeature([Member])],
  controllers: [MemberController],
  providers: [MemberService, MemberRepository],
  exports: [MemberRepository],
})
export class MemberModule {}
