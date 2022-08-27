package store

const (
	createStory = `INSERT INTO public.story
(id, name, mood, category, description, customer_id, status, created_at, updated_at)
VALUES($1, $2, $3, $4, $5, $6, $7, $8,$9);
`
	getStoryByID = `SELECT * FROM public.story WHERE id = $1`
VALUES($1, $2, $3, $4, $5, $6, $7, $8,$9)`

	getStories = `SELECT id, "name", mood, category, description, customer_id, status, created_at, updated_at
FROM story`
)
