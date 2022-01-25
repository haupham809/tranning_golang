-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Máy chủ: 127.0.0.1
-- Thời gian đã tạo: Th1 24, 2022 lúc 11:46 AM
-- Phiên bản máy phục vụ: 10.4.22-MariaDB
-- Phiên bản PHP: 8.0.13

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Cơ sở dữ liệu: `tranning_goalng`
--

-- --------------------------------------------------------

--
-- Cấu trúc bảng cho bảng `t_category`
--

CREATE TABLE `t_category` (
  `id` int(11) NOT NULL,
  `category_name` int(11) NOT NULL,
  `date_created` datetime NOT NULL,
  `date_update` datetime DEFAULT NULL,
  `is_delete` tinyint(4) NOT NULL DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Cấu trúc bảng cho bảng `t_coupon`
--

CREATE TABLE `t_coupon` (
  `id` int(11) NOT NULL,
  `code` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `discount` float NOT NULL,
  `date_expiry` datetime DEFAULT NULL,
  `note` text DEFAULT NULL,
  `type` tinyint(4) NOT NULL DEFAULT 0,
  `date_created` datetime NOT NULL,
  `date_update` datetime DEFAULT NULL,
  `is_delete` tinyint(4) NOT NULL DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Đang đổ dữ liệu cho bảng `t_coupon`
--

INSERT INTO `t_coupon` (`id`, `code`, `name`, `discount`, `date_expiry`, `note`, `type`, `date_created`, `date_update`, `is_delete`) VALUES
(1, 'abcdef123', 'giam gia nguoi moi', 10, '2022-01-26 04:32:55', NULL, 1, '2022-01-24 04:32:55', '2022-01-24 04:32:55', 0);

-- --------------------------------------------------------

--
-- Cấu trúc bảng cho bảng `t_credit_card`
--

CREATE TABLE `t_credit_card` (
  `id` int(11) NOT NULL,
  `cc_number` varchar(50) NOT NULL,
  `cc_expiry` datetime NOT NULL,
  `cc_type_payment` tinyint(4) NOT NULL DEFAULT 0,
  `date_created` datetime NOT NULL,
  `date_update` datetime NOT NULL,
  `is_delete` tinyint(4) NOT NULL DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Cấu trúc bảng cho bảng `t_order`
--

CREATE TABLE `t_order` (
  `id` int(11) NOT NULL,
  `product_id` int(11) NOT NULL,
  `product_topping` int(11) DEFAULT NULL,
  `user_id` int(11) NOT NULL,
  `payment_id` int(11) DEFAULT NULL,
  `quantity` int(11) NOT NULL DEFAULT 1,
  `amount` float NOT NULL DEFAULT 1,
  `note` text DEFAULT NULL,
  `size` tinyint(4) NOT NULL DEFAULT 0,
  `shipment_date` datetime NOT NULL,
  `status` tinyint(4) NOT NULL DEFAULT 0,
  `date_created` datetime NOT NULL,
  `date_update` datetime DEFAULT NULL,
  `is_delete` tinyint(4) NOT NULL DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Đang đổ dữ liệu cho bảng `t_order`
--

INSERT INTO `t_order` (`id`, `product_id`, `product_topping`, `user_id`, `payment_id`, `quantity`, `amount`, `note`, `size`, `shipment_date`, `status`, `date_created`, `date_update`, `is_delete`) VALUES
(1, 15, 1, 22323, 12312, 1, 1, '23123', 0, '0000-00-00 00:00:00', 0, '2022-01-24 10:20:18', '2022-01-24 10:20:18', 0);

-- --------------------------------------------------------

--
-- Cấu trúc bảng cho bảng `t_payment`
--

CREATE TABLE `t_payment` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `credit_card_id` int(11) NOT NULL,
  `coupon_id` int(11) NOT NULL,
  `total_amount` float NOT NULL,
  `payment_result_id` varchar(255) NOT NULL,
  `pament_date` datetime NOT NULL,
  `refund_date` datetime NOT NULL,
  `is_refund` tinyint(4) NOT NULL DEFAULT 0,
  `date_created` datetime NOT NULL,
  `date_update` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Cấu trúc bảng cho bảng `t_product`
--

CREATE TABLE `t_product` (
  `id` int(11) NOT NULL,
  `category_id` int(11) NOT NULL,
  `name` varchar(100) NOT NULL,
  `id_suggest` tinyint(4) NOT NULL DEFAULT 0,
  `price` float NOT NULL,
  `date_created` datetime NOT NULL,
  `date_update` datetime DEFAULT NULL,
  `is_delete` tinyint(4) NOT NULL DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Đang đổ dữ liệu cho bảng `t_product`
--

INSERT INTO `t_product` (`id`, `category_id`, `name`, `id_suggest`, `price`, `date_created`, `date_update`, `is_delete`) VALUES
(1, 1, 'tra lài', 0, 30000, '2022-01-24 05:03:29', '2022-01-24 05:03:29', 0);

-- --------------------------------------------------------

--
-- Cấu trúc bảng cho bảng `t_product_topping`
--

CREATE TABLE `t_product_topping` (
  `id` int(11) NOT NULL,
  `product_id` int(11) NOT NULL,
  `topping_id` int(11) NOT NULL,
  `date_created` datetime NOT NULL,
  `date_update` datetime DEFAULT NULL,
  `is_delete` tinyint(4) NOT NULL DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Đang đổ dữ liệu cho bảng `t_product_topping`
--

INSERT INTO `t_product_topping` (`id`, `product_id`, `topping_id`, `date_created`, `date_update`, `is_delete`) VALUES
(1, 15, 1, '2022-01-24 04:01:16', '2022-01-24 04:01:16', 0);

-- --------------------------------------------------------

--
-- Cấu trúc bảng cho bảng `t_size_product`
--

CREATE TABLE `t_size_product` (
  `id` int(11) NOT NULL,
  `name` varchar(11) NOT NULL,
  `product_id` int(11) NOT NULL,
  `price` float NOT NULL,
  `date_created` datetime NOT NULL,
  `date_update` datetime NOT NULL,
  `is_delete` tinyint(4) NOT NULL DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Đang đổ dữ liệu cho bảng `t_size_product`
--

INSERT INTO `t_size_product` (`id`, `name`, `product_id`, `price`, `date_created`, `date_update`, `is_delete`) VALUES
(1, 's', 1, 30000, '2022-01-24 10:43:36', '2022-01-24 10:43:36', 0),
(2, 'm', 1, 30000, '2022-01-24 10:43:36', '2022-01-24 10:43:36', 0),
(3, 'L', 1, 30000, '2022-01-24 10:43:36', '2022-01-24 10:43:36', 0);

-- --------------------------------------------------------

--
-- Cấu trúc bảng cho bảng `t_topping`
--

CREATE TABLE `t_topping` (
  `id` int(11) NOT NULL,
  `name` varchar(50) NOT NULL,
  `price` float NOT NULL,
  `date_created` datetime NOT NULL,
  `date_update` datetime DEFAULT NULL,
  `is_delete` tinyint(4) NOT NULL DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Đang đổ dữ liệu cho bảng `t_topping`
--

INSERT INTO `t_topping` (`id`, `name`, `price`, `date_created`, `date_update`, `is_delete`) VALUES
(1, 'trân châu trắng', 3000, '2022-01-24 04:00:17', '2022-01-24 04:00:17', 0),
(2, 'trân châu đen', 3000, '2022-01-24 04:00:17', '2022-01-24 04:00:17', 0),
(3, 'trân châu đỏ', 3000, '2022-01-24 04:00:17', '2022-01-24 04:00:17', 0);

-- --------------------------------------------------------

--
-- Cấu trúc bảng cho bảng `t_user`
--

CREATE TABLE `t_user` (
  `id` int(11) NOT NULL,
  `name` varchar(50) NOT NULL,
  `email` varchar(50) NOT NULL,
  `password` varchar(50) NOT NULL,
  `address` varchar(255) NOT NULL,
  `phone` varchar(50) NOT NULL,
  `date_created` datetime NOT NULL,
  `date_update` datetime DEFAULT NULL,
  `is_delete` tinyint(4) NOT NULL DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Chỉ mục cho các bảng đã đổ
--

--
-- Chỉ mục cho bảng `t_category`
--
ALTER TABLE `t_category`
  ADD UNIQUE KEY `id` (`id`);

--
-- Chỉ mục cho bảng `t_coupon`
--
ALTER TABLE `t_coupon`
  ADD UNIQUE KEY `id` (`id`),
  ADD UNIQUE KEY `code` (`code`);

--
-- Chỉ mục cho bảng `t_credit_card`
--
ALTER TABLE `t_credit_card`
  ADD UNIQUE KEY `id` (`id`);

--
-- Chỉ mục cho bảng `t_order`
--
ALTER TABLE `t_order`
  ADD UNIQUE KEY `id` (`id`);

--
-- Chỉ mục cho bảng `t_payment`
--
ALTER TABLE `t_payment`
  ADD UNIQUE KEY `payment_result_id` (`payment_result_id`),
  ADD UNIQUE KEY `UniqueID` (`id`);

--
-- Chỉ mục cho bảng `t_product`
--
ALTER TABLE `t_product`
  ADD UNIQUE KEY `id` (`id`),
  ADD UNIQUE KEY `name` (`name`);

--
-- Chỉ mục cho bảng `t_product_topping`
--
ALTER TABLE `t_product_topping`
  ADD UNIQUE KEY `id` (`id`);

--
-- Chỉ mục cho bảng `t_size_product`
--
ALTER TABLE `t_size_product`
  ADD PRIMARY KEY (`id`);

--
-- Chỉ mục cho bảng `t_topping`
--
ALTER TABLE `t_topping`
  ADD UNIQUE KEY `id` (`id`);

--
-- Chỉ mục cho bảng `t_user`
--
ALTER TABLE `t_user`
  ADD UNIQUE KEY `id` (`id`);

--
-- AUTO_INCREMENT cho các bảng đã đổ
--

--
-- AUTO_INCREMENT cho bảng `t_category`
--
ALTER TABLE `t_category`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT cho bảng `t_coupon`
--
ALTER TABLE `t_coupon`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT cho bảng `t_credit_card`
--
ALTER TABLE `t_credit_card`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT cho bảng `t_order`
--
ALTER TABLE `t_order`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT cho bảng `t_payment`
--
ALTER TABLE `t_payment`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT cho bảng `t_product`
--
ALTER TABLE `t_product`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT cho bảng `t_product_topping`
--
ALTER TABLE `t_product_topping`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT cho bảng `t_size_product`
--
ALTER TABLE `t_size_product`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT cho bảng `t_topping`
--
ALTER TABLE `t_topping`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT cho bảng `t_user`
--
ALTER TABLE `t_user`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
