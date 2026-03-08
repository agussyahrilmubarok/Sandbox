const actionType = {
  RECEIVE_CATEGORIES: 'RECEIVE_CATEGORIES',
  ADD_CATEGORY: 'ADD_CATEGORY'
}

function receiveCategories(categories) {
  return {
    type: actionType.RECEIVE_CATEGORIES,
    payload: {
      categories
    }
  }
}

function addCategory(category) {
  return {
    type: actionType.ADD_CATEGORY,
    payload: {
      category
    }
  }
}

export { actionType, receiveCategories, addCategory }
