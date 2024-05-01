CREATE TABLE IF NOT EXISTS users (
    user_id uuid default gen_random_uuid() not null constraint users_pk primary key,
    name varchar(255) not null,
    email varchar(255) not null,
    password varchar(500) not null,
    created_at timestamp default now(),
    updated_at timestamp default now(),
    deleted_at timestamp default null
);

CREATE UNIQUE INDEX unique_email_idx ON users (email)
WHERE
    (deleted_at IS NULL);

CREATE TABLE IF NOT EXISTS cats (
    cat_id uuid default gen_random_uuid() not null constraint cats_pk primary key,
    user_id uuid not null,
    name varchar(30) not null,
    race varchar(255) not null,
    sex varchar(10) not null,
    age_in_month int not null,
    description varchar(255) not null,
    has_matched boolean default false,
    image_urls text[] not null,
    created_at timestamp default now(),
    updated_at timestamp default now(),
    deleted_at timestamp default null,
    constraint user_id_fk foreign key (user_id) references users(user_id)
);

CREATE TABLE IF NOT EXISTS cat_matches(
    id uuid default gen_random_uuid() not null constraint cat_matches_id_pk primary key,
    issuer_user_id uuid not null,
    issuer_cat_id uuid not null,
    receiver_user_id uuid not null,
    receiver_cat_id uuid not null,
    message varchar(255) not null,
    is_approved boolean default null,
    created_at timestamp default now(),
    updated_at timestamp default now(),
    deleted_at timestamp default null,
    constraint cat_matches_issuer_id_fk foreign key (issuer_user_id) references users(user_id),
    constraint cat_matches_issuer_cat_id_fk foreign key (issuer_cat_id) references cats(cat_id),
    constraint cat_matches_receiver_id_fk foreign key (receiver_user_id) references users(user_id),
    constraint cat_matches_receiver_cat_id_fk foreign key (receiver_cat_id) references cats(cat_id)
)