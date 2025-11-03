package com.example.flighthub.flight.model.dto.request.flight;

import com.example.flighthub.common.model.dto.request.CustomPagingRequest;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;
import lombok.experimental.SuperBuilder;

/**
 * Request class used for paginated flight retrieval.
 * Inherits from the base class {@link CustomPagingRequest}, which provides pagination details such as page number and size.
 * This class is used to request a paginated list of flights.
 */
@Getter
@Setter
@SuperBuilder
@NoArgsConstructor
public class FlightPagingRequest extends CustomPagingRequest {

}
