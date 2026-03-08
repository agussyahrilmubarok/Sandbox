import React from 'react'
import { Navigate } from 'react-router-dom'

export default function FallbackPage() {
  return <Navigate to="/" />
}
