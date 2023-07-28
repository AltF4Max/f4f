-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Хост: 127.0.0.1:3306
-- Время создания: Июл 24 2023 г., 19:42
-- Версия сервера: 8.0.30
-- Версия PHP: 7.2.34

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- База данных: `something`
--

-- --------------------------------------------------------

--
-- Структура таблицы `messages`
--

CREATE TABLE `messages` (
  `id` int UNSIGNED NOT NULL,
  `login` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `header` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `message` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `created` datetime NOT NULL,
  `updated` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Дамп данных таблицы `messages`
--

INSERT INTO `messages` (`id`, `login`, `header`, `message`, `created`, `updated`) VALUES
(3, 'admin', '66f66666', '666f666', '2023-07-20 19:29:36', '2023-07-23 20:45:42'),
(4, 'admin', '1a', 'f5', '2023-07-20 19:29:44', '2023-07-20 19:29:44'),
(5, 'admin', '1a', 'f5', '2023-07-20 19:30:49', '2023-07-20 19:30:49'),
(6, 'admin', '1a', 'hhf5', '2023-07-20 19:31:08', '2023-07-20 19:31:08'),
(9, 'admin', '1sDDDDDa', 'f5AAAAAA', '2023-07-21 12:30:46', '2023-07-21 20:48:23'),
(10, 'admin', '1sa', 'fо5', '2023-07-21 12:31:31', '2023-07-21 12:31:31'),
(11, 'admin', '1sa', 'fо5', '2023-07-21 12:35:59', '2023-07-21 12:35:59'),
(12, 'admin', '1sa', 'fо5', '2023-07-21 12:37:08', '2023-07-21 12:37:08'),
(13, 'admin', '1sa', 'fо5', '2023-07-21 12:41:46', '2023-07-21 12:41:46'),
(14, 'admin', '1sa', 'fо5', '2023-07-21 12:43:20', '2023-07-21 12:43:20'),
(15, 'admin', '1sa', 'fо5', '2023-07-21 12:45:45', '2023-07-21 12:45:45'),
(16, 'admin', '1sa', 'fо5', '2023-07-21 12:51:15', '2023-07-21 12:51:15'),
(17, 'admin', '1sa', 'fо5', '2023-07-21 13:04:39', '2023-07-21 13:04:39'),
(18, 'admin', '1sa', 'fо5', '2023-07-21 16:24:49', '2023-07-21 16:24:49'),
(19, 'admin', '1sa', 'fо5', '2023-07-21 16:26:53', '2023-07-21 16:26:53'),
(20, 'admin', '1sa', 'fо5', '2023-07-21 16:27:30', '2023-07-21 16:27:30'),
(21, 'admin', '1sa', 'fо5', '2023-07-21 16:30:01', '2023-07-21 16:30:01'),
(22, 'admin', '1sa', 'fо5', '2023-07-21 16:32:17', '2023-07-21 16:32:17'),
(23, 'admin', '1sa', 'fо5', '2023-07-21 16:36:27', '2023-07-21 16:36:27'),
(24, 'admin', '1sa', 'fо5', '2023-07-21 16:37:39', '2023-07-21 16:37:39'),
(25, 'admin', '1sa', 'fо5', '2023-07-21 16:40:25', '2023-07-21 16:40:25'),
(26, 'admin', '1sa', 'fо5', '2023-07-21 16:50:20', '2023-07-21 16:50:20'),
(27, 'admin', '1sa', 'fо5', '2023-07-21 16:52:09', '2023-07-21 16:52:09'),
(28, 'admin', '1sa', 'fо5', '2023-07-21 16:52:35', '2023-07-21 16:52:35'),
(29, 'admin', '1sa', 'fо5', '2023-07-21 17:09:23', '2023-07-21 17:09:23'),
(30, 'admin', '1sDDDDDa', 'f5AAAAAA', '2023-07-21 18:36:28', '2023-07-21 19:08:19'),
(31, 'admin', '1sa', 'fо5', '2023-07-21 18:41:48', '2023-07-21 18:41:48'),
(32, 'admin', '1sDDDDDa', 'f5AAAAAA', '2023-07-21 18:43:57', '2023-07-21 18:44:03'),
(33, 'admin', '1sDDDDDa', 'f5AAAAAA', '2023-07-21 18:46:19', '2023-07-21 18:46:23'),
(34, 'admin', '1sa', 'fо5', '2023-07-21 18:51:35', '2023-07-21 18:51:35'),
(35, 'admin', '1sa', 'fо5', '2023-07-21 19:01:51', '2023-07-21 19:01:51'),
(36, 'admin', '1sa', 'fо5', '2023-07-21 19:02:33', '2023-07-21 19:02:33'),
(37, 'admin', '1sa', 'fо5', '2023-07-21 19:07:47', '2023-07-21 19:07:47'),
(38, 'admin', '1sa', 'fо5', '2023-07-21 19:21:31', '2023-07-21 19:21:31'),
(39, 'admin', '1sa', 'fо5', '2023-07-21 19:22:08', '2023-07-21 19:22:08'),
(40, 'admin', '1sa', 'f5', '2023-07-21 20:26:29', '2023-07-21 20:26:29'),
(41, 'admin', '1sa', 'f5', '2023-07-21 20:27:26', '2023-07-21 20:27:26'),
(42, 'admin', '1sa', 'ddddf5', '2023-07-21 20:28:47', '2023-07-21 20:28:47'),
(43, 'admin', '1sa', 'f5', '2023-07-21 20:30:41', '2023-07-21 20:30:41'),
(44, 'admin', '1sa', 'f5', '2023-07-21 20:30:47', '2023-07-21 20:30:47'),
(45, 'admin', '1sa', 'ddddf5', '2023-07-21 20:31:49', '2023-07-21 20:31:49'),
(46, 'admin', '1sa', 'ddddf5', '2023-07-21 20:31:55', '2023-07-21 20:31:55'),
(47, 'admin', '1sa', 'ddddf5', '2023-07-21 20:32:10', '2023-07-21 20:32:10'),
(48, 'admin', '1sa', 'ddddf5', '2023-07-21 20:34:02', '2023-07-21 20:34:02'),
(49, 'admin', '1sa', 'ddddf5', '2023-07-21 20:34:06', '2023-07-21 20:34:06'),
(50, 'admin', '1sa', 'ddddf5', '2023-07-21 20:34:10', '2023-07-21 20:34:10'),
(51, 'admin', '1ddddsa', 'ddddf5', '2023-07-21 20:35:50', '2023-07-21 20:35:50'),
(52, 'admin', '1ddddsa', 'ddddf5', '2023-07-21 20:36:35', '2023-07-21 20:36:35'),
(53, 'admin', '1sa', 'ddddf5', '2023-07-21 20:36:43', '2023-07-21 20:36:43'),
(54, 'admin', '1sa', 'ddddf5', '2023-07-21 20:43:22', '2023-07-21 20:43:22'),
(55, 'admdddin', '1sa', 'f5', '2023-07-21 22:46:15', '2023-07-21 22:46:15'),
(56, 'admdddin', '1sa', 'f5', '2023-07-21 22:58:01', '2023-07-21 22:58:01'),
(57, 'admin', '1sdda', 'f5', '2023-07-21 23:00:32', '2023-07-21 23:00:32'),
(58, 'admdin', '1sdda', 'f5', '2023-07-21 23:08:23', '2023-07-21 23:08:23'),
(59, 'admin', '1sdda', 'f5', '2023-07-21 23:08:26', '2023-07-21 23:08:26'),
(60, 'admdin', '1sdda', 'f5', '2023-07-23 17:52:44', '2023-07-23 17:52:44'),
(61, 'admin', '1sdda', 'f5', '2023-07-23 17:54:55', '2023-07-23 17:54:55'),
(62, 'admdin', '1sdda', 'f5', '2023-07-23 18:16:02', '2023-07-23 18:16:02'),
(64, 'admdddin', '111', '222', '2023-07-23 20:11:48', '2023-07-23 20:11:48'),
(65, 'addmdddin', '6666666', '6666626', '2023-07-23 20:14:27', '2023-07-23 20:49:51');

-- --------------------------------------------------------

--
-- Структура таблицы `users`
--

CREATE TABLE `users` (
  `id` int UNSIGNED NOT NULL,
  `login` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `password_hash` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `token_jwt` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `email` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `last_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `first_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Дамп данных таблицы `users`
--

INSERT INTO `users` (`id`, `login`, `password_hash`, `token_jwt`, `email`, `last_name`, `first_name`) VALUES
(1, 'admin', '$2a$10$s0BO.GcCwDawOOAQwNpoc.I/uKUCiVPFZ0hljtiBvcYtl1lY0km/6', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySUQiOiJhZG1pbiIsImV4cCI6MTcxMTQ3MjQ4OH0.UI79VMNx3uKlqhB3U_VKihm3cmOLDEy0Nrvh8wELmh8', 'koljionu@hydert@qder.uft', 'Volioni', 'Xio'),
(2, 'admdin', '$2a$10$0vxrZOAaEaLuQeeR0LRM1.oHm2EIMuiLKQWFiXiKHM5RCtpRIju1m', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySUQiOiJhZG1kaW4iLCJleHAiOjE3MTE1NjU1MDh9.bz1nRPLyrEklKek-L4Pmd5HG_BR7a48phYSvkN1_j4s', 'admin', 'admin', 'admin'),
(3, 'admdddin', '$2a$10$a9dkX3a9r7btMohuOivXuuqCvt0QsmPXa8B1YDpj.MJYx7DtxIcdm', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySUQiOiJhZG1kZGRpbiIsImV4cCI6MTcxMTU2ODM3Nn0.K_Qra_zTwtpQNfXX5jZKp8EWrejEhyS4F0npRIksVrE', 'admin', 'admin', 'admin'),
(4, 'admdddi', '$2a$10$OJIXogUEktN6xqFMwzqfU.EHH2tzY1lH/VfvchdPviX8lILzunsge', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySUQiOiJhZG1kZGRpIiwiZXhwIjoxNzExNzI5MzQ5fQ.97Gu4NGR4mpyWR7_MSF8Y3_sevjKfsvIA-OU1fQGCAM', 'admin', 'admin', 'admin'),
(5, 'admddi', '$2a$10$GzLxWTw0BbReH5/V7rzfv.EWdJoyEnQP7700nGPU8lBuEwWoMnWM2', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySUQiOiJhZG1kZGkiLCJleHAiOjE3MTE3MjkzNzF9.3KMX5pvtJ4X3oB1--2D6yp4vTfFtQH6j9qlL7stG5Mg', 'admin', 'admin', 'admin'),
(6, 'adddmdddi', '$2a$10$sHRW4Bs08nZwNV8RxLMCa.MZJ3sfmHPnn.mVY8FyXKgpL8hMvg5MK', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySUQiOiJhZGRkbWRkZGkiLCJleHAiOjE3MTE3Mjk3NTR9.6QPjNC4M7ghQBtcwj-g2KKzvJaL5ZkKkUz4qa_ZhNjw', 'admin', 'admin', 'admin'),
(7, 'addmdddin', '$2a$10$4ZsnhwLCdi/yN6sRs3fMAOsJ9nzazgKziOT0GlOf3u5u9oGpsc9ee', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySUQiOiJhZGRtZGRkaW4iLCJleHAiOjE3MTE3MzIwMDJ9.7UqxDBf0WkVL9VDWOVzFvA_vtj6jK5px0iFlEGFmmxY', 'admin', 'admin', 'admin');

--
-- Индексы сохранённых таблиц
--

--
-- Индексы таблицы `messages`
--
ALTER TABLE `messages`
  ADD PRIMARY KEY (`id`);

--
-- Индексы таблицы `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT для сохранённых таблиц
--

--
-- AUTO_INCREMENT для таблицы `messages`
--
ALTER TABLE `messages`
  MODIFY `id` int UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=66;

--
-- AUTO_INCREMENT для таблицы `users`
--
ALTER TABLE `users`
  MODIFY `id` int UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
