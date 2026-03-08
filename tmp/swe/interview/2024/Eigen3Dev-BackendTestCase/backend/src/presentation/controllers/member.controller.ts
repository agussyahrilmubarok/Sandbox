import { Controller, Get, UseFilters } from '@nestjs/common';
import { ApiOperation, ApiResponse, ApiTags } from '@nestjs/swagger';
import { MemberService } from '../../application/services/member.service';
import { HttpExceptionFilter } from '../filters/http-exception.filter';

@ApiTags('api/members')
@Controller('api/members')
@UseFilters(new HttpExceptionFilter())
export class MemberController {
  constructor(private readonly memberService: MemberService) {}

  @Get()
  @ApiOperation({ summary: 'Get all members' })
  @ApiResponse({ status: 200, description: 'Return all members.' })
  async getAllMembers() {
    const members = await this.memberService.getAllMembers();
    return members;
  }
}
