import { api } from '../../lib/api'
import toast from 'react-hot-toast'
import { showLoading, hideLoading } from 'react-redux-loading-bar'

const actionType = {
  RECEIVE_LEADERBOARDS: 'RECEIVE_LEADERBOARDS',
  CLEAR_LEADERBOARDS: 'CLEAR_LEADERBOARDS'
}

function receiveLeaderboards(leaderboards) {
  return {
    type: actionType.RECEIVE_LEADERBOARDS,
    payload: {
      leaderboards
    }
  }
}

function clearLeaderboards() {
  return {
    type: actionType.CLEAR_LEADERBOARDS
  }
}

function asyncReceiveLeaderboards() {
  return async (dispatch) => {
    dispatch(showLoading())
    try {
      const leaderboards = await api.getLeaderboards()
      dispatch(receiveLeaderboards(leaderboards))
    } catch (error) {
      toast.error(error.message)
    }
    dispatch(hideLoading())
  }
}

export { actionType, receiveLeaderboards, clearLeaderboards, asyncReceiveLeaderboards }
