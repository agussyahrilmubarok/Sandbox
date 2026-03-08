package com.example.matrix_plus_shape.service.impl;

import com.example.matrix_plus_shape.service.IMatrixService;
import org.springframework.stereotype.Service;

import java.util.ArrayList;
import java.util.List;
import java.util.Random;

@Service
public class MatrixServiceImpl implements IMatrixService {

    private static final int WIDTH = 25;
    private static final int HEIGHT = 15;

    @Override
    public char[][] generateMatrix() {
        char[][] matrix = new char[HEIGHT][WIDTH];
        Random random = new Random();
        for (int i = 0; i < HEIGHT; i++) {
            for (int j = 0; j < WIDTH; j++) {
                matrix[i][j] = random.nextInt(2) == 0 ? '0' : '1';
            }
        }
        return matrix;
    }

    @Override
    public List<int[]> findPlusShape(char[][] param) {
        List<int[]> plusCoordinates = new ArrayList<>();

        for (int i = 1; i < HEIGHT - 1; i++) {
            for (int j = 1; j < WIDTH - 1; j++) {
                if (param[i][j] == '1' &&
                        param[i-1][j] == '1' &&
                        param[i+1][j] == '1' &&
                        param[i][j-1] == '1' &&
                        param[i][j+1] == '1') {
                    plusCoordinates.add(new int[]{j, i});
                    param[i][j] = '*';
                }
            }
        }

        printMatrix(param);
        printPlusCoordinates(plusCoordinates);

        return plusCoordinates;
    }

    private void printMatrix(char[][] matrix) {
        for (int i = 0; i < HEIGHT; i++) {
            for (int j = 0; j < WIDTH; j++) {
                System.out.print(matrix[i][j] + " ");
            }
            System.out.println();
        }
    }

    private void printPlusCoordinates(List<int[]> plusCoordinates) {
        System.out.println("Jawaban:");
        for (int[] coord : plusCoordinates) {
            System.out.println("(" + coord[0] + ", " + coord[1] + ")");
        }
    }
}