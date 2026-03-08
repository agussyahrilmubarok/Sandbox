import React from 'react'
import { useDispatch } from 'react-redux'
import { asyncSetAuthUser } from '../states/authUser/action'
import LoginForm from '../components/auth/LoginForm'
import { Helmet } from 'react-helmet'

export default function LoginPage() {
  const dispatch = useDispatch()

  const login = ({ email, password }) => {
    dispatch(asyncSetAuthUser({ email, password }))
  }

  return (
    <>
      <Helmet>
        <title>Login</title>
      </Helmet>
      <section className="mt-8">
        <h1 className="text-3xl font-bold text-center mb-4">Login</h1>
        <LoginForm login={login} />
      </section>
    </>
  )
}
