 /// <reference types="cypress" />
 import Chance from 'chance';
 const chance = new Chance();

 describe('Tests dashboard Component', ()=> {
     // check if the dashboard component is loading
     it('Visits the Go organic Dashboard Component', () => {
       cy.visit('http://localhost:4200/dashboard')
     });
     // check if the dashboard component contains following fields
     it('Checks for the fields in dashoard component', () => {
      cy.visit('http://localhost:4200/dashboard')
      cy.contains('Card Content Here')
    });


    // it('Button test', () => {
    //   cy.visit('http://localhost:4200/dashboard')
    //   cy.get('body')
    //   .then($body => {
    //     if ($body.find('.Expand').length) {
    //       return 'Expand';
    //     }
    //     return '.Not found';
    //   })
    //   .then(selector => {
    //     cy.get(selector);
    //   });
    // });
 });


 