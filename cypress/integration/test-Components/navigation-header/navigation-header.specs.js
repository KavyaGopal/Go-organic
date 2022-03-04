 /// <reference types="cypress" />
 import Chance from 'chance';
 const chance = new Chance();

 describe('Tests navigation-header Component', ()=> {
     // check if the home component is loading
     it('Visits the Go organic Navigation-Header Component', () => {
       cy.visit('http://localhost:4200/home')
     })
 });