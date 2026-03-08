/**
 * skenario testing
 *
 * - asyncRegisterUser function
 *  - should dispatch actions and toast success correctly when register is successful
 *  - should dispatch actions and toast error correctly when register is failed
 */
import { describe, expect, it, vi, beforeEach, afterEach } from 'vitest'
import { api } from '../../lib/api'
import { asyncRegisterUser } from './action'
import { showLoading, hideLoading } from 'react-redux-loading-bar'
import toast from 'react-hot-toast'

const fakeErrorResponse = {
  response: {
    data: {
      message: 'register failed'
    }
  }
}

describe('asyncRegisterUser function', () => {
  const dispatch = vi.fn()

  beforeEach(() => {
    api._register = api.register
  })

  afterEach(() => {
    dispatch.mockClear()
    api.register = api._register
    delete api._register
  })

  it('should dispatch actions and toast success correctly when register is successful', async () => {
    const payload = {
      name: 'John Doe',
      email: 'john@example.com',
      password: 'password'
    }

    const spy = vi.spyOn(toast, 'success')
    api.register = () => Promise.resolve()

    const result = await asyncRegisterUser(payload)(dispatch)

    expect(dispatch).toHaveBeenCalledWith(showLoading())
    expect(dispatch).toHaveBeenCalledWith(hideLoading())
    expect(spy).toHaveBeenCalledWith('successfully registered')
    expect(result).toEqual({ error: false })
  })

  it('should dispatch actions and toast error correctly when register is failed', async () => {
    const payload = {
      name: 'John Doe',
      email: 'john@example.com',
      password: 'password'
    }

    const spy = vi.spyOn(toast, 'error')
    api.register = () => Promise.reject(fakeErrorResponse)

    const result = await asyncRegisterUser(payload)(dispatch)

    expect(spy).toHaveBeenCalledWith(fakeErrorResponse.response.data.message)
    expect(dispatch).toHaveBeenCalledWith(showLoading())
    expect(dispatch).toHaveBeenCalledWith(hideLoading())
    expect(result).toEqual({ error: true })
  })
})
