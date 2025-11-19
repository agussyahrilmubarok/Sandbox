package com.example.member.rest;

import com.example.member.model.MemberDTO;
import com.example.member.service.MemberService;
import io.swagger.v3.oas.annotations.Operation;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.UUID;


@RestController
@RequestMapping(value = "/api/v1/members", produces = MediaType.APPLICATION_JSON_VALUE)
@Slf4j
@RequiredArgsConstructor
public class MemberResource {

    private final MemberService memberService;

    @GetMapping
    @Operation(
            summary = "Get all members",
            description = "Retrieves all member data",
            tags = {"Members"}
    )
    public ResponseEntity<List<MemberDTO>> findAll() {
        List<MemberDTO> members = memberService.findAll();

        log.info("Successfully retrieved all members, count={}", members.size());
        return ResponseEntity.ok(members);
    }

    @GetMapping("/find")
    @Operation(
            summary = "Find member by code",
            description = "Retrieve a single member by their code",
            tags = {"Members"}
    )
    public ResponseEntity<MemberDTO> find(@RequestParam(value = "code", required = true) String memberCode) {
        MemberDTO member = memberService.findByCode(memberCode);

        log.info("Successfully retrieved member with code={}", memberCode);
        return ResponseEntity.ok(member);
    }

    @PostMapping("/init-dummy")
    @Operation(
            summary = "Initialize dummy member data",
            description = "Create dummy data for testing",
            tags = {"Dummy"}
    )
    public ResponseEntity<List<MemberDTO>> initDummy() {
        List<MemberDTO> dummies = List.of(
                MemberDTO.builder()
                        .id(UUID.randomUUID())
                        .code("MEMBER-1000")
                        .name("John Doe")
                        .email("johndoe@mail.com")
                        .build(),

                MemberDTO.builder()
                        .id(UUID.randomUUID())
                        .code("MEMBER-1001")
                        .name("Jane Smith")
                        .email("janesmith@mail.com")
                        .build()
        );
        for (MemberDTO dummy : dummies) {
            try {
                memberService.create(dummy);
                log.info("Created dummy member: code={} email={}", dummy.getCode(), dummy.getEmail());
            } catch (Exception e) {
                log.error("Failed to create dummy member: code={} email={}", dummy.getCode(), dummy.getEmail(), e);
            }
        }

        log.info("Dummy members initialization completed. Total: {}", dummies.size());
        return ResponseEntity.ok(dummies);
    }

    @DeleteMapping("/clean-dummy")
    @Operation(
            summary = "Remove all dummy members",
            description = "Delete all dummy member data",
            tags = {"Dummy"}
    )
    public ResponseEntity<Map<String, String>> cleanDummy() {
        List<MemberDTO> members = memberService.findAll();
        for (MemberDTO member : members) {
            try {
                memberService.deleteById(member.getId());
                log.info("Deleted dummy member: code={} email={}", member.getCode(), member.getEmail());
            } catch (Exception e) {
                log.error("Failed to delete dummy member: code={} email={}", member.getCode(), member.getEmail(), e);
            }
        }

        Map<String, String> response = new HashMap<>();
        response.put("message", "Dummy data cleaned successfully");
        log.info("Dummy members cleanup completed. Total deleted: {}", members.size());
        return ResponseEntity.ok(response);
    }
}
