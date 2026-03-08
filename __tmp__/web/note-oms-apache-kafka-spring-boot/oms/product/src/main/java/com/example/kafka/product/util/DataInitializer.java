package com.example.kafka.product.util;

import com.example.kafka.product.products.Product;
import com.example.kafka.product.products.ProductRepository;
import org.springframework.boot.CommandLineRunner;
import org.springframework.stereotype.Component;

import java.math.BigDecimal;
import java.time.OffsetDateTime;
import java.util.ArrayList;
import java.util.List;

@Component
public class DataInitializer implements CommandLineRunner {

    private final ProductRepository productRepository;

    public DataInitializer(ProductRepository productRepository) {
        this.productRepository = productRepository;
    }

    @Override
    public void run(String... args) {
        List<Product> productsToSave = new ArrayList<>();

        List<Product> products = List.of(
                createProduct("PROD-0000000001", "Laptop", new BigDecimal("15000.00"), 10),
                createProduct("PROD-0000000002", "Mouse", new BigDecimal("300.00"), 50),
                createProduct("PROD-0000000003", "Keyboard", new BigDecimal("500.00"), 40),
                createProduct("PROD-0000000004", "Monitor", new BigDecimal("2000.00"), 15),
                createProduct("PROD-0000000005", "Printer", new BigDecimal("2500.00"), 8),
                createProduct("PROD-0000000006", "Smartphone", new BigDecimal("12000.00"), 20),
                createProduct("PROD-0000000007", "Tablet", new BigDecimal("8000.00"), 12),
                createProduct("PROD-0000000008", "Headset", new BigDecimal("700.00"), 30),
                createProduct("PROD-0000000009", "Webcam", new BigDecimal("900.00"), 25),
                createProduct("PROD-0000000010", "External HDD", new BigDecimal("1500.00"), 0)
        );

        for (Product product : products) {
            if (productRepository.findById(product.getId()).isEmpty()) {
                productsToSave.add(product);
            }
        }

        if (!productsToSave.isEmpty()) {
            productRepository.saveAll(productsToSave);
        }
    }

    private Product createProduct(String id, String name, BigDecimal price, int stock) {
        Product product = new Product();
        product.setId(id);
        product.setName(name);
        product.setPrice(price);
        product.setStock(stock);
        product.setDateCreated(OffsetDateTime.now());
        product.setLastUpdated(OffsetDateTime.now());
        return product;
    }
}
