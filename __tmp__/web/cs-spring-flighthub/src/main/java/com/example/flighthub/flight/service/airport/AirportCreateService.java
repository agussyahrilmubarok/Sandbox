package com.example.flighthub.flight.service.airport;

import com.example.flighthub.flight.model.Airport;
import com.example.flighthub.flight.model.dto.request.airport.CreateAirportRequest;

/**
 * Service interface for creating an airport in the system.
 */
public interface AirportCreateService {

    /**
     * Create a new airport to the database.
     *
     * @param createAirportRequest the request object containing the details of the CreateAirportRequest to be saved.
     * @return the saved {@link Airport} domain object.
     */
    Airport createAirport(CreateAirportRequest createAirportRequest);

}
