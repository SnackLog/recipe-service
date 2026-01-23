CREATE INDEX idx_text_search ON recipes 
USING GIN (to_tsvector('german', name));
