import React from 'react'
import SkeletonThreadList from '../skeleton/SkeletonThreadList'
import ThreadItem from './ThreadItem'
import PropTypes from 'prop-types'

export default function ThreadList({ threads }) {
  return (
    <div className="space-y-3">
      {threads.length ? threads.map((thread, i) => <ThreadItem key={i} {...thread} />) : <SkeletonThreadList />}
    </div>
  )
}

const userShape = {
  id: PropTypes.string.isRequired,
  name: PropTypes.string.isRequired,
  avatar: PropTypes.string.isRequired
}

const threadShape = {
  id: PropTypes.string.isRequired,
  title: PropTypes.string.isRequired,
  body: PropTypes.string.isRequired,
  createdAt: PropTypes.string.isRequired,
  upVotesBy: PropTypes.arrayOf(PropTypes.string).isRequired,
  downVotesBy: PropTypes.arrayOf(PropTypes.string).isRequired,
  totalComments: PropTypes.number.isRequired,
  owner: PropTypes.shape(userShape).isRequired,
  authUser: PropTypes.shape(userShape)
}

ThreadList.propTypes = {
  threads: PropTypes.arrayOf(PropTypes.shape(threadShape)).isRequired
}
