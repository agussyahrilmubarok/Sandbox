import { actionType } from './action'

export default function threadsReducer(threads = [], action = {}) {
  switch (action.type) {
    case actionType.RECEIVE_THREADS:
      return action.payload.threads
    case actionType.ADD_THREAD:
      return [action.payload.thread, ...threads]
    case actionType.UPVOTE_THREAD:
      return threads.map((thread) => {
        if (thread.id === action.payload.threadId) {
          return {
            ...thread,
            upVotesBy: [...thread.upVotesBy, action.payload.userId],
            downVotesBy: thread.downVotesBy.filter((id) => id !== action.payload.userId)
          }
        }
        return thread
      })
    case actionType.DOWNVOTE_THREAD:
      return threads.map((thread) => {
        if (thread.id === action.payload.threadId) {
          return {
            ...thread,
            downVotesBy: [...thread.downVotesBy, action.payload.userId],
            upVotesBy: thread.upVotesBy.filter((id) => id !== action.payload.userId)
          }
        }
        return thread
      })
    case actionType.UNVOTE_THREAD:
      return threads.map((thread) => {
        if (thread.id === action.payload.threadId) {
          return {
            ...thread,
            upVotesBy: thread.upVotesBy.filter((id) => id !== action.payload.userId),
            downVotesBy: thread.downVotesBy.filter((id) => id !== action.payload.userId)
          }
        }
        return thread
      })
    default:
      return threads
  }
}
