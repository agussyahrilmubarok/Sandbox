package com.example.kafka.product.products;

import org.springframework.data.jpa.repository.JpaRepository;


public interface ProductRepository extends JpaRepository<Product, String> {

    boolean existsByIdIgnoreCase(String id);

}
