/*
 Navicat Premium Data Transfer

 Source Server         : Localhost
 Source Server Type    : MySQL
 Source Server Version : 50736
 Source Host           : localhost:3306
 Source Schema         : book-api

 Target Server Type    : MySQL
 Target Server Version : 50736
 File Encoding         : 65001

 Date: 28/07/2023 17:38:58
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for books
-- ----------------------------
DROP TABLE IF EXISTS `books`;
CREATE TABLE `books`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `title` longtext CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL,
  `author` longtext CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL,
  `description` longtext CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 5 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of books
-- ----------------------------
INSERT INTO `books` VALUES (1, 'The Girl who Fell Beneath the Seas', 'Axie Oh', 'Badai mematikan telah merusak tanah kelahiran Mina selama beberapa generasi. Banjir menyapu seluruh desa, sementara perang berdarah dikobarkan untuk memperebutkan sumber daya yang tersisa. Masyarakat di desa Mina memercayai bahwa Dewa Laut mengutuk mereka dengan kematian dan keputusasaan. Dalam upaya untuk menenangkan Dewa Laut, setiap tahun seorang gadis cantik dibuang ke laut untuk menjadi pengantin Dewa Laut, dengan harapan suatu hari \"pengantin sejati\" akan dipilih dan mengakhiri penderitaan mereka. Shim Cheong adalah gadis tercantik di desa, sekaligus kekasih Joon. Banyak yang percaya dialah pengantin sejati Dewa Laut Tapi pada malam Cheong hendak dikorbankan, Joon mengikuti Cheong, meski tahu bahwa ikut campur akan dihadiahi hukuman mati. Untuk menyelamatkan kakaknya, Mina terjun ke air menggantikan Cheong. Kecantikan Mina memang tidak sebanding dengan para pengantin terdahulu. Tapi dia sudah tersapu ke Alam Roh. Kini Mina tak punya banyak waktu untuk menemukan Dewa Laut sebelum dirinya sendiri berubah menjadi arwah. Berbekal kemampuan mendongeng dan bantuan para arwah, Mina harus berhasil menemukan Dewa Laut.');
INSERT INTO `books` VALUES (2, 'Xoxo', 'Axie Oh', 'Jenny mendedikasikan hidupnya bermain dan berlatih selo agar diterima di sekolah musik impiannya. Hingga dia bertemu cowok tampan dan misterius. Dalam momen spontan itu, Jenny membiarkan dirinya bebas dan menikmati petualangan semalam yang tak terlupakan… sebelum cowok itu menghilang tanpa kabar. Tiga bulan kemudian, Jenny kembali bertemu cowok itu. Bukan cuma bersekolah di tempat yang sama, Jaewoo ternyata juga anggota salah satu grup K-pop terbesar di dunia! Semakin mengenal Jaewoo, Jenny semakin tidak bisa memilih antara masa depan yang sudah dia rancang bertahun-tahun dan ledakan-ledakan manis penuh warna kehidupan Seoul.');
INSERT INTO `books` VALUES (3, 'Hate First, Love You Later', 'Nimas Disri', '“Mau sejauh apapun aku mengembara, yang namanya cinta itu akan selalu tahu jalan pulang. Jalanku itu rindu, rumahku itu kamu.”\r\n\r\nMenceritakan tentang Gwen, seorang karyawati kantoran dan juga Bara, bos Gwen yang (meskipun sebenarnya normal) suka nyuruh-nyuruh ini itu. Bagaimana jika suatu hari kamu punya bos yang nyebelin banget?. Suka banget memerintah, kayak gak ada habisnya?. Itulah yang terjadi pada sosok Gwen. Gwen ini benar-benar dibuat jengkel banget dengan kelakuan Bara, atasannya di kantor. Entah kenapa ada saja yang diperintahkan Bara untuk dilakukan Gwen meskipun hal-hal itu terkesan sepele saja. Seakan-akan Gwen itu tidak boleh nganggur sedikitpun di kantor, dengan kata lain \"GABUT\".\r\n\r\nNah, anehnya Bara kemudian tiba tiba berubah. Sejak Gwen bertemu bahkan dekat dengan Reno, anak Bara. Alih-alih memerintah Gwen kayak biasanya, sekarang malah Bara berubah haluan dari kesana menjadi kesini. Bara menunjukkan perhatian, padahal selama ini kan dia… ehem, ya begitulah.\r\n\r\nNamun, belakangan Bara sering menghampiri Gwen yang sedang menunggu ojek online, lalu menawarkan pulang bareng. Apa yang sebenarnya terjadi? Gwen pun bingung. Semakin Gwen berusaha menjauh, Bara malah memberinya tugas yang mengharuskan mereka berdua bertemu. Kenapa sih sebenarnya bos gue ini?. Pertanyaan yang berusaha Gwen cari tahu jawabannya. Jadi, kira-kira ada apa dengan Bara?. Bagaimana kelanjutan kisah mereka ya?.');
INSERT INTO `books` VALUES (4, 'The Girl Who Fell to Earth', 'Sophia Al-Maria', 'When Sophia Al-Maria\'s mother sends her away from rainy Washington State to stay with her husband\'s desert-dwelling Bedouin family in Qatar, she intends it to be a sort of teenage cultural boot camp. What her mother doesn\'t know is that there are some things about growing up that are universal. In Qatar, Sophia is faced with a new world she\'d only imagined as a child. She sets out to find her freedom, even in the most unlikely of places.\r\n\r\nBoth family saga and coming-of-age story, takes readers from the green valleys of the Pacific Northwest to the dunes of the Arabian Gulf and on to the sprawling chaos of Cairo. Struggling to adapt to her nomadic lifestyle, Sophia is haunted by the feeling that she is perpetually in exile: hovering somewhere between two families, two cultures, and two worlds. She must make a place for herself a complex journey that includes finding young love in the Arabian Gulf, rebellion in Cairo, and, finally, self-discovery in the mountains of Sinai.\r\n\r\nheralds the arrival of an electric new talent and takes us on the most personal of quests: the voyage home.');

SET FOREIGN_KEY_CHECKS = 1;
