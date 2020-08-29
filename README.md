# Rating Service

Clone this project
- `git clone https://github.com/abhinavmaury/ticket-service.git`

Import in IDE and run-
It will run on port 9000

This service is using MySQL-
- create user_ticket table-

``CREATE TABLE `user_ticket` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` varchar(45) NOT NULL,
  `phone` varchar(45) NOT NULL,
  `query` varchar(1000) NOT NULL,
  `status` enum('created','in_progress','completed','rejected') NOT NULL DEFAULT 'created',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`));``

- create user_ticket table-

``
CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_name` varchar(45) NOT NULL,
  `user_type` varchar(45) DEFAULT 'agent',
  `phone` varchar(10) NOT NULL,
  `is_online` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`));
``

It contains 2 REST POST APIs

- To generate Ticket -

**Request URL-  {{url}}/api/customer/ticket/generate  
Request Body- 
    {
        "user_id": "12345",
        "phone": "7654345123",
        "query": "hey there! I have a query"
    }**
    
- For User Login -

**Request URL-  {{url}}/api/agent/login  
 Request Body- 
     {
         "username": "abhinav",
         "password": "12345",
     }** 