/**
 * skenario testing
 *
 * - CreateCommentForm component
 *  - should render login required message when user not authenticated
 *  - should handle comment input correctly
 *  - should handle create comment button click correctly
 */
import React from 'react'
import { MemoryRouter } from 'react-router-dom'
import { describe, expect, it, afterEach } from 'vitest'
import { screen, cleanup } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import matchers from '@testing-library/jest-dom/matchers'
import CreateCommentForm from './CreateCommentForm'
import { renderWithProviders } from '../../lib/test-utils'
import { api } from '../../lib/api'

expect.extend(matchers)

const threadDetail = {
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
  comments: [
    {
      id: 'comment-1',
      content: 'Ini adalah komentar pertama',
      createdAt: '2021-06-21T07:00:00.000Z',
      owner: {
        id: 'users-1',
        name: 'John Doe',
        avatar: 'https://generated-image-url.jpg'
      },
      upVotesBy: [],
      downVotesBy: []
    }
  ]
}

describe('CreateCommentForm component', () => {
  afterEach(() => {
    cleanup()
  })

  it('should render login required message when user not authenticated', async () => {
    renderWithProviders(
      <MemoryRouter>
        <CreateCommentForm threadId="thread-1" />
      </MemoryRouter>
    )

    const loginRequiredMessage = screen.getByText('You need to login first to post a comment.')

    expect(loginRequiredMessage).toBeInTheDocument()
  })

  it('should handle comment input correctly', async () => {
    const preloadedState = {
      authUser: { id: 'users-1' },
      threadDetail
    }
    renderWithProviders(<CreateCommentForm threadId="thread-1" />, {
      preloadedState
    })
    const comment = 'test comment'

    const commentInput = screen.getByText((_, element) => element.attributes.contenteditable)
    await userEvent.type(commentInput, comment)

    expect(commentInput).toHaveTextContent(comment)
  })

  it('should handle create comment button click correctly', async () => {
    const preloadedState = {
      authUser: { id: 'users-1' },
      threadDetail
    }

    api.createComment = ({ id = 'comment-n', content }) => Promise.resolve({ id, content })

    const { store } = renderWithProviders(<CreateCommentForm threadId="thread-1" />, {
      preloadedState
    })
    const comment = 'test comment'

    const commentInput = screen.getByText((_, element) => element.attributes.contenteditable)
    await userEvent.type(commentInput, comment)
    const submitButton = screen.getByRole('button', { name: /post/i })
    await userEvent.click(submitButton)

    const states = store.getState()

    expect(states.threadDetail.comments).toHaveLength(2)
    expect(states.threadDetail.comments[1].content).toEqual(comment)
    expect(commentInput).toHaveTextContent('')
  })
})
