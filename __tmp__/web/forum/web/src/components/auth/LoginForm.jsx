import React from 'react'
import useInput from '../../hooks/useInput'
import Input from '../Input'
import PropTypes from 'prop-types'
import { Link } from 'react-router-dom'

export default function LoginForm({ login }) {
  const [email, onEmailChange] = useInput('')
  const [password, onPasswordChange] = useInput('')

  const submitHandler = (ev) => {
    ev.preventDefault()
    login({ email, password })
  }

  const fields = [
    {
      label: 'Email',
      type: 'email',
      name: 'email',
      placeholder: 'Enter your email',
      value: email,
      onChange: onEmailChange
    },
    {
      label: 'Password',
      type: 'password',
      name: 'password',
      placeholder: 'Enter your password',
      value: password,
      onChange: onPasswordChange
    }
  ]

  return (
    <form onSubmit={submitHandler} className="flex flex-col max-w-lg mx-auto gap-y-2.5">
      {fields.map((field, index) => (
        <Input
          key={index}
          label={field.label}
          type={field.type}
          placeholder={field.placeholder}
          name={field.name}
          value={field.value}
          onChange={field.onChange}
          required
        />
      ))}
      <p className="my-2 text-neutral-content text-sm">
        don{"'"}t have an account?{' '}
        <Link className="text-primary hover:underline" to="/register">
          Register
        </Link>
      </p>
      <button type="submit" className="btn btn-primary">
        Login
      </button>
    </form>
  )
}

LoginForm.propTypes = {
  login: PropTypes.func.isRequired
}
