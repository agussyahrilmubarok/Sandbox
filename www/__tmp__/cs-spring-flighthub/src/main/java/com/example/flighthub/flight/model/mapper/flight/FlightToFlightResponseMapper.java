package com.example.flighthub.flight.model.mapper.flight;

import com.example.flighthub.common.model.mapper.BaseMapper;
import com.example.flighthub.flight.model.Flight;
import com.example.flighthub.flight.model.dto.response.flight.FlightResponse;
import org.mapstruct.Mapper;
import org.mapstruct.factory.Mappers;

/**
 * Mapper interface for converting a {@link Flight} to a {@link FlightResponse}.
 * This interface extends the {@link BaseMapper} interface, enabling automatic mapping
 * between {@link Flight} and {@link FlightResponse} objects.
 */
@Mapper
public interface FlightToFlightResponseMapper extends BaseMapper<Flight, FlightResponse> {

    /**
     * Initializes and returns an instance of the {@link FlightToFlightResponseMapper}.
     * This method is used to obtain a mapper instance for converting a {@link Flight}
     * object into a {@link FlightResponse} object.
     *
     * @return an instance of the {@link FlightToFlightResponseMapper} mapper
     */
    static FlightToFlightResponseMapper initialize() {
        return Mappers.getMapper(FlightToFlightResponseMapper.class);
    }

}
