CREATE TABLE ingredients(
    id SERIAL PRIMARY KEY,
    recipe_id INT REFERENCES recipes(id),
    ingredient_id INT NOT NULL,
    quantity FLOAT NOT NULL
);

CREATE INDEX idx_ingredients_recipe_id
ON ingredients(recipe_id);

CREATE TABLE custom_ingredients(
    id SERIAL PRIMARY KEY,
    recipe_id INT REFERENCES recipes(id),
    custom_ingredient_id INT NOT NULL,
    quantity FLOAT NOT NULL
);

CREATE INDEX idx_custom_ingredients_recipe_id
ON custom_ingredients(recipe_id);
