CREATE TABLE IF NOT EXISTS goals
(
    book_id       UUID                     NOT NULL,
    created_at    TIMESTAMP WITH TIME ZONE DEFAULT now(),
    end_date      TIMESTAMP WITH TIME ZONE NOT NULL,
    goal_words    INT                      NOT NULL,
    id            UUID PRIMARY KEY         DEFAULT uuid_generate_v4(),
    is_finished   BOOLEAN                  DEFAULT FALSE,
    start_date    TIMESTAMP WITH TIME ZONE NOT NULL,
    title         VARCHAR(255)             NOT NULL,
    updated_at    TIMESTAMP WITH TIME ZONE DEFAULT now(),
    user_id       UUID                     NOT NULL,
    description   VARCHAR(255),
    written_words INT                      NOT NULL,
    words_per_day INT                      NOT NULL,

    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id),
    CONSTRAINT fk_book FOREIGN KEY (book_id) REFERENCES books (id)
);

CREATE INDEX idx_user_id ON goals (user_id);
CREATE INDEX idx_book_id ON goals (book_id);
