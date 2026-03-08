package com.example.matrix_plus_shape.controller;

import com.example.matrix_plus_shape.service.IMatrixService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.GetMapping;

import java.util.List;

@Controller
public class IndexController {

    @Autowired
    private IMatrixService matrixService;

    @GetMapping("/")
    public String index(Model model) {
        char[][] matrix = matrixService.generateMatrix();
        List<int[]> plusCoordinates = matrixService.findPlusShape(matrix);

        model.addAttribute("matrix", matrix);
        model.addAttribute("plusCoordinates", plusCoordinates);

        return "index";
    }

}
