-- +goose Up
CREATE TABLE tenants (
    id BIGINT PRIMARY KEY
);

-- +goose Down
DROP TABLE tenants;
