package com.finly.user.model;

import jakarta.validation.constraints.NotNull;
import jakarta.validation.constraints.Size;
import lombok.Getter;
import lombok.Setter;


@Getter
@Setter
public class UserDTO {

    private Long id;

    @NotNull(message = "Email cannot be empty")
    @Size(max = 255, message = "Email must not exceed 255 characters")
    @UserEmailUnique(message = "This email is already registered")
    private String email;

    @NotNull(message = "Password cannot be empty")
    @Size(max = 255, message = "Password must not exceed 255 characters")
    private String password;
}
