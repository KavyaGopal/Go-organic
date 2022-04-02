 /// <reference types="cypress" />
 import Chance from 'chance';
 const chance = new Chance();

 describe('Tests login Component', () => {
     // check if the login component is loading
     it('Visits the Login Component', () => {
         cy.visit('http://localhost:4200/login')
     });
     // check if the login component contains following fields
     it('Checks for the fields in login component', () => {
         cy.visit('http://localhost:4200/login')
         cy.contains('Sign In')
         cy.contains('Password')
         cy.contains('New to Go-Organic?')
         cy.contains('Sign Up')
     });

     it('Navigate to Sign Up from Login', () => {
         cy.visit('http://localhost:4200/login')
         cy.get('.signUp > a').click();
         cy.url().should('include', '/sign-up')
     });

     it('Sign in an user: should redirect to home', () => {
         cy.visit('http://localhost:4200/login')
         cy.get('#mat-input-0').type('orange@gmail.com');
         cy.get('#mat-input-1').type('testing');
         cy.get('form.ng-dirty > .mat-focus-indicator > .mat-button-wrapper').click();
         cy.url().should('include', '/home')
             //  User name should be updated in home page
         cy.contains('Orange')
     });
 });