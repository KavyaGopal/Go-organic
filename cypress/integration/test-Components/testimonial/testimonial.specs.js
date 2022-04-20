 /// <reference types="cypress" />
 import Chance from 'chance';
 const chance = new Chance();

 describe('Tests Testimonial Component', () => {
     // check if the home component is loading - Category:Groceries
     it('Clicking on Testimonial should redirect to Testimonial pahe', () => {
         cy.visit('http://localhost:4200/home')
         cy.contains('Testimonials')
         cy.get('.pr-2 > .mat-focus-indicator').click();
         //  Verify redirection to testimonial page
         cy.url().should('include', '/testimonial')
         cy.contains('Zack Wright')
             // Click arrow icon to view next set of testimonials
         cy.get('.arrNext > .mat-icon').click();
         cy.contains('Phoebe')


     });


 });