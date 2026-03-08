import { actionType } from './action'

export default function categoriesReducer(categories = [], action = {}) {
  switch (action.type) {
    case actionType.RECEIVE_CATEGORIES:
      return action.payload.categories
    case actionType.ADD_CATEGORY:
      return [action.payload.category, ...categories]
    default:
      return categories
  }
}
