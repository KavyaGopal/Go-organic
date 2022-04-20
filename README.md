# Go-organic
A one-stop place to buy all the organic and healthy food. It is a web application that can be used to place orders for all organic foods.  The website connects the end-user to the best local organic products like groceries, fruits, snacks, vegetables, and cosmetics.  The user can sign up and sign in and browse through the catalog. The items can be added to the cart and the quantity of the items can be modified. While the user has added some products to the cart, she can continue to go back and shop. Once shopping is done, the user can proceed to payment after adding the delivery address and providing the necessary details. The user can then choose to log out or continue shopping. The user experience has been given the utmost importance and every aspect of this web application reflects that.

**Table of Contents**
- [Features](#features)
- [Technology Stack](#technology-stack)
- [Setup Environment](#setup-and-install)
- [Testing](#testing)
- [License and authors](#authors)


## Features
- Login|Register|Logout

- Browse|Search|Add items to the cart

- Cart| Increase or decrease item quantity| Delete item from cart|add based on availability of products

- Checkout Cart| Proceed to payment

- PayPal Merchant API Integration | Sandbox Testing | Payment Modes: PayPal Cash, Card, Bank Transfer

- Testimonial| About us| Contact details

## Technology Stack
- Front end: Angular 13

- Back end: Go (1.17)

- DB: SQLite3 using Gorm

## Setup and Install

- FrontEnd: Install NodeJS, Angular

- Backend: Install Golang

- Database: Install SQLite 3 using Gorm

## Run instructions

#### Frontend
    In terminal 1
    cd <location to the folder Go-organic>
    npm install;
    ng serve -o
    
#### Backend

    In terminal 2
    cd <location to the folder Go-organic>/backend-go/pkg/api;
    go run api.go

## Documentation
[Go-Organic Wiki](https://github.com/KavyaGopal/Go-organic/wiki)


## Product Demo Short Videos

- Frontend

1. Sign In and Sign Up
![sprint3](https://user-images.githubusercontent.com/17459225/161350431-46a13433-72e0-4787-a60b-9c8689ae9aea.gif)

2. Logout

![logout](https://user-images.githubusercontent.com/17459225/164273519-ae72e05e-677f-4614-a0ac-8d2de6f3b502.gif)


3. Add products to cart -> Checkout Components -> Proceed to Payment

 ![sprint3demo](https://user-images.githubusercontent.com/30043582/161351344-9beb1c22-6d5a-4432-9849-4e42f0a8ac70.gif)
 
4. Make Payment using PayPal (PayPal Cash, Card, Bank Transfer)
![sprint4pay](https://user-images.githubusercontent.com/30043582/164334359-fc38a970-b424-44cc-9cd2-e6930cad2b32.gif)


6. Navigation bar -> Navigates within home pages -> Navigates from other page to homepage

![navigation](https://user-images.githubusercontent.com/17459225/164273785-f9b0e447-1a8f-4a7a-9ecc-e1e473201186.gif)


6. Testimonial page -> from home page

![testimonial](https://user-images.githubusercontent.com/17459225/164273895-f57fa8d5-dc83-46fc-b7aa-94db735963ea.gif)


7. Cart -> Delete items from cart(will update cart number) -> Add button is disabled when items are out of stocks

![cart](https://user-images.githubusercontent.com/17459225/164274307-0f308336-4cea-4275-a20d-34e124ad0afb.gif)




## Testing
For testing, we will do unit testing for both front and back end when developing, also end to end testing.

- end to end testing
- back end unit testing
- front end unit testing

## Cypress testing

SIGN UP

https://user-images.githubusercontent.com/17459225/161349777-ce77882c-c4cd-4895-ab66-a92686ff8130.mp4



LOGIN

https://user-images.githubusercontent.com/17459225/161349888-4a41fee9-2005-465e-8a80-bea8661ad1c1.mp4


LOGOUT:

https://user-images.githubusercontent.com/17459225/164269145-e7c5c76d-1745-4544-9e2b-70d10f78fbac.mp4


CART

https://user-images.githubusercontent.com/17459225/164269196-289f9504-5c51-43b2-8c3b-fbe1e0488dbe.mp4


CHECKOUT:

https://user-images.githubusercontent.com/30043582/161351242-28374fc1-ace2-4b2c-875b-b1be3f0af326.mp4


TESTIMONIAL:


https://user-images.githubusercontent.com/17459225/164269288-dcc1826a-62ef-4e21-af8e-aba55977ffd3.mp4


DASHBOARD:


https://user-images.githubusercontent.com/30043582/164331145-f1e6b40a-14db-4f69-b843-294318d7ed4d.mp4


HOME:

https://user-images.githubusercontent.com/30043582/164331297-58755b3d-d60e-4af6-b832-f2c5078470b9.mp4


PYPAL PAYMENT:

https://user-images.githubusercontent.com/30043582/164331422-67ed53f9-5874-418d-a7e8-00fe0776599d.mp4



## Backend unit testing


## API Documentation

## Project Board

    https://github.com/KavyaGopal/Go-organic/projects/1


## Sprint 4 Deliverables

## FrontEnd

- Payment
    - Clicking on checkout redirects the user to fill out delivery mode and address and then proceed to the payment section
    - Added integration with paypal where end user can pay for the items
    - Added a payment successful component which is shown on successful payment. It allows the user to continue shopping
    - Test cases are complete
    - Integration with backend is complete. The number of remaining items of the products are updated in db.
- Logout
    - UI is complete. Logout button appears once the user is logged in.
    - Clicking Logout redirects user to the sign in screen and ensures the session is still not shown to the user(cart elements)
    - Integration with backend is complete
    - Test cases are complete
- Testimonial
    - UI is complete. The user can redirect from homepage to testimonial and read through all the testimonials.
    - Integration with backend is complete
    - Test cases are complete
- Cart
    - Added functionality to add items to the cart only when they are in stock
    - Added functionality to delete the item from the cart, even if it has multiple quantities (on clicking delete)



## Backend

- Logout
    - API development completed
    - 
- Payment
    - 

## Bugs Fixed

    - Change home banner since the current one is blurred in different resolutions
    - Navigate to components(groceries, cosmetics and all) doesn't work clicking navigation bar
    - Update product description and fix styling



## Sprint 1

- https://github.com/KavyaGopal/Go-organic/blob/main/sprint1.md

## Sprint 2

- https://github.com/KavyaGopal/Go-organic/blob/main/sprint2.md

## Sprint 3

- https://github.com/KavyaGopal/Go-organic/blob/main/sprint3.md




## Authors

#### Authors
| Name | GitHub ID | Email ID | Work Type |
|------|-----------|---------------------|--------|
|Ashwin Rai|ashwin181|raiashwin018@gmail.com|Backend Developer|
|Kavya Gopal|KavyaGopal|95kavya@gmail.com|Frontend Developer|
|Vaibhav Kulkarni|KulkarniVaibhav|kulvaibhavp@gmail.com|Frontend Developer|
|Zubin Arya|zubin-arya|virtuoso02031995@gmail.com|Backend Developer|
