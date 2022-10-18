CREATE TYPE user_status AS ENUM('creator', 'administrator', 'member', 'restricted', 'left', 'kicked');

CREATE TYPE account_role AS ENUM('creator', 'admin', 'bot');

CREATE TYPE sub_status AS ENUM('expired', 'active', 'cancelled', 'inactive');


CREATE TABLE TgUser (
	user_id bigserial not null unique PRIMARY KEY,
  	telegram_id bigint not null,
  	username text not null,
  	status user_status not null
);

CREATE TABLE Account (
	account_id bigserial not null unique PRIMARY KEY,
  	password bytea not null,
  	email text unique not null,
	role account_role not null default 'creator',
  	chan_name text,
	user_id bigint
);

ALTER TABLE Account
	ADD FOREIGN KEY (user_id) REFERENCES TgUser (user_id);


CREATE TABLE Sub (
	user_id bigint not null,
  	account_id bigint not null,
  	activated_at timestamptz NOT NULL DEFAULT (now()),
  	expires_at timestamptz NOT NULL DEFAULT (now()),
  	status sub_status not null default 'inactive',
  	price integer default 0.0,
	PRIMARY KEY (user_id, account_id)
);

ALTER TABLE Sub
	ADD FOREIGN KEY (user_id) REFERENCES TgUser (user_id);
    
ALTER TABLE Sub
	ADD FOREIGN KEY (account_id) REFERENCES Account (account_id);

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
ADD CONSTRAINT account_id_fk 
FOREIGN KEY (account_id) REFERENCES Account(account_id) 
ON DELETE SET NULL;
