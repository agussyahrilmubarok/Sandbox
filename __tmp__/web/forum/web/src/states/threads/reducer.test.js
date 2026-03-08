/**
 * skenario testing
 *
 * - threadsReducer function
 *  - should return the initial state when no action is passed
 *  - should return the threads' data when RECEIVE_THREADS action is passed
 *  - should return the threads' data with new thread when ADD_THREAD action is passed
 *  - should return the threads with user id in upVotesBy when UPVOTE_THREAD action is passed
 *  - should return the threads with user id is removed from downVotesBy when UPVOTE_THREAD action is passed
 *  - should return the threads with user id in downVotesBy when DOWNVOTE_THREAD action is passed
 *  - should return the threads with user id is removed from upVotesBy when DOWNVOTE_THREAD action is passed
 *  - should return the threads with user id is removed from upVotesBy or downVotesBy when UNVOTE_THREAD action is passed
 */
import { describe, it, expect } from 'vitest'
import threadsReducer from './reducer'

describe('threadsReducer function', () => {
  it('should return the initial state when no action is passed', () => {
    const initialState = []
    const action = {
      type: 'ANY_ACTION'
    }

    const nextState = threadsReducer(initialState, action)

    expect(nextState).toEqual(initialState)
  })

  it("should return the threads' data when RECEIVE_THREADS action is passed", () => {
    const initialState = []
    const action = {
      type: 'RECEIVE_THREADS',
      payload: {
        threads: [
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
      }
    }

    const nextState = threadsReducer(initialState, action)

    expect(nextState).toEqual(action.payload.threads)
  })

  it("should return the threads' data with new thread when ADD_THREAD action is passed", () => {
    const initialState = [
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

    const action = {
      type: 'ADD_THREAD',
      payload: {
        thread: {
          id: 'thread-3',
          title: 'Thread Ketiga',
          body: 'Ini adalah thread ketiga',
          category: 'General',
          createdAt: '2021-06-21T07:00:00.000Z',
          ownerId: 'users-3',
          upVotesBy: [],
          downVotesBy: [],
          totalComments: 0
        }
      }
    }

    const nextState = threadsReducer(initialState, action)

    expect(nextState).toEqual([action.payload.thread, ...initialState])
  })

  it('should return the threads with user id in upVotesBy when UPVOTE_THREAD action is passed', () => {
    const initialState = [
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
      }
    ]

    const action = {
      type: 'UPVOTE_THREAD',
      payload: { threadId: 'thread-1', userId: 'users-2' }
    }

    const nextState = threadsReducer(initialState, action)

    expect(nextState).toEqual([
      {
        ...initialState[0],
        upVotesBy: [action.payload.userId]
      }
    ])
    expect(nextState[0].upVotesBy)
  })

  it('should return the threads with user id is removed from downVotesBy when UPVOTE_THREAD action is passed', () => {
    const initialState = [
      {
        id: 'thread-1',
        title: 'Thread Pertama',
        body: 'Ini adalah thread pertama',
        category: 'General',
        createdAt: '2021-06-21T07:00:00.000Z',
        ownerId: 'users-1',
        upVotesBy: [],
        downVotesBy: ['users-2'],
        totalComments: 0
      }
    ]

    const action = {
      type: 'UPVOTE_THREAD',
      payload: { threadId: 'thread-1', userId: 'users-2' }
    }

    const nextState = threadsReducer(initialState, action)

    expect(nextState).toEqual([
      {
        ...initialState[0],
        upVotesBy: [action.payload.userId],
        downVotesBy: []
      }
    ])
  })

  it('should return the threads with user id in downVotesBy when DOWNVOTE_THREAD action is passed', () => {
    const initialState = [
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
      }
    ]

    const action = {
      type: 'DOWNVOTE_THREAD',
      payload: { threadId: 'thread-1', userId: 'users-2' }
    }

    const nextState = threadsReducer(initialState, action)

    expect(nextState).toEqual([
      {
        ...initialState[0],
        downVotesBy: [action.payload.userId]
      }
    ])
  })

  it('should return the threads with user id is removed from upVotesBy when DOWNVOTE_THREAD action is passed', () => {
    const initialState = [
      {
        id: 'thread-1',
        title: 'Thread Pertama',
        body: 'Ini adalah thread pertama',
        category: 'General',
        createdAt: '2021-06-21T07:00:00.000Z',
        ownerId: 'users-1',
        upVotesBy: ['users-2'],
        downVotesBy: [],
        totalComments: 0
      }
    ]

    const action = {
      type: 'DOWNVOTE_THREAD',
      payload: { threadId: 'thread-1', userId: 'users-2' }
    }

    const nextState = threadsReducer(initialState, action)

    expect(nextState).toEqual([
      {
        ...initialState[0],
        upVotesBy: [],
        downVotesBy: [action.payload.userId]
      }
    ])
  })

  it('should return the threads with user id is removed from upVotesBy or downVotesBy when UNVOTE_THREAD action is passed', () => {
    const initialState = [
      {
        id: 'thread-1',
        title: 'Thread Pertama',
        body: 'Ini adalah thread pertama',
        category: 'General',
        createdAt: '2021-06-21T07:00:00.000Z',
        ownerId: 'users-1',
        upVotesBy: ['users-2'],
        downVotesBy: ['users-1'],
        totalComments: 0
      }
    ]

    const action = {
      type: 'UNVOTE_THREAD',
      payload: { threadId: 'thread-1', userId: 'users-2' }
    }

    const nextState = threadsReducer(initialState, action)

    const action2 = {
      type: 'UNVOTE_THREAD',
      payload: { threadId: 'thread-1', userId: 'users-1' }
    }

    const nextState2 = threadsReducer(nextState, action2)

    expect(nextState).toEqual([{ ...initialState[0], upVotesBy: [] }])
    expect(nextState2).toEqual([{ ...nextState[0], downVotesBy: [] }])
  })
})
