import { api } from '../../lib/api'
import toast from 'react-hot-toast'
import { showLoading, hideLoading } from 'react-redux-loading-bar'

const actionType = {
  RECEIVE_USERS: 'RECEIVE_USERS'
}

function receiveUsers(users) {
  return {
    type: actionType.RECEIVE_USERS,
    payload: {
      users
    }
  }
}

function asyncRegisterUser({ name, email, password }) {
  return async (dispatch) => {
    dispatch(showLoading())
    try {
      await api.register({ name, email, password })
      toast.success('successfully registered')
      dispatch(hideLoading())
      return { error: false }
    } catch (error) {
      toast.error(error.response.data.message)
      dispatch(hideLoading())
      return { error: true }
    }
  }
}

export { actionType, receiveUsers, asyncRegisterUser }
