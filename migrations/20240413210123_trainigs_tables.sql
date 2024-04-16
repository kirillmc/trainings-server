-- +goose Up
create table train_programs
(
    id         serial primary key,
    program_name   varchar(50)      not null,
    description      varchar(500)      not null,
    is_public boolean not null default false
);



create table train_days
(
    id         serial primary key,
    day_name   varchar(50)      not null,
    description      varchar(500)      not null
);



create table exercises
(
    id         serial primary key,
    exercise_name   varchar(50)      not null,
    pictures      varchar(500)      not null,
    description      varchar(500)      not null
);



create table sets
(
    id         serial primary key,
    quantity   int      not null,
    weight      double precision      not null
);



create table statistics
(
    id         serial primary key,
    time timestamp not null default now(),
    quantity   int      not null,
    weight      double precision      not null,
    set_id int not null,
    exercise_id int not null,
    trains_day_id int not null,
    program_id int not null
);

create table users_programs
(
    id         serial primary key,
    user_id      int      not null,
    program_id   int      not null,
    foreign key (program_id) references train_programs(id) on delete cascade
);

create table days_programs
(
    id         serial primary key,
    program_id   int      not null,
    day_id      int      not null,
    foreign key (program_id) references train_programs(id) on delete cascade,
    foreign key (day_id) references train_days(id) on delete cascade
);
create table exercises_days
(
    id         serial primary key,
    day_id   int      not null,
    exercise_id      int      not null,
    foreign key (day_id) references train_days(id) on delete cascade,
    foreign key (exercise_id) references exercises(id) on delete cascade
);

create table sets_exercises
(
    id         serial primary key,
    exercise_id      int      not null,
    set_id   int      not null,
    foreign key (exercise_id) references exercises(id) on delete cascade,
    foreign key (set_id) references sets(id) on delete cascade
);

create table statistics_sets
(
    id         serial primary key,
    set_id   int      not null,
    statistic_id      int      not null,
    foreign key (set_id) references sets(id) on delete cascade,
    foreign key (statistic_id) references statistics(id) on delete cascade
);


-- +goose Down
drop table statistics_sets;
drop table sets_exercises;
drop table exercises_days;
drop table days_programs;
drop table users_programs;
drop table statistics;
drop table sets;
drop table exercises;
drop table train_days;
drop table train_programs;