erDiagram
	users ||--o{ elections : references
	elections ||--o{ vote_variants : references
	users ||--|| votes : references
	vote_variants ||--o{ votes : references

	users {
		UUID id
		VARCHAR(255) nickname
		VARCHAR(255) password
		TIMESTAMP created_at
		TIMESTAMP updated_at
	}

	elections {
		UUID id
		UUID user_id
		VARCHAR(255) name
		VARCHAR(255) descriprion
	}

	vote_variants {
		UUID id
		UUID election_id
		VARCHAR(255) name
		BIGINT counter
	}

	votes {
		UUID id
		UUID user_id
		UUID variant_id
	}