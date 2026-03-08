package com.example.flighthub.flight.repository;


import com.example.flighthub.flight.model.entity.FlightEntity;
import org.springframework.data.mongodb.repository.MongoRepository;

/**
 * Repository interface for managing {@link FlightRepository} entities in the Couchbase database.
 * Provides CRUD operations and custom query methods for handling {@link FlightEntity} records.
 */
public interface FlightRepository extends MongoRepository<FlightEntity, String> {

}
