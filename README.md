# Golang Upvote gRC API

Code repository for the project of an API meant to serve as an upvoting system

----------------------------------
## Context

Upvoting systems have been widely used in certain social media such as reddit, facebook and youtube
The idea of this project is to implement an upvoting API for a Forum, considering that any user can vote
on any of the threads.

Whenever a thread receives a new upvote/downvote, its author will receive a notification
Informing him that his thread has been voted by someone.

---------------------------------
## Project structure

The project has been divided into packages ending with the name ending with `_service`

So far, three services are being used:

    user_service: Perform operations on a collection where the users are stored
    thread_service: Perform operations on a collection where the threads and comments are stored
    upvote_service: Perform operations on a collection where the votes are being registered are stored
    
The  [gRPC-Gateway package](https://github.com/grpc-ecosystem/grpc-gateway) has been used in order to perform
Create and Retrieve operations on the user and thread services

So far, voting has been implemented for the threads.

--------------------------------
## Usage

There is the option of running the mongo database on a docker container, just by the command:

    docker-compose up -d

https://docs.docker.com/compose/install/

The main database will be started at the usual port: **27017**

----------------------------------
## Possible improvements

This project can be considered as the first version of an MVP, there are a few aspects
on which it could be improved:

- #### Provide de possibility of voting on comments
- #### Implementation of an authentication system
- #### Writing integration tests
- #### Take transactions into account for some operation with the database
