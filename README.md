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

It also utilizes Flyway to be able to migrate and seed the database with the necessary tables and some initial data. If I would need to alter a table I can simply just add a migration file that handles that for me without having to worry about a production environment going down.

I've tried to keep the schema as general as possible, or at least so that it can be easily modified to handle many different products. A product is set up to just give the most minimal amount of information such as name and base price and what kind of product it is.

When a product has been added to an order, it does so through a Line Item. The Line Item has a few more fields to help specify other information about the product that has been added. For instance, a product size, product variation and in the case of pizza it can also have multiple extra ingredients. All of these extra pieces of information also carries a price. When it comes to calculating the price of the line item, it will take all of these values' prices and add them together forming a unit price for the line item.

In the case of a Pizza, these values would be:

```
                   Name,         Price
Product:           Margherita,   20kr
Product Size:      Large,        30kr
Product Variation: Thick Crust,  30kr
Extra Ingredients: Kebab Meat,   10kr
```

Or for a drink perhaps:

```
                   Name,         Price
Product:           Coca Cola,    0
Product Size:      33cl,         10kr
Product Variation: NULL,         NULL
```

