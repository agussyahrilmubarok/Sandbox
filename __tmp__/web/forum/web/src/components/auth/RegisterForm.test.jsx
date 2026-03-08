/**
 * skenario testing
 *
 * - RegisterForm component
 *  - should handle name input correctly
 *  - should handle email input correctly
 *  - should handle password input correctly
 *  - should handle register when register button is clicked
 */
import React from 'react'
import { describe, expect, it, vi, afterEach } from 'vitest'
import { render, screen, cleanup } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import matchers from '@testing-library/jest-dom/matchers'
import RegisterForm from './RegisterForm'

expect.extend(matchers)

describe('RegisterForm component', () => {
  afterEach(() => {
    cleanup()
  })

  it('should handle name input correctly', async () => {
    render(<RegisterForm register={() => {}} />)
    const name = 'John Doe'

    const nameInput = screen.getByPlaceholderText('Enter your Name')
    await userEvent.type(nameInput, name)

    expect(nameInput).toHaveValue(name)
  })

  it('should handle email input correctly', async () => {
    render(<RegisterForm register={() => {}} />)
    const email = 'john@example.com'

    const emailInput = screen.getByPlaceholderText('Enter your email')
    await userEvent.type(emailInput, email)

    expect(emailInput).toHaveValue(email)
  })

  it('should handle password input correctly', async () => {
    render(<RegisterForm register={() => {}} />)
    const password = 'password'

    const passwordInput = screen.getByPlaceholderText('Enter your password')
    await userEvent.type(passwordInput, password)

    expect(passwordInput).toHaveValue(password)
  })

  it('should handle register when register button is clicked', async () => {
    const registerHandler = vi.fn()
    render(<RegisterForm register={registerHandler} />)
    const inputValues = {
      name: 'John Doe',
      email: 'john@example.com',
      password: 'password'
    }

    const nameInput = screen.getByPlaceholderText('Enter your Name')
    await userEvent.type(nameInput, inputValues.name)
    const emailInput = screen.getByPlaceholderText('Enter your email')
    await userEvent.type(emailInput, inputValues.email)
    const passwordInput = screen.getByPlaceholderText('Enter your password')
    await userEvent.type(passwordInput, inputValues.password)
    const submitButton = screen.getByRole('button', { name: /register/i })
    await userEvent.click(submitButton)

    expect(registerHandler).toHaveBeenCalledWith(inputValues)
  })
})
