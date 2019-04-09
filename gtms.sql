/*
 Navicat Premium Data Transfer

 Source Server         : virtual_machine
 Source Server Type    : MySQL
 Source Server Version : 50725
 Source Host           : 172.16.230.130:3306
 Source Schema         : gtms

 Target Server Type    : MySQL
 Target Server Version : 50725
 File Encoding         : 65001

 Date: 09/04/2019 01:22:41
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for admin
-- ----------------------------
DROP TABLE IF EXISTS `admin`;
CREATE TABLE `admin` (
  `admin_id` int(10) NOT NULL COMMENT '管理员id',
  `pwd` varchar(255) NOT NULL COMMENT '管理员密码',
  `admin_name` varchar(100) DEFAULT NULL COMMENT '姓名',
  `admin_sex` enum('男','女') DEFAULT NULL COMMENT '性别',
  `phone` varchar(20) DEFAULT NULL COMMENT '手机',
  `email` varchar(100) DEFAULT NULL COMMENT '邮箱',
  PRIMARY KEY (`admin_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='管理员表';

-- ----------------------------
-- Table structure for notice
-- ----------------------------
DROP TABLE IF EXISTS `notice`;
CREATE TABLE `notice` (
  `nid` int(10) NOT NULL AUTO_INCREMENT COMMENT '公告id',
  `title` varchar(200) DEFAULT NULL COMMENT '公告标题',
  `content` text COMMENT '内容',
  `attachment` varchar(1000) DEFAULT NULL COMMENT '附件',
  `view` int(10) DEFAULT NULL COMMENT '阅读次数',
  `create_user` varchar(100) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`nid`)
) ENGINE=InnoDB AUTO_INCREMENT=32 DEFAULT CHARSET=utf8 COMMENT='公告表';

-- ----------------------------
-- Table structure for opening_time
-- ----------------------------
DROP TABLE IF EXISTS `opening_time`;
CREATE TABLE `opening_time` (
  `year` int(5) NOT NULL COMMENT '年份',
  `start_time` datetime DEFAULT NULL COMMENT '开始时间',
  `end_time` datetime DEFAULT NULL COMMENT '结束时间',
  `operate_uid` int(10) DEFAULT NULL COMMENT '操作人id',
  `operate_name` varchar(100) DEFAULT NULL COMMENT '操作人姓名',
  `operate_time` datetime DEFAULT NULL COMMENT '操作时间',
  PRIMARY KEY (`year`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='开放选题时间表';

-- ----------------------------
-- Table structure for selected_thesis
-- ----------------------------
DROP TABLE IF EXISTS `selected_thesis`;
CREATE TABLE `selected_thesis` (
  `sid` int(10) NOT NULL AUTO_INCREMENT COMMENT '选题id',
  `uid` bigint(20) DEFAULT NULL COMMENT '学生id',
  `tid` bigint(20) DEFAULT NULL COMMENT '论文id',
  `tech_id` bigint(20) DEFAULT NULL COMMENT '教师id',
  `confirm` enum('0','1','2') DEFAULT NULL COMMENT '教师确认(0代表未确认,1代表同意，2代表不同意)',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `confirm_time` datetime DEFAULT NULL COMMENT '教师确认时间',
  PRIMARY KEY (`sid`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8 COMMENT='选题表';

-- ----------------------------
-- Table structure for student
-- ----------------------------
DROP TABLE IF EXISTS `student`;
CREATE TABLE `student` (
  `stu_no` bigint(20) NOT NULL COMMENT '学号',
  `pwd` varchar(255) NOT NULL COMMENT '密码',
  `stu_name` varchar(100) DEFAULT NULL COMMENT '姓名',
  `stu_sex` enum('男','女') DEFAULT NULL COMMENT '性别',
  `id_card` varchar(18) DEFAULT NULL COMMENT '身份证',
  `birthplace` varchar(100) DEFAULT NULL COMMENT '籍贯',
  `department` varchar(100) DEFAULT NULL COMMENT '院系',
  `major` varchar(100) DEFAULT NULL COMMENT '专业',
  `class` varchar(10) DEFAULT NULL COMMENT '班级',
  `phone` varchar(20) DEFAULT NULL COMMENT '手机号',
  `qq` varchar(20) DEFAULT NULL COMMENT 'qq号',
  `email` varchar(100) DEFAULT NULL COMMENT '邮箱',
  `we_chat` varchar(100) DEFAULT NULL COMMENT '微信',
  `school_system` int(2) DEFAULT NULL COMMENT '学制',
  `entry_date` datetime DEFAULT NULL COMMENT '入学日期',
  `education` varchar(10) DEFAULT NULL COMMENT '学历',
  PRIMARY KEY (`stu_no`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='学生表';

-- ----------------------------
-- Table structure for teacher
-- ----------------------------
DROP TABLE IF EXISTS `teacher`;
CREATE TABLE `teacher` (
  `tech_id` bigint(20) NOT NULL COMMENT '教师id',
  `pwd` varchar(255) NOT NULL COMMENT '密码',
  `tech_name` varchar(100) DEFAULT NULL COMMENT '姓名',
  `tech_sex` enum('男','女') DEFAULT NULL COMMENT '性别',
  `education` varchar(100) DEFAULT NULL COMMENT '学历',
  `degree` varchar(100) DEFAULT NULL COMMENT '学位',
  `research_direction` varchar(100) DEFAULT NULL COMMENT '研究方向',
  `job_title` varchar(100) DEFAULT NULL COMMENT '职称',
  `job` varchar(100) DEFAULT NULL COMMENT '职务',
  `instruct_nums` int(5) DEFAULT NULL COMMENT '指导人数',
  `instruct_major` varchar(100) DEFAULT NULL COMMENT '指导专业',
  `email` varchar(100) DEFAULT NULL COMMENT '邮箱',
  `phone` varchar(20) DEFAULT NULL COMMENT '手机',
  `qq` varchar(20) DEFAULT NULL COMMENT 'qq号',
  `we_chat` varchar(100) DEFAULT NULL COMMENT '微信',
  PRIMARY KEY (`tech_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='教师表';

-- ----------------------------
-- Table structure for thesis
-- ----------------------------
DROP TABLE IF EXISTS `thesis`;
CREATE TABLE `thesis` (
  `tid` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '论文id',
  `subject` varchar(100) DEFAULT NULL COMMENT '论文题目',
  `subtopic` varchar(100) DEFAULT NULL COMMENT '副题目',
  `keyword` varchar(100) DEFAULT NULL COMMENT '关键字',
  `type` varchar(30) DEFAULT NULL COMMENT '题目类型',
  `source` varchar(30) DEFAULT NULL COMMENT '题目来源',
  `workload` varchar(10) DEFAULT NULL COMMENT '工作量',
  `degree_difficulty` varchar(10) DEFAULT NULL COMMENT '难易度',
  `research_direc` varchar(100) DEFAULT NULL COMMENT '研究方向',
  `content` text COMMENT '内容',
  `update_uid` bigint(20) DEFAULT NULL COMMENT '创建人uid(最后修改的教师id)',
  `update_user` varchar(100) DEFAULT NULL COMMENT '创建人(最后修改的教师)',
  `update_time` datetime DEFAULT NULL COMMENT '创建时间(最后修改的时间)',
  `status` enum('0','1') DEFAULT NULL COMMENT '状态，0代表可选，1代表已经被选',
  PRIMARY KEY (`tid`)
) ENGINE=InnoDB AUTO_INCREMENT=1012 DEFAULT CHARSET=utf8 COMMENT='论文表';

-- ----------------------------
-- Table structure for thesis_file
-- ----------------------------
DROP TABLE IF EXISTS `thesis_file`;
CREATE TABLE `thesis_file` (
  `uid` bigint(20) NOT NULL COMMENT '学生uid',
  `tid` bigint(20) DEFAULT NULL COMMENT '论文id',
  `opening_report` varchar(200) DEFAULT NULL COMMENT '开题报告',
  `report_time` datetime DEFAULT NULL COMMENT '开题报告更新时间',
  `thesis` varchar(200) DEFAULT NULL COMMENT '毕业论文',
  `thesis_time` datetime DEFAULT NULL COMMENT '论文更新时间',
  PRIMARY KEY (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='论文文件表';

-- ----------------------------
-- Table structure for user_session
-- ----------------------------
DROP TABLE IF EXISTS `user_session`;
CREATE TABLE `user_session` (
  `uid` bigint(20) NOT NULL COMMENT '账号',
  `token` varchar(40) NOT NULL COMMENT '令牌',
  `role` enum('admin','teacher','student') NOT NULL COMMENT '角色',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户session表';

SET FOREIGN_KEY_CHECKS = 1;
