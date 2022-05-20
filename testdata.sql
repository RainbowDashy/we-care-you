INSERT INTO user(username, location, password_hash)
VALUES("test", "d1", "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08");
INSERT INTO user(username, location, password_hash)
VALUES("test2", "d2", "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08");

INSERT INTO mall(user_id, begin_time, end_time, state)
VALUES (1, 0, 1000000000, 1);
INSERT INTO mall(user_id, begin_time, end_time, state)
VALUES (1, 0, 1000000000, 1);
INSERT INTO mall(user_id, begin_time, end_time, state)
VALUES (2, 0, 1000000000, 1);

INSERT INTO item(mall_id, total, price, name, description, data)
VALUES (1, 10, 20, "番茄", "好吃的番茄", "");
INSERT INTO item(mall_id, total, price, name, description, data)
VALUES (1, 20, 10, "土豆", "好吃的土豆", "");

INSERT INTO item(mall_id, total, price, name, description, data)
VALUES (2, 1, 20, "苹果", "好吃的苹果", "");
INSERT INTO item(mall_id, total, price, name, description, data)
VALUES (2, 3, 10, "橘子", "好吃的橘子", "");


INSERT INTO item(mall_id, total, price, name, description, data)
VALUES (3, 1, 20, "矿泉水", "好喝的水", "");
INSERT INTO item(mall_id, total, price, name, description, data)
VALUES (3, 3, 10, "纸巾", "好用的纸巾", "");

INSERT INTO mall_customer(mall_id, user_id, mall_item_id, buy_count)
VALUES (1, 1, 1, 3);
INSERT INTO mall_customer(mall_id, user_id, mall_item_id, buy_count)
VALUES (1, 1, 2, 2);
INSERT INTO mall_customer(mall_id, user_id, mall_item_id, buy_count)
VALUES (2, 1, 3, 1);
