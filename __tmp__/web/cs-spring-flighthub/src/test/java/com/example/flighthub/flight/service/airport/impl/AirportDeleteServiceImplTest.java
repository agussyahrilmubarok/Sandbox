package com.example.flighthub.flight.service.airport.impl;

import com.example.flighthub.base.AbstractBaseServiceTest;
import com.example.flighthub.builder.AirportEntityBuilder;
import com.example.flighthub.flight.exception.AirportNotFoundException;
import com.example.flighthub.flight.model.entity.AirportEntity;
import com.example.flighthub.flight.repository.AirportRepository;
import org.junit.jupiter.api.Test;
import org.mockito.InjectMocks;
import org.mockito.Mock;

import java.util.Optional;
import java.util.UUID;

import static org.junit.jupiter.api.Assertions.assertThrows;
import static org.mockito.Mockito.*;

/**
 * Unit test class for {@link AirportDeleteServiceImpl}.
 * This class verifies the correctness of the business logic in {@link AirportDeleteServiceImpl}.
 */
class AirportDeleteServiceImplTest extends AbstractBaseServiceTest {

    @InjectMocks
    private AirportDeleteServiceImpl airportDeleteService;

    @Mock
    private AirportRepository airportRepository;


    @Test
    void givenValidAirportId_whenDeleteAirportById_thenDeleteAirportSuccessfully() {

        // Given
        final String mockId = UUID.randomUUID().toString();
        final AirportEntity mockEntity = new AirportEntityBuilder()
                .withId(mockId)
                .build();

        // When
        when(airportRepository.findById(mockId)).thenReturn(java.util.Optional.of(mockEntity));

        // Then
        airportDeleteService.deleteAirportById(mockId);

        // Verify
        verify(airportRepository, times(1)).findById(mockId);
        verify(airportRepository, times(1)).delete(mockEntity);

    }

    @Test
    void givenInvalidAirportId_whenDeleteAirportById_thenThrowAirportNotFoundException() {

        // Given
        final String mockId = UUID.randomUUID().toString();

        // When
        when(airportRepository.findById(mockId)).thenReturn(Optional.empty());

        // Then
        assertThrows(AirportNotFoundException.class, () ->
                airportDeleteService.deleteAirportById(mockId)
        );

        // Verify
        verify(airportRepository, times(1)).findById(mockId);

    }

}