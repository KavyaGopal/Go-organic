 /// <reference types="cypress" />
 import Chance from 'chance';
 const chance = new Chance();

 describe('Tests home Component', ()=> {
     // check if the home component is loading
     it('Visits the Go organic Home Component', () => {
       cy.visit('http://localhost:4200/home')
     })
 });