 /// <reference types="cypress" />
 import Chance from 'chance';
 const chance = new Chance();

 describe('Tests dashboard Component', ()=> {
     // check if the dashboard component is loading
     it('Visits the Go organic Dashboard Component', () => {
       cy.visit('http://localhost:4200/home')
     })
 });