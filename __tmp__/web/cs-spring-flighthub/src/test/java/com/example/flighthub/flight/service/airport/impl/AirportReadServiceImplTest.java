package com.example.flighthub.flight.service.airport.impl;

import com.example.flighthub.base.AbstractBaseServiceTest;
import com.example.flighthub.builder.AirportEntityBuilder;
import com.example.flighthub.common.model.CustomPage;
import com.example.flighthub.common.model.CustomPaging;
import com.example.flighthub.flight.exception.AirportNotFoundException;
import com.example.flighthub.flight.model.Airport;
import com.example.flighthub.flight.model.dto.request.airport.AirportPagingRequest;
import com.example.flighthub.flight.model.entity.AirportEntity;
import com.example.flighthub.flight.model.mapper.airport.AirportEntityToAirportMapper;
import com.example.flighthub.flight.model.mapper.airport.ListAirportEntityToListAirportMapper;
import com.example.flighthub.flight.repository.AirportRepository;
import org.junit.jupiter.api.Test;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.PageImpl;
import org.springframework.data.domain.Pageable;

import java.util.Collections;
import java.util.List;
import java.util.Optional;
import java.util.UUID;

import static org.junit.jupiter.api.Assertions.*;
import static org.mockito.Mockito.*;

/**
 * Unit test class for {@link AirportReadServiceImpl}.
 * This class verifies the correctness of the business logic in {@link AirportReadServiceImpl}.
 */
class AirportReadServiceImplTest extends AbstractBaseServiceTest {

    private final AirportEntityToAirportMapper airportEntityToAirportMapper =
            AirportEntityToAirportMapper.initialize();
    private final ListAirportEntityToListAirportMapper listAirportEntityToListAirportMapper =
            ListAirportEntityToListAirportMapper.initialize();
    @InjectMocks
    private AirportReadServiceImpl airportReadService;
    @Mock
    private AirportRepository airportRepository;

    @Test
    void givenExistAirportId_whenGetAirportById_thenReturnAirport() {

        // Given
        final String mockId = UUID.randomUUID().toString();

        final AirportEntity airportEntity = new AirportEntityBuilder()
                .withId(mockId)
                .withValidFields();

        final Airport mockAirport = airportEntityToAirportMapper.map(airportEntity);

        // When
        when(airportRepository.findById(mockId)).thenReturn(Optional.of(airportEntity));

        // Then
        final Airport expected = airportReadService.getAirportById(mockId);

        assertNotNull(expected);
        assertEquals(mockAirport.getId(), expected.getId());
        assertEquals(mockAirport.getName(), expected.getName());

        // Verify
        verify(airportRepository, times(1)).findById(mockId);

    }

    @Test
    void givenNonExistId_whenGetById_thenThrowTaskNotFoundException() {

        // Given
        final String mockId = UUID.randomUUID().toString();

        // When
        when(airportRepository.findById(mockId)).thenReturn(Optional.empty());

        // Then
        assertThrows(AirportNotFoundException.class,
                () -> airportReadService.getAirportById(mockId));

        // Verify
        verify(airportRepository, times(1)).findById(mockId);

    }

    @Test
    void givenTaskPagingRequest_WhenTaskPageList_ThenReturnCustomPageTaskList() {

        // Given
        final AirportPagingRequest pagingRequest = AirportPagingRequest.builder()
                .pagination(
                        CustomPaging.builder()
                                .pageSize(1)
                                .pageNumber(1)
                                .build()
                ).build();

        final AirportEntity airportEntity = new AirportEntityBuilder().withValidFields();

        Page<AirportEntity> airportEntityPage = new PageImpl<>(Collections.singletonList(airportEntity));

        List<Airport> products = listAirportEntityToListAirportMapper.toAirportList(airportEntityPage.getContent());

        CustomPage<Airport> expected = CustomPage.of(products, airportEntityPage);

        // When
        when(airportRepository.findAll(any(Pageable.class))).thenReturn(airportEntityPage);

        // Then
        CustomPage<Airport> result = airportReadService.getAllAirports(pagingRequest);

        assertNotNull(result);
        assertFalse(result.getContent().isEmpty());
        assertEquals(expected.getPageNumber(), result.getPageNumber());
        assertEquals(expected.getContent().get(0).getId(), result.getContent().get(0).getId());
        assertEquals(expected.getTotalPageCount(), result.getTotalPageCount());
        assertEquals(expected.getTotalElementCount(), result.getTotalElementCount());

        // Verify
        verify(airportRepository, times(1)).findAll(any(Pageable.class));

    }

    @Test
    void givenTaskPagingRequest_WhenNoTaskPageList_ThenThrowTaskNotFoundException() {

        // Given
        final AirportPagingRequest pagingRequest = AirportPagingRequest.builder()
                .pagination(
                        CustomPaging.builder()
                                .pageSize(1)
                                .pageNumber(1)
                                .build()
                ).build();

        Page<AirportEntity> airportEntityPage = new PageImpl<>(Collections.emptyList());

        // When
        when(airportRepository.findAll(any(Pageable.class))).thenReturn(airportEntityPage);

        // Then
        assertThrows(AirportNotFoundException.class, () -> airportReadService.getAllAirports(pagingRequest));

        // Verify
        verify(airportRepository, times(1)).findAll(any(Pageable.class));

    }

}