package com.finly.controller;

import com.finly.security.model.SignUpDTO;
import com.finly.user.model.UserDTO;
import com.finly.user.service.UserService;
import jakarta.validation.Valid;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.validation.BindingResult;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.ModelAttribute;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.servlet.mvc.support.RedirectAttributes;

import java.util.Map;

@Controller
@Slf4j
@RequiredArgsConstructor
public class SignUpController extends BaseController {

    private final UserService userService;

    @GetMapping("/signup")
    public String signUp(Model model) {
        if (isLoggedIn()) return "redirect:/dashboard";

        model.addAttribute("dto", new SignUpDTO());

        return "home/signup";
    }

    @PostMapping("/signup")
    public String signUp(@ModelAttribute("dto") @Valid SignUpDTO signUpDTO,
                         BindingResult bindingResult,
                         RedirectAttributes redirectAttributes) {
        log.debug("Processing login for email: {}", signUpDTO.getEmail());

        if (bindingResult.hasErrors()) {
            log.warn("SignUp validation failed: {}", bindingResult.getAllErrors());
            return "home/signup";
        }

        UserDTO param = new UserDTO();
        param.setEmail(signUpDTO.getEmail());
        param.setPassword(signUpDTO.getPassword());
        Long result = userService.create(param);

        redirectAttributes.addFlashAttribute("info", Map.of("message", "Account registered successfully!", "type", "success"));

        log.info("SignUp successful for id: {}", result);
        return "redirect:/login";
    }
}
