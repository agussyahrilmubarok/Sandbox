package com.finly.invoice.repository;

import com.finly.customer.Customer;
import com.finly.invoice.Invoice;
import com.finly.user.User;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.query.Param;

import java.util.List;


public interface InvoiceRepository extends JpaRepository<Invoice, Long> {

    // Find invoices by user ID with a search term for the customer name
    @Query("SELECT i FROM Invoice i WHERE i.owner.id = :userId " +
            "AND LOWER(i.customer.name) LIKE LOWER(CONCAT('%', :search, '%'))")
    List<Invoice> searchInvoicesByUserId(@Param("userId") Long userId, @Param("search") String search);

    // Find all invoices by user ID
    List<Invoice> findByOwnerId(Long userId);

    // Find the first invoice by the owner (user)
    Invoice findFirstByOwner(User owner);

    // Find the first invoice by the customer
    Invoice findFirstByCustomer(Customer customer);

    // Count records by ownerId
    long countByOwnerId(Long ownerId);
}
