-- +goose Up 

create table posts (
  id UUID PRIMARY KEY,
  created_at TIMESTAMP NOT NULL,
  updated_At TIMESTAMP NOT NULL,
  title TEXT NOT NULL,
  description TEXT,
  published_at TIMESTAMP NOT NULL,
  url TEXT NOT NULL UNIQUEe,
  feed_id UUID NOT NULL references feeds(id) ON DELETE CASCADE
);




-- +goose Down

DROP TABLE posts;