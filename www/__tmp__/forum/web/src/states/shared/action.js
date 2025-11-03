import { api } from '../../lib/api'
import toast from 'react-hot-toast'
import { addThread, receiveThreads } from '../threads/action'
import { addCategory, receiveCategories } from '../categories/action'
import { receiveUsers } from '../users/action'
import { showLoading, hideLoading } from 'react-redux-loading-bar'

function asyncPopulateData() {
  return async (dispatch) => {
    dispatch(showLoading())
    try {
      const threads = await api.getAllThreads()
      const users = await api.getAllUsers()
      const categories = threads.map((thread) => thread.category)

      dispatch(receiveCategories([...new Set(categories)]))
      dispatch(receiveThreads(threads))
      dispatch(receiveUsers(users))
    } catch (error) {
      toast.error(error.message)
    }
    dispatch(hideLoading())
  }
}

function asyncCreateThreadAndCategory({ title, body, category }) {
  return async (dispatch, getState) => {
    dispatch(showLoading())
    const { categories } = getState()
    const isCategoryExist = categories.includes(category)
    try {
      const newThread = await api.createThread({ title, body, category })
      if (!isCategoryExist) {
        dispatch(addCategory(newThread.category))
      }
      dispatch(addThread(newThread))
      toast.success('Thread has been posted')
      dispatch(hideLoading())
      return { error: false }
    } catch (error) {
      toast.error(error.message)
      dispatch(hideLoading())
      return { error: true }
    }
  }
}

export { asyncPopulateData, asyncCreateThreadAndCategory }
