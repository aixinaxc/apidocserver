/*
Navicat MySQL Data Transfer

Source Server         : 192.168.2.150
Source Server Version : 50624
Source Host           : 192.168.2.150:3306
Source Database       : apidoc

Target Server Type    : MYSQL
Target Server Version : 50624
File Encoding         : 65001

Date: 2018-10-19 04:32:35
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for apidoc_api
-- ----------------------------
DROP TABLE IF EXISTS `apidoc_api`;
CREATE TABLE `apidoc_api` (
  `api_id` varchar(64) NOT NULL COMMENT 'api 的id',
  `sort_id` varchar(64) DEFAULT '0' COMMENT '分类id',
  `project_id` varchar(64) DEFAULT '0' COMMENT '项目id',
  `api_name` varchar(255) DEFAULT NULL COMMENT 'api的名称',
  `api_edit_content` text COMMENT 'api的编辑内容',
  `api_show_content` text COMMENT 'api的展示内容',
  `api_state` int(1) DEFAULT '1' COMMENT '状态（1正常，2冻结，3删除）',
  `created_at` int(13) DEFAULT '0' COMMENT '创建时间',
  `updated_at` int(13) DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`api_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='api表';

-- ----------------------------
-- Records of apidoc_api
-- ----------------------------
INSERT INTO `apidoc_api` VALUES ('141869d51380344edfc33af9cb867417', '1', '1', '测试2', '### 请求类型： get\n\n### 请求路径： http://127.0.0.1/login\n\n### 请求参数：\n|字段名|字段类型|备注|\n|-|-|-|\n|test|varchar(255)|测试|\n|name|varchar(255)|名称|\n|age|int|年龄|\n\n### 返回数据：\n```\n{\n    \"code\":200,\n    \"msg\":\"OK\",\n    \"total\":5,\n    \"data\":[\n        {\n            \"project_id\":\"1\",\n            \"project_name\":\"测试项目\",\n            \"project_state\":1,\n            \"created_at\":0,\n            \"updated_at\":0\n        },\n        {\n            \"project_id\":\"2\",\n            \"project_name\":\"测试项目2\",\n            \"project_state\":1,\n            \"created_at\":0,\n            \"updated_at\":0\n        }\n    ]\n}\n```', '<h3><a id=\"_get_0\"></a>请求类型： get</h3>\n<h3><a id=\"_http127001login_2\"></a>请求路径： http://127.0.0.1/login</h3>\n<h3><a id=\"_4\"></a>请求参数：</h3>\n<table>\n<thead>\n<tr>\n<th>字段名</th>\n<th>字段类型</th>\n<th>备注</th>\n</tr>\n</thead>\n<tbody>\n<tr>\n<td>test</td>\n<td>varchar(255)</td>\n<td>测试</td>\n</tr>\n<tr>\n<td>name</td>\n<td>varchar(255)</td>\n<td>名称</td>\n</tr>\n<tr>\n<td>age</td>\n<td>int</td>\n<td>年龄</td>\n</tr>\n</tbody>\n</table>\n<h3><a id=\"_11\"></a>返回数据：</h3>\n<pre><code class=\"lang-\">{\n    &quot;code&quot;:200,\n    &quot;msg&quot;:&quot;OK&quot;,\n    &quot;total&quot;:5,\n    &quot;data&quot;:[\n        {\n            &quot;project_id&quot;:&quot;1&quot;,\n            &quot;project_name&quot;:&quot;测试项目&quot;,\n            &quot;project_state&quot;:1,\n            &quot;created_at&quot;:0,\n            &quot;updated_at&quot;:0\n        },\n        {\n            &quot;project_id&quot;:&quot;2&quot;,\n            &quot;project_name&quot;:&quot;测试项目2&quot;,\n            &quot;project_state&quot;:1,\n            &quot;created_at&quot;:0,\n            &quot;updated_at&quot;:0\n        }\n    ]\n}\n</code></pre>\n', '1', '1537249382', '0');
INSERT INTO `apidoc_api` VALUES ('440dbd346fd5ea44ff0ef74b45b1e3dd', '1', '1', '测试22', '### 请求类型： get\n\n### 请求路径： http://127.0.0.1/login\n\n### 请求参数：\n|字段名|字段类型|备注|\n|-|-|-|\n|test|varchar(255)|测试|\n|name|varchar(255)|名称|\n|age|int|年龄|\n\n### 返回数据：\n```\n{\n    \"code\":200,\n    \"msg\":\"OK\",\n    \"total\":5,\n    \"data\":[\n        {\n            \"project_id\":\"1\",\n            \"project_name\":\"测试项目\",\n            \"project_state\":1,\n            \"created_at\":0,\n            \"updated_at\":0\n        },\n        {\n            \"project_id\":\"2\",\n            \"project_name\":\"测试项目2\",\n            \"project_state\":1,\n            \"created_at\":0,\n            \"updated_at\":0\n        }\n    ]\n}\n```', '<h3><a id=\"_get_0\"></a>请求类型： get</h3>\n<h3><a id=\"_http127001login_2\"></a>请求路径： http://127.0.0.1/login</h3>\n<h3><a id=\"_4\"></a>请求参数：</h3>\n<table>\n<thead>\n<tr>\n<th>字段名</th>\n<th>字段类型</th>\n<th>备注</th>\n</tr>\n</thead>\n<tbody>\n<tr>\n<td>test</td>\n<td>varchar(255)</td>\n<td>测试</td>\n</tr>\n<tr>\n<td>name</td>\n<td>varchar(255)</td>\n<td>名称</td>\n</tr>\n<tr>\n<td>age</td>\n<td>int</td>\n<td>年龄</td>\n</tr>\n</tbody>\n</table>\n<h3><a id=\"_11\"></a>返回数据：</h3>\n<pre><code class=\"lang-\">{\n    &quot;code&quot;:200,\n    &quot;msg&quot;:&quot;OK&quot;,\n    &quot;total&quot;:5,\n    &quot;data&quot;:[\n        {\n            &quot;project_id&quot;:&quot;1&quot;,\n            &quot;project_name&quot;:&quot;测试项目&quot;,\n            &quot;project_state&quot;:1,\n            &quot;created_at&quot;:0,\n            &quot;updated_at&quot;:0\n        },\n        {\n            &quot;project_id&quot;:&quot;2&quot;,\n            &quot;project_name&quot;:&quot;测试项目2&quot;,\n            &quot;project_state&quot;:1,\n            &quot;created_at&quot;:0,\n            &quot;updated_at&quot;:0\n        }\n    ]\n}\n</code></pre>\n', '1', '1537496150', '0');
INSERT INTO `apidoc_api` VALUES ('58abea630967791da4e78665891beb7f', '1', '1', '测试44', '### 请求类型： get\n\n### 请求路径： http://127.0.0.1/login\n\n### 请求参数：\n|字段名|字段类型|备注|\n|-|-|-|\n|test|varchar(255)|测试|\n|name|varchar(255)|名称|\n|age|int|年龄|\n\n### 返回数据：\n```\n{\n    \"code\":200,\n    \"msg\":\"OK\",\n    \"total\":5,\n    \"data\":[\n        {\n            \"project_id\":\"1\",\n            \"project_name\":\"测试项目\",\n            \"project_state\":1,\n            \"created_at\":0,\n            \"updated_at\":0\n        },\n        {\n            \"project_id\":\"2\",\n            \"project_name\":\"测试项目2\",\n            \"project_state\":1,\n            \"created_at\":0,\n            \"updated_at\":0\n        }\n    ]\n}\n```', '<h3><a id=\"_get_0\"></a>请求类型： get</h3>\n<h3><a id=\"_http127001login_2\"></a>请求路径： http://127.0.0.1/login</h3>\n<h3><a id=\"_4\"></a>请求参数：</h3>\n<table>\n<thead>\n<tr>\n<th>字段名</th>\n<th>字段类型</th>\n<th>备注</th>\n</tr>\n</thead>\n<tbody>\n<tr>\n<td>test</td>\n<td>varchar(255)</td>\n<td>测试</td>\n</tr>\n<tr>\n<td>name</td>\n<td>varchar(255)</td>\n<td>名称</td>\n</tr>\n<tr>\n<td>age</td>\n<td>int</td>\n<td>年龄</td>\n</tr>\n</tbody>\n</table>\n<h3><a id=\"_11\"></a>返回数据：</h3>\n<pre><code class=\"lang-\">{\n    &quot;code&quot;:200,\n    &quot;msg&quot;:&quot;OK&quot;,\n    &quot;total&quot;:5,\n    &quot;data&quot;:[\n        {\n            &quot;project_id&quot;:&quot;1&quot;,\n            &quot;project_name&quot;:&quot;测试项目&quot;,\n            &quot;project_state&quot;:1,\n            &quot;created_at&quot;:0,\n            &quot;updated_at&quot;:0\n        },\n        {\n            &quot;project_id&quot;:&quot;2&quot;,\n            &quot;project_name&quot;:&quot;测试项目2&quot;,\n            &quot;project_state&quot;:1,\n            &quot;created_at&quot;:0,\n            &quot;updated_at&quot;:0\n        }\n    ]\n}\n</code></pre>\n', '1', '1537496175', '0');
INSERT INTO `apidoc_api` VALUES ('68140de37250b7b5b6cd550745cfd924', 'cae14b5d565f78a64787ed1437de2473', '2', '啦啦啦', '### 请求类型： post\n\n### 请求路径： http://127.0.0.1/login\n\n### 请求参数：\n|字段名|字段类型|备注|\n|-|-|-|\n|test|varchar(255)|测试|\n|name|varchar(255)|名称|\n|age|int|年龄|\n\n### 返回数据：\n```\n{\n    \"code\":200,\n    \"msg\":\"OK\",\n    \"total\":5,\n    \"data\":[\n        {\n            \"project_id\":\"1\",\n            \"project_name\":\"测试项目\",\n            \"project_state\":1,\n            \"created_at\":0,\n            \"updated_at\":0\n        },\n        {\n            \"project_id\":\"2\",\n            \"project_name\":\"测试项目2\",\n            \"project_state\":1,\n            \"created_at\":0,\n            \"updated_at\":0\n        }\n    ]\n}\n```', '<h3><a id=\"_post_0\"></a>请求类型： post</h3>\n<h3><a id=\"_http127001login_2\"></a>请求路径： http://127.0.0.1/login</h3>\n<h3><a id=\"_4\"></a>请求参数：</h3>\n<table>\n<thead>\n<tr>\n<th>字段名</th>\n<th>字段类型</th>\n<th>备注</th>\n</tr>\n</thead>\n<tbody>\n<tr>\n<td>test</td>\n<td>varchar(255)</td>\n<td>测试</td>\n</tr>\n<tr>\n<td>name</td>\n<td>varchar(255)</td>\n<td>名称</td>\n</tr>\n<tr>\n<td>age</td>\n<td>int</td>\n<td>年龄</td>\n</tr>\n</tbody>\n</table>\n<h3><a id=\"_11\"></a>返回数据：</h3>\n<pre><code class=\"lang-\">{\n    &quot;code&quot;:200,\n    &quot;msg&quot;:&quot;OK&quot;,\n    &quot;total&quot;:5,\n    &quot;data&quot;:[\n        {\n            &quot;project_id&quot;:&quot;1&quot;,\n            &quot;project_name&quot;:&quot;测试项目&quot;,\n            &quot;project_state&quot;:1,\n            &quot;created_at&quot;:0,\n            &quot;updated_at&quot;:0\n        },\n        {\n            &quot;project_id&quot;:&quot;2&quot;,\n            &quot;project_name&quot;:&quot;测试项目2&quot;,\n            &quot;project_state&quot;:1,\n            &quot;created_at&quot;:0,\n            &quot;updated_at&quot;:0\n        }\n    ]\n}\n</code></pre>\n', '1', '1537932225', '0');
INSERT INTO `apidoc_api` VALUES ('80bb88bb5b8e29f908c11ad8a0303f25', '1', '1', '测试55', '### 请求类型： get\n\n### 请求路径： http://127.0.0.1/login\n\n### 请求参数：\n|字段名|字段类型|备注|\n|-|-|-|\n|test|varchar(255)|测试|\n|name|varchar(255)|名称|\n|age|int|年龄|\n\n### 返回数据：\n```\n{\n    \"code\":200,\n    \"msg\":\"OK\",\n    \"total\":5,\n    \"data\":[\n        {\n            \"project_id\":\"1\",\n            \"project_name\":\"测试项目\",\n            \"project_state\":1,\n            \"created_at\":0,\n            \"updated_at\":0\n        },\n        {\n            \"project_id\":\"2\",\n            \"project_name\":\"测试项目2\",\n            \"project_state\":1,\n            \"created_at\":0,\n            \"updated_at\":0\n        }\n    ]\n}\n```', '<h3><a id=\"_get_0\"></a>请求类型： get</h3>\n<h3><a id=\"_http127001login_2\"></a>请求路径： http://127.0.0.1/login</h3>\n<h3><a id=\"_4\"></a>请求参数：</h3>\n<table>\n<thead>\n<tr>\n<th>字段名</th>\n<th>字段类型</th>\n<th>备注</th>\n</tr>\n</thead>\n<tbody>\n<tr>\n<td>test</td>\n<td>varchar(255)</td>\n<td>测试</td>\n</tr>\n<tr>\n<td>name</td>\n<td>varchar(255)</td>\n<td>名称</td>\n</tr>\n<tr>\n<td>age</td>\n<td>int</td>\n<td>年龄</td>\n</tr>\n</tbody>\n</table>\n<h3><a id=\"_11\"></a>返回数据：</h3>\n<pre><code class=\"lang-\">{\n    &quot;code&quot;:200,\n    &quot;msg&quot;:&quot;OK&quot;,\n    &quot;total&quot;:5,\n    &quot;data&quot;:[\n        {\n            &quot;project_id&quot;:&quot;1&quot;,\n            &quot;project_name&quot;:&quot;测试项目&quot;,\n            &quot;project_state&quot;:1,\n            &quot;created_at&quot;:0,\n            &quot;updated_at&quot;:0\n        },\n        {\n            &quot;project_id&quot;:&quot;2&quot;,\n            &quot;project_name&quot;:&quot;测试项目2&quot;,\n            &quot;project_state&quot;:1,\n            &quot;created_at&quot;:0,\n            &quot;updated_at&quot;:0\n        }\n    ]\n}\n</code></pre>\n', '1', '1537496190', '0');
INSERT INTO `apidoc_api` VALUES ('9daa355af458c8a126bab215a4a68d89', '1', '1', '测试11', '### 请求类型： get\n\n### 请求路径： http://127.0.0.1/login\n\n### 请求参数：\n|字段名|字段类型|备注|\n|-|-|-|\n|test|varchar(255)|测试|\n|name|varchar(255)|名称|\n|age|int|年龄|\n\n### 返回数据：\n```\n{\n    \"code\":200,\n    \"msg\":\"OK\",\n    \"total\":5,\n    \"data\":[\n        {\n            \"project_id\":\"1\",\n            \"project_name\":\"测试项目\",\n            \"project_state\":1,\n            \"created_at\":0,\n            \"updated_at\":0\n        },\n        {\n            \"project_id\":\"2\",\n            \"project_name\":\"测试项目2\",\n            \"project_state\":1,\n            \"created_at\":0,\n            \"updated_at\":0\n        }\n    ]\n}\n```', '<h3><a id=\"_get_0\"></a>请求类型： get</h3>\n<h3><a id=\"_http127001login_2\"></a>请求路径： http://127.0.0.1/login</h3>\n<h3><a id=\"_4\"></a>请求参数：</h3>\n<table>\n<thead>\n<tr>\n<th>字段名</th>\n<th>字段类型</th>\n<th>备注</th>\n</tr>\n</thead>\n<tbody>\n<tr>\n<td>test</td>\n<td>varchar(255)</td>\n<td>测试</td>\n</tr>\n<tr>\n<td>name</td>\n<td>varchar(255)</td>\n<td>名称</td>\n</tr>\n<tr>\n<td>age</td>\n<td>int</td>\n<td>年龄</td>\n</tr>\n</tbody>\n</table>\n<h3><a id=\"_11\"></a>返回数据：</h3>\n<pre><code class=\"lang-\">{\n    &quot;code&quot;:200,\n    &quot;msg&quot;:&quot;OK&quot;,\n    &quot;total&quot;:5,\n    &quot;data&quot;:[\n        {\n            &quot;project_id&quot;:&quot;1&quot;,\n            &quot;project_name&quot;:&quot;测试项目&quot;,\n            &quot;project_state&quot;:1,\n            &quot;created_at&quot;:0,\n            &quot;updated_at&quot;:0\n        },\n        {\n            &quot;project_id&quot;:&quot;2&quot;,\n            &quot;project_name&quot;:&quot;测试项目2&quot;,\n            &quot;project_state&quot;:1,\n            &quot;created_at&quot;:0,\n            &quot;updated_at&quot;:0\n        }\n    ]\n}\n</code></pre>\n', '1', '1537496137', '0');
INSERT INTO `apidoc_api` VALUES ('c8949f4688484f6ad6fc0b6dbaebeb7c', '1', '1', '测试33', '### 请求类型： get\n\n### 请求路径： http://127.0.0.1/login\n\n### 请求参数：\n|字段名|字段类型|备注|\n|-|-|-|\n|test|varchar(255)|测试|\n|name|varchar(255)|名称|\n|age|int|年龄|\n\n### 返回数据：\n```\n{\n    \"code\":200,\n    \"msg\":\"OK\",\n    \"total\":5,\n    \"data\":[\n        {\n            \"project_id\":\"1\",\n            \"project_name\":\"测试项目\",\n            \"project_state\":1,\n            \"created_at\":0,\n            \"updated_at\":0\n        },\n        {\n            \"project_id\":\"2\",\n            \"project_name\":\"测试项目2\",\n            \"project_state\":1,\n            \"created_at\":0,\n            \"updated_at\":0\n        }\n    ]\n}\n```', '<h3><a id=\"_get_0\"></a>请求类型： get</h3>\n<h3><a id=\"_http127001login_2\"></a>请求路径： http://127.0.0.1/login</h3>\n<h3><a id=\"_4\"></a>请求参数：</h3>\n<table>\n<thead>\n<tr>\n<th>字段名</th>\n<th>字段类型</th>\n<th>备注</th>\n</tr>\n</thead>\n<tbody>\n<tr>\n<td>test</td>\n<td>varchar(255)</td>\n<td>测试</td>\n</tr>\n<tr>\n<td>name</td>\n<td>varchar(255)</td>\n<td>名称</td>\n</tr>\n<tr>\n<td>age</td>\n<td>int</td>\n<td>年龄</td>\n</tr>\n</tbody>\n</table>\n<h3><a id=\"_11\"></a>返回数据：</h3>\n<pre><code class=\"lang-\">{\n    &quot;code&quot;:200,\n    &quot;msg&quot;:&quot;OK&quot;,\n    &quot;total&quot;:5,\n    &quot;data&quot;:[\n        {\n            &quot;project_id&quot;:&quot;1&quot;,\n            &quot;project_name&quot;:&quot;测试项目&quot;,\n            &quot;project_state&quot;:1,\n            &quot;created_at&quot;:0,\n            &quot;updated_at&quot;:0\n        },\n        {\n            &quot;project_id&quot;:&quot;2&quot;,\n            &quot;project_name&quot;:&quot;测试项目2&quot;,\n            &quot;project_state&quot;:1,\n            &quot;created_at&quot;:0,\n            &quot;updated_at&quot;:0\n        }\n    ]\n}\n</code></pre>\n', '1', '1537496162', '0');
INSERT INTO `apidoc_api` VALUES ('d440f8fc5830b7c2f2bbfb1f602fa40c', '1', '1', '测试', '### 请求类型： get\n\n### 请求路径： http://127.0.0.1/login\n\n### 请求参数：\n|字段名|字段类型|备注|\n|-|-|-|\n|test|varchar(255)|测试|\n|name|varchar(255)|名称|\n|age|int|年龄|\n\n### 返回数据：\n```\n{\n    \"code\":200,\n    \"msg\":\"OK\",\n    \"total\":5,\n    \"data\":[\n        {\n            \"project_id\":\"1\",\n            \"project_name\":\"测试项目\",\n            \"project_state\":1,\n            \"created_at\":0,\n            \"updated_at\":0\n        },\n        {\n            \"project_id\":\"2\",\n            \"project_name\":\"测试项目2\",\n            \"project_state\":1,\n            \"created_at\":0,\n            \"updated_at\":0\n        }\n    ]\n}\n```', '<h3><a id=\"_get_0\"></a>请求类型： get</h3>\n<h3><a id=\"_http127001login_2\"></a>请求路径： http://127.0.0.1/login</h3>\n<h3><a id=\"_4\"></a>请求参数：</h3>\n<table>\n<thead>\n<tr>\n<th>字段名</th>\n<th>字段类型</th>\n<th>备注</th>\n</tr>\n</thead>\n<tbody>\n<tr>\n<td>test</td>\n<td>varchar(255)</td>\n<td>测试</td>\n</tr>\n<tr>\n<td>name</td>\n<td>varchar(255)</td>\n<td>名称</td>\n</tr>\n<tr>\n<td>age</td>\n<td>int</td>\n<td>年龄</td>\n</tr>\n</tbody>\n</table>\n<h3><a id=\"_11\"></a>返回数据：</h3>\n<pre><code class=\"lang-\">{\n    &quot;code&quot;:200,\n    &quot;msg&quot;:&quot;OK&quot;,\n    &quot;total&quot;:5,\n    &quot;data&quot;:[\n        {\n            &quot;project_id&quot;:&quot;1&quot;,\n            &quot;project_name&quot;:&quot;测试项目&quot;,\n            &quot;project_state&quot;:1,\n            &quot;created_at&quot;:0,\n            &quot;updated_at&quot;:0\n        },\n        {\n            &quot;project_id&quot;:&quot;2&quot;,\n            &quot;project_name&quot;:&quot;测试项目2&quot;,\n            &quot;project_state&quot;:1,\n            &quot;created_at&quot;:0,\n            &quot;updated_at&quot;:0\n        }\n    ]\n}\n</code></pre>\n', '1', '1537249335', '0');
INSERT INTO `apidoc_api` VALUES ('e6a1849c79c7c330b0bd3ef473636fcc', '253d8c0b00f990b7548328b2de559a7b', '2', '测试1', '### 请求类型： get\n\n### 请求路径： http://127.0.0.1/login\n\n### 请求参数：\n|字段名|字段类型|备注|\n|-|-|-|\n|test|varchar(255)|测试|\n|name|varchar(255)|名称|\n|age|int|年龄|\n\n### 返回数据：\n```\n{\n    \"code\":200,\n    \"msg\":\"OK\",\n    \"total\":5,\n    \"data\":[\n        {\n            \"project_id\":\"1\",\n            \"project_name\":\"测试项目\",\n            \"project_state\":1,\n            \"created_at\":0,\n            \"updated_at\":0\n        },\n        {\n            \"project_id\":\"2\",\n            \"project_name\":\"测试项目2\",\n            \"project_state\":1,\n            \"created_at\":0,\n            \"updated_at\":0\n        }\n    ]\n}\n```', '<h3><a id=\"_get_0\"></a>请求类型： get</h3>\n<h3><a id=\"_http127001login_2\"></a>请求路径： http://127.0.0.1/login</h3>\n<h3><a id=\"_4\"></a>请求参数：</h3>\n<table>\n<thead>\n<tr>\n<th>字段名</th>\n<th>字段类型</th>\n<th>备注</th>\n</tr>\n</thead>\n<tbody>\n<tr>\n<td>test</td>\n<td>varchar(255)</td>\n<td>测试</td>\n</tr>\n<tr>\n<td>name</td>\n<td>varchar(255)</td>\n<td>名称</td>\n</tr>\n<tr>\n<td>age</td>\n<td>int</td>\n<td>年龄</td>\n</tr>\n</tbody>\n</table>\n<h3><a id=\"_11\"></a>返回数据：</h3>\n<pre><code class=\"lang-\">{\n    &quot;code&quot;:200,\n    &quot;msg&quot;:&quot;OK&quot;,\n    &quot;total&quot;:5,\n    &quot;data&quot;:[\n        {\n            &quot;project_id&quot;:&quot;1&quot;,\n            &quot;project_name&quot;:&quot;测试项目&quot;,\n            &quot;project_state&quot;:1,\n            &quot;created_at&quot;:0,\n            &quot;updated_at&quot;:0\n        },\n        {\n            &quot;project_id&quot;:&quot;2&quot;,\n            &quot;project_name&quot;:&quot;测试项目2&quot;,\n            &quot;project_state&quot;:1,\n            &quot;created_at&quot;:0,\n            &quot;updated_at&quot;:0\n        }\n    ]\n}\n</code></pre>\n', '1', '1537853247', '0');

-- ----------------------------
-- Table structure for apidoc_project
-- ----------------------------
DROP TABLE IF EXISTS `apidoc_project`;
CREATE TABLE `apidoc_project` (
  `project_id` varchar(64) NOT NULL COMMENT '项目id',
  `project_name` varchar(255) DEFAULT NULL COMMENT '项目名称',
  `project_state` int(1) DEFAULT '1' COMMENT '状态（1正常，2冻结，3删除）',
  `created_at` int(13) DEFAULT '0' COMMENT '创建时间',
  `updated_at` int(13) DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`project_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='项目表';

-- ----------------------------
-- Records of apidoc_project
-- ----------------------------
INSERT INTO `apidoc_project` VALUES ('2', '测试项目2', '1', '0', '0');
INSERT INTO `apidoc_project` VALUES ('3d92f7981f0ad3892524391f5388c7cf', '测试项目7', '1', '1536715489', '0');
INSERT INTO `apidoc_project` VALUES ('40d647ea3d13de2492909d2713b0a849', '测试项目6', '1', '1536715273', '0');
INSERT INTO `apidoc_project` VALUES ('4495954a3c80f1b670c3faba7f948b15', '大家好大家好大家好大家好大家好', '1', '1539066759', '0');
INSERT INTO `apidoc_project` VALUES ('b743b68c6dbb6824ba9e19a5592a324d', '测试数据5', '1', '1536310398', '0');
INSERT INTO `apidoc_project` VALUES ('c14611b7fb00e43fd4013d8b3e4d7de6', '测试数据4', '1', '1536310166', '0');
INSERT INTO `apidoc_project` VALUES ('d4fc216036f928bb307cde968f6c899e', '测试项目x', '1', '1536310019', '1536889646');

-- ----------------------------
-- Table structure for apidoc_sort
-- ----------------------------
DROP TABLE IF EXISTS `apidoc_sort`;
CREATE TABLE `apidoc_sort` (
  `sort_id` varchar(64) NOT NULL COMMENT '分类id',
  `sort_top_id` varchar(64) DEFAULT '0' COMMENT '分类最上级id （默认0）',
  `sort_pid` varchar(64) DEFAULT '0' COMMENT '分类上级id （默认0）',
  `project_id` varchar(64) DEFAULT '0' COMMENT '项目id',
  `sort_name` varchar(255) DEFAULT '无' COMMENT '分类名称',
  `sort_seq` int(11) DEFAULT '1000' COMMENT '分类排序',
  `status` int(1) DEFAULT '1' COMMENT '状态（1正常，2冻结，3删除）',
  `created_at` int(13) DEFAULT '0' COMMENT '创建时间',
  `updated_at` int(13) DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`sort_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='分类表';

-- ----------------------------
-- Records of apidoc_sort
-- ----------------------------
INSERT INTO `apidoc_sort` VALUES ('1', '0', '0', '1', '测试管理', '1000', '1', '0', '0');
INSERT INTO `apidoc_sort` VALUES ('2', '0', '0', '1', '其他管理', '1000', '1', '0', '0');
INSERT INTO `apidoc_sort` VALUES ('253d8c0b00f990b7548328b2de559a7b', '', '', '2', '测试分类', '0', '1', '1536903977', '0');
INSERT INTO `apidoc_sort` VALUES ('3bb85d952bb5a14c1faab9fbcd9e02f3', '', '', '1', '项目管理', '0', '1', '1536715863', '0');
INSERT INTO `apidoc_sort` VALUES ('5cc83eff13f87c3cf38df3d1398f6d72', '', '', '1', '测试1', '0', '1', '1537496094', '0');
INSERT INTO `apidoc_sort` VALUES ('7458b6cdbabe5f52a47f157884bdacbc', '', '', '1', '测试3', '0', '1', '1537496107', '0');
INSERT INTO `apidoc_sort` VALUES ('86a74b14bf9a59da22a85928bc29e466', '', '', '2', '大家好大家好大家好', '0', '1', '1537854653', '0');
INSERT INTO `apidoc_sort` VALUES ('cae14b5d565f78a64787ed1437de2473', '', '', '2', '测试', '0', '1', '1537853098', '0');
INSERT INTO `apidoc_sort` VALUES ('ec1ee2fbb0d42339d550936024706213', '', '', '1', '测试2', '0', '1', '1537496101', '0');
INSERT INTO `apidoc_sort` VALUES ('ffdf64719cc14afc05468abb7c4ae5a5', '', '', 'b743b68c6dbb6824ba9e19a5592a324d', '新填分类', '0', '1', '1537868098', '0');

-- ----------------------------
-- Table structure for apidoc_user
-- ----------------------------
DROP TABLE IF EXISTS `apidoc_user`;
CREATE TABLE `apidoc_user` (
  `user_id` varchar(64) NOT NULL COMMENT '用户id',
  `user_username` varchar(255) DEFAULT NULL COMMENT '用户名',
  `user_password` varchar(255) DEFAULT NULL COMMENT '密码',
  `created_at` int(13) DEFAULT '0' COMMENT '创建时间',
  `updated_at` int(13) DEFAULT '0' COMMENT '修改时间',
  `user_state` int(1) DEFAULT '1' COMMENT '状态（1正常，2冻结，3删除）',
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

-- ----------------------------
-- Records of apidoc_user
-- ----------------------------
INSERT INTO `apidoc_user` VALUES ('1', 'admin', '1ee40133a1d9cb7c52ae50845558c3f6', '0', '1536628262', '1');
INSERT INTO `apidoc_user` VALUES ('3ecd7ea7f3c8a77f9eb8d1c71003dd9e', 'kslamp', '1ee40133a1d9cb7c52ae50845558c3f6', '1536712161', '0', '1');
INSERT INTO `apidoc_user` VALUES ('97130659c466682c14e692c3db4e555d', 'hades', '1ee40133a1d9cb7c52ae50845558c3f6', '1536731453', '0', '1');
INSERT INTO `apidoc_user` VALUES ('fa562196c8c659bd3dd30dede7fc68cf', 'test', '1ee40133a1d9cb7c52ae50845558c3f6', '1536630488', '0', '1');

-- ----------------------------
-- Table structure for apidoc_user_project
-- ----------------------------
DROP TABLE IF EXISTS `apidoc_user_project`;
CREATE TABLE `apidoc_user_project` (
  `user_id` varchar(64) DEFAULT '0' COMMENT '用户id',
  `project_id` varchar(64) DEFAULT '0' COMMENT '项目id'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户项目表';

-- ----------------------------
-- Records of apidoc_user_project
-- ----------------------------
INSERT INTO `apidoc_user_project` VALUES ('1', '1');
INSERT INTO `apidoc_user_project` VALUES ('1', '2');
INSERT INTO `apidoc_user_project` VALUES ('1', 'd4fc216036f928bb307cde968f6c899e');
INSERT INTO `apidoc_user_project` VALUES ('1', 'c14611b7fb00e43fd4013d8b3e4d7de6');
INSERT INTO `apidoc_user_project` VALUES ('1', 'b743b68c6dbb6824ba9e19a5592a324d');
INSERT INTO `apidoc_user_project` VALUES ('fa562196c8c659bd3dd30dede7fc68cf', '1');
INSERT INTO `apidoc_user_project` VALUES ('fa562196c8c659bd3dd30dede7fc68cf', '2');
INSERT INTO `apidoc_user_project` VALUES ('fa562196c8c659bd3dd30dede7fc68cf', 'd4fc216036f928bb307cde968f6c899e');
INSERT INTO `apidoc_user_project` VALUES ('1', '40d647ea3d13de2492909d2713b0a849');
INSERT INTO `apidoc_user_project` VALUES ('1', '3d92f7981f0ad3892524391f5388c7cf');
INSERT INTO `apidoc_user_project` VALUES ('1', '058e9375c2e893bb47bab15320e9640a');
INSERT INTO `apidoc_user_project` VALUES ('fa562196c8c659bd3dd30dede7fc68cf', '4495954a3c80f1b670c3faba7f948b15');

