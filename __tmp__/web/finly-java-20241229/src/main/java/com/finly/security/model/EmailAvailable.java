package com.finly.security.model;

import com.finly.user.repository.UserRepository;
import jakarta.validation.Constraint;
import jakarta.validation.ConstraintValidator;
import jakarta.validation.ConstraintValidatorContext;
import jakarta.validation.Payload;

import java.lang.annotation.*;

/**
 * Custom annotation to check if the email is already in use.
 */
@Target(ElementType.FIELD) // Only for fields
@Retention(RetentionPolicy.RUNTIME)
@Documented
@Constraint(validatedBy = EmailAvailable.EmailValidator.class)
public @interface EmailAvailable {

    String message() default "Email is already in use.";

    Class<?>[] groups() default {};

    Class<? extends Payload>[] payload() default {};

    // Validator class for checking email usage
    class EmailValidator implements ConstraintValidator<EmailAvailable, String> {

        private final UserRepository userRepository;

        public EmailValidator(UserRepository userRepository) {
            this.userRepository = userRepository;
        }

        @Override
        public boolean isValid(String email, ConstraintValidatorContext context) {
            // Check if the email is empty or null; if so, it's valid (other validations may handle this)
            if (email == null || email.isEmpty()) {
                return true;
            }

            boolean isValid = true;
            boolean exists = userRepository.existsByEmailIgnoreCase(email);
            if (exists) return !isValid;
            else return isValid;
        }
    }
}
