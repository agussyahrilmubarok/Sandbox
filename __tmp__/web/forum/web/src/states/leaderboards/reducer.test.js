/**
 * skenario testing
 *
 * - leaderboardsReducer function
 *  - should return the initial state when no action is passed
 *  - should return the leaderboards' data when RECEIVE_LEADERBOARDS action is passed
 *  - should return an empty array when CLEAR_LEADERBOARDS action is passed
 */
import { describe, it, expect } from 'vitest'
import leaderboardsReducer from './reducer'

describe('leaderboardsReducer function', () => {
  it('should return the initial state when no action is passed', () => {
    const initialState = []
    const action = {
      type: 'ANY_ACTION'
    }

    const nextState = leaderboardsReducer(initialState, action)

    expect(nextState).toEqual(initialState)
  })

  it("should return the leaderboards' data when RECEIVE_LEADERBOARDS action is passed", () => {
    const initialState = []
    const action = {
      type: 'RECEIVE_LEADERBOARDS',
      payload: {
        leaderboards: [
          {
            user: {
              id: 'users-1',
              name: 'John Doe',
              email: 'john@example.com',
              avatar: 'https://generated-image-url.jpg'
            },
            score: 10
          },
          {
            user: {
              id: 'users-2',
              name: 'Jane Doe',
              email: 'jane@example.com',
              avatar: 'https://generated-image-url.jpg'
            },
            score: 5
          }
        ]
      }
    }

    const nextState = leaderboardsReducer(initialState, action)

    expect(nextState).toEqual(action.payload.leaderboards)
  })

  it('should return an empty array when CLEAR_LEADERBOARDS action is passed', () => {
    const initialState = [
      {
        user: {
          id: 'users-1',
          name: 'John Doe',
          email: 'john@example.com',
          avatar: 'https://generated-image-url.jpg'
        },
        score: 10
      },
      {
        user: {
          id: 'users-2',
          name: 'Jane Doe',
          email: 'jane@example.com',
          avatar: 'https://generated-image-url.jpg'
        },
        score: 5
      }
    ]

    const action = {
      type: 'CLEAR_LEADERBOARDS'
    }

    const nextState = leaderboardsReducer(initialState, action)

    expect(nextState).toEqual([])
  })
})
