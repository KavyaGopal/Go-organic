 /// <reference types="cypress" />
 import Chance from 'chance';
 const chance = new Chance();

 describe('Tests Cart Component', () => {
     // check if the cart component is loading
     it('Visits the Go organic PayPal Payment Component', () => {
         cy.visit('http://localhost:4200/paypalpayment')
     })

     // check if the payment overview has following elements
     it('Checks if payment methods window has required fields ', () => {
         cy.visit('http://localhost:4200/paypalpayment')
         cy.contains('Overview')
         cy.contains('Total')
         cy.contains('Make Payment')
         cy.get('.btn').click();

     });

 });