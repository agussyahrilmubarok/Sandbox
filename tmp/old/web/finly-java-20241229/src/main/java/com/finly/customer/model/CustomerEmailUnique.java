package com.finly.customer.model;

import com.finly.customer.service.CustomerService;
import jakarta.servlet.http.HttpServletRequest;
import jakarta.validation.Constraint;
import jakarta.validation.ConstraintValidator;
import jakarta.validation.ConstraintValidatorContext;
import jakarta.validation.Payload;
import org.springframework.web.servlet.HandlerMapping;

import java.lang.annotation.Documented;
import java.lang.annotation.Retention;
import java.lang.annotation.RetentionPolicy;
import java.lang.annotation.Target;
import java.util.Map;

import static java.lang.annotation.ElementType.*;


/**
 * Validate that the email value isn't taken yet.
 */
@Target({ FIELD, METHOD, ANNOTATION_TYPE })
@Retention(RetentionPolicy.RUNTIME)
@Documented
@Constraint(validatedBy = CustomerEmailUnique.CustomerEmailUniqueValidator.class)
public @interface CustomerEmailUnique {

    String message() default "{Exists.customer.email}";

    Class<?>[] groups() default {};

    Class<? extends Payload>[] payload() default {};

    class CustomerEmailUniqueValidator implements ConstraintValidator<CustomerEmailUnique, String> {

        private final CustomerService customerService;
        private final HttpServletRequest request;

        public CustomerEmailUniqueValidator(final CustomerService customerService,
                final HttpServletRequest request) {
            this.customerService = customerService;
            this.request = request;
        }

        @Override
        public boolean isValid(final String value, final ConstraintValidatorContext cvContext) {
            if (value == null) {
                // no value present
                return true;
            }
            @SuppressWarnings("unchecked") final Map<String, String> pathVariables =
                    ((Map<String, String>)request.getAttribute(HandlerMapping.URI_TEMPLATE_VARIABLES_ATTRIBUTE));
            final String currentId = pathVariables.get("id");
            if (currentId != null && value.equalsIgnoreCase(customerService.get(Long.parseLong(currentId)).getEmail())) {
                // value hasn't changed
                return true;
            }
            return !customerService.emailExists(value);
        }
    }
}
