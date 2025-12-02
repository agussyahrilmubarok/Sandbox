package com.example.member.service;

import com.example.member.domain.Member;
import com.example.member.model.MemberDTO;
import com.example.member.repos.MemberRepository;
import com.example.member.util.NotFoundException;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.data.domain.Sort;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.UUID;

@Service
@Slf4j
@RequiredArgsConstructor
public class MemberServiceImpl implements MemberService {

    private final MemberRepository memberRepository;

    @Override
    public List<MemberDTO> findAll() {
        final List<Member> members = memberRepository.findAll(Sort.by("code"));

        log.info("Successfully fetched {} members", members.size());
        return members.stream()
                .map(member -> mapToDTO(member, new MemberDTO()))
                .toList();
    }

    @Override
    public MemberDTO findById(UUID id) {
        return memberRepository.findById(id)
                .map(member -> {
                    log.info("Member found for id={}", id);
                    return mapToDTO(member, new MemberDTO());
                })
                .orElseThrow(() -> {
                    log.warn("Member not found for id={}", id);
                    return new NotFoundException("Member not found");
                });
    }

    @Override
    public MemberDTO findByCode(String code) {
        return memberRepository.findByCode(code)
                .map(member -> {
                    log.info("Member found for code={}", code);
                    return mapToDTO(member, new MemberDTO());
                })
                .orElseThrow(() -> {
                    log.warn("Member not found for code={}", code);
                    return new NotFoundException("Member not found");
                });
    }

    @Override
    public MemberDTO findByEmail(String email) {
        return memberRepository.findByEmail(email)
                .map(member -> {
                    log.info("Member found for email={}", email);
                    return mapToDTO(member, new MemberDTO());
                })
                .orElseThrow(() -> {
                    log.warn("Member not found for email={}", email);
                    return new NotFoundException("Member not found");
                });
    }

    @Override
    public MemberDTO create(MemberDTO memberDTO) {
        if (codeExists(memberDTO.getCode())) {
            log.warn("Failed creating member — code={} already exists", memberDTO.getCode());
            throw new IllegalArgumentException("Code already exists");
        }

        if (emailExists(memberDTO.getEmail())) {
            log.warn("Failed creating member — email={} already exists", memberDTO.getEmail());
            throw new IllegalArgumentException("Email already exists");
        }

        Member member = new Member();
        mapToEntity(memberDTO, member);
        member = memberRepository.save(member);

        log.info("Successfully created member id={} code={}", member.getId(), member.getCode());
        return mapToDTO(member, new MemberDTO());
    }

    @Override
    public MemberDTO update(UUID id, MemberDTO memberDTO) {
        Member member = memberRepository.findById(id)
                .orElseThrow(() -> {
                    log.warn("Update failed — member not found for id={}", id);
                    return new NotFoundException("Member not found");
                });

        mapToEntity(memberDTO, member);
        member = memberRepository.save(member);

        log.info("Successfully updated member id={}", id);
        return mapToDTO(member, new MemberDTO());
    }

    @Override
    public void deleteById(UUID id) {
        final Member member = memberRepository.findById(id)
                .orElseThrow(() -> {
                    log.warn("Delete failed — member not found for id={}", id);
                    return new NotFoundException("Member not found");
                });

        memberRepository.delete(member);
        log.info("Successfully deleted member id={}", id);
    }

    private MemberDTO mapToDTO(final Member member, final MemberDTO memberDTO) {
        memberDTO.setId(member.getId());
        memberDTO.setCode(member.getCode());
        memberDTO.setName(member.getName());
        memberDTO.setEmail(member.getEmail());
        memberDTO.setCreatedAt(member.getCreatedAt());
        memberDTO.setUpdatedAt(member.getUpdatedAt());
        return memberDTO;
    }

    private Member mapToEntity(final MemberDTO memberDTO, final Member member) {
        member.setCode(memberDTO.getCode());
        member.setName(memberDTO.getName());
        member.setEmail(memberDTO.getEmail());
        return member;
    }

    private boolean codeExists(final String code) {
        return memberRepository.existsByCodeIgnoreCase(code);
    }

    private boolean emailExists(final String email) {
        return memberRepository.existsByEmailIgnoreCase(email);
    }
}
