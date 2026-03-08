import { actionType } from './action'

export default function threadDetailReducer(threadDetail = null, action = {}) {
  switch (action.type) {
    case actionType.RECEIVE_THREAD_DETAIL:
      return action.payload.threadDetail
    case actionType.CLEAR_THREAD_DETAIL:
      return null
    case actionType.UPVOTE_THREAD_DETAIL:
      return {
        ...threadDetail,
        upVotesBy: [...threadDetail.upVotesBy, action.payload.userId],
        downVotesBy: threadDetail.downVotesBy.filter((id) => id !== action.payload.userId)
      }
    case actionType.DOWNVOTE_THREAD_DETAIL:
      return {
        ...threadDetail,
        downVotesBy: [...threadDetail.downVotesBy, action.payload.userId],
        upVotesBy: threadDetail.upVotesBy.filter((id) => id !== action.payload.userId)
      }
    case actionType.UNVOTE_THREAD_DETAIL:
      return {
        ...threadDetail,
        upVotesBy: threadDetail.upVotesBy.filter((id) => id !== action.payload.userId),
        downVotesBy: threadDetail.downVotesBy.filter((id) => id !== action.payload.userId)
      }
    case actionType.ADD_THREAD_DETAIL_COMMENT:
      return {
        ...threadDetail,
        comments: [...threadDetail.comments, action.payload.comment]
      }
    case actionType.UPVOTE_THREAD_DETAIL_COMMENT:
      return {
        ...threadDetail,
        comments: threadDetail.comments.map((comment) => {
          if (comment.id === action.payload.commentId) {
            return {
              ...comment,
              upVotesBy: [...comment.upVotesBy, action.payload.userId],
              downVotesBy: comment.downVotesBy.filter((id) => id !== action.payload.userId)
            }
          }
          return comment
        })
      }
    case actionType.DOWNVOTE_THREAD_DETAIL_COMMENT:
      return {
        ...threadDetail,
        comments: threadDetail.comments.map((comment) => {
          if (comment.id === action.payload.commentId) {
            return {
              ...comment,
              downVotesBy: [...comment.downVotesBy, action.payload.userId],
              upVotesBy: comment.upVotesBy.filter((id) => id !== action.payload.userId)
            }
          }
          return comment
        })
      }
    case actionType.UNVOTE_THREAD_DETAIL_COMMENT:
      return {
        ...threadDetail,
        comments: threadDetail.comments.map((comment) => {
          if (comment.id === action.payload.commentId) {
            return {
              ...comment,
              upVotesBy: comment.upVotesBy.filter((id) => id !== action.payload.userId),
              downVotesBy: comment.downVotesBy.filter((id) => id !== action.payload.userId)
            }
          }
          return comment
        })
      }
    default:
      return threadDetail
  }
}
