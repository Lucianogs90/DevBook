USE golang;

INSERT INTO usuarios (nome, nick, email, senha)
VALUES
("Luciano", "Luço", "luco.familia@gmail.com", "$2a$10$oLtRHiDQMwC3sMlpBoH14.zCnIWv9ZMnFTH.VCmCmmT5sNm0cEoxO"),
("Fernando", "Fer", "fer.familia@gmail.com", "$2a$10$oLtRHiDQMwC3sMlpBoH14.zCnIWv9ZMnFTH.VCmCmmT5sNm0cEoxO"),
("Carlos Eduardo", "Edu", "edu.familia@gmail.com", "$2a$10$oLtRHiDQMwC3sMlpBoH14.zCnIWv9ZMnFTH.VCmCmmT5sNm0cEoxO"),
("Claudeir Rogério", "Craudin", "claudeir.familia@gmail.com", "$2a$10$oLtRHiDQMwC3sMlpBoH14.zCnIWv9ZMnFTH.VCmCmmT5sNm0cEoxO"),
("Socorro", "Corrinha", "socorro.familia@gmail.com", "$2a$10$oLtRHiDQMwC3sMlpBoH14.zCnIWv9ZMnFTH.VCmCmmT5sNm0cEoxO"),
("José Antônio", "Zé", "pai.familia@gmail.com", "$2a$10$oLtRHiDQMwC3sMlpBoH14.zCnIWv9ZMnFTH.VCmCmmT5sNm0cEoxO"),
("Maria Virgínia", "Donmaria", "mae.familia@gmail.com", "$2a$10$oLtRHiDQMwC3sMlpBoH14.zCnIWv9ZMnFTH.VCmCmmT5sNm0cEoxO"),
("Carlos André", "Dé", "de.familia@gmail.com", "$2a$10$oLtRHiDQMwC3sMlpBoH14.zCnIWv9ZMnFTH.VCmCmmT5sNm0cEoxO"),
("Maria Vanuza", "Noza", "noza.familia@gmail.com", "$2a$10$oLtRHiDQMwC3sMlpBoH14.zCnIWv9ZMnFTH.VCmCmmT5sNm0cEoxO"), 
("Antônio Messias", "Tõin", "toin.familia@gmail.com", "$2a$10$oLtRHiDQMwC3sMlpBoH14.zCnIWv9ZMnFTH.VCmCmmT5sNm0cEoxO"),
("José Carlos", "Guila", "guila.familia@gmail.com", "$2a$10$oLtRHiDQMwC3sMlpBoH14.zCnIWv9ZMnFTH.VCmCmmT5sNm0cEoxO"),
("Vanessa", "Bula", "bula.familia@gmail.com", "$2a$10$oLtRHiDQMwC3sMlpBoH14.zCnIWv9ZMnFTH.VCmCmmT5sNm0cEoxO");

INSERT INTO seguidores (usuario_id, seguidor_id)
VALUES
(1, 2),
(1, 3),
(1, 4),
(1, 5),
(2, 1),
(2, 2),
(2, 3),
(3, 1),
(3, 2),
(4, 6),
(5, 7),
(10, 1);
