package com.example.matrix_plus_shape.service;

import java.util.List;

public interface IMatrixService {
    char[][] generateMatrix();
    List<int[]> findPlusShape(char[][] param);
}
