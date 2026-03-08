package com.finly.customer.service;

import com.finly.customer.Customer;
import com.finly.customer.model.CustomerDTO;
import com.finly.customer.repository.CustomerRepository;
import com.finly.exception.NotFoundException;
import com.finly.invoice.Invoice;
import com.finly.invoice.repository.InvoiceRepository;
import com.finly.user.repository.UserRepository;
import com.finly.util.ReferencedWarning;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.data.domain.Sort;
import org.springframework.stereotype.Service;

import java.util.List;


@Service
@Slf4j
@RequiredArgsConstructor
public class CustomerService {

    private final CustomerRepository customerRepository;
    private final UserRepository userRepository;
    private final InvoiceRepository invoiceRepository;

    // Retrieve all customers, sorted by ID
    public List<CustomerDTO> findAll() {
        log.info("Fetching all customers...");
        return customerRepository.findAll(Sort.by("id")).stream()
                .map(customer -> mapToDTO(customer, new CustomerDTO()))
                .toList();
    }

    // Search customers by user ID and search term
    public List<CustomerDTO> findCustomersByOwnerAndSearch(Long userId, String search) {
        log.info("Searching customers for user {} with term '{}'", userId, search);
        List<Customer> customers = (search != null && !search.isEmpty())
                ? customerRepository.searchCustomersByUserId(userId, search)
                : customerRepository.findByUserId(userId);
        return customers.stream()
                .map(customer -> mapToDTO(customer, new CustomerDTO()))
                .toList();
    }

    // Retrieve customers by user ID
    public List<CustomerDTO> findByUserId(Long userId) {
        log.info("Fetching customers for user {}", userId);
        return customerRepository.findByUserId(userId).stream()
                .map(customer -> mapToDTO(customer, new CustomerDTO()))
                .toList();
    }

    // Get a single customer by ID or throw exception
    public CustomerDTO get(Long id) {
        log.info("Fetching customer with id {}", id);
        return customerRepository.findById(id)
                .map(customer -> mapToDTO(customer, new CustomerDTO()))
                .orElseThrow(() -> {
                    log.error("Customer with id {} not found", id);
                    return new NotFoundException();
                });
    }

    // Create a new customer and return its ID
    public Long create(CustomerDTO customerDTO) {
        log.info("Creating new customer with email {}", customerDTO.getEmail());
        Customer customer = mapToEntity(customerDTO, new Customer());
        return customerRepository.save(customer).getId();
    }

    // Update an existing customer by ID
    public void update(Long id, CustomerDTO customerDTO) {
        log.info("Updating customer with id {}", id);
        Customer customer = customerRepository.findById(id)
                .orElseThrow(() -> {
                    log.error("Customer with id {} not found for update", id);
                    return new NotFoundException();
                });
        customerRepository.save(mapToEntity(customerDTO, customer));
        log.info("Customer with id {} updated successfully", id);
    }

    // Delete a customer by ID
    public void delete(Long id) {
        log.info("Deleting customer with id {}", id);
        customerRepository.deleteById(id);
    }

    // Count customer by user
    public long countByUser(Long userId) {
        return customerRepository.countByUserId(userId);
    }

    // Check if email exists
    public boolean emailExists(String email) {
        log.info("Checking if email {} exists", email);
        boolean exists = customerRepository.existsByEmailIgnoreCase(email);
        log.info("Email {} {}", email, exists ? "exists" : "does not exist");
        return exists;
    }

    // Check if phone exists
    public boolean phoneExists(String phone) {
        log.info("Checking if phone {} exists", phone);
        boolean exists = customerRepository.existsByPhoneIgnoreCase(phone);
        log.info("Phone {} {}", phone, exists ? "exists" : "does not exist");
        return exists;
    }

    // Check for referenced warnings for a customer
    public ReferencedWarning getReferencedWarning(Long id) {
        log.info("Checking referenced warnings for customer with id {}", id);
        Customer customer = customerRepository.findById(id)
                .orElseThrow(() -> {
                    log.error("Customer with id {} not found for referenced warning check", id);
                    return new NotFoundException();
                });

        ReferencedWarning warning = new ReferencedWarning();
        Invoice referencedInvoice = invoiceRepository.findFirstByCustomer(customer);
        if (referencedInvoice != null) {
            warning.setKey("customer.invoice.customer.referenced");
            warning.addParam(referencedInvoice.getId());
            log.warn("Customer with id {} has a referenced invoice id {}", id, referencedInvoice.getId());
        }
        return warning;
    }

    // Map Customer entity to DTO
    private CustomerDTO mapToDTO(Customer customer, CustomerDTO customerDTO) {
        customerDTO.setId(customer.getId());
        customerDTO.setName(customer.getName());
        customerDTO.setEmail(customer.getEmail());
        customerDTO.setPhone(customer.getPhone());
        customerDTO.setAddress(customer.getAddress());
        customerDTO.setUserId(customer.getUser() == null ? null : customer.getUser().getId());
        return customerDTO;
    }

    // Map DTO to Customer entity
    private Customer mapToEntity(CustomerDTO customerDTO, Customer customer) {
        customer.setName(customerDTO.getName());
        customer.setEmail(customerDTO.getEmail());
        customer.setPhone(customerDTO.getPhone());
        customer.setAddress(customerDTO.getAddress());
        if (customerDTO.getUserId() != null) {
            customer.setUser(userRepository.findById(customerDTO.getUserId())
                    .orElseThrow(() -> {
                        log.error("User with id {} not found", customerDTO.getUserId());
                        return new NotFoundException("User not found");
                    }));
        }
        return customer;
    }
}