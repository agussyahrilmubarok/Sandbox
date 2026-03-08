import React, { useEffect } from 'react'
import { useSelector, useDispatch } from 'react-redux'
import useSelectedCategory from '../hooks/useSelectedCategory'
import { asyncPopulateData } from '../states/shared/action'
import ThreadList from '../components/home/ThreadList'
import CategoryList from '../components/home/CategoryList'
import { Helmet } from 'react-helmet'

export default function HomePage() {
  const { threads, users, categories, authUser } = useSelector((state) => state)

  const dispatch = useDispatch()
  const { selectedCategory } = useSelectedCategory()

  useEffect(() => {
    dispatch(asyncPopulateData())
  }, [dispatch])

  const threadList = threads.map((thread) => ({
    ...thread,
    owner: users.find((user) => user.id === thread.ownerId),
    authUser
  }))

  const filteredThreadList = threadList.filter((thread) => thread.category === selectedCategory)
  const threadsToRender = filteredThreadList.length ? filteredThreadList : threadList

  return (
    <>
      <Helmet>
        <title>Forum - Where Ideas Gather</title>
      </Helmet>
      <section className="mt-8 pb-24">
        <CategoryList categories={categories} />
        <div className="divider my-6"></div>
        <ThreadList threads={threadsToRender} />
      </section>
    </>
  )
}
