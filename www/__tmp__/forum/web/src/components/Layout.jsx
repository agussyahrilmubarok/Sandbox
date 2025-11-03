import React from 'react'
import AppHeader from './app/AppHeader'
import AppNavigation from './app/AppNavigation'
import AppToaster from './app/AppToaster'
import PropTypes from 'prop-types'
import AppLoadingBar from './app/AppLoadingBar'

export default function Layout({ children }) {
  return (
    <>
      <AppLoadingBar />
      <AppToaster />
      <main>
        <AppHeader />
        <div className="max-w-2xl m-auto pt-16 min-h-dvh px-2.5 border border-neutral">{children}</div>
        <AppNavigation />
      </main>
    </>
  )
}

Layout.propTypes = {
  children: PropTypes.node.isRequired
}
