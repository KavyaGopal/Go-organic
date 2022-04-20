 /// <reference types="cypress" />
 import Chance from 'chance';
 const chance = new Chance();

 describe('Tests logout Component', () => {

     it('Sign in an user and sign out: should redirect to login once user signs out', () => {
         cy.visit('http://localhost:4200/login')
         cy.get('#mat-input-0').type('kavyagopal@gmail.com');
         cy.get('#mat-input-1').type('tester');
         cy.get('form.ng-dirty > .mat-focus-indicator > .mat-button-wrapper').click();
         cy.url().should('include', '/home')
             //   User name should be updated in home page
         cy.contains('kavya')
         cy.contains('Logout')
             // Click logout
         cy.get('form.ng-dirty > .mat-focus-indicator > .mat-button-wrapper').click();
         // Should redirect to login screen
         cy.url().should('include', '/login')

     });


 });