/**
 * skenario testing
 *
 * - LoginForm component
 *  - should handle email input correctly
 *  - should handle password input correctly
 *  - should handle login when login button is clicked
 *  - should redirect to /register when register link is clicked
 */
import React from 'react'
import { describe, expect, it, vi, afterEach } from 'vitest'
import { render, screen, cleanup } from '@testing-library/react'
import { MemoryRouter, Route, Routes } from 'react-router-dom'
import userEvent from '@testing-library/user-event'
import matchers from '@testing-library/jest-dom/matchers'
import LoginForm from './LoginForm'

expect.extend(matchers)

describe('LoginForm component', () => {
  afterEach(() => {
    cleanup()
  })

  it('should handle email input correctly', async () => {
    render(
      <MemoryRouter>
        <LoginForm login={() => {}} />
      </MemoryRouter>
    )
    const email = 'john@example.com'

    const emailInput = screen.getByPlaceholderText('Enter your email')
    await userEvent.type(emailInput, email)

    expect(emailInput).toHaveValue(email)
  })

  it('should handle password input correctly', async () => {
    render(
      <MemoryRouter>
        <LoginForm login={() => {}} />
      </MemoryRouter>
    )
    const password = 'password'

    const passwordInput = screen.getByPlaceholderText('Enter your password')
    await userEvent.type(passwordInput, password)

    expect(passwordInput).toHaveValue(password)
  })

  it('should handle login when login button is clicked', async () => {
    const inputValues = {
      email: 'john@example.com',
      password: 'password'
    }

    const loginHandler = vi.fn()
    render(
      <MemoryRouter>
        <LoginForm login={loginHandler} />
      </MemoryRouter>
    )

    const emailInput = screen.getByPlaceholderText('Enter your email')
    await userEvent.type(emailInput, inputValues.email)
    const passwordInput = screen.getByPlaceholderText('Enter your password')
    await userEvent.type(passwordInput, inputValues.password)
    const submitButton = screen.getByRole('button', { name: /login/i })
    await userEvent.click(submitButton)

    expect(loginHandler).toHaveBeenCalled(inputValues)
  })

  it('should redirect to /register when register link is clicked', async () => {
    render(
      <MemoryRouter initialEntries={['/login']}>
        <Routes>
          <Route path="/login" element={<LoginForm login={() => {}} />} />
          <Route path="/register" element={<div data-testid="register-page">Register Page</div>} />
        </Routes>
      </MemoryRouter>
    )

    const registerLink = screen.getByRole('link', { name: /register/i })
    await userEvent.click(registerLink)
    const registerPage = screen.getByTestId('register-page')

    expect(registerPage).toBeInTheDocument()
  })
})
