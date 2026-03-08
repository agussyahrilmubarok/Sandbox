/**
 * - createThread spec
 *  - should display a form to create a new thread correctly
 *  - should focus on the title input when the form is submitted without any input
 *  - should focus on the category input when the form is submitted without any input
 *  - should display toast error when the form is submitted without any input
 *  - should display toast success when the form is submitted with valid input
 */
const inputData = {
  title: 'Awesome thread',
  category: 'testing-e2e',
  content: 'This is an awesome thread'
}

describe('createThread spec', () => {
  beforeEach(() => {
    cy.login()
    cy.visit('/')
    cy.get('a[href="/threads/new"]').click()
  })

  it('should display a form to create a new thread correctly', () => {
    cy.get('input[placeholder*="title" i]').should('be.visible')
    cy.get('input[placeholder*="category" i]').should('be.visible')
    cy.get('div[contentEditable]').should('be.visible')
    cy.get('button[type="submit"]').should('be.visible')
  })

  it('should focus on the title input when the form is submitted without any input', () => {
    cy.get('button[type="submit"]').click()

    cy.get("input[placeholder*='title' i]").should('have.focus')
  })

  it('should focus on the category input when the form is submitted without any input', () => {
    cy.get("input[placeholder*='title' i]").type(inputData.title)
    cy.get('button[type="submit"]').click()

    cy.get("input[placeholder*='category' i]").should('have.focus')
  })

  it('should display toast error when the form is submitted without any input', () => {
    cy.get("input[placeholder*='title' i]").type(inputData.title)
    cy.get('input[placeholder*="category" i]').type(inputData.category)
    cy.get('button[type="submit"]').click()

    cy.get('.app-toast').should('be.visible')
  })

  it('should display toast success when the form is submitted with valid input', () => {
    cy.get("input[placeholder*='title' i]").type(inputData.title)
    cy.get('input[placeholder*="category" i]').type(inputData.category)
    cy.get('div[contentEditable]').type(inputData.content)
    cy.get('button[type="submit"]').click()

    cy.get('.app-toast').should('be.visible')
    cy.get('.app-toast').should('contain.text', 'Thread has been posted')
  })
})
