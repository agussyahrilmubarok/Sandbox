import { useState } from 'react'

export default function useInput(initialValue = '') {
  const [value, setValue] = useState(initialValue)

  const onChange = (e) => {
    e.target.value ? setValue(e.target.value) : setValue(e.target.innerHTML)
  }

  return [value, onChange]
}
