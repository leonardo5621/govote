# Golang Upvote gRC API

Code repository for the project of an API meant to serve as an upvoting system

----------------------------------
## Context

Upvoting systems have been widely used in certain social media such as reddit, facebook and youtube.
The idea behind this project is to implement an upvoting API for Forum Threads, considering that any user can vote
on any thread.

Whenever a thread receives a new upvote/downvote, its author will receive a notification
Informing him that his thread has been voted by someone. Something similar occurs in some social platforms,
where one receives an email or notification after there was a reaction to one of its posts or comments.

---------------------------------
## Project structure

The project has been divided into packages ending with the name ending with `_service`

So far, three services are being used:

    user_service: Perform operations on a collection where the users are stored
    thread_service: Perform operations on a collection where the threads and comments are stored
    upvote_service: Perform operations on a collection where the votes are being registered are stored
    
The  [gRPC-Gateway package](https://github.com/grpc-ecosystem/grpc-gateway) has been used in order to perform some
CRUD like operations. In this project, it has been used to execute Create and Retrieve operations on the user and thread services

It is possible to upvote/downvote on the threads of the forum, specifying the threadId and the userId of the voter.

--------------------------------
## Usage

There is the option of running the mongo database on a docker container, just by the command:

    docker-compose up -d

https://docs.docker.com/compose/install/

The main database will be started at the usual port: **27017**

The server and client pair are in their respective folders in the repository

The server will run on the port **5005**, whereas the gateway will be set on the port **8081** in the localhost

The following routes have been registered to the gateway:

  - `/user` (create and get user)
  - `/thread` (create and get thread)
  - `/thread/comment` (create comment)

The voting endpoint consists of a bi-directional streaming route. Where the client is able to send the
votes to be registered and the server sends the notifications back.



----------------------------------
## Possible improvements

This project can be considered as the first version of a MVP, there are a few aspects
on which it could be improved:

- #### Provide de possibility of voting on comments
- #### Implementation of an authentication system for the voting
- #### Writing integration tests
- #### Take transactions into account for some operation with the database
