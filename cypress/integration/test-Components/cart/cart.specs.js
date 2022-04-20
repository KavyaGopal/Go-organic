 /// <reference types="cypress" />
 import Chance from 'chance';
 const chance = new Chance();

 describe('Tests Cart Component', () => {
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
         cy.contains('Proceed to Payment')
     });

     // checks if button is clickable and it's routing to the next component
     it('Visits the Go organic Cart Component', () => {
         cy.visit('http://localhost:4200/cart')
         cy.get('.center > .mat-focus-indicator > .mat-button-wrapper').click();
         cy.url().should('include', '/checkout')
     })

     it('Tests adding and removing element from cart', () => {
         cy.visit('http://localhost:4200/home')
             // Add items to cart
         cy.get('#slick-slide00 > .card > .card-body > .justify-content-between > .mat-focus-indicator').click();
         cy.get('#slick-slide01 > .card > .card-body > .justify-content-between > .mat-focus-indicator').click();
         cy.visit('http://localhost:4200/cart')
             // Increase the number of items to 3
         cy.get(':nth-child(1) > .cdk-column-itemDesc > div > [aria-label="add"]').click();
         cy.get(':nth-child(1) > .cdk-column-itemDesc > div > [aria-label="add"]').click();
         cy.contains('3')
             // Decrease the number of items to 2
         cy.get(' :nth-child(1) > .cdk-column-itemDesc > div > [aria-label="subtract"] > .mat-button-wrapper > .mat-icon').click();
         cy.contains('2')
             // Delete an item
         cy.get(':nth-child(1) > .cdk-column-star > .iconContainer > .mat-icon').click();
         // Proceed to Payment
         cy.get('.center > .mat-focus-indicator').click();
         cy.url().should('include', '/checkout')






     })
 });