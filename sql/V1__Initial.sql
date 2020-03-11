CREATE TABLE pizza (
  id int NOT NULL AUTO_INCREMENT,
  name varchar(255) NOT NULL,
  price decimal(5,2) NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE pizza_option_type (
  id int NOT NULL AUTO_INCREMENT,
  name varchar(255) NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE pizza_option (
  id int NOT NULL AUTO_INCREMENT,
  name varchar(255) NOT NULL,
  type_id int NOT NULL,
  price decimal(5,2) NOT NULL,
  PRIMARY KEY (id),
  CONSTRAINT fk_option_pizza_option_types FOREIGN KEY (type_id)
    REFERENCES pizza_option_type(id)
);

CREATE TABLE pizzas_pizza_options (
  id int NOT NULL AUTO_INCREMENT,
  pizza_id INT NOT NULL,
  pizza_option_id INT NOT NULL,
  PRIMARY KEY (id),
  CONSTRAINT fk_pizzas_options_pizza FOREIGN KEY (pizza_id)
    REFERENCES pizza(id),
  CONSTRAINT fk_pizzas_options_pizza_option FOREIGN KEY (pizza_option_id)
    REFERENCES pizza_option(id)
);

-- Seeds

INSERT INTO pizza (name, price) 
  VALUES 
    ("Super Cheesy Margherita", 90.99),
    ("Funghi", 75.00),
    ("Creamy Sucuk", 85.00);

INSERT INTO pizza_option_type (name) 
  VALUES 
    ("Base"),
    ("Topping"),
    ("Crust"),
    ("Dough"),
    ("Size");

INSERT INTO pizza_option (name, type_id, price)
  VALUES
    ("Tomato sauce", 1, 10),
    ("Mozarella", 2, 10),
    ("Thin", 3, 10),
    ("Thick", 3, 15),
    ("Cheese-filled", 3, 20),
    ("Mushrooms", 2, 10),
    ("Creme Fraiche", 1, 10),
    ("Spinach", 2, 10),
    ("Sucuk", 2, 10),
    ("Red Onion", 2, 10),
    ("Garlic Olive Oil", 3, 25),
    ("Gluten Free", 4, 15),
    ("Sourdough", 4, 20),
    ("Classic", 4, 10),
    ("Small", 5, -10),
    ("Medium", 5, 0),
    ("Large", 5, 25);

INSERT INTO pizzas_pizza_options (pizza_id, pizza_option_id)
  VALUES
    (1, 1),
    (1, 2),
    (1, 5),
    (1, 14),
    (1, 16),
    (2, 1),
    (2, 2),
    (2, 6),
    (2, 3),
    (2, 12),
    (2, 16),
    (3, 7),
    (3, 8),
    (3, 9),
    (3, 10),
    (3, 11),
    (3, 13),
    (3, 16);