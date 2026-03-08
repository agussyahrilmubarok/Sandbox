package com.example.demo.flight.job;

import com.example.demo.base.AbstractBaseServiceTest;
import com.example.demo.flight.model.entity.AirportEntity;
import com.example.demo.flight.repository.AirportRepository;
import com.example.demo.flight.repository.FlightRepository;
import org.junit.jupiter.api.Test;
import org.mockito.InjectMocks;
import org.mockito.Mock;

import java.util.List;
import java.util.UUID;

import static org.mockito.Mockito.*;

/**
 * Unit test for {@link FlightDataLoader}.
 * This test verifies that the dummy flight data is properly loaded
 * and that interactions with the {@link FlightRepository} and {@link AirportRepository}
 * occur as expected.
 */
class FlightDataLoaderTest extends AbstractBaseServiceTest {

    @InjectMocks
    private FlightDataLoader flightDataLoader;

    @Mock
    private FlightRepository flightRepository;

    @Mock
    private AirportRepository airportRepository;

    @Test
    void shouldLoadFlightDummyData() {

        // Given
        AirportEntity airport1 = AirportEntity.builder()
                .id(UUID.randomUUID().toString())
                .name("Bandara Internasional Soekarno-Hatta")
                .cityName("Jakarta")
                .build();

        AirportEntity airport2 = AirportEntity.builder()
                .id(UUID.randomUUID().toString())
                .name("Bandara Internasional Ngurah Rai")
                .cityName("Denpasar")
                .build();

        AirportEntity airport3 = AirportEntity.builder()
                .id(UUID.randomUUID().toString())
                .name("Bandara Internasional Juanda")
                .cityName("Surabaya")
                .build();

        List<AirportEntity> mockAirports = List.of(airport1, airport2, airport3);

        // When
        when(airportRepository.findAll()).thenReturn(mockAirports);

        // Then
        flightDataLoader.loadFlightDumpyData();

        // Verify
        verify(airportRepository, times(1)).saveAll(anyList());
        verify(flightRepository, times(1)).saveAll(anyList());

    }

}