package com.example.flighthub.flight.job;

import com.example.flighthub.flight.model.entity.AirportEntity;
import com.example.flighthub.flight.model.entity.FlightEntity;
import com.example.flighthub.flight.repository.AirportRepository;
import com.example.flighthub.flight.repository.FlightRepository;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Component;

import java.time.LocalDateTime;
import java.util.List;
import java.util.UUID;

/**
 * FlightDataLoader is responsible for initializing dummy airport and flight data
 * in the database. It creates and saves sample airport entities and flight entities.
 */
@Component
@RequiredArgsConstructor
public class FlightDataLoader {

    private final FlightRepository flightRepository;
    private final AirportRepository airportRepository;

    /**
     * Loads sample airport and flight data into the database.
     * Creates airport entities and associated flights with dynamic schedules.
     */
    public void loadFlightDumpyData() {

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

        airportRepository.saveAll(List.of(airport1, airport2, airport3));

        List<AirportEntity> airports = airportRepository.findAll();

        // Define dynamic departure times
        LocalDateTime departureFlight1 = LocalDateTime.now().plusDays(1)
                .withHour(10).withMinute(0); // Tomorrow at 10:00 AM
        LocalDateTime arrivalFlight1 = departureFlight1.plusHours(4); // Arrives 4 hours later (2 PM)

        LocalDateTime departureFlight2 = departureFlight1.plusDays(1)
                .withHour(9).withMinute(0); // Return flight next day at 9:00 AM
        LocalDateTime arrivalFlight2 = departureFlight2.plusHours(4); // Arrives back at 1:00 PM

        LocalDateTime departureFlight3 = departureFlight1.plusDays(2)
                .withHour(15).withMinute(30); // Denpasar → Surabaya (Next day 3:30 PM)
        LocalDateTime arrivalFlight3 = departureFlight3.plusHours(3); // Arrives 3 hours later (6:30 PM)

        // Create dummy flights
        FlightEntity flight1 = FlightEntity.builder()
                .id(UUID.randomUUID().toString())
                .fromAirport(airports.get(0)) // Jakarta
                .toAirport(airports.get(1)) // Denpasar
                .departureTime(departureFlight1)
                .arrivalTime(arrivalFlight1)
                .price(250.0)
                .build();

        FlightEntity flight2 = FlightEntity.builder()
                .id(UUID.randomUUID().toString())
                .fromAirport(airports.get(1)) // Denpasar
                .toAirport(airports.get(0)) // Jakarta (Return flight)
                .departureTime(departureFlight2)
                .arrivalTime(arrivalFlight2)
                .price(260.0)
                .build();

        FlightEntity flight3 = FlightEntity.builder()
                .id(UUID.randomUUID().toString())
                .fromAirport(airports.get(1)) // Denpasar
                .toAirport(airports.get(2)) // Surabaya
                .departureTime(departureFlight3)
                .arrivalTime(arrivalFlight3)
                .price(180.0)
                .build();

        flightRepository.saveAll(List.of(flight1, flight2, flight3));

    }

}
