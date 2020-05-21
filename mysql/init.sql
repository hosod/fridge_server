use test;

CREATE TABLE users (
  id int(10) unsigned not null auto_increment,
  name varchar(255),
  email varchar(255),
  primary key (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

CREATE TABLE fridges (
  id int(10) unsigned not null auto_increment,
  name varchar(255),
  primary key (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- CREATE TABLE user_fridge_relations (
--   user_id int(10) unsigned not null,
--   fridge_id int(10) unsigned not null,
--   primary key (user_id, fridge_id),
--   foreign key (user_id) 
--     references users (id) match full
--     on delete cascade
--     on update no action,
--   foreign key (fridge_id)
--     references fridges (id) match full
--     on delete cascade
--     on update no action
-- ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

INSERT INTO users (name, email) VALUES ('yamada', 'yamada@mail.com');
INSERT INTO users (name,email) VALUES ('tanaka', 'tanaka@mail.com');
INSERT INTO users (name, email) VALUES ('sato', 'sato@mail.com');
INSERT INTO fridges (name) VALUES ('Yamada home');
INSERT INTO fridges (name) VALUES ('Dormitory');
-- INSERT INTO user_fridge_relations values (1,1);
-- INSERT INTO user_fridge_relations values (2,2);
-- INSERT INTO user_fridge_relations values (3,2)
