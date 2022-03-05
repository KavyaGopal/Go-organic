 /// <reference types="cypress" />
 import Chance from 'chance';
 const chance = new Chance();

 describe('Tests Cart Component', ()=> {
     // check if the cart component is loading
     it('Visits the Go organic Cart Component', () => {
       cy.visit('http://localhost:4200/cart')
     })

     // check if the cart component has following elements
     it('Checks if the cart has following elements', () => {
      cy.visit('http://localhost:4200/cart')
      cy.contains('ID')
      cy.contains('Name')
      cy.contains('Item')
      cy.contains('Cost')
      cy.contains('Quantity')
      cy.contains('Sub Total')
      cy.contains('Actions')
      cy.contains('Checkout')
    });
 });