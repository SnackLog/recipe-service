TRUNCATE TABLE ingredients;
TRUNCATE TABLE custom_ingredients;
TRUNCATE TABLE recipes CASCADE;

ALTER SEQUENCE recipes_id_seq RESTART WITH 4;

INSERT INTO recipes (id, name, unit, username) VALUES 
(1, 'Fluffy Pancakes', 'serving', 'foo'),
(2, 'Spicy Chicken Stir Fry', 'bowl', 'bar'),
(3, 'Classic Lemonade', 'pitcher', 'foo');

INSERT INTO ingredients (recipe_id, ingredient_id, quantity) VALUES 
(1, 101, 200.0),
(1, 102, 2.0),
(1, 103, 300.0);

INSERT INTO ingredients (recipe_id, ingredient_id, quantity) VALUES 
(2, 201, 500.0),
(2, 202, 2.0),
(2, 203, 15.0);

INSERT INTO ingredients (recipe_id, ingredient_id, quantity) VALUES 
(3, 301, 5.0),
(3, 302, 100.0),
(3, 303, 1.5);

INSERT INTO custom_ingredients (recipe_id, custom_ingredient_id, quantity) VALUES 
(1, 901, 50.0);

INSERT INTO custom_ingredients (recipe_id, custom_ingredient_id, quantity) VALUES 
(2, 902, 10.0);
