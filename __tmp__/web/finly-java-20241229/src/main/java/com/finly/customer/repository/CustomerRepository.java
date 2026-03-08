package com.finly.customer.repository;

import com.finly.customer.Customer;
import com.finly.user.User;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;

import java.util.List;

public interface CustomerRepository extends JpaRepository<Customer, Long> {

    // Search customers by user ID and a keyword in name, email, phone, or address
    @Query("SELECT c FROM Customer c WHERE c.user.id = :userId " +
            "AND (LOWER(c.name) LIKE LOWER(CONCAT('%', :search, '%')) " +
            "OR LOWER(c.email) LIKE LOWER(CONCAT('%', :search, '%')) " +
            "OR LOWER(c.phone) LIKE LOWER(CONCAT('%', :search, '%')) " +
            "OR LOWER(c.address) LIKE LOWER(CONCAT('%', :search, '%')))")
    List<Customer> searchCustomersByUserId(Long userId, String search);

    // Retrieve all customers by user ID
    List<Customer> findByUserId(Long userId);

    // Retrieve the first customer by the associated user
    Customer findFirstByUser(User user);

    // Check if an email exists (case-insensitive)
    boolean existsByEmailIgnoreCase(String email);

    // Check if a phone number exists (case-insensitive)
    boolean existsByPhoneIgnoreCase(String phone);

    // Count records by userId
    long countByUserId(Long userId);
}
