package com.finly.controller;

import com.finly.customer.service.CustomerService;
import com.finly.invoice.model.InvoiceDTO;
import com.finly.invoice.service.InvoiceService;
import jakarta.validation.Valid;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.validation.BindingResult;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.servlet.mvc.support.RedirectAttributes;

import java.util.Map;

@Controller
@RequestMapping("/dashboard/invoices")
@Slf4j
@RequiredArgsConstructor
public class InvoiceController extends BaseController {

    private final InvoiceService invoiceService;
    private final CustomerService customerService;

    @GetMapping
    public String index(@RequestParam(value = "search", required = false) String search,
                        final Model model) {
        log.info("Fetching all invoices for the user with ID: {}", getUserId());

        model.addAttribute("title", "Invoices");
        model.addAttribute("typePage", "data");
        model.addAttribute("invoices", invoiceService.findInvoiceByOwnerAndSearch(getUserId(), search));

        log.info("Invoices fetched successfully");
        return "invoice/index";
    }

    @GetMapping("/create")
    public String create(final Model model) {
        log.info("Rendering the create invoice form for user with ID: {}", getUserId());

        model.addAttribute("title", "Create Invoice");
        model.addAttribute("typePage", "form");
        model.addAttribute("customers", customerService.findByUserId(getUserId()));
        model.addAttribute("invoice", new InvoiceDTO());

        log.info("Customers data and invoice form initialized");
        return "invoice/index";
    }

    @PostMapping
    public String save(@ModelAttribute("invoice") @Valid InvoiceDTO invoiceDTO,
                       BindingResult bindingResult,
                       RedirectAttributes redirectAttributes,
                       final Model model) {
        log.info("Attempting to save a new invoice for user with ID: {}", getUserId());

        if (bindingResult.hasErrors()) {
            log.warn("Validation errors occurred while saving the invoice: {}", bindingResult.getAllErrors());
            model.addAttribute("title", "Create Invoice");
            model.addAttribute("typePage", "form");
            model.addAttribute("customers", customerService.findByUserId(getUserId()));
            model.addAttribute("invoice", new InvoiceDTO());
            return "invoice/index";
        }

        invoiceDTO.setOwner(getUserId());
        invoiceService.create(invoiceDTO);
        redirectAttributes.addFlashAttribute("info", Map.of("message", "Invoice created successfully!", "type", "success"));

        log.info("Invoice created successfully with ID: {}", invoiceDTO.getId());
        return "redirect:/dashboard/invoices";
    }

    @GetMapping("/{id}/edit")
    public String edit(@PathVariable("id") Long id, final Model model) {
        log.info("Fetching invoice details for invoice ID: {}", id);

        model.addAttribute("title", "Edit Invoice");
        model.addAttribute("typePage", "form");
        model.addAttribute("customers", customerService.findByUserId(getUserId()));
        model.addAttribute("invoice", invoiceService.getDetail(id));

        log.info("Invoice details fetched successfully for invoice ID: {}", id);
        return "invoice/index";
    }

    @PostMapping("/{id}")
    public String update(@PathVariable("id") Long id,
                         @ModelAttribute("invoice") @Valid InvoiceDTO invoiceDTO,
                         BindingResult bindingResult,
                         RedirectAttributes redirectAttributes,
                         final Model model) {
        log.info("Attempting to update invoice with ID: {}", id);

        if (bindingResult.hasErrors()) {
            log.warn("Validation errors occurred while updating the invoice: {}", bindingResult.getAllErrors());
            model.addAttribute("title", "Edit Invoice");
            model.addAttribute("typePage", "form");
            model.addAttribute("invoice", invoiceDTO);
            return "invoice/index";
        }

        invoiceDTO.setOwner(getUserId());
        invoiceService.update(id, invoiceDTO);
        redirectAttributes.addFlashAttribute("info", Map.of("message", "Invoice updated successfully!", "type", "success"));
        log.info("Invoice updated successfully with ID: {}", id);

        return "redirect:/dashboard/invoices";
    }

    @PostMapping("/{id}/delete")
    public String delete(@PathVariable("id") Long id, RedirectAttributes redirectAttributes) {
        try {
            log.info("Attempting to delete invoice with ID: {}", id);
            invoiceService.delete(id);
            log.info("Invoice with ID {} has been deleted.", id);
            redirectAttributes.addFlashAttribute("info", Map.of("message", "Invoice deleted successfully!", "type", "success"));
        } catch (Exception e) {
            log.error("Error deleting invoice with ID {}", id, e);
            redirectAttributes.addFlashAttribute("info", Map.of("message", "Error deleting invoice. Please try again.", "type", "error"));
        }
        return "redirect:/dashboard/invoices";
    }
}
