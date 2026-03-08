package com.example.flighthub.flight.model.mapper.airport;

import com.example.flighthub.common.model.mapper.BaseMapper;
import com.example.flighthub.flight.model.Airport;
import com.example.flighthub.flight.model.entity.AirportEntity;
import org.mapstruct.Mapper;
import org.mapstruct.factory.Mappers;


/**
 * Mapper interface for converting a {@link AirportEntity} to a {@link Airport}.
 * This interface extends the {@link BaseMapper} interface, enabling automatic mapping
 * between {@link AirportEntity} and {@link Airport} objects.
 */
@Mapper
public interface AirportEntityToAirportMapper extends BaseMapper<AirportEntity, Airport> {

    /**
     * Initializes and returns an instance of the {@link AirportEntityToAirportMapper}.
     *
     * @return an instance of the mapper
     */
    static AirportEntityToAirportMapper initialize() {
        return Mappers.getMapper(AirportEntityToAirportMapper.class);
    }

}
