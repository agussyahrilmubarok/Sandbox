import React from 'react'
import CommentItem from './CommentItem'
import PropTypes from 'prop-types'

export default function CommentList({ threadId, comments }) {
  return (
    <>
      <h2 className="ml-2.5 text-xl font-bold">Comments[{comments.length}]</h2>
      <div>
        {comments.map(({ id, owner, content, createdAt, upVotesBy, downVotesBy }) => (
          <CommentItem
            key={id}
            threadId={threadId}
            commentId={id}
            owner={owner}
            content={content}
            createdAt={createdAt}
            upVotesBy={upVotesBy}
            downVotesBy={downVotesBy}
          />
        ))}
      </div>
    </>
  )
}

const ownerShape = {
  id: PropTypes.string.isRequired,
  name: PropTypes.string.isRequired,
  avatar: PropTypes.string.isRequired
}

const commentShape = {
  id: PropTypes.string.isRequired,
  owner: PropTypes.shape(ownerShape).isRequired,
  content: PropTypes.string.isRequired,
  createdAt: PropTypes.string.isRequired,
  upVotesBy: PropTypes.arrayOf(PropTypes.string).isRequired,
  downVotesBy: PropTypes.arrayOf(PropTypes.string).isRequired
}

CommentList.propTypes = {
  threadId: PropTypes.string.isRequired,
  comments: PropTypes.arrayOf(PropTypes.shape(commentShape)).isRequired
}
