package com.example.demo.rest;

import com.example.demo.exception.AlreadyExistsException;
import com.example.demo.exception.JwtTokenException;
import com.example.demo.exception.UserNotFoundException;
import com.example.demo.rest.payload.ErrorPayload;
import lombok.extern.slf4j.Slf4j;
import org.springframework.context.support.DefaultMessageSourceResolvable;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpStatus;
import org.springframework.http.HttpStatusCode;
import org.springframework.http.ResponseEntity;
import org.springframework.http.converter.HttpMessageNotReadableException;
import org.springframework.web.bind.MethodArgumentNotValidException;
import org.springframework.web.bind.annotation.ControllerAdvice;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.context.request.WebRequest;
import org.springframework.web.servlet.mvc.method.annotation.ResponseEntityExceptionHandler;

import java.util.ArrayList;
import java.util.Collections;
import java.util.List;

@ControllerAdvice
@Slf4j
public class ExceptionResource extends ResponseEntityExceptionHandler {

    @Override
    protected ResponseEntity<Object> handleHttpMessageNotReadable(HttpMessageNotReadableException ex,
                                                                  HttpHeaders headers,
                                                                  HttpStatusCode status,
                                                                  WebRequest request) {
        log.error(ex.getMessage(), ex);

        List<String> errors = new ArrayList<>();
        errors.add(ex.getMessage());

        ErrorPayload body = new ErrorPayload();
        body.setMessage("Malformed JSON request");
        body.setErrors(errors);

        return ResponseEntity.status(status).body(body);
    }


    @Override
    protected ResponseEntity<Object> handleMethodArgumentNotValid(MethodArgumentNotValidException ex,
                                                                  HttpHeaders headers,
                                                                  HttpStatusCode status,
                                                                  WebRequest request) {
        log.error(ex.getMessage(), ex);

        List<String> errors = ex.getBindingResult()
                .getFieldErrors()
                .stream()
                .map(DefaultMessageSourceResolvable::getDefaultMessage)
                .toList();

        ErrorPayload body = new ErrorPayload();
        body.setMessage("Validation Errors");
        body.setErrors(errors);

        return ResponseEntity.status(status).body(body);
    }


    @ExceptionHandler(UserNotFoundException.class)
    protected ResponseEntity<ErrorPayload> handleUserNotFoundException(UserNotFoundException ex) {
        log.error(ex.getMessage(), ex);
        List<String> errors = new ArrayList<>();
        errors.add(ex.getMessage());

        ErrorPayload body = new ErrorPayload();
        body.setMessage("User Not Found");
        body.setErrors(errors);

        return ResponseEntity.status(HttpStatus.BAD_REQUEST).body(body);
    }

    @ExceptionHandler(value = {AlreadyExistsException.class})
    protected ResponseEntity<ErrorPayload> handleAlreadyExistsException(final AlreadyExistsException ex) {
        log.error(ex.getMessage(), ex);

        List<String> errors = new ArrayList<>();
        errors.add(ex.getMessage());

        ErrorPayload body = new ErrorPayload();
        body.setMessage("Already Exists Error");
        body.setErrors(errors);

        return ResponseEntity.status(HttpStatus.CONFLICT).body(body);
    }

    @ExceptionHandler(JwtTokenException.class)
    protected ResponseEntity<ErrorPayload> handleInvalidJwtTokenException(JwtTokenException ex) {
        log.error(ex.getMessage(), ex);

        ErrorPayload body = new ErrorPayload();
        body.setMessage("Invalid Token");
        body.setErrors(Collections.singletonList(ex.getMessage()));

        return ResponseEntity.status(HttpStatus.BAD_REQUEST).body(body);
    }
}
