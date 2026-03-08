package com.example.demo.rest.payload;

import com.example.demo.common.Constants;
import jakarta.validation.constraints.NotEmpty;
import jakarta.validation.constraints.Pattern;
import jakarta.validation.constraints.Size;
import lombok.Data;

@Data
public class LoginPayload {
    @NotEmpty(message = Constants.USERNAME_REQUIRED)
    @Size(min = 3, max = 50, message = Constants.USERNAME_LENGTH)
    private String username;

    @NotEmpty(message = Constants.PASSWORD_REQUIRED)
    @Size(min = 8, message = Constants.PASSWORD_MIN_LENGTH)
    @Pattern(regexp = "^(?=.*[0-9])(?=.*[a-zA-Z]).*$", message = Constants.PASSWORD_FORMAT)
    private String password;
}
