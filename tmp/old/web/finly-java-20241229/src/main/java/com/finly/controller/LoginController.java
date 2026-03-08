package com.finly.controller;

import com.finly.security.model.LoginDTO;
import jakarta.validation.Valid;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.context.SecurityContextHolder;
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
public class LoginController extends BaseController {

    private final AuthenticationManager authenticationManager;

    @GetMapping("/login")
    public String login(Model model) {
        if (isLoggedIn()) return "redirect:/dashboard";

        model.addAttribute("dto", new LoginDTO());

        return "home/login";
    }

    @PostMapping("/login")
    public String login(@ModelAttribute("dto") @Valid LoginDTO loginDTO,
                        BindingResult bindingResult,
                        RedirectAttributes redirectAttributes) {
        log.debug("Processing login for email: {}", loginDTO.getEmail());

        if (bindingResult.hasErrors()) {
            log.warn("Login validation failed: {}", bindingResult.getAllErrors());
            return "home/login";
        }

        UsernamePasswordAuthenticationToken authToken =
                new UsernamePasswordAuthenticationToken(loginDTO.getEmail(), loginDTO.getPassword());
        Authentication auth = authenticationManager.authenticate(authToken);
        SecurityContextHolder.getContext().setAuthentication(auth);
        redirectAttributes.addFlashAttribute("info", Map.of("message", "Account login successfully!", "type", "success"));

        log.info("Login successful");
        return "redirect:/dashboard";
    }
}
