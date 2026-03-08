package com.example.flighthub.builder;

import lombok.SneakyThrows;

/**
 * A generic base builder class that can be used to build objects of type {@code T}.
 * The builder initializes the object through reflection and provides a {@link #build()} method
 * to return the constructed instance.
 *
 * @param <T> the type of object this builder creates
 */
public abstract class BaseBuilder<T> {

    T data;

    @SneakyThrows
    public BaseBuilder(Class<T> clazz) {
        this.data = clazz.getDeclaredConstructor().newInstance();
    }

    public T build() {
        return data;
    }

}
