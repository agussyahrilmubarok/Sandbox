import React from 'react'
import { useDispatch } from 'react-redux'
import { asyncRegisterUser } from '../states/users/action'
import { useNavigate } from 'react-router-dom'
import RegisterForm from '../components/auth/RegisterForm'
import { Helmet } from 'react-helmet'

export default function RegisterPage() {
  const dispatch = useDispatch()
  const navigate = useNavigate()

  const register = ({ name, email, password }) => {
    dispatch(asyncRegisterUser({ name, email, password })).then((res) => {
      if (!res.error) {
        navigate('/login')
      }
    })
  }

  return (
    <>
      <Helmet>
        <title>Register</title>
      </Helmet>
      <section className="mt-8">
        <h1 className="text-3xl font-bold text-center mb-4">Register</h1>
        <RegisterForm register={register} />
      </section>
    </>
  )
}
