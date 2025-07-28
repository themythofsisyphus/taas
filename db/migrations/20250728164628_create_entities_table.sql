-- +goose Up
-- +goose StatementBegin
CREATE TABLE entities (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    tenant_id BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    CONSTRAINT uniq_entity UNIQUE (tenant_id, name),
    CONSTRAINT fk_tenant FOREIGN KEY (tenant_id) REFERENCES tenants(id) ON DELETE CASCADE
);
-- +goose StatementEnd
CREATE INDEX idx_entities_tenant_id ON tags (tenant_id);

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_entities_tenant_id;
DROP TABLE IF EXISTS entities;
-- +goose StatementEnd
