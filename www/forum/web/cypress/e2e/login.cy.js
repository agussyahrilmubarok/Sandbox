/**
 * - Login spec
 *  - should display login page correctly
 *  - should display toast error when email is not valid
 *  - should display toast error when email or password is wrong
 *  - should display toast success and redirect to home page when login succeeds
 */
const inputData = {
  email: 'testuser@example.com',
  password: 'test4321'
}

describe('Login spec', () => {
  beforeEach(() => {
    cy.visit('/login')
  })

  it('should display login page correctly', () => {
    cy.get('input[name="email"]').should('be.visible')
    cy.get('input[name="password"]').should('be.visible')
    cy.get('button[type="submit"]').should('be.visible')
  })

  it('should display toast error when email is not valid', () => {
    cy.get('input[name="email"]').type('testuser@example')
    cy.get('input[name="password"]').type('password')
    cy.get('button[type="submit"]').click()

    cy.get('.app-toast').should('be.visible')
    cy.get('.app-toast').should('contain.text', '"email" must be a valid email')
  })

  it('should display toast error when email or password is wrong', () => {
    cy.get('input[name="email"]').type(inputData.email)
    cy.get('input[name="password"]').type('wrongpassword')
    cy.get('button[type="submit"]').click()

    cy.get('.app-toast').should('be.visible')
    cy.get('.app-toast').should('contain.text', 'email or password is wrong')
  })

  it('should display toast success and redirect to home page when login succeeds', () => {
    cy.get('input[name="email"]').type(inputData.email)
    cy.get('input[name="password"]').type(inputData.password)
    cy.get('button[type="submit"]').click()

    cy.get('.app-toast').should('be.visible')
    cy.get('.app-toast').should('contain.text', 'Logged in')
    cy.location('pathname').should('eq', '/')
  })
})
