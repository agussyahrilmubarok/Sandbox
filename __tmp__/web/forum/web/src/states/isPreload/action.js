import { api } from '../../lib/api'
import { setAuthUser } from '../authUser/action'
import { showLoading, hideLoading } from 'react-redux-loading-bar'

const actionType = {
  SET_IS_PRELOAD: 'SET_IS_PRELOAD'
}

function setIsPreload(isPreload) {
  return {
    type: actionType.SET_IS_PRELOAD,
    payload: {
      isPreload
    }
  }
}

function asyncPreloadProcess() {
  return async (dispatch) => {
    dispatch(showLoading())
    try {
      const authUser = await api.getUserOwnProfile()
      dispatch(setAuthUser(authUser))
    } catch (error) {
      dispatch(setAuthUser(null))
    } finally {
      dispatch(setIsPreload(false))
    }
    dispatch(hideLoading())
  }
}

export { actionType, setIsPreload, asyncPreloadProcess }
