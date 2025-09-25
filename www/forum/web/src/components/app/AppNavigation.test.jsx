/**
 * skenario testing
 *
 * - AppNavigation component
 *  - should render navigation correctly when user not authenticated
 *  - should render navigation correctly when user authenticated
 *  - should handle home navigation click correctly
 *  - should handle leaderboard navigation click correctly
 *  - should handle login navigation click correctly
 *  - should handle create thread navigation click correctly
 *  - should handle logout button click correctly
 */
import React from 'react'
import { describe, expect, it, afterEach } from 'vitest'
import { MemoryRouter, Route, Routes } from 'react-router-dom'
import { screen, cleanup } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import matchers from '@testing-library/jest-dom/matchers'
import AppNavigation from './AppNavigation'
import { renderWithProviders } from '../../lib/test-utils'

expect.extend(matchers)

const parsePathname = (url) => new URL(url).pathname

describe('AppNavigation component', () => {
  afterEach(() => {
    cleanup()
  })

  it('should render navigation correctly when user not authenticated', async () => {
    const expectedPaths = ['/', '/leaderboards', '/login']

    renderWithProviders(
      <MemoryRouter>
        <AppNavigation />
      </MemoryRouter>
    )

    const navButtons = screen.getAllByRole('button')
    const navigationPathnames = navButtons.map((button) => parsePathname(button.closest('a').href))

    expect(navButtons).toHaveLength(3)
    for (const pathname of expectedPaths) {
      expect(navigationPathnames).toContain(pathname)
    }
  })

  it('should render navigation correctly when user authenticated', async () => {
    const preloadedState = {
      authUser: { id: 'users-1' }
    }
    const expectedPaths = ['/', '/leaderboards', '/threads/new']

    renderWithProviders(
      <MemoryRouter>
        <AppNavigation />
      </MemoryRouter>,
      { preloadedState }
    )

    const navButtons = screen.getAllByRole('button')
    const logoutButton = navButtons.at(-1)
    const navigationPathnames = navButtons.map(
      (button) => button.closest('a') && parsePathname(button.closest('a').href)
    )

    expect(navButtons).toHaveLength(4)
    expect(logoutButton.dataset.tip).toEqual('Logout')
    for (const pathname of expectedPaths) {
      expect(navigationPathnames).toContain(pathname)
    }
  })

  it('should handle home navigation click correctly', async () => {
    renderWithProviders(
      <MemoryRouter initialEntries={['/leaderboards']}>
        <Routes>
          <Route path="/" element={<div data-testid="home">Home</div>} />
          <Route path="/leaderboards" element={<div>Leaderboards</div>} />
        </Routes>
        <AppNavigation />
      </MemoryRouter>
    )

    const navButtons = screen.getAllByRole('button')
    const homeButton = navButtons.find((button) => button.dataset.tip === 'Home')
    await userEvent.click(homeButton)

    const homePage = screen.getByTestId('home')

    expect(homePage).toBeInTheDocument()
  })

  it('should handle leaderboard navigation click correctly', async () => {
    renderWithProviders(
      <MemoryRouter initialEntries={['/']}>
        <Routes>
          <Route path="/" element={<div>Home</div>} />
          <Route path="/leaderboards" element={<div data-testid="leaderboards">Leaderboards</div>} />
        </Routes>
        <AppNavigation />
      </MemoryRouter>
    )

    const navButtons = screen.getAllByRole('button')
    const leaderboardButton = navButtons.find((button) => button.dataset.tip === 'Leaderboards')
    await userEvent.click(leaderboardButton)

    const leaderboardPage = screen.getByTestId('leaderboards')

    expect(leaderboardPage).toBeInTheDocument()
  })

  it('should handle login navigation click correctly', async () => {
    renderWithProviders(
      <MemoryRouter initialEntries={['/']}>
        <Routes>
          <Route path="/" element={<div>Home</div>} />
          <Route path="/login" element={<div data-testid="login">Login Page</div>} />
        </Routes>
        <AppNavigation />
      </MemoryRouter>
    )

    const navButtons = screen.getAllByRole('button')
    const loginButton = navButtons.find((button) => button.dataset.tip === 'Login')
    await userEvent.click(loginButton)

    const loginPage = screen.getByTestId('login')

    expect(loginPage).toBeInTheDocument()
  })

  it('should handle create thread navigation click correctly', async () => {
    const preloadedState = {
      authUser: { id: 'users-1' }
    }

    renderWithProviders(
      <MemoryRouter initialEntries={['/']}>
        <Routes>
          <Route path="/" element={<div>Home</div>} />
          <Route path="/threads/new" element={<div data-testid="create-thread">Create Thread Page</div>} />
        </Routes>
        <AppNavigation />
      </MemoryRouter>,
      { preloadedState }
    )

    const navButtons = screen.getAllByRole('button')
    const createThreadButton = navButtons.find((button) => button.dataset.tip === 'Create Thread')
    await userEvent.click(createThreadButton)

    const createThreadPage = screen.getByTestId('create-thread')

    expect(createThreadPage).toBeInTheDocument()
  })

  it('should handle logout button click correctly', async () => {
    const preloadedState = {
      authUser: { id: 'users-1' }
    }

    const { store } = renderWithProviders(
      <MemoryRouter>
        <AppNavigation />
      </MemoryRouter>,
      { preloadedState }
    )

    const navButtons = screen.getAllByRole('button')
    const logoutButton = navButtons.find((button) => button.dataset.tip === 'Logout')
    await userEvent.click(logoutButton)

    const nextNavButtons = screen.getAllByRole('button')
    const loginButton = nextNavButtons.find((button) => button.dataset.tip === 'Login')

    const state = store.getState()

    expect(state.authUser).toBeNull()
    expect(logoutButton).not.toBeInTheDocument()
    expect(loginButton).toBeInTheDocument()
  })
})
