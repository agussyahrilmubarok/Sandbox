import React from 'react'
import { Link } from 'react-router-dom'
import PropTypes from 'prop-types'

export default function NavigationButton({ title, path, icon, onClick }) {
  if (path) {
    return (
      <Link to={path}>
        <button className="btn btn-square btn-ghost tooltip flex lg:tooltip-right" data-tip={title}>
          {icon}
        </button>
      </Link>
    )
  }

  return (
    <button onClick={onClick} className="btn btn-square btn-ghost tooltip flex lg:tooltip-right" data-tip={title}>
      {icon}
    </button>
  )
}

NavigationButton.propTypes = {
  title: PropTypes.string.isRequired,
  icon: PropTypes.node.isRequired,
  path: PropTypes.string,
  onClick: PropTypes.func
}
