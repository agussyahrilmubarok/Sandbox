import { actionType } from './action'

export default function leaderboardsReducer(leaderboards = [], action = {}) {
  switch (action.type) {
    case actionType.RECEIVE_LEADERBOARDS:
      return action.payload.leaderboards
    case actionType.CLEAR_LEADERBOARDS:
      return []
    default:
      return leaderboards
  }
}
