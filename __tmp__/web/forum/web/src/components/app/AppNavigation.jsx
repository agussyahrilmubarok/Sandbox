import React from 'react'
import { useSelector, useDispatch } from 'react-redux'
import { HomeIcon, BarChart2Icon, LogInIcon, LogOutIcon, SquarePenIcon } from 'lucide-react'
import NavigationButton from '../NavigationButton'
import { asyncUnsetAuthUser } from '../../states/authUser/action'

export default function AppNavigation() {
  const authUser = useSelector((state) => state.authUser)
  const dispatch = useDispatch()

  const navigationList = [
    {
      title: 'Home',
      path: '/',
      icon: <HomeIcon size={22} />
    },
    {
      title: 'Leaderboards',
      path: '/leaderboards',
      icon: <BarChart2Icon size={22} />
    },
    {
      title: 'Login',
      path: '/login',
      icon: <LogInIcon size={22} />
    },
    {
      title: 'Create Thread',
      path: '/threads/new',
      icon: <SquarePenIcon size={22} />
    },
    {
      title: 'Logout',
      clickHandler: () => dispatch(asyncUnsetAuthUser()),
      icon: <LogOutIcon size={22} />
    }
  ]

  const filteredNavigationList = navigationList.filter((nav) => {
    const excludeList = authUser ? ['Login'] : ['Create Thread', 'Logout']
    return !excludeList.includes(nav.title)
  })

  return (
    <nav className="navbar fixed w-full py-1 bg-base-100 min-h-3 bottom-0 right-0 left-0 backdrop-blur-lg border-t border-neutral lg:border-t-0 lg:border-r lg:w-16 lg:h-full lg:p-0">
      <div className="gap-6 m-auto lg:flex-col lg:items-center lg:justify-center">
        {filteredNavigationList.map(({ title, path, icon, clickHandler }, i) => {
          return <NavigationButton key={i} title={title} icon={icon} path={path} onClick={clickHandler} />
        })}
      </div>
    </nav>
  )
}
