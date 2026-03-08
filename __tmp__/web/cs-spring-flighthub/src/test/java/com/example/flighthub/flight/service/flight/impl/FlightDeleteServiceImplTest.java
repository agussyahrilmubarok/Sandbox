package com.example.flighthub.flight.service.flight.impl;

import com.example.flighthub.base.AbstractBaseServiceTest;
import com.example.flighthub.builder.FlightEntityBuilder;
import com.example.flighthub.flight.exception.FlightNotFoundException;
import com.example.flighthub.flight.model.entity.FlightEntity;
import com.example.flighthub.flight.repository.FlightRepository;
import org.junit.jupiter.api.Test;
import org.mockito.InjectMocks;
import org.mockito.Mock;

import java.util.Optional;
import java.util.UUID;

import static org.junit.jupiter.api.Assertions.assertThrows;
import static org.mockito.Mockito.*;

/**
 * Unit test class for {@link FlightDeleteServiceImpl}.
 * This class verifies the correctness of the business logic in {@link FlightDeleteServiceImpl}.
 */
class FlightDeleteServiceImplTest extends AbstractBaseServiceTest {

    @InjectMocks
    private FlightDeleteServiceImpl flightDeleteService;

    @Mock
    private FlightRepository flightRepository;

    @Test
    void givenValidFlightId_whenDeleteFlightById_thenDeleteFlightSuccessfully() {

        // Given
        final String mockId = UUID.randomUUID().toString();
        final FlightEntity mockEntity = new FlightEntityBuilder()
                .withId(mockId)
                .withValidFields()
                .build();

        // When
        when(flightRepository.findById(mockId)).thenReturn(Optional.of(mockEntity));

        // Then
        flightDeleteService.deleteFlightById(mockId);

        // Verify
        verify(flightRepository, times(1)).findById(mockId);
        verify(flightRepository, times(1)).delete(mockEntity);

    }

    @Test
    void givenInvalidFlightId_whenDeleteFlightById_thenThrowFlightNotFoundException() {

        // Given
        final String mockId = UUID.randomUUID().toString();

        // When
        when(flightRepository.findById(mockId)).thenReturn(Optional.empty());

        // Then
        assertThrows(FlightNotFoundException.class, () ->
                flightDeleteService.deleteFlightById(mockId)
        );

        // Verify
        verify(flightRepository, times(1)).findById(mockId);
        verify(flightRepository, never()).delete(any());

    }

}