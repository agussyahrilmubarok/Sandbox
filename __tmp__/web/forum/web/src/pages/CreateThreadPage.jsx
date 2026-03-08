import React from 'react'
import CreateThreadForm from '../components/CreateThreadForm'
import { Helmet } from 'react-helmet'

export default function CreateThreadPage() {
  return (
    <>
      <Helmet>
        <title>Create a new thread</title>
      </Helmet>
      <div className="mt-20">
        <CreateThreadForm />
      </div>
    </>
  )
}
