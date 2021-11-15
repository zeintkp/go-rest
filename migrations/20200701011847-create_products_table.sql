
-- +migrate Up
CREATE TABLE IF NOT EXISTS products (
    id character(36) PRIMARY KEY DEFAULT uuid_generate_v4() NOT NULL,
    product_name character varying(50) NOT NULL,
    is_active boolean DEFAULT true NOT NULL,
    created_at timestamp without time zone NOT NULL DEFAULT timezone('utc', now()),
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone
);

-- +migrate Down
DROP TABLE IF EXISTS products;