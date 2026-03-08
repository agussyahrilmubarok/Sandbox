/**
 * skenario testing
 * - asyncPopulateData thunk
 *  - should dispatch action correctly when fetching data is successful
 *  - should dispatch action correctly when fetching data is failed
 *
 * - asyncCreateThreadAndCategory thunk
 *  - should dispatch action correctly when creating thread with existing category is successful
 *  - should dispatch action correctly when creating thread with new category is successful
 *  - should dispatch action correctly when creating thread is failed
 */
import { describe, expect, it, vi, afterEach, beforeEach } from 'vitest'
import { api } from '../../lib/api'
import { asyncPopulateData, asyncCreateThreadAndCategory } from './action'
import toast from 'react-hot-toast'
import { showLoading, hideLoading } from 'react-redux-loading-bar'
import { receiveCategories, addCategory } from '../categories/action'
import { receiveThreads, addThread } from '../threads/action'
import { receiveUsers } from '../users/action'

const fakeThreadsResponse = [
  {
    id: 'thread-1',
    title: 'Thread Pertama',
    body: 'Ini adalah thread pertama',
    category: 'General',
    createdAt: '2021-06-21T07:00:00.000Z',
    ownerId: 'users-1',
    upVotesBy: [],
    downVotesBy: [],
    totalComments: 0
  },
  {
    id: 'thread-2',
    title: 'Thread Kedua',
    body: 'Ini adalah thread kedua',
    category: 'General',
    createdAt: '2021-06-21T07:00:00.000Z',
    ownerId: 'users-2',
    upVotesBy: [],
    downVotesBy: [],
    totalComments: 0
  }
]

const createThreadObject = ({ title, body, category }) => ({
  id: 'thread-n',
  title,
  body,
  category,
  createdAt: '2021-06-21T07:00:00.000Z',
  ownerId: 'john_doe',
  upVotesBy: [],
  downVotesBy: [],
  totalComments: 0
})

const fakeUsersResponse = [
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

const fakeErrorResponse = new Error('fail')

describe('asyncPopulateData thunk', () => {
  const dispatch = vi.fn()

  beforeEach(() => {
    api._getAllThreads = api.getAllThreads
    api._getAllUsers = api.getAllUsers
  })

  afterEach(() => {
    dispatch.mockClear()
    api.getAllThreads = api._getAllThreads
    api.getAllUsers = api._getAllUsers
    delete api._getAllThreads
    delete api._getAllUsers
  })

  it('should dispatch action correctly when fetching data is successful', async () => {
    api.getAllThreads = () => Promise.resolve(fakeThreadsResponse)
    api.getAllUsers = () => Promise.resolve(fakeUsersResponse)

    await asyncPopulateData()(dispatch)

    expect(dispatch).toHaveBeenCalledWith(showLoading())
    expect(dispatch).toHaveBeenCalledWith(hideLoading())
    expect(dispatch).toHaveBeenCalledWith(receiveCategories(['General']))
    expect(dispatch).toHaveBeenCalledWith(receiveThreads(fakeThreadsResponse))
    expect(dispatch).toHaveBeenCalledWith(receiveUsers(fakeUsersResponse))
  })

  it('should dispatch action correctly when fetching data is failed', async () => {
    api.getAllThreads = () => Promise.reject(fakeErrorResponse)
    api.getAllUsers = () => Promise.reject(fakeErrorResponse)
    const spy = vi.spyOn(toast, 'error')

    await asyncPopulateData()(dispatch)

    expect(dispatch).toHaveBeenCalledWith(showLoading())
    expect(dispatch).toHaveBeenCalledWith(hideLoading())
    expect(spy).toHaveBeenCalledWith(fakeErrorResponse.message)
  })
})

describe('asyncCreateThreadAndCategory thunk', () => {
  const dispatch = vi.fn()

  beforeEach(() => {
    api._createThread = api.createThread
  })

  afterEach(() => {
    dispatch.mockClear()
    api.createThread = api._createThread
    delete api._createThread
  })

  it('should dispatch action correctly when creating thread with existing category is successful', async () => {
    const payload = {
      title: 'Thread Baru',
      body: 'Ini adalah thread baru',
      category: 'General'
    }

    api.createThread = ({ title, body, category }) => Promise.resolve(createThreadObject({ title, body, category }))
    const getState = () => ({ categories: ['General'] })

    const result = await asyncCreateThreadAndCategory(payload)(dispatch, getState)

    expect(dispatch).toHaveBeenCalledWith(showLoading())
    expect(dispatch).toHaveBeenCalledWith(hideLoading())
    expect(dispatch).toHaveBeenCalledWith(addThread(createThreadObject(payload)))
    expect(dispatch).not.toHaveBeenCalledWith(addCategory(payload.category))
    expect(result).toEqual({ error: false })
  })

  it('should dispatch action correctly when creating thread with new category is successful', async () => {
    const payload = {
      title: 'Thread Baru',
      body: 'Ini adalah thread baru',
      category: 'Technology'
    }

    const getState = () => ({ categories: ['General'] })

    const spy = vi.spyOn(toast, 'success')
    api.createThread = ({ title, body, category }) => Promise.resolve(createThreadObject({ title, body, category }))

    const result = await asyncCreateThreadAndCategory(payload)(dispatch, getState)

    expect(dispatch).toHaveBeenCalledWith(showLoading())
    expect(dispatch).toHaveBeenCalledWith(hideLoading())
    expect(dispatch).toHaveBeenCalledWith(addThread(createThreadObject(payload)))
    expect(dispatch).toHaveBeenCalledWith(addCategory(payload.category))
    expect(spy).toHaveBeenCalled()
    expect(result).toEqual({ error: false })
  })

  it('should dispatch action correctly when creating thread is failed', async () => {
    const payload = {
      title: 'Thread Baru',
      body: 'Ini adalah thread baru',
      category: 'General'
    }
    const getState = () => ({ categories: ['General'] })

    const spy = vi.spyOn(toast, 'error')
    api.createThread = () => Promise.reject(fakeErrorResponse)

    const result = await asyncCreateThreadAndCategory(payload)(dispatch, getState)

    expect(dispatch).toHaveBeenCalledWith(showLoading())
    expect(dispatch).toHaveBeenCalledWith(hideLoading())
    expect(spy).toHaveBeenCalledWith(fakeErrorResponse.message)
    expect(result).toEqual({ error: true })
  })
})
