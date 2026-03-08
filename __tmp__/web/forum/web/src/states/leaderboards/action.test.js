/**
 * skenario testing
 *
 * - asyncReceiveLeaderboards thunk
 *  - should dispatch actions correctly when data is fetched successfully
 *  - should dispatch actions and toast error correctly when data fetching failed
 */
import { describe, expect, it, vi, beforeEach, afterEach } from 'vitest'
import { api } from '../../lib/api'
import { asyncReceiveLeaderboards, receiveLeaderboards } from './action'
import { showLoading, hideLoading } from 'react-redux-loading-bar'
import toast from 'react-hot-toast'

const fakeleaderboardsResponse = [
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

const fakeErrorResponse = new Error('Failed to fetch leaderboards')

describe('asyncReceiveLeaderboards thunk', () => {
  const dispatch = vi.fn()

  beforeEach(() => {
    api._getLeaderboards = api.getLeaderboards
  })

  afterEach(() => {
    dispatch.mockClear()
    api.getLeaderboards = api._getLeaderboards
    delete api._getLeaderboards
  })

  it('should dispatch actions correctly when data is fetched successfully', async () => {
    api.getLeaderboards = () => Promise.resolve(fakeleaderboardsResponse)

    await asyncReceiveLeaderboards()(dispatch)

    expect(dispatch).toHaveBeenCalledWith(receiveLeaderboards(fakeleaderboardsResponse))
    expect(dispatch).toHaveBeenCalledWith(showLoading())
    expect(dispatch).toHaveBeenCalledWith(hideLoading())
  })

  it('should dispatch actions and toast error correctly when data fetching failed', async () => {
    api.getLeaderboards = () => Promise.reject(fakeErrorResponse)
    const spy = vi.spyOn(toast, 'error')

    await asyncReceiveLeaderboards()(dispatch)

    expect(dispatch).toHaveBeenCalledWith(showLoading())
    expect(dispatch).toHaveBeenCalledWith(hideLoading())
    expect(spy).toHaveBeenCalledWith(fakeErrorResponse.message)
  })
})
