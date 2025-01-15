CREATE TABLE IF NOT EXISTS reports
(
    book_id      UUID                                                       NOT NULL,
    goal_id      UUID                                                       NOT NULL,
    created_at   timestamp(0)     default (now())::timestamp with time zone not null,
    words_amount INT                                                        NOT NULL,
    id           UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    title        VARCHAR(255)                                               NOT NULL,
    updated_at   timestamp(0)     default NULL::timestamp without time zone NOT NULL,
    user_id      UUID                                                       NOT NULL,
    description  VARCHAR(255),

    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id),
    CONSTRAINT fk_book FOREIGN KEY (book_id) REFERENCES books (id),
    CONSTRAINT fk_goal FOREIGN KEY (goal_id) REFERENCES goals (id)
);


CREATE INDEX idx_reports_book_id ON reports (book_id);
CREATE INDEX idx_reports_goal_id ON reports (goal_id);
CREATE INDEX idx_reports_user_id ON reports (user_id);