package com.example.member.repos;

import com.example.member.domain.Member;
import org.springframework.data.jpa.repository.JpaRepository;

import java.util.Optional;
import java.util.UUID;

public interface MemberRepository extends JpaRepository<Member, UUID> {

    Optional<Member> findByCode(String code);

    Optional<Member> findByEmail(String email);

    boolean existsByCodeIgnoreCase(String code);

    boolean existsByEmailIgnoreCase(String email);
}
