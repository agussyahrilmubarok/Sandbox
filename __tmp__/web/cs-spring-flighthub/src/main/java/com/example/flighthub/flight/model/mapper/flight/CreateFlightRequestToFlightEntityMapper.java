package com.example.flighthub.flight.model.mapper.flight;

import com.example.flighthub.common.model.mapper.BaseMapper;
import com.example.flighthub.flight.model.dto.request.flight.CreateFlightRequest;
import com.example.flighthub.flight.model.entity.AirportEntity;
import com.example.flighthub.flight.model.entity.FlightEntity;
import com.example.flighthub.flight.model.mapper.airport.CreateAirportRequestToAirportEntityMapper;
import org.mapstruct.Mapper;
import org.mapstruct.Mapping;
import org.mapstruct.Mappings;
import org.mapstruct.Named;
import org.mapstruct.factory.Mappers;

/**
 * Mapper interface for converting a {@link CreateFlightRequest} to a {@link FlightEntity}.
 * This interface extends the {@link BaseMapper} interface, allowing automatic mapping
 * between the {@link CreateFlightRequest} and {@link FlightEntity} objects.
 */
@Mapper
public interface CreateFlightRequestToFlightEntityMapper extends BaseMapper<CreateFlightRequest, FlightEntity> {

    /**
     * Initializes and returns an instance of the {@link CreateAirportRequestToAirportEntityMapper}.
     *
     * @return an instance of the mapper
     */
    static CreateFlightRequestToFlightEntityMapper initialize() {
        return Mappers.getMapper(CreateFlightRequestToFlightEntityMapper.class);
    }

    /**
     * Converts a {@link CreateFlightRequest} to a {@link FlightEntity}.
     * Ignores fields that are not present in the request.
     */
    @Override
    @Mappings({
            @Mapping(target = "id", ignore = true),
            @Mapping(target = "createdAt", ignore = true),
            @Mapping(target = "createdBy", ignore = true),
            @Mapping(target = "updatedAt", ignore = true),
            @Mapping(target = "updatedBy", ignore = true),
            @Mapping(target = "fromAirport", ignore = true),
            @Mapping(target = "toAirport", ignore = true)
    })
    FlightEntity map(CreateFlightRequest source);

    /**
     * Converts a {@link CreateFlightRequest} to a {@link FlightEntity}.
     * This method is used for saving a flight, mapping only the necessary fields from the
     * {@link CreateFlightRequest} to the {@link FlightEntity}.
     *
     * @param request     the {@link CreateFlightRequest} to be mapped
     * @param fromAirport the from airport entity
     * @param toAirport   the to airport entity
     * @return the resulting {@link FlightEntity} containing the mapped fields
     */
    @Named("mapForSaving")
    default FlightEntity mapForSaving(CreateFlightRequest request, AirportEntity fromAirport, AirportEntity toAirport) {
        return FlightEntity.builder()
                .fromAirport(fromAirport)
                .toAirport(toAirport)
                .departureTime(request.getDepartureTime())
                .arrivalTime(request.getArrivalTime())
                .price(request.getPrice())
                .build();
    }


}
