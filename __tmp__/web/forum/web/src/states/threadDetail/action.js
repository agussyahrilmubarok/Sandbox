import { api } from '../../lib/api'
import toast from 'react-hot-toast'
import { showLoading, hideLoading } from 'react-redux-loading-bar'

const actionType = {
  RECEIVE_THREAD_DETAIL: 'RECEIVE_THREAD_DETAIL',
  CLEAR_THREAD_DETAIL: 'CLEAR_THREAD_DETAIL',
  UPVOTE_THREAD_DETAIL: 'UPVOTE_THREAD_DETAIL',
  DOWNVOTE_THREAD_DETAIL: 'DOWNVOTE_THREAD_DETAIL',
  UNVOTE_THREAD_DETAIL: 'UNVOTE_THREAD_DETAIL',
  ADD_THREAD_DETAIL_COMMENT: 'ADD_THREAD_DETAIL_COMMENT',
  UPVOTE_THREAD_DETAIL_COMMENT: 'UPVOTE_THREAD_DETAIL_COMMENT',
  DOWNVOTE_THREAD_DETAIL_COMMENT: 'DOWNVOTE_THREAD_DETAIL_COMMENT',
  UNVOTE_THREAD_DETAIL_COMMENT: 'UNVOTE_THREAD_DETAIL_COMMENT'
}

const voteType = {
  UPVOTE: 1,
  NEUTRAL: 0,
  DOWNVOTE: -1
}

function receiveThreadDetail(threadDetail) {
  return {
    type: actionType.RECEIVE_THREAD_DETAIL,
    payload: {
      threadDetail
    }
  }
}

function clearThreadDetail() {
  return {
    type: actionType.CLEAR_THREAD_DETAIL
  }
}

function upVoteThreadDetail(userId) {
  return {
    type: actionType.UPVOTE_THREAD_DETAIL,
    payload: {
      userId
    }
  }
}

function unVoteThreadDetail(userId) {
  return {
    type: actionType.UNVOTE_THREAD_DETAIL,
    payload: {
      userId
    }
  }
}

function downVoteThreadDetail(userId) {
  return {
    type: actionType.DOWNVOTE_THREAD_DETAIL,
    payload: {
      userId
    }
  }
}

function addThreadDetailComment(comment) {
  return {
    type: actionType.ADD_THREAD_DETAIL_COMMENT,
    payload: {
      comment
    }
  }
}

function upVoteThreadDetailComment({ commentId, userId }) {
  return {
    type: actionType.UPVOTE_THREAD_DETAIL_COMMENT,
    payload: {
      commentId,
      userId
    }
  }
}

function downVoteThreadDetailComment({ commentId, userId }) {
  return {
    type: actionType.DOWNVOTE_THREAD_DETAIL_COMMENT,
    payload: {
      commentId,
      userId
    }
  }
}

function unVoteThreadDetailComment({ commentId, userId }) {
  return {
    type: actionType.UNVOTE_THREAD_DETAIL_COMMENT,
    payload: {
      commentId,
      userId
    }
  }
}

function asyncReceiveThreadDetail(threadId) {
  return async (dispatch) => {
    dispatch(showLoading())
    try {
      const threadDetail = await api.getThreadDetail(threadId)
      dispatch(receiveThreadDetail(threadDetail))
    } catch (error) {
      if (error.response.status === 404) {
        toast.error('Thread not found')
      } else {
        toast.error(error.message)
      }
    }
    dispatch(hideLoading())
  }
}

function asyncUpVoteThreadDetail(threadId) {
  return async (dispatch, getState) => {
    const { authUser } = getState()
    if (!authUser) {
      toast.error('You need to login first')
      return
    }
    dispatch(showLoading())
    dispatch(upVoteThreadDetail(authUser.id))
    try {
      await api.upVoteThread(threadId)
    } catch (error) {
      toast.error(error.message)
      dispatch(unVoteThreadDetail({ userId: authUser.id, previousVoteType: voteType.UPVOTE }))
    }
    dispatch(hideLoading())
  }
}

function asyncDownVoteThreadDetail(threadId) {
  return async (dispatch, getState) => {
    const { authUser } = getState()
    if (!authUser) {
      toast.error('You need to login first')
      return
    }
    dispatch(showLoading())
    dispatch(downVoteThreadDetail(authUser.id))
    try {
      await api.downVoteThread(threadId)
    } catch (error) {
      toast.error(error.message)
      dispatch(unVoteThreadDetail({ userId: authUser.id, previousVoteType: voteType.DOWNVOTE }))
    }
    dispatch(hideLoading())
  }
}

function asyncUnVoteThreadDetail({ threadId, previousVoteType }) {
  return async (dispatch, getState) => {
    const { authUser } = getState()
    dispatch(showLoading())
    dispatch(unVoteThreadDetail(authUser.id))
    try {
      await api.unvoteThread(threadId)
    } catch (error) {
      toast.error(error.message)
      if (previousVoteType === voteType.UPVOTE) {
        dispatch(upVoteThreadDetail(authUser.id))
      }
      if (previousVoteType === voteType.DOWNVOTE) {
        dispatch(downVoteThreadDetail(authUser.id))
      }
    }
    dispatch(hideLoading())
  }
}

function asyncAddThreadDetailComment({ threadId, content }) {
  return async (dispatch) => {
    dispatch(showLoading())
    try {
      const newComment = await api.createComment({ threadId, content })
      dispatch(addThreadDetailComment(newComment))
      toast.success('Comment added successfully')
    } catch (error) {
      toast.error(error.message)
    }
    dispatch(hideLoading())
  }
}

function asyncUpVoteThreadDetailComment({ threadId, commentId }) {
  return async (dispatch, getState) => {
    const { authUser } = getState()
    if (!authUser) {
      toast.error('You need to login first')
      return
    }
    dispatch(showLoading())
    dispatch(upVoteThreadDetailComment({ commentId, userId: authUser.id }))
    try {
      await api.upVoteComment(threadId, commentId)
    } catch (error) {
      toast.error(error.message)
      dispatch(unVoteThreadDetailComment({ commentId, userId: authUser.id }))
    }
    dispatch(hideLoading())
  }
}

function asyncDownVoteThreadDetailComment({ threadId, commentId }) {
  return async (dispatch, getState) => {
    const { authUser } = getState()
    if (!authUser) {
      toast.error('You need to login first')
      return
    }
    dispatch(showLoading())
    dispatch(downVoteThreadDetailComment({ commentId, userId: authUser.id }))
    try {
      await api.downVoteComment(threadId, commentId)
    } catch (error) {
      toast.error(error.message)
      dispatch(unVoteThreadDetailComment({ commentId, userId: authUser.id }))
    }
    dispatch(hideLoading())
  }
}

function asyncUnVoteThreadDetailComment({ threadId, commentId, previousVoteType }) {
  return async (dispatch, getState) => {
    const { authUser } = getState()
    dispatch(showLoading())
    dispatch(unVoteThreadDetailComment({ commentId, userId: authUser.id }))
    try {
      await api.unvoteComment(threadId, commentId)
    } catch (error) {
      toast.error(error.message)
      if (previousVoteType === voteType.UPVOTE) {
        dispatch(upVoteThreadDetailComment({ commentId, userId: authUser.id }))
      }
      if (previousVoteType === voteType.DOWNVOTE) {
        dispatch(downVoteThreadDetailComment({ commentId, userId: authUser.id }))
      }
    }
    dispatch(hideLoading())
  }
}

export {
  actionType,
  receiveThreadDetail,
  addThreadDetailComment,
  upVoteThreadDetailComment,
  downVoteThreadDetailComment,
  unVoteThreadDetailComment,
  asyncReceiveThreadDetail,
  asyncAddThreadDetailComment,
  asyncUpVoteThreadDetailComment,
  asyncDownVoteThreadDetailComment,
  asyncUnVoteThreadDetailComment,
  clearThreadDetail,
  upVoteThreadDetail,
  downVoteThreadDetail,
  unVoteThreadDetail,
  asyncUpVoteThreadDetail,
  asyncDownVoteThreadDetail,
  asyncUnVoteThreadDetail
}
