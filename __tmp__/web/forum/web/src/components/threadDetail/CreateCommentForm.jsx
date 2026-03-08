import React, { useRef } from 'react'
import { useSelector, useDispatch } from 'react-redux'
import { asyncAddThreadDetailComment } from '../../states/threadDetail/action'
import { Link } from 'react-router-dom'
import useInput from '../../hooks/useInput'
import PropTypes from 'prop-types'

export default function CreateCommentForm({ threadId }) {
  const authUser = useSelector((state) => state.authUser)
  const dispatch = useDispatch()

  const [content, onContentChange] = useInput()
  const contentRef = useRef()

  const onAddComment = () => {
    dispatch(asyncAddThreadDetailComment({ threadId, content }))
    contentRef.current.innerHTML = ''
  }

  if (!authUser) {
    return (
      <div className="bg-base-200 text-sm p-3">
        <p>
          You need to login first to post a comment.{' '}
          <Link to="/login">
            <button className="text-primary hover:underline focus:outline-none">Login</button>
          </Link>
        </p>
      </div>
    )
  }

  return (
    <div className="flex flex-col gap-y-4 p-4 rounded-md bg-base-200">
      <div className="flex gap-x-2.5 items-center text-neutral-content">
        <img className="rounded-full size-8" src={authUser.avatar} alt={authUser.name} />
        <p className="text-sm">{authUser.name}</p>
      </div>
      <div
        className="w-full min-h-24 p-2 rounded-md bg-base-100 outline-none"
        ref={contentRef}
        onInput={onContentChange}
        contentEditable
        suppressContentEditableWarning
      ></div>
      <button onClick={onAddComment} className="btn btn-sm btn-primary w-1/3">
        Post
      </button>
    </div>
  )
}

CreateCommentForm.propTypes = {
  threadId: PropTypes.string.isRequired
}
