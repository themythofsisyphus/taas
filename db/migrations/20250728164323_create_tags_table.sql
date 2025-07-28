-- +goose Up
CREATE TABLE tags (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    tenant_id BIGINT NOT NULL,
    name_tsv TSVECTOR GENERATED ALWAYS AS (to_tsvector('english', name)) STORED,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    CONSTRAINT uniq_tag UNIQUE (tenant_id, name),
    CONSTRAINT fk_tenant FOREIGN KEY (tenant_id) REFERENCES tenants(id) ON DELETE CASCADE
);

CREATE INDEX idx_tags_name_tsv ON tags USING GIN (name_tsv);
CREATE INDEX idx_tags_tenant_id ON tags (tenant_id);
CREATE INDEX idx_tags_tenant_name ON tags (tenant_id, name);

-- +goose Down
DROP INDEX IF EXISTS idx_tags_tenant_name;
DROP INDEX IF EXISTS idx_tags_tenant_id;
DROP INDEX IF EXISTS idx_tags_name_tsv;
DROP TABLE IF EXISTS tags;