CREATE TYPE user_status AS ENUM('creator', 'administrator', 'member', 'restricted', 'left', 'kicked');

CREATE TYPE sub_status AS ENUM('expired', 'active', 'cancelled');

CREATE TABLE TgUser (
	user_id bigserial not null unique PRIMARY KEY,
  	telegram_id bigint not null,
  	username text not null,
  	status user_status not null
);

CREATE TABLE Creator (
	creator_id bigserial not null unique PRIMARY KEY,
  	telegram_id bigint not null,
  	username text not null,
  	password bytea,
  	email text,
  	chan_name text
);

CREATE TABLE Sub (
	user_id bigint not null,
  	creator_id bigint not null,
  	activated_at timestamptz NOT NULL DEFAULT (now()),
  	expires_at timestamptz NOT NULL DEFAULT (now()),
  	status sub_status NOT NULL,
  	price integer default 0.0,
	PRIMARY KEY (user_id, creator_id)
);

ALTER TABLE Sub
	ADD FOREIGN KEY (user_id) REFERENCES TgUser (user_id);
    
ALTER TABLE Sub
	ADD FOREIGN KEY (creator_id) REFERENCES Creator (creator_id);

CREATE TABLE Sub_History (
       LIKE Sub
       including defaults
);

ALTER TABLE Sub_History
      ADD COLUMN sub_hist_id BIGSERIAL PRIMARY KEY;

ALTER TABLE Sub_History 
ADD CONSTRAINT user_id_fk 
FOREIGN KEY (user_id) REFERENCES TgUser(user_id) 
ON DELETE SET NULL;

ALTER TABLE Sub_History 
ADD CONSTRAINT creator_id_fk 
FOREIGN KEY (creator_id) REFERENCES Creator(creator_id) 
ON DELETE SET NULL;
