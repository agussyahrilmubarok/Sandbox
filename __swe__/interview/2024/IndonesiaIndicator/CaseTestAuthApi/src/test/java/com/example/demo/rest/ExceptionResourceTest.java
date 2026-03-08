package com.example.demo.rest;

import com.example.demo.common.Constants;
import com.example.demo.exception.AlreadyExistsException;
import com.example.demo.exception.JwtTokenException;
import com.example.demo.exception.UserNotFoundException;
import com.example.demo.rest.payload.ErrorPayload;
import lombok.extern.slf4j.Slf4j;
import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;
import org.mockito.InjectMocks;
import org.mockito.Mockito;
import org.springframework.core.MethodParameter;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpInputMessage;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.http.converter.HttpMessageNotReadableException;
import org.springframework.validation.BeanPropertyBindingResult;
import org.springframework.validation.BindingResult;
import org.springframework.web.bind.MethodArgumentNotValidException;

import java.lang.reflect.Method;
import java.util.ArrayList;
import java.util.List;
import java.util.Objects;

@Slf4j
class ExceptionResourceTest extends BaseResourceTest {

    @InjectMocks
    private ExceptionResource exceptionResource;

    @Test
    void givenMalformedJsonRequest_whenThrowHttpMessageNotReadableException_thenReturnErrorPayload() {
        HttpInputMessage httpInputMessage = Mockito.mock(HttpInputMessage.class);
        HttpMessageNotReadableException mockException =
                new HttpMessageNotReadableException("Malformed JSON request", httpInputMessage);

        List<String> details = new ArrayList<>();
        details.add(mockException.getMessage());

        ErrorPayload expected = new ErrorPayload();
        expected.setMessage("Malformed JSON request");
        expected.setErrors(details);

        ResponseEntity<Object> result = exceptionResource.handleHttpMessageNotReadable(
                mockException, new HttpHeaders(), HttpStatus.BAD_REQUEST, null);

        assert result != null;
        Assertions.assertEquals(HttpStatus.BAD_REQUEST, result.getStatusCode());
        Assertions.assertEquals(expected.getMessage(),
                ((ErrorPayload) Objects.requireNonNull(result.getBody())).getMessage());
    }

    @Test
    void givenValidationErrors_whenThrowMethodArgumentNotValidException_thenReturnErrorPayload() throws NoSuchMethodException {
        Method method = ExceptionResourceTest.class
                .getDeclaredMethod("givenValidationErrors_whenThrowMethodArgumentNotValidException_thenReturnErrorPayload");
        int parameterIndex = -1;

        MethodParameter mockParameter = new MethodParameter(method, parameterIndex);
        BindingResult mockBindingResult = new BeanPropertyBindingResult(null, null);
        MethodArgumentNotValidException mockException = new MethodArgumentNotValidException(mockParameter, mockBindingResult);

        List<String> errors = mockException.getBindingResult()
                .getFieldErrors()
                .stream()
                .map(error -> error.getDefaultMessage())
                .toList();

        ErrorPayload expected = new ErrorPayload();
        expected.setMessage("Validation Errors");
        expected.setErrors(errors);

        ResponseEntity<Object> result = exceptionResource.handleMethodArgumentNotValid(
                mockException, new HttpHeaders(), HttpStatus.BAD_REQUEST, null);

        assert result != null;
        Assertions.assertEquals(HttpStatus.BAD_REQUEST, result.getStatusCode());
        Assertions.assertEquals(expected.getMessage(),
                ((ErrorPayload) Objects.requireNonNull(result.getBody())).getMessage());
    }

    @Test
    void handleAlreadyExistsException() {
        AlreadyExistsException mockException = new AlreadyExistsException(Constants.EMAIL_ALREADY_EXISTS);

        List<String> errors = new ArrayList<>();
        errors.add(mockException.getMessage());

        ErrorPayload expected = new ErrorPayload();
        expected.setMessage("Already Exists Error");
        expected.setErrors(errors);

        ResponseEntity<ErrorPayload> result = exceptionResource.handleAlreadyExistsException(mockException);

        Assertions.assertEquals(HttpStatus.CONFLICT, result.getStatusCode());
        Assertions.assertEquals(expected.getMessage(),
                Objects.requireNonNull(result.getBody()).getMessage());
    }

    @Test
    void handleUserNotFoundException() {
        UserNotFoundException mockException = new UserNotFoundException("User with username " + "johndoe" + " is not not registered");

        List<String> errors = new ArrayList<>();
        errors.add(mockException.getMessage());

        ErrorPayload expected = new ErrorPayload();
        expected.setMessage("User Not Found");
        expected.setErrors(errors);

        ResponseEntity<ErrorPayload> result = exceptionResource.handleUserNotFoundException(mockException);

        Assertions.assertEquals(HttpStatus.BAD_REQUEST, result.getStatusCode());
        Assertions.assertEquals(expected.getMessage(),
                Objects.requireNonNull(result.getBody()).getMessage());
    }

    @Test
    void handleInvalidJwtTokenException() {
        JwtTokenException mockException = new JwtTokenException("Invalid Token");

        List<String> errors = new ArrayList<>();
        errors.add(mockException.getMessage());

        ErrorPayload expected = new ErrorPayload();
        expected.setMessage("Invalid Token");
        expected.setErrors(errors);

        ResponseEntity<ErrorPayload> result = exceptionResource.handleInvalidJwtTokenException(mockException);

        Assertions.assertEquals(HttpStatus.BAD_REQUEST, result.getStatusCode());
        Assertions.assertEquals(expected.getMessage(),
                Objects.requireNonNull(result.getBody()).getMessage());
    }
}