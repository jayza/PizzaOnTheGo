/*
* Products
*/

/**
* Example:
* Pizza,
* Drink
*/
CREATE TABLE `product_type` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` tinytext NOT NULL,
  PRIMARY KEY (`id`)
);

/**
* Example:
* Small,
* Medium,
* Large
*/
CREATE TABLE `product_size` (
  `id` int NOT NULL AUTO_INCREMENT,
  `product_type_id` int NOT NULL,
  `name` tinytext NOT NULL,
  `price` decimal(5,2) NOT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT fk_product_size_product_type FOREIGN KEY (`product_type_id`)
    REFERENCES `product_type` (`id`)
);

/**
* Example:
* Crust,
* Calzone Style
*/
CREATE TABLE `product_variation` (
  `id` int NOT NULL AUTO_INCREMENT,
  `product_type_id` int NOT NULL,
  `name` tinytext NOT NULL,
  `price` decimal(5,2) NOT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT fk_product_variation_product_type FOREIGN KEY (`product_type_id`)
    REFERENCES `product_type` (`id`)
);

CREATE TABLE `product` (
  `id` int NOT NULL AUTO_INCREMENT,
  `product_type_id` int NOT NULL,
  `name` tinytext NULL DEFAULT NULL,
  `custom` boolean DEFAULT 0,
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated` timestamp NULL ON UPDATE CURRENT_TIMESTAMP,
  `deleted` timestamp NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT fk_product_product_ FOREIGN KEY (`product_type_id`)
    REFERENCES `product_type`(`id`)
);

CREATE TABLE `ingredient_type` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` tinytext NOT NULL,
  PRIMARY KEY (`id`)
);

CREATE TABLE `ingredient_category` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` tinytext NOT NULL,
  PRIMARY KEY (`id`)
);
/**
* Example:
* Tomato sauce,
* Mozarella
*/
CREATE TABLE `ingredient` (
  `id` int NOT NULL AUTO_INCREMENT,
  `ingredient_type_id` int NOT NULL,
  `ingredient_category_id` int NULL DEFAULT NULL,
  `name` tinytext NOT NULL,
  `price` decimal(5,2) NOT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT fk_ingredient_ingredient_type FOREIGN KEY (`ingredient_type_id`)
    REFERENCES `ingredient_type` (`id`),
  CONSTRAINT fk_ingredient_ingredient_category FOREIGN KEY (`ingredient_category_id`)
    REFERENCES `ingredient_category` (`id`)
);

CREATE TABLE `product_ingredients` (
  `id` int NOT NULL AUTO_INCREMENT,
  `product_id` int NOT NULL,
  `ingredient_id` int NOT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT fk_product_ingredients_product FOREIGN KEY (`product_id`)
    REFERENCES `product`(`id`),
  CONSTRAINT fk_product_ingredients_ingredient FOREIGN KEY (`ingredient_id`)
    REFERENCES `ingredient`(`id`)
);

/**
* Orders
*/
CREATE TABLE `orders` (
  `id` int NOT NULL AUTO_INCREMENT,
  `status` int NOT NULL DEFAULT(0),
  `user_id` int NOT NULL,
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated` timestamp NULL ON UPDATE CURRENT_TIMESTAMP,
  `deleted` timestamp NULL,
  PRIMARY KEY(`id`)
);

CREATE TABLE `product_line_item` (
  `id` int NOT NULL AUTO_INCREMENT,
  `order_id` int NOT NULL,
  `product_id` int NOT NULL,
  `product_size_id` int NOT NULL,
  `product_variation_id` int NOT NULL,
  `quantity` int NOT NULL DEFAULT(1),
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated` timestamp NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY(`id`),
  CONSTRAINT fk_product_line_item_orders FOREIGN KEY (`order_id`)
    REFERENCES `orders`(`id`),
  CONSTRAINT fk_product_line_item_product FOREIGN KEY (`product_id`)
    REFERENCES `product`(`id`),
  CONSTRAINT fk_product_line_item_product_size FOREIGN KEY (`product_size_id`)
    REFERENCES `product_size`(`id`),
  CONSTRAINT fk_product_line_item_product_variation FOREIGN KEY (`product_variation_id`)
    REFERENCES `product_variation`(`id`)
);

CREATE TABLE `product_special_instruction` (
  `id` int NOT NULL AUTO_INCREMENT,
  `product_line_item_id` int NOT NULL,
  `description` tinytext NOT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT fk_product_special_instruction_product_line_item FOREIGN KEY (`product_line_item_id`)
    REFERENCES `product_line_item` (`id`)
);

/**
* Seeds
*/
INSERT INTO `product_type` (id, name)
  VALUES
    (1, "Pizza"),
    (2, "Drink");

INSERT INTO `product_size` (product_type_id, name, price) 
  VALUES 
    (1, "Small", 30),
    (1, "Medium", 40),
    (1, "Large", 50),
    (1, "X-Large", 60),
    (2, "33cl", 15),
    (2, "50cl", 25);

INSERT INTO `product_variation` (product_type_id, name, price) 
  VALUES 
    (1, "Thin Crust", 30),
    (1, "Thick Crust", 40),
    (1, "Cheesy Crust", 50);

INSERT INTO `ingredient_type` (id, name)
  VALUES
    (1, "Base"),
    (2, "Topping"),
    (3, "Dough");

INSERT INTO `ingredient_category` (id, name)
  VALUES
    (1, "Meat"),
    (2, "Vegetable"),
    (3, "Cheese");

INSERT INTO `ingredient` (ingredient_type_id, name, price, ingredient_category_id)
  VALUES
    (1, "Creme Fraiche", 10, NULL),
    (1, "Tomato sauce", 10, NULL),
    (2, "Mozarella", 10, 3),
    (2, "Mushrooms", 10, 2),
    (2, "Spinach", 10, 2),
    (2, "Sucuk", 10, 1),
    (2, "Kebab", 10, 1),
    (2, "Red Onion", 10, 2),
    (3, "Gluten Free", 15, NULL),
    (3, "Sourdough", 20, NULL),
    (3, "Classic", 10, NULL);

INSERT INTO `product` (product_type_id, name, custom)
  VALUES
    (1, "Margherita", 0),
    (1, "Custom Pizza", 1);

INSERT INTO `product_ingredients` (product_id, ingredient_id)
  VALUES
    (1, 2),
    (1, 3),
    (1, 11),
    (2, 2),
    (2, 3),
    (2, 11),
    (2, 6);
    