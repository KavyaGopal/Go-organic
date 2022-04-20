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

- Testimonial
    - Fetch the testimonials of the users from the database as a GET request
    - New Entries can be added in database to without need for code change 
    - Integration with frontend is complete
    - Test cases are complete
- Payment
    - Get the user specific cart items and checks against the inventory items of all products
    - Initiates third party API invocation passing all params of cart items as product key mapping in the database
    - DB update on the item quantity is reduced after successful payment
    - Integration with frontend is complete. The number of remaining items are correctly reflected in the database
    - Test cases are complete
- Reset Password
    - Added a functionality to reset the password in case user has forgotten the password
    - New password is successfully updated in the database
    - Relogin attempt with old password does not allow user to login
    - Test cases are complete
- Cart Checkout
    - Added functionality to keep a check on item values from cart items in the frontend against inventory in the database
    - Reflect which product is out of stock on items exceeding the inventory value in the database
    - Test cases are complete

## Bugs Fixed

    - Change home banner since the current one is blurred in different resolutions
    - Navigate to components(groceries, cosmetics and all) doesn't work clicking navigation bar
    - Update product description and fix styling
