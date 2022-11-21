CREATE TABLE
    IF NOT EXISTS characters (
        id char(21) DEFAULT nanoid(),
        name jsonb,
        description jsonb,
        birth_at DATE,
        created_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp(0),
        updated_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp(0),
        deleted_at TIMESTAMPTZ,
        is_draft BOOLEAN NOT NULL DEFAULT false,
        PRIMARY KEY (id)
    );

CREATE INDEX idx_characters_pagination ON characters(created_at, id);

CREATE TABLE
    IF NOT EXISTS figures (
        id char(21) DEFAULT nanoid(),
        character_id char(21),
        name VARCHAR(255) NOT NULL,
        description jsonb,
        type VARCHAR(255),
        version VARCHAR(255),
        size VARCHAR(30),
        height SMALLINT,
        materials VARCHAR(100) [],
        release_date DATE,
        manufacturer VARCHAR(150),
        links jsonb,
        price VARCHAR(50) [],
        created_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp(0),
        updated_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp(0),
        deleted_at TIMESTAMPTZ,
        is_nsfw BOOLEAN NOT NULL DEFAULT false,
        is_draft BOOLEAN NOT NULL DEFAULT false,
        PRIMARY KEY (id),
        CONSTRAINT fk_characters FOREIGN KEY (character_id) REFERENCES characters (id)
    );

CREATE INDEX idx_figures_pagination ON figures(created_at, id);

CREATE TABLE
    IF NOT EXISTS figures_images (
        id char(21) DEFAULT nanoid(),
        image_path VARCHAR(255) NOT NULL,
        blurhash VARCHAR(100) NOT NULL,
        preview BOOLEAN NOT NULL DEFAULT false,
        figure_id char(21) NOT NULL,
        created_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp(0),
        updated_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp(0),
        PRIMARY KEY (id),
        CONSTRAINT fk_figures_images FOREIGN KEY (figure_id) REFERENCES figures (id)
    );

CREATE INDEX idx_fig_id_img ON figures_images(figure_id);

CREATE TABLE
    IF NOT EXISTS characters_images (
        id char(21) DEFAULT nanoid(),
        image_path VARCHAR(255) NOT NULL,
        blurhash VARCHAR(100) NOT NULL,
        preview BOOLEAN NOT NULL DEFAULT false,
        character_id char(21) NOT NULL,
        created_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp(0),
        updated_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp(0),
        PRIMARY KEY (id),
        CONSTRAINT fk_characters_images FOREIGN KEY (character_id) REFERENCES characters (id)
    );

CREATE INDEX idx_char_id_img ON characters_images(character_id);
