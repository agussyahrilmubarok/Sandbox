import React from 'react'
import { useSelector, useDispatch } from 'react-redux'
import { useNavigate } from 'react-router-dom'
import useInput from '../hooks/useInput'
import { asyncCreateThreadAndCategory } from '../states/shared/action'

export default function CreateThreadForm() {
  const categories = useSelector((state) => state.categories)
  const dispatch = useDispatch()
  const navigate = useNavigate()

  const [title, onTitleChange] = useInput()
  const [body, onBodyChange] = useInput()
  const [category, onCategoryChange] = useInput()

  const submitHandler = (ev) => {
    ev.preventDefault()
    dispatch(asyncCreateThreadAndCategory({ title, body, category })).then((res) => {
      !res.error && navigate('/')
    })
  }

  return (
    <form className="space-y-4 relative" onSubmit={submitHandler}>
      <input
        type="text"
        className="input text-xl font-semibold input-ghost rounded-none w-full border-0 focus:outline-0 px-0.5 py-0 focus:border-b focus:border-primary"
        placeholder="An Interest Title"
        value={title}
        onChange={onTitleChange}
        required
      />
      <input
        type="text"
        placeholder="A Cool Category"
        className="input input-xs input-ghost rounded-none w-full border-0 focus:outline-0 px-0.5 py-0 focus:border-b focus:border-primary"
        list="categories"
        value={category}
        onChange={onCategoryChange}
        required
      />
      <datalist id="categories">
        {categories.map((category, i) => {
          return <option key={i} value={category} />
        })}
      </datalist>
      <div
        className="min-h-28 outline outline-1 outline-neutral-content/10 rounded-lg text-neutral-content/70 overflow-auto p-1.5 focus:outline-primary"
        onInput={onBodyChange}
        contentEditable
        suppressContentEditableWarning
      ></div>
      <button type="submit" className="btn btn-primary btn-sm text-xs px-8 mt-4 absolute right-0 -top-14">
        Post
      </button>
    </form>
  )
}
