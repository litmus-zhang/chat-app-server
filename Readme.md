# Chat application

**Functional Requirements**
User authentication
Chat between two users



**API**
- Authentication 
  - Login : username, password =>token 
  - Sign up : email, full name, password => stored user
- Sending messages
    - Chat: Content, senderId, receiverId

### Class Model

**User :**

- Id : uuid
- username: string 
- Password : string 
- Chats : Messages[]

**Messages:**
- Id : ulid 
- Message: string 
- SenderId: references user(Id)
- ReceiverId: references user(id)
- TimeofMessage: timestamp

  - **N.B** : 
  For the Message Id, ulid is preferred as it will make sorting of messages easier compared to using timestamp. For example sorting the messages in a chat, when the time difference between two messages is in the microseconds range, ulid makes it easier to sort.
For the Chats, this is gotten through a sql query shown below,

```
SELECT *
FROM message
WHERE (senderId = <user1_id> AND receiverId = <user2_id>)
   OR (senderId = <user2_id> AND receiverId = <user1_id>)
ORDER BY id;
```




**Frontend**

Main Components 
- Message card 
  - Message content 
  - Time
- User profile 
  - Full name 
  - Last time for sent message*


