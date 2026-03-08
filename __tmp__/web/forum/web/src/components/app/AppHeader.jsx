import React from 'react'
import { MessagesSquare } from 'lucide-react'

export default function AppHeader() {
  return (
    <header className="navbar fixed z-10 top-0 left-0 right-0 bg-base-100 backdrop-blur-lg shadow-sm border-b border-neutral">
      <div className="container flex justify-center max-w-2xl gap-2">
        <MessagesSquare size={24} className="text-primary" />
        <h1 className="text-xl font-bold">Forum</h1>
      </div>
    </header>
  )
}
