/**
 * skenario testing
 *
 * - categoriesReducer function
 *  - should return the initial state when no action is passed
 *  - should return the categories data when RECEIVE_CATEGORIES action is passed
 *  - should return the categories' data with the new category when ADD_CATEGORY action is passed
 */
import { describe, it, expect } from 'vitest'
import categoriesReducer from './reducer'

describe('categoriesReducer function', () => {
  it('should return the initial state when no action is passed', () => {
    const initialState = []

    const action = {
      type: 'ANY_ACTION'
    }

    const nextState = categoriesReducer(initialState, action)

    expect(nextState).toEqual(initialState)
  })

  it('should return the categories data when RECEIVE_CATEGORIES action is passed', () => {
    const initialState = []

    const action = {
      type: 'RECEIVE_CATEGORIES',
      payload: {
        categories: ['Category 1', 'Category 2']
      }
    }

    const nextState = categoriesReducer(initialState, action)

    expect(nextState).toEqual(action.payload.categories)
  })

  it("should return the categories' data with the new category when ADD_CATEGORY action is passed", () => {
    const initialState = ['Category 1', 'Category 2']

    const action = {
      type: 'ADD_CATEGORY',
      payload: {
        category: 'Category 3'
      }
    }

    const nextState = categoriesReducer(initialState, action)

    expect(nextState).toEqual([action.payload.category, ...initialState])
  })
})
