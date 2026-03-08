package com.example.flighthub.flight.service.flight.impl;

import com.example.flighthub.flight.exception.FlightNotFoundException;
import com.example.flighthub.flight.model.entity.FlightEntity;
import com.example.flighthub.flight.repository.FlightRepository;
import com.example.flighthub.flight.service.flight.FlightDeleteService;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;

/**
 * Service implementation for deleting a flight in the system.
 */
@Service
@RequiredArgsConstructor
public class FlightDeleteServiceImpl implements FlightDeleteService {

    private final FlightRepository flightRepository;

    /**
     * Deletes a flight by its ID.
     *
     * @param id the ID of the flight to be deleted.
     */
    @Override
    public void deleteFlightById(String id) {
        FlightEntity flightEntityToBeDeleted = flightRepository.findById(id)
                .orElseThrow(() -> new FlightNotFoundException("With given id = " + id));

        flightRepository.delete(flightEntityToBeDeleted);
    }

}
