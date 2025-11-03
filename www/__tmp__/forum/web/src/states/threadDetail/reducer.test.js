/**
 * skenario testing
 *
 * - threadDetailReducer function
 *  - should return the initial state when no action is passed
 *  - should return the thread detail when RECEIVE_THREAD_DETAIL action is passed
 *  - should return null when CLEAR_THREAD_DETAIL action is passed
 *  - should return the thread detail with user id in upVotesBy when UPVOTE_THREAD_DETAIL action is passed
 *  - should return the thread detail with user id in downVotesBy when DOWNVOTE_THREAD_DETAIL action is passed
 *  - should return the thread detail with user id removed from upVotesBy or downVotesBy when UNVOTE_THREAD_DETAIL action is passed
 *  - should return the thread detail with added comment when ADD_THREAD_DETAIL_COMMENT action is passed
 *  - should return the thread detail with user id in comment upVotesBy when UPVOTE_THREAD_DETAIL_COMMENT action is passed
 *  - should return the thread detail with user id in comment downVotesBy when DOWNVOTE_THREAD_DETAIL_COMMENT action is passed
 *  - should return the thread detail with user id removed from comment upVotesBy or downVotesBy when UNVOTE_THREAD_DETAIL_COMMENT action is passed
 */
import { describe, it, expect } from 'vitest'
import threadDetailReducer from './reducer'

describe('threadDetailReducer function', () => {
  it('should return the initial state when no action is passed', () => {
    const initialState = null
    const action = {
      type: 'ANY_ACTION'
    }

    const nextState = threadDetailReducer(initialState, action)

    expect(nextState).toEqual(initialState)
  })

  it('should return the thread detail when RECEIVE_THREAD_DETAIL action is passed', () => {
    const initialState = null
    const action = {
      type: 'RECEIVE_THREAD_DETAIL',
      payload: {
        threadDetail: {
          id: 'thread-1',
          title: 'Thread Pertama',
          body: 'Ini adalah thread pertama',
          category: 'General',
          createdAt: '2021-06-21T07:00:00.000Z',
          owner: {
            id: 'users-1',
            name: 'John Doe',
            avatar: 'https://generated-image-url.jpg'
          },
          upVotesBy: [],
          downVotesBy: [],
          comments: []
        }
      }
    }

    const nextState = threadDetailReducer(initialState, action)

    expect(nextState).toEqual(action.payload.threadDetail)
  })

  it('should return null when CLEAR_THREAD_DETAIL action is passed', () => {
    const initialState = {
      id: 'thread-1',
      title: 'Thread Pertama',
      body: 'Ini adalah thread pertama',
      category: 'General',
      createdAt: '2021-06-21T07:00:00.000Z',
      ownerId: 'users-1',
      upVotesBy: [],
      downVotesBy: [],
      totalComments: 0,
      comments: []
    }
    const action = {
      type: 'CLEAR_THREAD_DETAIL'
    }

    const nextState = threadDetailReducer(initialState, action)

    expect(nextState).toBeNull()
  })

  it('should return the thread detail with user id in upVotesBy when UPVOTE_THREAD_DETAIL action is passed', () => {
    const initialState = {
      id: 'thread-1',
      title: 'Thread Pertama',
      body: 'Ini adalah thread pertama',
      category: 'General',
      createdAt: '2021-06-21T07:00:00.000Z',
      ownerId: 'users-1',
      upVotesBy: [],
      downVotesBy: [],
      totalComments: 0,
      comments: []
    }
    const action = {
      type: 'UPVOTE_THREAD_DETAIL',
      payload: { userId: 'users-2' }
    }

    const nextState = threadDetailReducer(initialState, action)

    expect(nextState).toEqual({
      ...initialState,
      upVotesBy: ['users-2']
    })
  })

  it('should return the thread detail with user id in downVotesBy when DOWNVOTE_THREAD_DETAIL action is passed', () => {
    const initialState = {
      id: 'thread-1',
      title: 'Thread Pertama',
      body: 'Ini adalah thread pertama',
      category: 'General',
      createdAt: '2021-06-21T07:00:00.000Z',
      ownerId: 'users-1',
      upVotesBy: [],
      downVotesBy: [],
      totalComments: 0,
      comments: []
    }
    const action = {
      type: 'DOWNVOTE_THREAD_DETAIL',
      payload: { userId: 'users-2' }
    }

    const nextState = threadDetailReducer(initialState, action)

    expect(nextState).toEqual({
      ...initialState,
      downVotesBy: ['users-2']
    })
  })

  it('should return the thread detail with user id removed from upVotesBy or downVotesBy when UNVOTE_THREAD_DETAIL action is passed', () => {
    const initialState = {
      id: 'thread-1',
      title: 'Thread Pertama',
      body: 'Ini adalah thread pertama',
      category: 'General',
      createdAt: '2021-06-21T07:00:00.000Z',
      ownerId: 'users-1',
      upVotesBy: ['users-2'],
      downVotesBy: ['users-3'],
      totalComments: 0,
      comments: []
    }
    const action = {
      type: 'UNVOTE_THREAD_DETAIL',
      payload: { userId: 'users-2' }
    }

    const nextState = threadDetailReducer(initialState, action)

    const action2 = {
      type: 'UNVOTE_THREAD_DETAIL',
      payload: { userId: 'users-3' }
    }

    const nextState2 = threadDetailReducer(nextState, action2)

    expect(nextState).toEqual({
      ...initialState,
      upVotesBy: []
    })
    expect(nextState2).toEqual({
      ...nextState,
      downVotesBy: []
    })
  })

  it('should return the thread detail with added comment when ADD_THREAD_DETAIL_COMMENT action is passed', () => {
    const initialState = {
      id: 'thread-1',
      title: 'Thread Pertama',
      body: 'Ini adalah thread pertama',
      category: 'General',
      createdAt: '2021-06-21T07:00:00.000Z',
      ownerId: 'users-1',
      upVotesBy: [],
      downVotesBy: [],
      comments: []
    }
    const action = {
      type: 'ADD_THREAD_DETAIL_COMMENT',
      payload: {
        comment: {
          id: 'comment-1',
          content: 'Ini adalah komentar pertama',
          createdAt: '2021-06-21T07:00:00.000Z',
          upVotesBy: [],
          downVotesBy: [],
          owner: {
            id: 'users-1',
            name: 'John Doe',
            email: 'john@example.com'
          }
        }
      }
    }

    const nextState = threadDetailReducer(initialState, action)

    expect(nextState).toEqual({
      ...initialState,
      comments: [action.payload.comment, ...initialState.comments]
    })
  })

  it('should return the thread detail with user id in comment upVotesBy when UPVOTE_THREAD_DETAIL_COMMENT action is passed', () => {
    const initialState = {
      id: 'thread-1',
      title: 'Thread Pertama',
      body: 'Ini adalah thread pertama',
      category: 'General',
      createdAt: '2021-06-21T07:00:00.000Z',
      ownerId: 'users-1',
      upVotesBy: [],
      downVotesBy: [],
      comments: [
        {
          id: 'comment-1',
          content: 'Ini adalah komentar pertama',
          createdAt: '2021-06-21T07:00:00.000Z',
          owner: {
            id: 'users-1',
            name: 'John Doe',
            email: 'john@example.com'
          },
          upVotesBy: [],
          downVotesBy: []
        }
      ]
    }

    const action = {
      type: 'UPVOTE_THREAD_DETAIL_COMMENT',
      payload: { commentId: 'comment-1', userId: 'users-2' }
    }

    const nextState = threadDetailReducer(initialState, action)

    expect(nextState).toEqual({
      ...initialState,
      comments: [{ ...initialState.comments[0], upVotesBy: ['users-2'] }]
    })
  })

  it('should return the thread detail with user id in comment downVotesBy when DOWNVOTE_THREAD_DETAIL_COMMENT action is passed', () => {
    const initialState = {
      id: 'thread-1',
      title: 'Thread Pertama',
      body: 'Ini adalah thread pertama',
      category: 'General',
      createdAt: '2021-06-21T07:00:00.000Z',
      ownerId: 'users-1',
      upVotesBy: [],
      downVotesBy: [],
      comments: [
        {
          id: 'comment-1',
          content: 'Ini adalah komentar pertama',
          createdAt: '2021-06-21T07:00:00.000Z',
          owner: {
            id: 'users-1',
            name: 'John Doe',
            email: 'john@example.com'
          },
          upVotesBy: [],
          downVotesBy: []
        }
      ]
    }

    const action = {
      type: 'DOWNVOTE_THREAD_DETAIL_COMMENT',
      payload: { commentId: 'comment-1', userId: 'users-2' }
    }

    const nextState = threadDetailReducer(initialState, action)

    expect(nextState).toEqual({
      ...initialState,
      comments: [{ ...initialState.comments[0], downVotesBy: ['users-2'] }]
    })
  })

  it('should return the thread detail with user id removed from comment upVotesBy or downVotesBy when UNVOTE_THREAD_DETAIL_COMMENT action is passed', () => {
    const initialState = {
      id: 'thread-1',
      title: 'Thread Pertama',
      body: 'Ini adalah thread pertama',
      category: 'General',
      createdAt: '2021-06-21T07:00:00.000Z',
      ownerId: 'users-1',
      upVotesBy: [],
      downVotesBy: [],
      comments: [
        {
          id: 'comment-1',
          content: 'Ini adalah komentar pertama',
          createdAt: '2021-06-21T07:00:00.000Z',
          owner: {
            id: 'users-1',
            name: 'John Doe',
            email: 'john@example.com'
          },
          upVotesBy: ['users-2'],
          downVotesBy: ['users-3']
        }
      ]
    }

    const action = {
      type: 'UNVOTE_THREAD_DETAIL_COMMENT',
      payload: { commentId: 'comment-1', userId: 'users-2' }
    }

    const nextState = threadDetailReducer(initialState, action)

    const action2 = {
      type: 'UNVOTE_THREAD_DETAIL_COMMENT',
      payload: { commentId: 'comment-1', userId: 'users-3' }
    }

    const nextState2 = threadDetailReducer(nextState, action2)

    expect(nextState).toEqual({
      ...initialState,
      comments: [{ ...initialState.comments[0], upVotesBy: [] }]
    })
    expect(nextState2).toEqual({
      ...nextState,
      comments: [{ ...nextState.comments[0], downVotesBy: [] }]
    })
  })
})
