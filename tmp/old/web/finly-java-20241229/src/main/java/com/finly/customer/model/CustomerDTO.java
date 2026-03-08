package com.finly.customer.model;

import jakarta.validation.constraints.NotNull;
import jakarta.validation.constraints.Size;
import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
public class CustomerDTO {

    private Long id;

    @NotNull(message = "Name cannot be empty")
    @Size(max = 255, message = "Name must not exceed 255 characters")
    private String name;

    @NotNull(message = "Email cannot be empty")
    @Size(max = 255, message = "Email must not exceed 255 characters")
    @CustomerEmailUnique(message = "This email is already in use")
    private String email;

    @NotNull(message = "Phone number cannot be empty")
    @Size(max = 15, message = "Phone number must not exceed 15 characters")
    @CustomerPhoneUnique(message = "This phone number is already in use")
    private String phone;

    @NotNull(message = "Address cannot be empty")
    private String address;

    private Long userId;
}
