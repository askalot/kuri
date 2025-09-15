- Use Go Fiber for routes

"github.com/gofiber/fiber/v2"
"github.com/gofiber/fiber/v2/middleware/logger"

- Group "/notes" routes
https://go-chi.io/#/pages/routing

- Add uuid as noteId?
ID          uuid.UUID `json:"id"`
ID:          uuid.New(),
"github.com/google/uuid"

- Format note as markdown with frontmatter

- [DOING] Add Tags (CSV list)
	- add tags to frontmatter?
