import { actionType } from './action'

export default function authUserReducer(authUser = null, action = {}) {
  switch (action.type) {
    case actionType.SET_AUTH_USER:
      return action.payload.user
    case actionType.UNSET_AUTH_USER:
      return null
    default:
      return authUser
  }
}
