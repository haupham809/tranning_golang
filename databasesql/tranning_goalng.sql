-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Máy chủ: 127.0.0.1
-- Thời gian đã tạo: Th1 22, 2022 lúc 04:04 PM
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
  `shipment_date` int(11) NOT NULL,
  `status` tinyint(4) NOT NULL DEFAULT 0,
  `date_created` datetime NOT NULL,
  `date_update` datetime DEFAULT NULL,
  `is_delete` tinyint(4) NOT NULL DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

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

-- --------------------------------------------------------

--
-- Cấu trúc bảng cho bảng `t_product_topping`
--

CREATE TABLE `t_product_topping` (
  `id` int(11) NOT NULL,
  `product-id` int(11) NOT NULL,
  `topping_id` int(11) NOT NULL,
  `date_created` datetime NOT NULL,
  `date_update` datetime DEFAULT NULL,
  `is_delete` tinyint(4) NOT NULL DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

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
  ADD PRIMARY KEY (`product_id`),
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
  ADD UNIQUE KEY `id` (`id`),
  ADD UNIQUE KEY `topping_id` (`topping_id`),
  ADD UNIQUE KEY `product-id` (`product-id`);

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
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT cho bảng `t_credit_card`
--
ALTER TABLE `t_credit_card`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT cho bảng `t_order`
--
ALTER TABLE `t_order`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT cho bảng `t_payment`
--
ALTER TABLE `t_payment`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT cho bảng `t_product`
--
ALTER TABLE `t_product`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT cho bảng `t_product_topping`
--
ALTER TABLE `t_product_topping`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT cho bảng `t_topping`
--
ALTER TABLE `t_topping`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT cho bảng `t_user`
--
ALTER TABLE `t_user`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
