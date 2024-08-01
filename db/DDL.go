package db

const (
	CreateTasksTableQuery = `CREATE TABLE IF NOT EXISTS tasks
							(
								id          SERIAL PRIMARY KEY,
								title       VARCHAR NOT NULL,
								description TEXT,
								is_done     BOOLEAN DEFAULT FALSE,
								is_deleted  BOOLEAN DEFAULT FALSE,
								priority    INT DEFAULT 3,    
								created_at  TIMESTAMP DEFAULT NOW()
							);`
)
