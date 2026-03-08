package com.finly.invoice.service;

import com.finly.customer.Customer;
import com.finly.customer.model.CustomerDTO;
import com.finly.customer.repository.CustomerRepository;
import com.finly.exception.NotFoundException;
import com.finly.invoice.Invoice;
import com.finly.invoice.model.InvoiceDTO;
import com.finly.invoice.model.InvoiceDetailDTO;
import com.finly.invoice.repository.InvoiceRepository;
import com.finly.user.User;
import com.finly.user.model.UserDTO;
import com.finly.user.repository.UserRepository;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.data.domain.Sort;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;


@Service
@Slf4j
@RequiredArgsConstructor
public class InvoiceService {

    private final InvoiceRepository invoiceRepository;
    private final UserRepository userRepository;
    private final CustomerRepository customerRepository;

    // Fetch all invoices with sorting by ID
    public List<InvoiceDTO> findAll() {
        log.info("Fetching all invoices...");
        return invoiceRepository.findAll(Sort.by("id"))
                .stream()
                .map(invoice -> mapToDTO(invoice, new InvoiceDTO()))
                .toList();
    }

    // Fetch invoices by user ID and optional search term
    @Transactional
    public List<InvoiceDetailDTO> findInvoiceByOwnerAndSearch(Long userId, String search) {
        log.info("Searching invoices for user {} with term '{}'", userId, search);
        List<Invoice> invoices;
        invoices = (search != null && !search.isEmpty())
                ? invoiceRepository.searchInvoicesByUserId(userId, search)
                : invoiceRepository.findByOwnerId(userId);
        return invoices.stream()
                .map(invoice -> mapToDetailDTO(invoice, new InvoiceDetailDTO()))
                .toList();
    }

    // Retrieve an invoice by ID
    public InvoiceDTO get(Long id) {
        log.info("Fetching invoice with ID {}", id);
        return invoiceRepository.findById(id)
                .map(invoice -> mapToDTO(invoice, new InvoiceDTO()))
                .orElseThrow(() -> new NotFoundException("Invoice not found"));
    }

    // Retrieve detailed invoice information by ID
    @Transactional
    public InvoiceDetailDTO getDetail(Long id) {
        log.info("Fetching detailed invoice for ID {}", id);
        return invoiceRepository.findById(id)
                .map(invoice -> mapToDetailDTO(invoice, new InvoiceDetailDTO()))
                .orElseThrow(() -> new NotFoundException("Invoice details not found"));
    }

    // Create a new invoice
    public Long create(InvoiceDTO invoiceDTO) {
        log.info("Creating a new invoice...");
        Invoice invoice = mapToEntity(invoiceDTO, new Invoice());
        return invoiceRepository.save(invoice).getId();
    }

    // Update an existing invoice by ID
    public void update(Long id, InvoiceDTO invoiceDTO) {
        log.info("Updating invoice with ID {}", id);
        Invoice invoice = invoiceRepository.findById(id)
                .orElseThrow(() -> new NotFoundException("Invoice not found for update"));
        invoiceRepository.save(mapToEntity(invoiceDTO, invoice));
        log.info("Invoice with ID {} updated successfully", id);
    }

    // Delete an invoice by ID
    public void delete(Long id) {
        log.info("Deleting invoice with ID {}", id);
        invoiceRepository.deleteById(id);
        log.info("Invoice with ID {} deleted successfully", id);
    }

    // Count invoice by owner
    public long countByOwner(Long ownerId) {
        return invoiceRepository.countByOwnerId(ownerId);
    }

    // Map Invoice to Detail DTO
    private InvoiceDTO mapToDTO(final Invoice invoice, final InvoiceDTO invoiceDTO) {
        log.debug("Mapping Invoice with ID {} to DTO.", invoice.getId());
        invoiceDTO.setId(invoice.getId());
        invoiceDTO.setAmount(invoice.getAmount());
        invoiceDTO.setDueDate(invoice.getDueDate());
        invoiceDTO.setStatus(invoice.getStatus());
        invoiceDTO.setOwner(invoice.getOwner() == null ? null : invoice.getOwner().getId());
        invoiceDTO.setCustomer(invoice.getCustomer() == null ? null : invoice.getCustomer().getId());

        return invoiceDTO;
    }

    // Map Invoice to Detail DTO
    private InvoiceDetailDTO mapToDetailDTO(final Invoice invoice, final InvoiceDetailDTO invoiceDetailDTO) {
        log.debug("Mapping detailed Invoice with ID {} to DetailDTO.", invoice.getId());
        invoiceDetailDTO.setId(invoice.getId());
        invoiceDetailDTO.setAmount(invoice.getAmount());
        invoiceDetailDTO.setDueDate(invoice.getDueDate());
        invoiceDetailDTO.setStatus(invoice.getStatus());

        if (invoice.getOwner() != null) {
            invoiceDetailDTO.setOwner(mapToDTO(invoice.getOwner(), new UserDTO()));
        }

        if (invoice.getCustomer() != null) {
            invoiceDetailDTO.setCustomer(mapToDTO(invoice.getCustomer(), new CustomerDTO()));
        }

        return invoiceDetailDTO;
    }

    // Map User to DTO
    private UserDTO mapToDTO(final User user, final UserDTO userDTO) {
        log.debug("Mapping User with ID {} to DTO.", user.getId());
        userDTO.setId(user.getId());
        userDTO.setEmail(user.getEmail());
        userDTO.setPassword(user.getPassword());
        return userDTO;
    }

    // Map Customer to DTO
    private CustomerDTO mapToDTO(final Customer customer, final CustomerDTO customerDTO) {
        log.debug("Mapping Customer with ID {} to DTO.", customer.getId());
        customerDTO.setId(customer.getId());
        customerDTO.setName(customer.getName());
        customerDTO.setEmail(customer.getEmail());
        customerDTO.setPhone(customer.getPhone());
        customerDTO.setAddress(customer.getAddress());
        customerDTO.setUserId(customer.getUser() == null ? null : customer.getUser().getId());
        return customerDTO;
    }

    // Map DTO to Invoice entity
    private Invoice mapToEntity(InvoiceDTO invoiceDTO, Invoice invoice) {
        invoice.setAmount(invoiceDTO.getAmount());
        invoice.setDueDate(invoiceDTO.getDueDate());
        invoice.setStatus(invoiceDTO.getStatus());

        invoice.setOwner(fetchUser(invoiceDTO.getOwner()));
        invoice.setCustomer(fetchCustomer(invoiceDTO.getCustomer()));

        return invoice;
    }

    // Helper method to fetch User by ID
    private User fetchUser(Long userId) {
        if (userId == null) return null;
        return userRepository.findById(userId)
                .orElseThrow(() -> new NotFoundException("User not found"));
    }

    // Helper method to fetch Customer by ID
    private Customer fetchCustomer(Long customerId) {
        if (customerId == null) return null;
        return customerRepository.findById(customerId)
                .orElseThrow(() -> new NotFoundException("Customer not found"));
    }


}
