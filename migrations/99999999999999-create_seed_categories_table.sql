
-- +migrate Up
INSERT INTO products ("product_name") VALUES 
    ('Product 1'),
    ('Product 2'),
    ('Product 3'),
    ('Product 4'),
    ('Product 5'),
    ('Product 6'),
    ('Product 7'),
    ('Product 8'),
    ('Product 9'),
    ('Product 10'),
    ('Product 11'),
    ('Product 12'),
    ('Product 13'),
    ('Product 14'),
    ('Product 15');

-- +migrate Down
