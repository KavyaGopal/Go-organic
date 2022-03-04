 /// <reference types="cypress" />
 import Chance from 'chance';
 const chance = new Chance();

 describe('Tests Cart Component', ()=> {
     // check if the cart component is loading
     it('Visits the Go organic Cart Component', () => {
       cy.visit('http://localhost:4200/cart')
     })
 });