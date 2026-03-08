package com.example.flighthub.flight.model.mapper.flight;

import com.example.flighthub.common.model.mapper.BaseMapper;
import com.example.flighthub.flight.model.Flight;
import com.example.flighthub.flight.model.entity.FlightEntity;
import org.mapstruct.Mapper;
import org.mapstruct.factory.Mappers;

/**
 * Mapper interface for converting a {@link FlightEntity} to a {@link Flight}.
 * This interface extends the {@link BaseMapper} interface, enabling automatic mapping
 * between {@link FlightEntity} and {@link Flight} objects.
 */
@Mapper
public interface FlightEntityToFlightMapper extends BaseMapper<FlightEntity, Flight> {

    /**
     * Initializes and returns an instance of the {@link FlightEntityToFlightMapper}.
     *
     * @return an instance of the mapper
     */
    static FlightEntityToFlightMapper initialize() {
        return Mappers.getMapper(FlightEntityToFlightMapper.class);
    }

}
