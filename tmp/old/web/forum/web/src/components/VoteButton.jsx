import React from 'react'
import { ChevronDownIcon, ChevronUpIcon } from 'lucide-react'
import { cn } from '../lib/utils'
import PropTypes from 'prop-types'

export default function VoteButton({ type, onClick, isVoted, voteCount, disabled }) {
  const voteTypes = {
    up: <ChevronUpIcon size={24} />,
    down: <ChevronDownIcon size={24} />
  }

  return (
    <button
      disabled={disabled}
      className={cn(
        'btn btn-ghost btn-square min-w-0 min-h-0 h-full flex justify-center items-center gap-x-0.5 font-semibold px-0',
        isVoted && 'text-primary'
      )}
      onClick={onClick}
    >
      {voteTypes[type]}
      {voteCount}
    </button>
  )
}

VoteButton.propTypes = {
  type: PropTypes.oneOf(['up', 'down']).isRequired,
  onClick: PropTypes.func.isRequired,
  isVoted: PropTypes.bool.isRequired,
  voteCount: PropTypes.number.isRequired,
  disabled: PropTypes.bool
}
