// ***********************************************
// This example commands.js shows you how to
// create various custom commands and overwrite
// existing commands.
//
// For more comprehensive examples of custom
// commands please read more here:
// https://on.cypress.io/custom-commands
// ***********************************************
//
//
// -- This is a parent command --
Cypress.Commands.add('login', (email = 'testuser@example.com', password = 'test4321') => {
  cy.session([email, password], () => {
    cy.request({
      method: 'POST',
      url: 'https://forum-api.dicoding.dev/v1/login',
      body: { email, password }
    }).then(({ body }) => {
      window.localStorage.setItem('accessToken', body.data.token)
    })
  })
})

// -- This is a child command --
// Cypress.Commands.add('drag', { prevSubject: 'element'}, (subject, options) => { ... })
//
//
// -- This is a dual command --
// Cypress.Commands.add('dismiss', { prevSubject: 'optional'}, (subject, options) => { ... })
//
//
// -- This will overwrite an existing command --
// Cypress.Commands.overwrite('visit', (originalFn, url, options) => { ... })
