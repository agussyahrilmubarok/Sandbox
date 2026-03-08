package com.example.flighthub.flight.service.airport.impl;

import com.example.flighthub.flight.exception.AirportNotFoundException;
import com.example.flighthub.flight.model.entity.AirportEntity;
import com.example.flighthub.flight.repository.AirportRepository;
import com.example.flighthub.flight.service.airport.AirportDeleteService;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;

/**
 * Service implementation for deleting an airport in the system.
 */
@Service
@RequiredArgsConstructor
public class AirportDeleteServiceImpl implements AirportDeleteService {

    private final AirportRepository airportRepository;

    /**
     * Deletes an airport by its ID.
     *
     * @param id the ID of the airport to be deleted.
     */
    @Override
    public void deleteAirportById(String id) {
        AirportEntity airportEntityToBeDeleted = airportRepository.findById(id)
                .orElseThrow(() -> new AirportNotFoundException("With given id = " + id));

        airportRepository.delete(airportEntityToBeDeleted);
    }

}
