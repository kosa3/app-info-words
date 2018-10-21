CREATE TABLE
IF NOT EXISTS `appli_info_words`
(
  `id` int
(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `appli_category_id` int
(11) NOT NULL COMMENT 'カテゴリーID',
  `name` varchar
(256) NOT NULL COMMENT '表示名',
  `code` varchar
(256) NOT NULL COMMENT 'コード',
  `description` varchar
(256) NOT NULL COMMENT '説明文',
  PRIMARY KEY
(`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='情報単語集';

CREATE TABLE IF NOT EXISTS `appli_info_categories`
(
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` varchar(256) NOT NULL COMMENT 'カテゴリ名',
  PRIMARY KEY
(`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='情報カテゴリー';
