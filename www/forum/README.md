# Forum

## Endpoint API

| No  | Endpoint                                                | Method | Description                     |
| --- | ------------------------------------------------------- | ------ | ------------------------------- |
| 1   | `/auth/sign-up`                                         | POST   | Register a new user             |
| 2   | `/auth/sign-in`                                         | POST   | User login and token generation |
| 3   | `/users`                                                | GET    | Retrieve all users              |
| 4   | `/users/me`                                             | GET    | Retrieve own profile            |
| 5   | `/threads`                                              | POST   | Create a new thread             |
| 6   | `/threads`                                              | GET    | Retrieve all threads            |
| 7   | `/threads/<id>`                                         | GET    | Retrieve a specific thread      |
| 8   | `/threads/<threadId>/comments`                          | POST   | Create a comment on a thread    |
| 9   | `/threads/<threadId>/up-vote`                           | POST   | Up-vote a thread                |
| 10  | `/threads/<threadId>/down-vote`                         | POST   | Down-vote a thread              |
| 11  | `/threads/<threadId>/neutral-vote`                      | POST   | Neutralize a thread vote        |
| 12  | `/threads/<threadId>/comments/<commentId>/up-vote`      | POST   | Up-vote a comment               |
| 13  | `/threads/<threadId>/comments/<commentId>/down-vote`    | POST   | Down-vote a comment             |
| 14  | `/threads/<threadId>/comments/<commentId>/neutral-vote` | POST   | Neutralize a comment vote       |
| 15  | `/leaderboards`                                         | GET    | See user leaderboards           |

---

[![App CI](https://forum.id/actions/workflows/app-ci.yml/badge.svg)](https://forum.id/actions/workflows/app-ci.yml)


## How to run
## References