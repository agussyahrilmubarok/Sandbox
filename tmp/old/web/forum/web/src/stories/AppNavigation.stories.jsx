import React from 'react'
import Wrapper from './Wrapper'
import AppNavigation from '../components/app/AppNavigation'

const stories = {
  title: 'AppNavigation',
  component: AppNavigation
}

const TemplateStory = (args) => (
  <Wrapper states={{ authUser: args.login }}>
    <AppNavigation />
  </Wrapper>
)

const LoggedIn = TemplateStory.bind({})
const LoggedOut = TemplateStory.bind({})

LoggedIn.args = {
  login: true
}

LoggedOut.args = {
  login: false
}

export { LoggedIn, LoggedOut }
export default stories
