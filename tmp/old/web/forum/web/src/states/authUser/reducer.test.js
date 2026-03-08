/**
 * skenario testing
 *
 * - authUserReducer function
 *  - should return the initial state when no action is passed
 *  - should return the user's data when SET_AUTH_USER action is passed
 *  - should return null when UNSET_AUTH_USER action is passed
 */
import { describe, it, expect } from 'vitest'
import authUserReducer from './reducer'

describe('authUserReducer function', () => {
  it('should return the initial state when no action is passed', () => {
    const initialState = null
    const action = {
      type: 'ANY_ACTION'
    }

    const nextState = authUserReducer(initialState, action)

    expect(nextState).toEqual(initialState)
  })

  it("should return the user's data when SET_AUTH_USER action is passed", () => {
    const action = {
      type: 'SET_AUTH_USER',
      payload: {
        user: {
          id: 1,
          name: 'John Doe',
          email: 'john@example.com',
          avatar: 'https://ui-avatars.com/api/?name=John+doe'
        }
      }
    }

    const nextState = authUserReducer(null, action)

    expect(nextState).toEqual(action.payload.user)
  })

  it('should return null when UNSET_AUTH_USER action is passed', () => {
    const initialState = {
      id: 1,
      name: 'John Doe',
      email: 'john@example.com',
      avatar: 'https://generated-image-url.jpg'
    }

    const action = {
      type: 'UNSET_AUTH_USER'
    }

    const nextState = authUserReducer(initialState, action)

    expect(nextState).toBeNull()
  })
})
