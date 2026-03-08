package com.example.flighthub.flight.service.airport.impl;

import com.example.flighthub.base.AbstractBaseServiceTest;
import com.example.flighthub.builder.AirportEntityBuilder;
import com.example.flighthub.builder.UpdateAirportRequestBuilder;
import com.example.flighthub.flight.exception.AirportNameAlreadyExistException;
import com.example.flighthub.flight.exception.AirportNotFoundException;
import com.example.flighthub.flight.model.Airport;
import com.example.flighthub.flight.model.dto.request.airport.UpdateAirportRequest;
import com.example.flighthub.flight.model.entity.AirportEntity;
import com.example.flighthub.flight.model.mapper.airport.AirportEntityToAirportMapper;
import com.example.flighthub.flight.repository.AirportRepository;
import org.junit.jupiter.api.Test;
import org.mockito.InjectMocks;
import org.mockito.Mock;

import java.util.Optional;
import java.util.UUID;

import static org.junit.jupiter.api.Assertions.*;
import static org.mockito.Mockito.*;

/**
 * Unit test class for {@link AirportUpdateServiceImpl}.
 * This class verifies the correctness of the business logic in {@link AirportUpdateServiceImpl}.
 */
class AirportUpdateServiceImplTest extends AbstractBaseServiceTest {

    private final AirportEntityToAirportMapper airportEntityToAirportMapper =
            AirportEntityToAirportMapper.initialize();
    @InjectMocks
    private AirportUpdateServiceImpl airportUpdateService;
    @Mock
    private AirportRepository airportRepository;

    @Test
    void givenUpdateAirportRequest_whenUpdateAirportById_thenSuccess() {

        // Given
        final String mockId = UUID.randomUUID().toString();

        final UpdateAirportRequest mockUpdateAirportRequest = new UpdateAirportRequestBuilder()
                .withValidFields()
                .build();

        final AirportEntity existingAirportEntity = new AirportEntityBuilder()
                .withId(mockId)
                .withValidFields();

        final AirportEntity updatedAirportEntity = new AirportEntityBuilder()
                .withId(mockId)
                .withName(mockUpdateAirportRequest.getName())
                .withCityName(mockUpdateAirportRequest.getCityName())
                .build();

        final Airport expected = airportEntityToAirportMapper.map(updatedAirportEntity);

        // When
        when(airportRepository.findById(mockId)).thenReturn(Optional.of(existingAirportEntity));
        when(airportRepository.save(existingAirportEntity)).thenReturn(updatedAirportEntity);

        Airport result = airportUpdateService.updateAirportById(mockId, mockUpdateAirportRequest);

        // Then
        assertNotNull(result);
        assertEquals(expected.getName(), result.getName());
        assertEquals(expected.getCityName(), result.getCityName());

        // Verify
        verify(airportRepository, times(1)).findById(mockId);
        verify(airportRepository, times(1)).save(existingAirportEntity);

    }

    @Test
    void givenNonexistentAirportId_whenUpdateAirportById_thenThrowAirportNotFoundException() {

        // Given
        final String airportId = UUID.randomUUID().toString();

        final UpdateAirportRequest updateRequest = new UpdateAirportRequestBuilder()
                .withValidFields()
                .build();

        // When
        when(airportRepository.findById(airportId)).thenReturn(Optional.empty());

        // Then
        assertThrows(AirportNotFoundException.class, () ->
                airportUpdateService.updateAirportById(airportId, updateRequest));

        // Verify
        verify(airportRepository, times(1)).findById(airportId);

    }

    @Test
    void givenDuplicateAirportName_whenUpdateAirportById_thenThrowIllegalArgumentException() {

        // Given
        final String airportId = UUID.randomUUID().toString();
        final UpdateAirportRequest updateRequest = new UpdateAirportRequestBuilder()
                .withName("Duplicate Name")
                .withCityName("New Location")
                .build();

        // When
        when(airportRepository.existsByName(updateRequest.getName())).thenReturn(true);

        // Then
        AirportNameAlreadyExistException exception = assertThrows(AirportNameAlreadyExistException.class, () ->
                airportUpdateService.updateAirportById(airportId, updateRequest));

        assertEquals("Airport with this name already exist\n" +
                " With given airport name = Duplicate Name", exception.getMessage());

        // Verify
        verify(airportRepository, times(0)).findById(airportId);
        verify(airportRepository, times(1)).existsByName(updateRequest.getName());

    }

}