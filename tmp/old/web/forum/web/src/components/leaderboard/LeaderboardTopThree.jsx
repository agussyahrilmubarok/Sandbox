import React from 'react'
import { useSelector } from 'react-redux'
import { cn } from '../../lib/utils'
import SkeletonTopThree from '../skeleton/SkeletonTopThree'
import PropTypes from 'prop-types'

export default function LeaderboardTopThree({ leaderboards }) {
  const authUser = useSelector((state) => state.authUser)

  return (
    <>
      <ol className="flex justify-around items-end gap-x-3">
        {leaderboards.length ? (
          leaderboards.map(({ user, score }, i) => (
            <li
              key={i}
              className={cn(
                'flex flex-col items-center text-center rounded-md bg-base-200 py-5 px-2 w-1/3 h-fit',
                i === 0 && 'py-12',
                i === 1 && 'order-first'
              )}
            >
              <img className="avatar size-18 rounded-full" src={user.avatar} alt={user.name} />
              <span className="block relative bg-neutral size-6 text-neutral-content m-auto -top-3 rounded-full">
                {i + 1}
              </span>
              <p className={cn('text-neutral-content mb-0.5', user.id === authUser?.id && 'text-primary font-medium')}>
                {user.name}
              </p>
              <span className="font-bold">{score} pts</span>
            </li>
          ))
        ) : (
          <SkeletonTopThree />
        )}
      </ol>
    </>
  )
}

const userShape = {
  id: PropTypes.string.isRequired,
  name: PropTypes.string.isRequired,
  avatar: PropTypes.string.isRequired
}

LeaderboardTopThree.propTypes = {
  leaderboards: PropTypes.arrayOf(
    PropTypes.shape({
      user: PropTypes.shape(userShape).isRequired,
      score: PropTypes.number.isRequired
    })
  )
}
