import React from 'react'
import { Toaster } from 'react-hot-toast'

export default function AppToaster() {
  return (
    <Toaster
      position="bottom-center"
      containerClassName="mb-16 lg:mb-4"
      toastOptions={{
        className: 'app-toast bg-neutral text-neutral-content'
      }}
    />
  )
}
