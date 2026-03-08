package com.example.kafka.product.products;

import java.util.List;

public interface ProductService {

    List<ProductDTO> findAll();

    ProductDTO create(ProductDTO dto);
}
