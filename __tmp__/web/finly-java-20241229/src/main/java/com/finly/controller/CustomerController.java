package com.finly.controller;

import com.finly.customer.model.CustomerDTO;
import com.finly.customer.service.CustomerService;
import jakarta.validation.Valid;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.validation.BindingResult;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.servlet.mvc.support.RedirectAttributes;

import java.util.List;
import java.util.Map;

@Controller
@RequestMapping("/dashboard/customers")
@Slf4j
@RequiredArgsConstructor
public class CustomerController extends BaseController {

    private final CustomerService customerService;

    @GetMapping
    public String index(@RequestParam(value = "search", required = false) String search,
                        final Model model) {
        log.info("Search query received: {}", search);

        List<CustomerDTO> customers = customerService.findCustomersByOwnerAndSearch(getUserId(), search);
        model.addAttribute("title", "Customers");
        model.addAttribute("typePage", "data");
        model.addAttribute("customers", customers);

        return "customer/index";
    }

    @GetMapping("/create")
    public String create(final Model model) {
        model.addAttribute("title", "Create Customer");
        model.addAttribute("typePage", "form");
        model.addAttribute("customer", new CustomerDTO());

        return "customer/index";
    }

    @PostMapping
    public String save(@ModelAttribute("customer") @Valid CustomerDTO customerDTO,
                       BindingResult bindingResult,
                       RedirectAttributes redirectAttributes,
                       final Model model) {

        if (bindingResult.hasErrors()) {
            model.addAttribute("title", "Create Customer");
            model.addAttribute("typePage", "form");
            model.addAttribute("customer", new CustomerDTO());
            return "customer/index";
        }

        customerDTO.setUserId(getUserId());
        customerService.create(customerDTO);
        redirectAttributes.addFlashAttribute("info", Map.of("message", "Customer created successfully!", "type", "success"));

        return "redirect:/dashboard/customers";
    }

    @GetMapping("/{id}/edit")
    public String edit(@PathVariable("id") Long id,
                       final Model model) {
        model.addAttribute("title", "Edit Customer");
        model.addAttribute("typePage", "form");
        model.addAttribute("customer", customerService.get(id));

        return "customer/index";
    }

    @PostMapping("/{id}")
    public String update(@PathVariable("id") Long id,
                         @ModelAttribute("customer") @Valid CustomerDTO customerDTO,
                         BindingResult bindingResult,
                         RedirectAttributes redirectAttributes,
                         final Model model) {

        if (bindingResult.hasErrors()) {
            model.addAttribute("title", "Edit Customer");
            model.addAttribute("typePage", "form");
            model.addAttribute("customer", customerDTO);
            return "customer/index";
        }

        customerDTO.setUserId(getUserId());
        customerService.update(id, customerDTO);
        redirectAttributes.addFlashAttribute("info", Map.of("message", "Customer updated successfully!", "type", "success"));

        return "redirect:/dashboard/customers";
    }

    @PostMapping("/{id}/delete")
    public String delete(@PathVariable("id") Long id,
                         RedirectAttributes redirectAttributes) {
        try {
            customerService.delete(id);
            log.info("Customer with ID {} has been deleted.", id);
            redirectAttributes.addFlashAttribute("info", Map.of("message", "Customer deleted successfully!", "type", "success"));
        } catch (Exception e) {
            log.error("Error deleting customer with ID {}", id, e);
            redirectAttributes.addFlashAttribute("info", Map.of("message", "Error deleting customer. Please try again.", "type", "error"));
        }

        return "redirect:/dashboard/customers";
    }
}
