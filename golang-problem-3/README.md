# Golang-Problem-3

### *Description*

You are to implement *Bale* interface.\
\
The `AddUser` function create new user. It has two
arguments. The first one is the username of the new user,
and the second one is the boolean value shows the type of
the user. username is unique, and it should have length larger than 3, and
it should contain both letters and digits. If one of these roles 
violated, This function returns an error with a message: **invalid username** .\
In successful situation, it returns an id of the user. Ids of users are
begins from 1 and increase by one in each new creation.
\
\
The `AddChat` function create new chat. It has four
arguments. The first one is the chat name. The second one is
the type of the chat. The third one is the id of the creator
of the chat. The last one is the array of ids that represents
the admins of the chat. The creator of the chat cannot be bot, and
in this situation the function returns an error with message: **could not create chat** .
\
In successful situation, it returns an id of the created chat. It works just like the ids
for the users. \
\
The `SendMessage` function sends message to the specific chat.
As the first parameter, it gets the id of the user who wants to send message.
The second parameter is the id of the chat that the user wants to send message to.
Finally, the last parameter is the text of the message.
In the chats with type channels, only admin users could send message, and
if a non admin user tries to send message, the function returns an error with
the message: **user could not send message** .

The `SendLike` function likes a specific message.
each user can like each message only once. If this situation 
occurs, the function returns an error with the message: **this user has liked this message before** .
\
If the message doesn't exist, the function returns an error with the message: **message not found** .
\
\
The `GetNumberOfLikes` function returns the number of likes
of the message.
\
\
The `SetChatAdmin` function sets the specific user the admin of
the specific chat. If the user is already admin of that chat,
the function returns an error with the message: **user is already admin** .\
\
The `GetLastMessage` function returns the text and the id of the last
message that sends to the chat.

The `GetLastUserMessage` function returns the text and the id of the last
message that the user sends.