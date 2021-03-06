# Status Update for Sprint 3
We are developing organic food products web application to place the orders of the organic foods. We are using the tech stack Angular 13 for frontend, Golang 1.17 for the backend and SQLite for database. For sprint 1, we have completed the home page with
high level order of the categories to navigate and developed the mock APIs to get the desired products according to the category of the organic food.  

**Table of Contents**
* [Task Completed](#task-completed)  
    - [Frontend](#frontend)  
    - [Backend](#backend)
* [Bug Fixed](#bugs-fixed)
* [Video Recording](#product-demo-short-videos)
* [Descriptive Recording](#product-demo-videos-with-audio)

## Task Completed
- Sign Up and Sign in functionality
- Checkout component to make payment after adding items to cart
- Unit tests for front end in Cypress and Jasmine
- Unit tests for backend


## FrontEnd

- Sign In
    - UI is complete
    - Integration with backend is complete
    - Test cases are complete
    - Sign in redirects to home component and updates the user name in navigation bar
- Sign Up
    - UI is complete
    - Integration with backend is complete
    - Test cases are complete
    - Redirects to sign in upon successful sign up.
- Checkout
    - Added routing from cart component to Checkout component
    - Personal information and Delivery information fields created
    - Payment Methods/Modes of Payment options added
    - Summary of order and 'Make Payment' section added
    - Test Cases for Checkout component and routing through Cart component are added
    - Integration with backend is complete


## Backend

- Sign Up
    - API development completed
    - Password encryption and decryption done using bcrypt
    - Handles already registered user to sign in
    - Test cases for success and failure done
    - Integrated with frontend complete
- Sign In
    - API development completed
    - Authentication check added when user prompts for wrong password
    - Provide data to front end on the API response so user specific details can be displayed
    - Test cases for success and failure done
    - Integrated with frontend complete
- Checkout
    - API development completed
    - Checks for the item quantity sent by user against the inventory completed
    - Update the item inventory in the database post product checkout completed

## Bugs Fixed


## Product Demo Short Videos

- Frontend

1. Sign In and Sign Up
![sprint3](https://user-images.githubusercontent.com/17459225/161350431-46a13433-72e0-4787-a60b-9c8689ae9aea.gif)

2. Add products to cart -> Checkout Components -> Proceed to Payment

 ![sprint3demo](https://user-images.githubusercontent.com/30043582/161351344-9beb1c22-6d5a-4432-9849-4e42f0a8ac70.gif)


UNIT TESTS:

SIGN UP

https://user-images.githubusercontent.com/17459225/161349777-ce77882c-c4cd-4895-ab66-a92686ff8130.mp4



LOGIN

https://user-images.githubusercontent.com/17459225/161349888-4a41fee9-2005-465e-8a80-bea8661ad1c1.mp4

CART


https://user-images.githubusercontent.com/30043582/161351280-290124fa-2f79-42de-83f4-f5ef8506c710.mp4



CHECKOUT:

https://user-images.githubusercontent.com/30043582/161351242-28374fc1-ace2-4b2c-875b-b1be3f0af326.mp4


- Backend

REGISTER API :

https://user-images.githubusercontent.com/20243138/161363799-a7d4e2a5-1003-470f-af42-cf57726606d4.mp4

LOGIN API:

https://user-images.githubusercontent.com/20243138/161363858-5f516562-7636-4b5e-9f5a-d53b0118dbaf.mp4

CHECKOUT CART API:

https://user-images.githubusercontent.com/20243138/161363877-3357a8ae-0c4d-4e25-9426-82a349b425c8.mp4


UNIT TEST COVERED:

https://user-images.githubusercontent.com/20243138/161363899-9a74cc89-f434-4e47-bdb0-6e04d3e8f720.mp4


