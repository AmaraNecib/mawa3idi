up:
	@goose postgres -dir=migrations  $(DATABASE_URL)  up