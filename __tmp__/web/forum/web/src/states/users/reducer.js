import { actionType } from './action'

export default function usersReducer(users = [], action = {}) {
  switch (action.type) {
    case actionType.RECEIVE_USERS:
      return action.payload.users
    default:
      return users
  }
}
