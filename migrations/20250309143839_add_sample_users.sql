-- +goose Up
-- +goose StatementBegin

INSERT INTO users (id, name, last_name, email, password, created_at, updated_at, account_role) VALUES
    ('74f5e34d-a966-420a-8427-f3abb843ee0f', 'Adil', 'Zhumashev', 'adil.zhumashev@mail.com', '$2a$10$ABC12345xyz', NOW(), NOW(), 'user'),
    ('8b864659-a44a-41a6-8bfe-0821d5dbd0dc', 'Aigul', 'Serikbayeva', 'aigul.serikbayeva@mail.com', '$2a$10$DEF67890uvw', NOW(), NOW(), 'user'),
    ('4ad4a8f2-5743-4c51-ba05-79147ed01a14', 'Bauyrzhan', 'Myrzaliev', 'bauyrzhan.myrzaliev@mail.com', '$2a$10$GHI13579rst', NOW(), NOW(), 'user'),
    ('a0f389a1-f340-42da-9d8f-e54bbb01a83d', 'Erlan', 'Sagyndykov', 'erlan.sagyndykov@mail.com', '$2a$10$JKL24680mno', NOW(), NOW(), 'user'),
    ('8a9508d6-17ff-4582-b108-d5f03d906f3a', 'Alina', 'Tulegenova', 'alina.tulegenova@mail.com', '$2a$10$MNO97531xyz', NOW(), NOW(), 'user'),
    ('b2026668-3ed4-49cf-adde-89eba5f827e4', 'Damir', 'Alimkhanov', 'damir.alimkhanov@mail.com', '$2a$10$PQR86420uvw', NOW(), NOW(), 'user'),
    ('f17b3266-2cfc-4973-9b22-b058582df736', 'Samat', 'Orazbayev', 'samat.orazbayev@mail.com', '$2a$10$STU15937rst', NOW(), NOW(), 'user'),
    ('5b268d42-dc4a-42fb-ad63-8d0c1422ae59', 'Aliya', 'Ramazanova', 'aliya.ramazanova@mail.com', '$2a$10$VWX26849mno', NOW(), NOW(), 'user'),
    ('da99333d-5f39-4a38-8bdc-1375d3092cb5', 'Nursultan', 'Bekmuratov', 'nursultan.bekmuratov@mail.com', '$2a$10$YZA35768xyz', NOW(), NOW(), 'user'),
    ('5236b356-aad4-45ea-91d5-023a5917e148', 'Aruzhan', 'Beisekova', 'aruzhan.beisekova@mail.com', '$2a$10$BCD12345uvw', NOW(), NOW(), 'user'),
    ('c3e99151-af51-4ce4-9d30-fcd976162ed7', 'Madina', 'Yesenzhanova', 'madina.yesenzhanova@mail.com', '$2a$10$EFG67890rst', NOW(), NOW(), 'user'),
    ('62b61f53-3d68-481c-9894-767e34b9d831', 'Zhandos', 'Togzhanov', 'zhandos.togzhanov@mail.com', '$2a$10$HIJ13579mno', NOW(), NOW(), 'user'),
    ('53b7c283-5ca4-4f70-8dc2-fd0c69479d12', 'Dana', 'Tuleshova', 'dana.tuleshova@mail.com', '$2a$10$KLM24680xyz', NOW(), NOW(), 'user'),
    ('58e83dd0-8806-434f-8c49-86fd028129e3', 'Zhanbolat', 'Serikov', 'zhanbolat.serikov@mail.com', '$2a$10$NOP97531uvw', NOW(), NOW(), 'user'),
    ('bce6d0e5-2d07-4a61-8d24-7b67840eac24', 'Gulnaz', 'Utegenova', 'gulnaz.utegenova@mail.com', '$2a$10$QRS86420rst', NOW(), NOW(), 'user'),
    ('d873c134-9762-46e1-969a-32b5b9387ea1', 'Aibek', 'Karabayev', 'aibek.karabayev@mail.com', '$2a$10$TUV15937mno', NOW(), NOW(), 'user'),
    ('a558bfdd-6cee-445f-b7fd-a2215c4e1c6a', 'Ruslan', 'Talgatov', 'ruslan.talgatov@mail.com', '$2a$10$WXY26849xyz', NOW(), NOW(), 'user'),
    ('d850d3ae-3875-4d62-a205-e57031c2a341', 'Kamila', 'Zhandarova', 'kamila.zhandarova@mail.com', '$2a$10$ZAB35768uvw', NOW(), NOW(), 'user'),
    ('082d9f06-c8c9-4977-9d3e-52fb16a57667', 'Miras', 'Sagynov', 'miras.sagynov@mail.com', '$2a$10$CDE12345rst', NOW(), NOW(), 'user'),
    ('f6f55eba-e94c-47e5-b30a-0aa4cfc477d3', 'Serik', 'Bekbolatov', 'serik.bekbolatov@mail.com', '$2a$10$FGH67890mno', NOW(), NOW(), 'user'),
    ('26ac60f0-8fe5-471f-a489-5902741f51ac', 'Aruzhan', 'Saparova', 'aruzhan.saparova@mail.com', '$2a$10$IJK13579xyz', NOW(), NOW(), 'user'),
    ('4712ce26-7052-4a6c-9877-1b745063c94e', 'Amina', 'Tuleshova', 'amina.tuleshova@mail.com', '$2a$10$LMN24680uvw', NOW(), NOW(), 'user');


INSERT INTO users (id, name, last_name, email, password, created_at, updated_at, account_role) VALUES
    ('f315df8e-b8cd-4c3e-b2eb-c287acaf8c31', 'Dias', 'Nurzhanov', 'dias.nurzhanov@mail.com', '$2a$10$XyZT8aF1L93J8A1KwZUPuOTIVb1oPZlUZrR3FG9d7zFyyMKo1P5Qe', NOW(), NOW(), 'user'),
    ('0a49a92a-694c-429f-b77a-6994ff4441e0', 'Kamila', 'Omarova', 'kamila.omarova@mail.com', '$2a$10$P6VvJ6/91KQYf9L2FlpdZuqoxIQ5/p13AqHEgPfbMRZsGR9RXN8GC', NOW(), NOW(), 'user'),
    ('639fd406-6af2-4e50-a090-0d92d428749d', 'Rustem', 'Alpysov', 'rustem.alpysov@mail.com', '$2a$10$ePoOeC6VzIZRpLfeg2ft/u1tPpFb/kbYXe7m9sErT6S1XyY35OuG6', NOW(), NOW(), 'user'),
    ('b9cf3d64-2b19-4d02-812f-cbebbb437bbc', 'Madina', 'Bekzhanova', 'madina.bekzhanova@mail.com', '$2a$10$zJ72eQdZdUfwW3RsOZptm.r5buvP/PXMGpDeHpI7S7UjLKQ7fCdmW', NOW(), NOW(), 'user'),
    ('2de42bcb-a2ea-40cf-8d86-6498c09b3d1f', 'Sanzhar', 'Tursynbayev', 'sanzhar.tursynbayev@mail.com', '$2a$10$DE3G7T5V/hJvP5Q./MllhWflFwGrECq9v6m0spUETADuE9CU3m0qC', NOW(), NOW(), 'user'),
    ('1dbf9a0e-aa30-4002-ac80-6bb6d8e19495', 'Laura', 'Karimova', 'laura.karimova@mail.com', '$2a$10$PCjLfdf8ae4AB6ZqG9hBVuoU3izQFQ/hYOA4tcX2g7p2XhTS.B1Qq', NOW(), NOW(), 'user'),
    ('ec9f58d5-82d5-442e-a664-ad21f2d0c17b', 'Zhanbolat', 'Ryskulov', 'zhanbolat.ryskulov@mail.com', '$2a$10$8z/PoOTXbHdkeGyOC6tUGeX33IUp69dw82G9VvbhqHGJx9nnl1y96', NOW(), NOW(), 'user'),
    ('3c5456ff-eef4-4c30-9f37-e52566fc7032', 'Aliya', 'Tulegenova', 'aliya.tulegenova2@mail.com', '$2a$10$NjF7iE3VkgY3UdF/k/TNQu96YZsDO7XHzF1E9kq4IHyCZ9kOkKJQG', NOW(), NOW(), 'user'),
    ('13e42502-d51a-42c8-a8a8-b92e1104bdd4', 'Nursultan', 'Bektayev', 'nursultan.bektayev2@mail.com', '$2a$10$V.Ro9dtR/Va/FJGzHHi/MuPbw3bnMNyTXzOYKNrQZsH6sXTLiEFJq', NOW(), NOW(), 'user'),
    ('964b13c7-3798-4551-8f5e-db7c181e8320', 'Samat', 'Esenov', 'samat.esenov2@mail.com', '$2a$10$A96LQoFQ7Gk0v0PxWGP/sTg1fHNkAP10XVrAcR0Ry9OyF9r0HheF6', NOW(), NOW(), 'user'),
    ('106ddee2-e2a9-4035-921f-bd5cbe7a628e', 'Aruzhan', 'Amankeldi', 'aruzhan.amankeldi2@mail.com', '$2a$10$67LudOeQH3XPi/T34P3TZeVh6lqNOYO06d41xDD/WwZVcAq8zy9Bi', NOW(), NOW(), 'user'),
    ('c992013c-5cde-48cf-8af0-00184338baf9', 'Yerlan', 'Bekturov', 'yerlan.bekturov2@mail.com', '$2a$10$3DTSkJLV/zs/n/JD9Vj5r9J1Dkp76xv1cQ8aVFYBTeFvTVAfnVuV6', NOW(), NOW(), 'user'),
    ('9fa1f519-d805-42c7-b9b1-a752365d1c56', 'Marat', 'Bekmuratov', 'marat.bekmuratov2@mail.com', '$2a$10$xQMT9T0KpdM/Rd7BnvNfOQpkZmHyXAVq8H5xPe28ldP0cRQT6UWxS', NOW(), NOW(), 'user'),
    ('d131a22f-0b3d-4767-9806-6a277dc0744d', 'Dana', 'Tuleshova', 'dana.tuleshova2@mail.com', '$2a$10$yJ63ImfI0pSp5VpMncZk8nE5zyD2iK1WJ7vPPUuXzgfVRMlDl6J62', NOW(), NOW(), 'user'),
    ('30bd5baf-d603-4b23-936b-c3633df709fe', 'Alisher', 'Zhanabayev', 'alisher.zhanabayev2@mail.com', '$2a$10$5NL6WvlKq/RFGl1FQvS18.eCYy/rhVbg1MvnZxzzPImwP8AlyJd5a', NOW(), NOW(), 'user'),
    ('636e0dc6-ce94-4f26-bd39-b8fb32b4b378', 'Kamila', 'Sadyk', 'kamila.sadyk2@mail.com', '$2a$10$h9HDebQzCdt/ZD2KLU9G2u7PbvVhKzLFLW8cA4pPQMQVqFBw/VF.G', NOW(), NOW(), 'user'),
    ('cb2d6442-ec58-4fb6-aea6-23dcd2a6b574', 'Nurbek', 'Tulegenov', 'nurbek.tulegenov2@mail.com', '$2a$10$4azRWtZKquc3FaLp0Tk/FjShuJde0E6tFQ4/zJ/KD5S68J9kOHffCa', NOW(), NOW(), 'user'),
    ('13d450c5-6e2c-4bd8-bc5a-9096f59497b2', 'Amina', 'Ramazanova', 'amina.ramazanova2@mail.com', '$2a$10$wCD2/PQXNk8q7G6jP0KaMC7eN12LoYIuS88MbXM1UhZtI5nMl2C1a', NOW(), NOW(), 'user'),
    ('a58335b7-8e7c-4276-8c15-e4e16ca894ec', 'Miras', 'Togzhanov', 'miras.togzhanov2@mail.com', '$2a$10$DUO6B6Z/ltHcBaWoXFn84U3deV6ohB2xWy34A3ZRfbcVll8CpXdfm', NOW(), NOW(), 'user'),
    ('65148bce-f576-4e3a-b30b-4951cc1e3621', 'Serik', 'Bekbolatov', 'serik.bekbolatov2@mail.com', '$2a$10$7eOxj2oI.pCf7lSMYIT57mRQ2X0XHlfK4l.XF9W0LzPb8AzVh04lC', NOW(), NOW(), 'user'),
    ('cfd3b85c-2e4a-4fe6-8bfe-4eafe6612759', 'Dias', 'Nurtas', 'dias.nurtas@mail.com', '$2a$10$ABCD12345xyz', NOW(), NOW(), 'user'),
('8d44290e-c670-42ee-bbc1-bf5917dd80de', 'Aruzhan', 'Omarova', 'aruzhan.omarova@mail.com', '$2a$10$EFGH67890uvw', NOW(), NOW(), 'user');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
