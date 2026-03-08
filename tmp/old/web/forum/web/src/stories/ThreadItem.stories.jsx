import React from 'react'
import Wrapper from './Wrapper'
import ThreadItem from '../components/home/ThreadItem'

const stories = {
  title: 'ThreadItem',
  component: ThreadItem
}

const TemplateStory = (args) => (
  <Wrapper>
    <div className="max-w-xl">
      <ThreadItem {...args} />
    </div>
  </Wrapper>
)

const withProps = TemplateStory.bind({})

withProps.args = {
  id: 'thread-1',
  title: 'Thread Title',
  category: 'General',
  body: 'lorem ipsum dolor sit amet',
  createdAt: '2024-03-01',
  upVotesBy: [],
  downVotesBy: [],
  totalComments: 0,
  owner: { name: 'User 1' },
  authUser: { name: 'User 1' }
}

export { withProps }
export default stories
