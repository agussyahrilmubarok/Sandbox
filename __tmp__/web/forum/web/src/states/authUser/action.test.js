/**
 * skenario testing
 *
 * - asyncSetAuthUser thunk
 *  - should dispatch actions and toast success correctly when login is successful
 *  - should dispatch actions and toast error correctly when login is failed
 *
 * - asyncUnsetAuthUser thunk
 *  - should dispatch actions, remove access token, and toast success correctly when user is logged out
 */
import { describe, it, expect, vi, beforeEach, afterEach } from 'vitest'
import { api } from '../../lib/api'
import { asyncSetAuthUser, asyncUnsetAuthUser, setAuthUser, unsetAuthUser } from './action'
import { showLoading, hideLoading } from 'react-redux-loading-bar'
import toast from 'react-hot-toast'

const fakeUserProfileResponse = {
  id: 'john_doe',
  name: 'John Doe',
  email: 'john@example.com',
  avatar: 'https://generated-image-url.jpg'
}

const fakeLoginResponse = 'userAccessToken'

const fakeErrorResponse = {
  response: {
    data: {
      message: 'login failed'
    }
  }
}

describe('asyncSetAuthUser thunk', () => {
  const dispatch = vi.fn()

  beforeEach(() => {
    api._getuserOwnProfile = api.getUserOwnProfile
    api._login = api.login
    api._putAccessToken = api.putAccessToken
    api._removeAccessToken = api.removeAccessToken
  })

  afterEach(() => {
    dispatch.mockClear()
    api.getUserOwnProfile = api._getuserOwnProfile
    api.login = api._login
    api.putAccessToken = api._putAccessToken
    api.removeAccessToken = api._removeAccessToken
    delete api._getuserOwnProfile
    delete api._login
    delete api._putAccessToken
    delete api._removeAccessToken
  })

  it('should dispatch actions and toast success correctly when login is successful', async () => {
    let localAccessToken
    const payload = {
      email: 'john@example.com',
      password: 'password'
    }

    const spy = vi.spyOn(toast, 'success')
    api.login = () => Promise.resolve(fakeLoginResponse)
    api.getUserOwnProfile = () => Promise.resolve(fakeUserProfileResponse)
    api.putAccessToken = (token) => (localAccessToken = token)

    await asyncSetAuthUser(payload)(dispatch)

    expect(dispatch).toHaveBeenCalledWith(setAuthUser(fakeUserProfileResponse))
    expect(dispatch).toHaveBeenCalledWith(showLoading())
    expect(dispatch).toHaveBeenCalledWith(hideLoading())
    expect(localAccessToken).toEqual(fakeLoginResponse)
    expect(spy).toHaveBeenCalledWith('Logged in')
  })

  it('should dispatch actions and toast error correctly when login is failed', async () => {
    const spy = vi.spyOn(toast, 'error')
    api.login = vi.fn().mockImplementation(() => Promise.reject(fakeErrorResponse))

    await asyncSetAuthUser({ email: '', password: '' })(dispatch)

    expect(dispatch).toHaveBeenCalledWith(showLoading())
    expect(dispatch).toHaveBeenCalledWith(hideLoading())
    expect(spy).toHaveBeenCalledWith(fakeErrorResponse.response.data.message)
  })
})

describe('asyncUnsetAuthUser thunk', () => {
  beforeEach(() => {
    api._removeAccessToken = api.removeAccessToken
  })

  afterEach(() => {
    api.removeAccessToken = api._removeAccessToken
    delete api._removeAccessToken
  })

  it('should dispatch actions, remove access token, and toast success correctly when user is logged out', async () => {
    let localAccessToken = 'userAccessToken'
    api.removeAccessToken = () => (localAccessToken = null)
    const spy = vi.spyOn(toast, 'success')
    const dispatch = vi.fn()

    await asyncUnsetAuthUser()(dispatch)

    expect(dispatch).toHaveBeenCalledWith(unsetAuthUser())
    expect(dispatch).toHaveBeenCalledWith(showLoading())
    expect(dispatch).toHaveBeenCalledWith(hideLoading())
    expect(localAccessToken).toBeNull()
    expect(spy).toHaveBeenCalledWith('Logged out')
  })
})
