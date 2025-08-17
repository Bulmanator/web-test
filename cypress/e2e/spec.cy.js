describe('baisc button click', () => {
  it('passes', () => {
    cy.visit('http://localhost:8080')
		cy.get('#btn_show').then(($btn) => { $btn.click() })
		cy.contains("Hello, World!")
  })
})
