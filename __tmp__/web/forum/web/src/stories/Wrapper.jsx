import React from 'react'
import { Provider } from 'react-redux'
import { MemoryRouter } from 'react-router-dom'
import PropTypes from 'prop-types'
import { setupStore } from '../states'
import '../style/index.css'

export default function Wrapper({ children, states }) {
  return (
    <div className="min-h-screen">
      <Provider store={setupStore(states)}>
        <MemoryRouter>{children}</MemoryRouter>
      </Provider>
    </div>
  )
}

Wrapper.propTypes = {
  children: PropTypes.node.isRequired,
  states: PropTypes.object
}
