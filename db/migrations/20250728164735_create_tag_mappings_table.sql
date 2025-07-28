-- +goose Up
CREATE TABLE tag_mappings (
    id SERIAL PRIMARY KEY,
    tag_id BIGINT NOT NULL,
    entity_id BIGINT NOT NULL,
    entity_type BIGINT NOT NULL,
    tenant_id BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    CONSTRAINT fk_tag FOREIGN KEY (tag_id) REFERENCES tags(id) ON DELETE CASCADE,
    CONSTRAINT fk_entity FOREIGN KEY (entity_type) REFERENCES entities(id) ON DELETE CASCADE,
    CONSTRAINT uniq_tag_mapping UNIQUE (tenant_id, tag_id, entity_id, entity_type)
);

-- Add indexes for performance
CREATE INDEX idx_tag_mappings_tenant_id ON tag_mappings (tenant_id);
CREATE INDEX idx_tag_mappings_tenant_entity ON tag_mappings (tenant_id, entity_id, entity_type);

-- +goose Down
DROP INDEX IF EXISTS idx_tag_mappings_tenant_entity;
DROP INDEX IF EXISTS idx_tag_mappings_tenant_id;
DROP TABLE IF EXISTS tag_mappings;