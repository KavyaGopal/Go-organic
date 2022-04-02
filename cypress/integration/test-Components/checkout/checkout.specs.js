 /// <reference types="cypress" />
 import Chance from 'chance';
 const chance = new Chance();

 describe('Tests checkout Component', () => {
     // check if the checkout component is loading : Personal Information Section
     it('Checks if contact information window has required fields ', () => {
         cy.visit('http://localhost:4200/checkout')
         cy.contains('Contact info')
         cy.contains('First name')
         cy.contains('Last name')
         cy.contains('Phone')
         cy.contains('Email')

     });

     // check if the checkout component is loading : Delivery Information Section
     it('Checks if delivery information window has required fields ', () => {
         cy.visit('http://localhost:4200/checkout')
         cy.get(':nth-child(1) > .js-check > input').click();
         cy.get(':nth-child(2) > .js-check > input').click();
         cy.get(':nth-child(2) > :nth-child(1) > .form-control').select('Gainesville');
         cy.contains('Delivery')
         cy.contains('Pick-up')
         cy.contains('City*')
         cy.contains('Area*')
         cy.contains('Street*')
         cy.contains('Building')
         cy.contains('House')
         cy.contains('Postal code')
         cy.contains('Contact')

     });

     // check if the checkout component is loading : Payment Method Selection Section
     it('Checks if payment methods window has required fields ', () => {
         cy.visit('http://localhost:4200/checkout')
         cy.contains('Paypal')
         cy.contains('Credit Card')
         cy.contains('Bank Transfer')

     });

     // check if the checkout component is loading : Payment Method Selection Section
     it('Checks if payment methods window has required fields ', () => {
         cy.visit('http://localhost:4200/checkout')
         cy.contains('Overview')
         cy.contains('Total')
         cy.contains('Make Payment')
         cy.get('.btn').click();
         //cy.url().should('include', '/checkout')

     });

     // checks if button is clickable payment modes are loading correctly
     it('Visits the Go organic Cart Component', () => {
         cy.visit('http://localhost:4200/checkout')
         cy.get(':nth-child(1) > .card-header > .form-check > .form-check-input').click();
         cy.get(':nth-child(2) > .card-header > .form-check > .form-check-input').click();
         cy.get(':nth-child(3) > .card-header > .form-check > .form-check-input').click();
     })
 });