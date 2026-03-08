package com.example.flighthub.flight.model.mapper.airport;

import com.example.flighthub.common.model.mapper.BaseMapper;
import com.example.flighthub.flight.model.dto.request.airport.CreateAirportRequest;
import com.example.flighthub.flight.model.entity.AirportEntity;
import org.mapstruct.Mapper;
import org.mapstruct.Mapping;
import org.mapstruct.Mappings;
import org.mapstruct.Named;
import org.mapstruct.factory.Mappers;

/**
 * Mapper interface for converting a {@link CreateAirportRequest} to a {@link AirportEntity}.
 * This interface extends the {@link BaseMapper} interface, allowing automatic mapping
 * between the {@link CreateAirportRequest} and {@link AirportEntity} objects.
 */
@Mapper
public interface CreateAirportRequestToAirportEntityMapper extends BaseMapper<CreateAirportRequest, AirportEntity> {

    /**
     * Initializes and returns an instance of the {@link CreateAirportRequestToAirportEntityMapper}.
     *
     * @return an instance of the mapper
     */
    static CreateAirportRequestToAirportEntityMapper initialize() {
        return Mappers.getMapper(CreateAirportRequestToAirportEntityMapper.class);
    }

    /**
     * Converts a {@link CreateAirportRequest} to a {@link AirportEntity}.
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
    AirportEntity map(CreateAirportRequest source);

    /**
     * Converts a {@link CreateAirportRequest} to a {@link AirportEntity}.
     * This method is used for saving an airport, mapping only the necessary fields from the
     * {@link CreateAirportRequest} to the {@link AirportEntity}.
     *
     * @param request the {@link CreateAirportRequest} to be mapped
     * @return the resulting {@link AirportEntity} containing the mapped fields
     */
    @Named("mapForSaving")
    default AirportEntity mapForSaving(CreateAirportRequest request) {
        return AirportEntity.builder()
                .name(request.getName())
                .cityName(request.getCityName())
                .build();
    }

}
