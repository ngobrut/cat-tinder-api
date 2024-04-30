CREATE TABLE
    IF NOT EXISTS cats (
        cat_id uuid default gen_random_uuid () not null constraint cats_pk primary key,
        user_id uuid not null,
        name varchar(30) not null,
        race varchar(255) not null,
        sex varchar(10) not null,
        age_in_month int not null,
        description varchar(255) not null,
        has_matched boolean default false,
        image_url text[] not null,
        created_at timestamp default now (),
        updated_at timestamp default now (),
        deleted_at timestamp default null,
        constraint user_id_fk FOREIGN key (user_id) REFERENCES users(user_id)
    );

CREATE TABLE 
    IF NOT EXISTS cat_match (
        id serial4 constraint cat_match_id_pk primary key,
        user_cat_id uuid not null,
        match_cat_id uuid not null,
        is_approved boolean default null,
        created_at timestamp default now (),
        updated_at timestamp default now (),
        constraint user_cat_id_fk FOREIGN key (user_cat_id) REFERENCES cats(cat_id),
        constraint match_cat_id_fk FOREIGN KEY (match_cat_id) REFERENCES cats(cat_id)
    );
    

CREATE UNIQUE INDEX cat_id_idx ON cats (cat_id)
WHERE (deleted_at IS NULL);
CREATE UNIQUE INDEX cat_match_id_idx ON cat_match (id);