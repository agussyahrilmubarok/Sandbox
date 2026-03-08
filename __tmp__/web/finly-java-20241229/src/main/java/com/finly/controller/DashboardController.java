package com.finly.controller;

import com.finly.customer.service.CustomerService;
import com.finly.invoice.model.InvoiceDetailDTO;
import com.finly.invoice.service.InvoiceService;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.GetMapping;

import java.text.NumberFormat;
import java.time.LocalDate;
import java.util.*;
import java.util.stream.Collectors;

@Controller
@Slf4j
@RequiredArgsConstructor
public class DashboardController extends BaseController {

    private final CustomerService customerService;
    private final InvoiceService invoiceService;

    @GetMapping("/dashboard")
    public String index(final Model model) {
        if (!isLoggedIn()) return "redirect:/login";

        var userId = getUserId();
        var customerCount = customerService.countByUser(userId);
        var invoiceCount = invoiceService.countByOwner(userId);
        log.debug("Counts retrieved for dashboard", Map.of(
                "userId", userId,
                "invoiceCount", invoiceCount,
                "customerCount", customerCount
        ));

        List<InvoiceDetailDTO> allInvoices = invoiceService.findInvoiceByOwnerAndSearch(userId, "");
        log.debug("Invoices retrieved and populated", Map.of(
                "userId", userId,
                "invoiceCount", allInvoices.size()
        ));

        double totalPaid = allInvoices.stream()
                .filter(invoice -> "paid".equalsIgnoreCase(invoice.getStatus()))
                .mapToDouble(i -> i.getAmount().doubleValue())
                .sum();
        double totalPending = allInvoices.stream()
                .filter(invoice -> "pending".equalsIgnoreCase(invoice.getStatus()))
                .mapToDouble(i -> i.getAmount().doubleValue())
                .sum();
        log.debug("Calculated totals", Map.of(
                "userId", userId,
                "totalPaid", totalPaid,
                "totalPending", totalPending
        ));

        List<InvoiceDetailDTO> latestInvoices = allInvoices.stream()
                .sorted(Comparator.comparing(InvoiceDetailDTO::getDueDate).reversed())
                .limit(5)
                .collect(Collectors.toList());

        // Count revenue from 6 month latest
        List<Map<String, Object>> revenueData = new ArrayList<>();
        for (int i = 0; i < 6; i++) {
            LocalDate today = LocalDate.now().minusMonths(i);
            String month = today.getMonth().toString().substring(0, 3);
            double revenueForMonth = allInvoices.stream()
                    .filter(invoice -> {
                        LocalDate invoiceDueDate = convertToLocalDate(invoice.getDueDate());
                        return invoiceDueDate.getMonth() == today.getMonth();
                    })
                    .mapToDouble(a -> a.getAmount().doubleValue())
                    .sum();

            Map<String, Object> revenueEntry = new HashMap<>();
            revenueEntry.put("month", month);
            revenueEntry.put("revenue", revenueForMonth);
            revenueData.add(0, revenueEntry);
        }
        log.info(revenueData.toString());

        model.addAttribute("title", "Dashboard");
        model.addAttribute("latestInvoices", latestInvoices);
        model.addAttribute("revenueData", revenueData);
        model.addAttribute("invoiceCount", invoiceCount);
        model.addAttribute("customerCount", customerCount);
        model.addAttribute("totalPaid", totalPaid);
        model.addAttribute("totalPending", totalPending);

        return "dashboard/index";
    }


    private LocalDate convertToLocalDate(Date date) {
        return date.toInstant()
                .atZone(java.time.ZoneId.systemDefault())
                .toLocalDate();
    }

    private String formatUSD(double amount) {
        NumberFormat usdFormatter = NumberFormat.getCurrencyInstance(Locale.US);
        return usdFormatter.format(amount);
    }
}
