package com.finly.user.repository;

import com.finly.user.User;
import org.springframework.data.jpa.repository.JpaRepository;

import java.util.Optional;


public interface UserRepository extends JpaRepository<User, Long> {

    // Find a user by email
    Optional<User> findByEmail(String email);

    // Check if an email exists (case-insensitive)
    boolean existsByEmailIgnoreCase(String email);
}
