package com.example.flighthub.flight.service.flight.impl;

import com.example.flighthub.flight.exception.AirportNotFoundException;
import com.example.flighthub.flight.exception.FlightNotFoundException;
import com.example.flighthub.flight.model.Flight;
import com.example.flighthub.flight.model.dto.request.flight.UpdateFlightRequest;
import com.example.flighthub.flight.model.entity.AirportEntity;
import com.example.flighthub.flight.model.entity.FlightEntity;
import com.example.flighthub.flight.model.mapper.flight.FlightEntityToFlightMapper;
import com.example.flighthub.flight.model.mapper.flight.UpdateFlightRequestToFlightEntityMapper;
import com.example.flighthub.flight.repository.AirportRepository;
import com.example.flighthub.flight.repository.FlightRepository;
import com.example.flighthub.flight.service.flight.FlightUpdateService;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;

/**
 * Service implementation for updating a flight in the system.
 */
@Service
@RequiredArgsConstructor
public class FlightUpdateServiceImpl implements FlightUpdateService {

    private final FlightRepository flightRepository;
    private final AirportRepository airportRepository;
    private final UpdateFlightRequestToFlightEntityMapper updateFlightRequestToFlightEntityMapper =
            UpdateFlightRequestToFlightEntityMapper.initialize();
    private final FlightEntityToFlightMapper flightEntityToFlightMapper =
            FlightEntityToFlightMapper.initialize();

    /**
     * Updates an existing a flight by its ID.
     *
     * @param id                  the ID of the flight to be updated.
     * @param updateFlightRequest the request object containing the updated details of the flight.
     * @return the updated {@link Flight} entity.
     */
    @Override
    public Flight updateFlightById(String id, UpdateFlightRequest updateFlightRequest) {
        FlightEntity flightEntity = flightRepository.findById(id)
                .orElseThrow(() -> new FlightNotFoundException("Flight given id cant found " + id));

        AirportEntity departureAirportEntity = airportRepository.findById(updateFlightRequest.getFromAirportId())
                .orElseThrow(() -> new AirportNotFoundException("Departure airport not found with id " + updateFlightRequest.getFromAirportId()));

        AirportEntity arrivalAirportEntity = airportRepository.findById(updateFlightRequest.getToAirportId())
                .orElseThrow(() -> new AirportNotFoundException("Arrival airport not found with id " + updateFlightRequest.getToAirportId()));

        updateFlightRequestToFlightEntityMapper.updateFlightMapper(flightEntity, updateFlightRequest,
                departureAirportEntity, arrivalAirportEntity);

        FlightEntity updatedFlightEntity = flightRepository.save(flightEntity);

        return flightEntityToFlightMapper.map(updatedFlightEntity);

    }

}
