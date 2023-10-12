CREATE TABLE  users
(
    user_id serial not null unique,
    login varchar(255) not null unique,
    password_hash varchar(255) not null,
    username varchar(255) not null,
    email_confirmation BOOLEAN NOT NULL
);
--нужна для защиты от перебора
--когда number_of_inputs достигает определённого порога то нужно запретить вход до unlock_time
--например на 4 попытку нужно ждать 5 минут на 7 - 20 минут на 10 бан
CREATE TABLE login_attempt 
(
    login_attempt_id serial not null unique,
    email TEXT not null unique,
    number_of_inputs int NOT NULL,
    unlock_time TIMESTAMP not null,
    login_method TEXT not null
);


--таблица которая хранит временные данные для подтверждения по почте
CREATE TABLE  email_confirmation
(
    email_confirmation_id serial not null unique,
    email TEXT not null,
    user_id int REFERENCES users (user_id) ON DELETE CASCADE NOT NULL unique,
    code_email_confirmation int,
    time_end TIMESTAMP not null,
    assignment TEXT not null,   --регистрация, 2fa,смена почты
    device TEXT unique
);



CREATE TABLE  auth_data
(
    auth_data_id serial not null unique,
    token TEXT,
    user_id int REFERENCES users (user_id) ON DELETE CASCADE NOT NULL,
    time_end TIMESTAMP not null,
    device TEXT
);


--------------------------------------------------------------------------------------------------------------------------

CREATE TABLE event_items
(
    event_items_id serial not null unique,
    admin int REFERENCES users(user_id) on delete cascade not null,
    time_start TIMESTAMP not null,
    time_end TIMESTAMP not null,
    place varchar(255) not null,
    game varchar(255) not null,
    description varchar (255) not null,
    public BOOLEAN not null
);


CREATE TABLE event_invitations
(
    event_invitations_id serial not null unique,
    event_id int REFERENCES event_items(event_items_id) on delete cascade not null,
    user_id int REFERENCES users(user_id) on delete cascade not null,
    status VARCHAR(20) DEFAULT 'Pending'
);


CREATE TABLE friends (
    friends_id SERIAL PRIMARY KEY,
    user_id int REFERENCES users(user_id) on delete cascade not null,
    friend_id int REFERENCES users(user_id) on delete cascade not null,
    status VARCHAR(20) DEFAULT 'Pending'
);



--------------------------------------------------------------------------------------------------------------------------
CREATE TABLE jwt_blacklist (
    id SERIAL PRIMARY KEY,
    user_id int REFERENCES users(user_id) on delete cascade not null,
    --время с которого начинается валидные токены
    token_valid_from TIMESTAMP not null,
    --время удаления этой записи
    cleanup_time TIMESTAMP not null
);

CREATE TABLE push_notification (
    id SERIAL PRIMARY KEY,
    user_id int REFERENCES users(user_id) on delete cascade not null,
    device TEXT NOT NULL UNIQUE,
    --токен полученный от firebase
    push_token TEXT not null
);


