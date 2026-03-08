/**
 * skenario testing
 *
 * - ThreadDetail component
 *  - should render thread detail correctly
 *  - vote buttons should be disabled when user not authenticated
 *  - vote buttons should be not disabled when user authenticated
 *  - should handle up vote button click correctly
 *  - should handle down vote button click correctly
 */
import React from 'react'
import { describe, expect, it, afterEach } from 'vitest'
import { screen, cleanup } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import matchers from '@testing-library/jest-dom/matchers'
import ThreadDetail from './ThreadDetail'
import { postedAt } from '../../lib/utils'
import { renderWithProviders } from '../../lib/test-utils'

expect.extend(matchers)

const threadDetailProps = {
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

describe('ThreadDetail component', () => {
  afterEach(() => {
    cleanup()
  })

  it('should render thread detail correctly', () => {
    const preloadedState = {
      threadDetail: threadDetailProps
    }

    renderWithProviders(<ThreadDetail {...threadDetailProps} />, {
      preloadedState
    })

    const buttons = screen.getAllByRole('button')

    const contents = [
      screen.getByText(`#${threadDetailProps.category}`),
      screen.getByText(postedAt(threadDetailProps.createdAt)),
      screen.getByText(threadDetailProps.owner.name),
      screen.getByAltText(threadDetailProps.owner.name),
      screen.getByText(threadDetailProps.title),
      screen.getByText(threadDetailProps.body),
      screen.getByText(/[0-9] comments/i),
      ...buttons
    ]

    contents.forEach((content) => {
      expect(content).toBeInTheDocument()
    })
  })

  it('vote buttons should be disabled when user not authenticated', async () => {
    renderWithProviders(<ThreadDetail {...threadDetailProps} />)

    const [upVoteButton, downVoteButton] = screen.getAllByRole('button')

    expect(upVoteButton).toBeDisabled()
    expect(downVoteButton).toBeDisabled()
  })

  it('vote buttons should be not disabled when user authenticated', async () => {
    const preloadedState = {
      authUser: { id: 'users-1' }
    }

    renderWithProviders(<ThreadDetail {...threadDetailProps} />, {
      preloadedState
    })

    const [upVoteButton, downVoteButton] = screen.getAllByRole('button')

    expect(upVoteButton).not.toBeDisabled()
    expect(downVoteButton).not.toBeDisabled()
  })

  it('should handle up vote button click correctly', async () => {
    const preloadedState = {
      authUser: { id: 'users-1' },
      threadDetail: threadDetailProps
    }

    const { store } = renderWithProviders(<ThreadDetail {...threadDetailProps} />, {
      preloadedState
    })

    const [upVoteButton] = screen.getAllByRole('button')
    await userEvent.click(upVoteButton)

    const { threadDetail, authUser } = store.getState()

    expect(threadDetail.upVotesBy).toContainEqual(authUser.id)
  })

  it('should handle down vote button click correctly', async () => {
    const preloadedState = {
      authUser: { id: 'users-1' },
      threadDetail: threadDetailProps
    }

    const { store } = renderWithProviders(<ThreadDetail {...threadDetailProps} />, {
      preloadedState
    })

    const [, downVoteButton] = screen.getAllByRole('button')
    await userEvent.click(downVoteButton)

    const { threadDetail, authUser } = store.getState()

    expect(threadDetail.downVotesBy).toContainEqual(authUser.id)
  })
})
