 /// <reference types="cypress" />
 import Chance from 'chance';
 const chance = new Chance();

 describe('Tests home Component', () => {
     // check if the home component is loading - Category:Groceries
     it('Checks the category: Groceries', () => {
         cy.visit('http://localhost:4200/home')
         cy.contains('Groceries')
         cy.contains('Rice')
         cy.contains('Wheat')
         cy.contains('Sugar')
         cy.contains('Salt')

     });

     //Category:Fruits
     it('Checks the category: Fruits', () => {
         cy.visit('http://localhost:4200/home')
         cy.contains('Fruits')
         cy.contains('Apple')
         cy.contains('Cherry')
         cy.contains('Orange')
         cy.contains('Pineapple')
     });

     //Category:Snacks
     it('Checks the category: Snacks', () => {
         cy.visit('http://localhost:4200/home')
         cy.contains('Snacks')
         cy.contains('Roasted Cashew')
         cy.contains('Roasted Peanut')
         cy.contains('Carrot Sticks')
         cy.contains('Dried Mango')

     });

     //Category:Vegetables
     it('Checks the category: Vegetables', () => {
         cy.visit('http://localhost:4200/home')
         cy.contains('Vegetables')
         cy.contains('Potato')
         cy.contains('Tomato')
         cy.contains('Cauliflower')
         cy.contains('Brinjal')
     });

     //Category:Cosmetics
     it('Visits the Go organic Home Component', () => {
         cy.visit('http://localhost:4200/home')
         cy.contains('Cosmetics')
         cy.contains('Cleanser')
         cy.contains('Body Balm')
         cy.contains('Hair Oil')
         cy.contains('Face Mask')
     });

     //Check if "Shop Now" button is clickable
     it('Check if "Shop Now" button is clickable', () => {
         cy.visit('http://localhost:4200/home')
         cy.contains('Shop Now').click()
     });
     //Check if "Add to Cart" button is clickable
     // it('Check if "Shop Now" button is clickable', () => {
     //   cy.visit('http://localhost:4200/home')
     //   cy.contains('Add to Cart').click()
     // });
 });