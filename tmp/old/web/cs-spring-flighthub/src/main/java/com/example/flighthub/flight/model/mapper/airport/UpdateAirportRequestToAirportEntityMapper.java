package com.example.flighthub.flight.model.mapper.airport;

import com.example.flighthub.common.model.mapper.BaseMapper;
import com.example.flighthub.flight.model.dto.request.airport.UpdateAirportRequest;
import com.example.flighthub.flight.model.entity.AirportEntity;
import org.mapstruct.Mapper;
import org.mapstruct.Mapping;
import org.mapstruct.Mappings;
import org.mapstruct.Named;
import org.mapstruct.factory.Mappers;

/**
 * Mapper interface for converting and updating a {@link AirportEntity} using data from an {@link UpdateAirportRequest}.
 * This interface extends the {@link BaseMapper} interface, enabling automatic mapping
 * between {@link UpdateAirportRequest} and {@link AirportEntity}.
 */
@Mapper
public interface UpdateAirportRequestToAirportEntityMapper extends BaseMapper<UpdateAirportRequest, AirportEntity> {

    /**
     * Initializes and returns an instance of the {@link UpdateAirportRequestToAirportEntityMapper}.
     * This method is used to obtain a mapper instance for mapping an {@link UpdateAirportRequest}
     * to a {@link AirportEntity}, including updating an existing entity.
     *
     * @return an instance of the {@link UpdateAirportRequestToAirportEntityMapper} mapper
     */
    static UpdateAirportRequestToAirportEntityMapper initialize() {
        return Mappers.getMapper(UpdateAirportRequestToAirportEntityMapper.class);
    }

    /**
     * Converts a {@link UpdateAirportRequest} to a {@link AirportEntity}.
     * Ignores fields that are not present in the request.
     */
    @Override
    @Mappings({
            @Mapping(target = "id", ignore = true),
            @Mapping(target = "createdAt", ignore = true),
            @Mapping(target = "createdBy", ignore = true),
            @Mapping(target = "updatedAt", ignore = true),
            @Mapping(target = "updatedBy", ignore = true)
    })
    AirportEntity map(UpdateAirportRequest source);

    /**
     * Updates the fields of an existing {@link AirportEntity} with data from an {@link UpdateAirportRequest}.
     * This method is used to map the fields from the request to the entity, specifically
     * updating the name of the task.
     *
     * @param airportEntity     the {@link AirportEntity} to be updated
     * @param updateTaskRequest the {@link UpdateAirportRequest} containing the new data
     */
    @Named("mapForUpdate")
    default void updateAirportMapper(final AirportEntity airportEntity, final UpdateAirportRequest updateTaskRequest) {

        airportEntity.setName(updateTaskRequest.getName());
        airportEntity.setCityName(updateTaskRequest.getCityName());

    }

}
