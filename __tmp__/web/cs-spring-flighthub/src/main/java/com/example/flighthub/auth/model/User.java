package com.example.flighthub.auth.model;

import com.example.flighthub.auth.model.enums.UserStatus;
import com.example.flighthub.auth.model.enums.UserType;
import com.example.flighthub.common.model.BaseDomainModel;
import lombok.EqualsAndHashCode;
import lombok.Getter;
import lombok.Setter;
import lombok.experimental.SuperBuilder;

/**
 * Represents an access and refresh token used for authentication and authorization.
 * This class contains information about the access token, its expiration time, and the refresh token.
 * It also provides utility methods for handling tokens in the form of "Bearer" tokens, which are commonly used
 * in HTTP Authorization headers.
 */
@Getter
@Setter
@SuperBuilder
@EqualsAndHashCode(callSuper = true)
public class User extends BaseDomainModel {

    private String id;

    private String email;

    private String firstName;

    private String lastName;

    private String phoneNumber;

    private UserType userType;

    private UserStatus userStatus;
}
