import { useSearchParams } from 'react-router-dom'

export default function useSelectedCategory() {
  const [searchParams, setSearchParams] = useSearchParams()
  const selectedCategory = searchParams.get('category')
  const setSelectedCategory = (category) => {
    if (category === selectedCategory) return setSearchParams({})
    setSearchParams({ category })
  }

  return { selectedCategory, setSelectedCategory }
}
