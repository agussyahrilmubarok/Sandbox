import React from 'react'
import { cn } from '../../lib/utils'
import PropTypes from 'prop-types'

export default function CategoryItem({ isSelected, clickHandler, category }) {
  return (
    <span
      className={cn('badge badge-lg cursor-pointer', isSelected ? 'badge-primary' : 'badge-outline')}
      onClick={clickHandler}
    >
      {'#'}
      {category}
    </span>
  )
}

CategoryItem.propTypes = {
  isSelected: PropTypes.bool.isRequired,
  clickHandler: PropTypes.func.isRequired,
  category: PropTypes.string.isRequired
}
