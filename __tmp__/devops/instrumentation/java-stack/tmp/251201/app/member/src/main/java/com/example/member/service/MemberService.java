package com.example.member.service;

import com.example.member.model.MemberDTO;

import java.util.List;
import java.util.UUID;

public interface MemberService {

    List<MemberDTO> findAll();

    MemberDTO findById(UUID id);

    MemberDTO findByCode(String code);

    MemberDTO findByEmail(String email);

    MemberDTO create(MemberDTO memberDTO);

    MemberDTO update(UUID id, MemberDTO memberDTO);

    void deleteById(UUID id);
}
