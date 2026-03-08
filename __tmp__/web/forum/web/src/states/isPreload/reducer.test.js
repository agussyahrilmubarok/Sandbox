/**
 * skenario testing
 *
 * - isPreloadReducer function
 *  - should return the initial state when no action is passed
 *  - should return the new isPreload value when SET_IS_PRELOAD action is passed
 */
import { describe, it, expect } from 'vitest'
import isPreloadReducer from './reducer'

describe('isPreloadReducer function', () => {
  it('should return the initial state when no action is passed', () => {
    const initialState = true
    const action = {
      type: 'ANY_ACTION'
    }

    const nextState = isPreloadReducer(initialState, action)

    expect(nextState).toEqual(initialState)
  })

  it('should return the new isPreload value when SET_IS_PRELOAD action is passed', () => {
    const initialState = true
    const action = {
      type: 'SET_IS_PRELOAD',
      payload: {
        isPreload: false
      }
    }

    const nextState = isPreloadReducer(initialState, action)

    expect(nextState).toEqual(action.payload.isPreload)
  })
})
