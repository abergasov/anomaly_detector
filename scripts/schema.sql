create table events_count
(
    ec_id int auto_increment,
    entity_id int null,
    event_date timestamp default current_timestamp,
    event_label char(50) null,
    event_counter int null,
    constraint events_count_pk
        primary key (ec_id)
);

create index events_count_entity_id_index
    on events_count (entity_id);

create index events_count_event_date_index
    on events_count (event_date);

create index events_count_event_label_index
    on events_count (event_label);

