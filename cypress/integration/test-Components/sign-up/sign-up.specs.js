 /// <reference types="cypress" />
 import Chance from 'chance';
 const chance = new Chance();

 describe('Tests Sign-Up Component', () => {
     // check if the login component is loading
     it('Visits the Login Component', () => {
         cy.visit('http://localhost:4200/sign-up')
     });
     // check if the login component contains following fields
     it('Checks for the fields in login component', () => {
         cy.visit('http://localhost:4200/sign-up')
         cy.contains('Sign In')
         cy.contains('User Name')
         cy.contains('Phone Number')
         cy.contains('Address')
         cy.contains('Email')
         cy.contains('Password')
         cy.contains('Re-enter Password')
         cy.contains('Already a member?')
         cy.contains('Sign Up')
     });

     it('Clicking on sign in should redirect to sign in', () => {
         cy.visit('http://localhost:4200/sign-up')
         cy.get('.signUp > a').click();
         cy.url().should('include', '/login')
     });

     it('Signs up an user and tests redirection to login page', () => {
         cy.visit('http://localhost:4200/sign-up')
         cy.get('#mat-input-0').type('Test');
         cy.get('#mat-input-1').type('1212121212');
         cy.get('#mat-input-2').type('3800 SW, 34th St');
         cy.get('#mat-input-3').type('abc@gmail.com');
         cy.get('#mat-input-4').type('tester');
         cy.get('#mat-input-5').type('tester');
         cy.get('form.ng-dirty > .mat-focus-indicator > .mat-button-wrapper').click();
         cy.url().should('include', '/login')
     });
 });