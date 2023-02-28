-- +goose Up
CREATE TABLE cards
(
    id          UUID                 DEFAULT uuid_generate_v4() PRIMARY KEY,
    name        VARCHAR     NOT NULL,
    code        VARCHAR,
    description TEXT                 DEFAULT NULL,
    image       VARCHAR                 DEFAULT NULL,
    sort        INT                  DEFAULT 500,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

INSERT INTO cards
VALUES (uuid_generate_v4(), 'Карточка №1', 'item_1', 'Lorem Ipsum - это текст-рыба, часто используемый в печати и вэб-дизайне. Lorem Ipsum является стандартной рыбой для текстов на латинице с начала XVI века. В то время некий безымянный печатник создал большую коллекцию размеров и форм шрифтов, используя Lorem Ipsum для распечатки образцов.', 'http://placehold.it/350x200', 500),
       (uuid_generate_v4(), 'Карточка №2', 'item_2', 'Lorem Ipsum - это текст-рыба, часто используемый в печати и вэб-дизайне. Lorem Ipsum является стандартной рыбой для текстов на латинице с начала XVI века. В то время некий безымянный печатник создал большую коллекцию размеров и форм шрифтов, используя Lorem Ipsum для распечатки образцов.', 'http://placehold.it/350x200', 600),
       (uuid_generate_v4(), 'Карточка №3', 'item_3', 'Lorem Ipsum - это текст-рыба, часто используемый в печати и вэб-дизайне. Lorem Ipsum является стандартной рыбой для текстов на латинице с начала XVI века. В то время некий безымянный печатник создал большую коллекцию размеров и форм шрифтов, используя Lorem Ipsum для распечатки образцов.', 'http://placehold.it/350x200', 700),
       (uuid_generate_v4(), 'Карточка №4', 'item_4', 'Lorem Ipsum - это текст-рыба, часто используемый в печати и вэб-дизайне. Lorem Ipsum является стандартной рыбой для текстов на латинице с начала XVI века. В то время некий безымянный печатник создал большую коллекцию размеров и форм шрифтов, используя Lorem Ipsum для распечатки образцов.', 'http://placehold.it/350x200', 800),
       (uuid_generate_v4(), 'Карточка №5', 'item_5', 'Lorem Ipsum - это текст-рыба, часто используемый в печати и вэб-дизайне. Lorem Ipsum является стандартной рыбой для текстов на латинице с начала XVI века. В то время некий безымянный печатник создал большую коллекцию размеров и форм шрифтов, используя Lorem Ipsum для распечатки образцов.', 'http://placehold.it/350x200', 900),
       (uuid_generate_v4(), 'Карточка №6', 'item_6', 'Lorem Ipsum - это текст-рыба, часто используемый в печати и вэб-дизайне. Lorem Ipsum является стандартной рыбой для текстов на латинице с начала XVI века. В то время некий безымянный печатник создал большую коллекцию размеров и форм шрифтов, используя Lorem Ipsum для распечатки образцов.', 'http://placehold.it/350x200', 1000),
       (uuid_generate_v4(), 'Карточка №7', 'item_7', 'Lorem Ipsum - это текст-рыба, часто используемый в печати и вэб-дизайне. Lorem Ipsum является стандартной рыбой для текстов на латинице с начала XVI века. В то время некий безымянный печатник создал большую коллекцию размеров и форм шрифтов, используя Lorem Ipsum для распечатки образцов.', 'http://placehold.it/350x200', 1100),
       (uuid_generate_v4(), 'Карточка №8', 'item_8', 'Lorem Ipsum - это текст-рыба, часто используемый в печати и вэб-дизайне. Lorem Ipsum является стандартной рыбой для текстов на латинице с начала XVI века. В то время некий безымянный печатник создал большую коллекцию размеров и форм шрифтов, используя Lorem Ipsum для распечатки образцов.', 'http://placehold.it/350x200', 1200),
       (uuid_generate_v4(), 'Карточка №9', 'item_9', 'Lorem Ipsum - это текст-рыба, часто используемый в печати и вэб-дизайне. Lorem Ipsum является стандартной рыбой для текстов на латинице с начала XVI века. В то время некий безымянный печатник создал большую коллекцию размеров и форм шрифтов, используя Lorem Ipsum для распечатки образцов.', 'http://placehold.it/350x200', 1300),
       (uuid_generate_v4(), 'Карточка №10', 'item_10', 'Lorem Ipsum - это текст-рыба, часто используемый в печати и вэб-дизайне. Lorem Ipsum является стандартной рыбой для текстов на латинице с начала XVI века. В то время некий безымянный печатник создал большую коллекцию размеров и форм шрифтов, используя Lorem Ipsum для распечатки образцов.', 'http://placehold.it/350x200', 1400),
       (uuid_generate_v4(), 'Карточка №11', 'item_11', 'Lorem Ipsum - это текст-рыба, часто используемый в печати и вэб-дизайне. Lorem Ipsum является стандартной рыбой для текстов на латинице с начала XVI века. В то время некий безымянный печатник создал большую коллекцию размеров и форм шрифтов, используя Lorem Ipsum для распечатки образцов.', 'http://placehold.it/350x200', 1500);

-- +goose Down
DROP TABLE cards;