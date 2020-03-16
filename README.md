# PizzaOnTheGo

## Frontend
React

## Backend
For the backend I chose to create an API with an MVC structure, except without any views.

I made sure to add Swagger as a way to create automatic generated API specifications for my endpoints. By adding annotations to each respective method, it will generate a nice interface that shows how to use the API. 

For routing I use the Mux router handler, which allows me to specify all my routes with ease and even add some middleware, i.e. check that the user is authenticated for the request.

After an order has completed, I use GofPDF to generate a PDF of the order receipt. This can later be downloaded through an endpoint in the API.

The way it's designed is that a user request an endpoint and that endpoint will lead to calling a controller which handles the request. When data is requested or posted I then use the repositories to handle the data and either fetch it or insert it into the database.

**Tests**

## Database
This project is using MariaDB.

It also utilizes Flyway to be able to migrate and seed the database with the necessary tables and some initial data.
