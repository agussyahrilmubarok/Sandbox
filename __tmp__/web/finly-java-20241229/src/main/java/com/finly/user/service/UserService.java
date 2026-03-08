package com.finly.user.service;

import com.finly.customer.Customer;
import com.finly.customer.repository.CustomerRepository;
import com.finly.invoice.Invoice;
import com.finly.invoice.repository.InvoiceRepository;
import com.finly.user.User;
import com.finly.user.model.UserDTO;
import com.finly.user.repository.UserRepository;
import com.finly.exception.NotFoundException;
import com.finly.util.ReferencedWarning;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.data.domain.Sort;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;

import java.util.List;


@Service
@Slf4j
@RequiredArgsConstructor
public class UserService {

    private final UserRepository userRepository;
    private final CustomerRepository customerRepository;
    private final InvoiceRepository invoiceRepository;
    private final PasswordEncoder passwordEncoder;

    // Retrieve all users and log results
    public List<UserDTO> findAll() {
        log.info("Fetching all users...");
        return userRepository.findAll(Sort.by("id"))
                .stream()
                .map(user -> mapToDTO(user, new UserDTO()))
                .toList();
    }

    // Retrieve user by ID or throw exception
    public UserDTO get(Long id) {
        log.info("Fetching user with id {}", id);
        return userRepository.findById(id)
                .map(user -> mapToDTO(user, new UserDTO()))
                .orElseThrow(() -> {
                    log.error("User with id {} not found", id);
                    return new NotFoundException("User is not found");
                });
    }

    // Create a new user and return its ID
    public Long create(UserDTO userDTO) {
        log.info("Creating user with email {}", userDTO.getEmail());
        User user = mapToEntity(userDTO, new User());
        return userRepository.save(user).getId();
    }

    // Update an existing user by ID
    public void update(Long id, UserDTO userDTO) {
        log.info("Updating user with id {}", id);
        User user = userRepository.findById(id)
                .orElseThrow(() -> {
                    log.error("User with id {} not found for update", id);
                    return new NotFoundException();
                });
        userRepository.save(mapToEntity(userDTO, user));
        log.info("User with id {} updated successfully", id);
    }

    // Delete a user by ID
    public void delete(Long id) {
        log.info("Deleting user with id {}", id);
        userRepository.deleteById(id);
    }

    // Check if an email exists and log the result
    public boolean emailExists(String email) {
        log.info("Checking if email {} exists", email);
        boolean exists = userRepository.existsByEmailIgnoreCase(email);
        log.info("Email {} {}", email, exists ? "exists" : "does not exist");
        return exists;
    }

    // Check for referenced warnings related to a user
    public ReferencedWarning getReferencedWarning(Long id) {
        log.info("Checking referenced warnings for user with id {}", id);
        User user = userRepository.findById(id)
                .orElseThrow(() -> new NotFoundException());

        ReferencedWarning warning = new ReferencedWarning();
        Customer referencedCustomer = customerRepository.findFirstByUser(user);
        if (referencedCustomer != null) {
            warning.setKey("user.customer.user.referenced");
            warning.addParam(referencedCustomer.getId());
            log.warn("User with id {} has referenced customer id {}", id, referencedCustomer.getId());
            return warning;
        }

        Invoice referencedInvoice = invoiceRepository.findFirstByOwner(user);
        if (referencedInvoice != null) {
            warning.setKey("user.invoice.owner.referenced");
            warning.addParam(referencedInvoice.getId());
            log.warn("User with id {} has referenced invoice id {}", id, referencedInvoice.getId());
            return warning;
        }

        log.info("No referenced warnings for user with id {}", id);
        return null;
    }

    // Map User entity to DTO
    private UserDTO mapToDTO(User user, UserDTO userDTO) {
        userDTO.setId(user.getId());
        userDTO.setEmail(user.getEmail());
        userDTO.setPassword(user.getPassword());
        return userDTO;
    }

    // Map UserDTO to entity and encode password
    private User mapToEntity(UserDTO userDTO, User user) {
        user.setEmail(userDTO.getEmail());
        user.setPassword(passwordEncoder.encode(userDTO.getPassword()));
        return user;
    }
}
