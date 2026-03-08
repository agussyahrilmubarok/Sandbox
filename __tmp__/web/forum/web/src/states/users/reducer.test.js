/**
 * skenario testing
 *
 * - usersReducer function
 *  - should return the initial state when no action is passed
 *  - should return the users data when RECEIVE_USERS action is passed
 */
import { describe, it, expect } from 'vitest'
import usersReducer from './reducer'

describe('usersReducer function', () => {
  it('should return the initial state when no action is passed', () => {
    const initialState = []
    const action = {
      type: 'ANY_ACTION'
    }

    const nextState = usersReducer(initialState, action)

    expect(nextState).toEqual(initialState)
  })

  it('should return the users data when RECEIVE_USERS action is passed', () => {
    const initialState = []
    const action = {
      type: 'RECEIVE_USERS',
      payload: {
        users: [
          {
            id: 'john_doe',
            name: 'John Doe',
            email: 'john@example.com',
            avatar: 'https://generated-image-url.jpg'
          },
          {
            id: 'jane_doe',
            name: 'Jane Doe',
            email: 'jane@example.com',
            avatar: 'https://generated-image-url.jpg'
          }
        ]
      }
    }

    const nextState = usersReducer(initialState, action)

    expect(nextState).toEqual(action.payload.users)
  })
})
