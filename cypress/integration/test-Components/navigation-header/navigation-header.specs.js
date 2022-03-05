 /// <reference types="cypress" />
 import Chance from 'chance';
 const chance = new Chance();

 describe('Tests navigation-header Component', ()=> {
     // check if the home component is loading
     it('Visits the Go organic Navigation-Header Component', () => {
       cy.visit('http://localhost:4200/navigation-header')
     });

     // check if the footer contains following fields
     it('Checks the footer of navigation header', () => {
      cy.visit('http://localhost:4200/navigation-header')
      cy.contains('Quick Links')
      cy.contains('Shop Now')
      cy.contains('Contact us')
      cy.contains('Visit Go-organic')
      cy.contains('About Go-organic')
    });
 });