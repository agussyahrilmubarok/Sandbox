/**
 * skenario testing
 *
 * - asyncPreloadProcess thunk
 *  - should dispatch actions correctly when getUserOwnProfile is successful
 *  - should dispatch actions correctly when getUserOwnProfile is failed
 */
import { describe, expect, it, vi, beforeEach, afterEach } from 'vitest'
import { api } from '../../lib/api'
import { asyncPreloadProcess, setIsPreload } from './action'
import { setAuthUser } from '../authUser/action'
import { showLoading, hideLoading } from 'react-redux-loading-bar'

const fakeSuccessResponse = {
  id: 'john_doe',
  name: 'John Doe',
  email: 'john@example.com',
  avatar: 'https://generated-image-url.jpg'
}

const fakeErrorResponse = new Error('Failed to fetch user profile')

describe('asyncPreloadProcess thunk', () => {
  const dispatch = vi.fn()

  beforeEach(() => {
    api._getUserOwnProfile = api.getUserOwnProfile
  })

  afterEach(() => {
    dispatch.mockClear()
    api.getUserOwnProfile = api._getUserOwnProfile
    delete api._getUserOwnProfile
  })

  it('should dispatch actions correctly when getUserOwnProfile is successful', async () => {
    api.getUserOwnProfile = () => Promise.resolve(fakeSuccessResponse)

    await asyncPreloadProcess()(dispatch)

    expect(dispatch).toHaveBeenCalledWith(setAuthUser(fakeSuccessResponse))
    expect(dispatch).toHaveBeenCalledWith(setIsPreload(false))
    expect(dispatch).toHaveBeenCalledWith(showLoading())
    expect(dispatch).toHaveBeenCalledWith(hideLoading())
  })

  it('should dispatch actions correctly when getUserOwnProfile is failed', async () => {
    api.getUserOwnProfile = () => Promise.reject(fakeErrorResponse)

    await asyncPreloadProcess()(dispatch)

    expect(dispatch).toHaveBeenCalledWith(setAuthUser(null))
    expect(dispatch).toHaveBeenCalledWith(setIsPreload(false))
    expect(dispatch).toHaveBeenCalledWith(showLoading())
    expect(dispatch).toHaveBeenCalledWith(hideLoading())
  })
})
