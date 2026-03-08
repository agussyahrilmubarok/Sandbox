import React from 'react'
import LoadingBar from 'react-redux-loading-bar'

export default function AppLoadingBar() {
  return (
    <div className="sticky top-0 w-full z-50">
      <LoadingBar updateTime={100} maxProgress={98} className="absolute h-0.5 bg-primary z-50" />
    </div>
  )
}
