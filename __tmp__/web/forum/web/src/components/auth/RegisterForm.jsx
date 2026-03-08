import React from 'react'
import Input from '../Input'
import useInput from '../../hooks/useInput'
import PropTypes from 'prop-types'

export default function RegisterForm({ register }) {
  const [name, onNameChange] = useInput()
  const [email, onEmailChange] = useInput()
  const [password, onPasswordChange] = useInput()

  const fields = [
    {
      label: 'Name',
      type: 'text',
      placeholder: 'Enter your Name',
      name: 'name',
      value: name,
      onChange: onNameChange
    },
    {
      label: 'Email',
      type: 'email',
      placeholder: 'Enter your email',
      name: 'email',
      value: email,
      onChange: onEmailChange
    },
    {
      label: 'Password',
      type: 'password',
      placeholder: 'Enter your password',
      name: 'password',
      value: password,
      onChange: onPasswordChange
    }
  ]

  const submitHandler = (ev) => {
    ev.preventDefault()
    register({ name, email, password })
  }

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

      <button type="submit" className="btn mt-2.5 btn-primary">
        register
      </button>
    </form>
  )
}

RegisterForm.propTypes = {
  register: PropTypes.func.isRequired
}
