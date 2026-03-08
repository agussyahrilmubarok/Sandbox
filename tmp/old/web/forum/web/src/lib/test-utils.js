import React from 'react'
import { render } from '@testing-library/react'
// import { configureStore } from '@reduxjs/toolkit'
import { Provider } from 'react-redux'
import { setupStore } from '../states'

export function renderWithProviders(
  ui,
  { preloadedState = {}, store = setupStore(preloadedState), ...renderOptions } = {}
) {
  // eslint-disable-next-line react/prop-types
  function Wrapper({ children }) {
    return React.createElement(Provider, { store }, children)
  }

  return { store, ...render(ui, { wrapper: Wrapper, ...renderOptions }) }
}
