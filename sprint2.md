# Status Update for Sprint 2
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
- The backend and frontend is integrated
- Unit tests for front end in Cypress and Jasmine
- Unit tests for backend


## FrontEnd

- The Api's are used to display the products in front end. Integration is complete
- Add to cart component
    - Adding products updates the number of elements in cart in home page
    - Routing is added to cart icon in homepage to cart page
    - Cart page contains list of all elements added to cart, their weight, quantity subtotal.
    - The total money is shown along with an option to checkout

## Backend

- Api's are created to fetch all products and required information
- Mock API setup done using gorilla/mux for fetching product details

## Bugs Fixed
https://github.com/KavyaGopal/Go-organic/issues/43  Remove additional "Add to cart item" 
https://github.com/KavyaGopal/Go-organic/issues/47 Code Clean up: Remove cypress files that are not required




## Product Demo Short Videos

- Frontend




- Backend

https://user-images.githubusercontent.com/20243138/156867388-2adcb84f-9ca7-4895-8cf7-05170ea0dad0.mp4



