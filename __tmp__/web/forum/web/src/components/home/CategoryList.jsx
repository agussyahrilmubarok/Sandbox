import React from 'react'
import useSelectedCategory from '../../hooks/useSelectedCategory'
import SkeletonCategoryList from '../skeleton/SkeletonCategoryList'
import CategoryItem from './CategoryItem'
import PropTypes from 'prop-types'

export default function CategoriesList({ categories }) {
  const { selectedCategory, setSelectedCategory } = useSelectedCategory()
  return (
    <>
      <h1 className="text-xl font-medium">Popular Categories</h1>
      <div className="flex flex-wrap my-4 gap-3">
        {categories.length ? (
          categories.map((category, i) => (
            <CategoryItem
              key={i}
              category={category}
              isSelected={selectedCategory === category}
              clickHandler={() => setSelectedCategory(category)}
            />
          ))
        ) : (
          <SkeletonCategoryList />
        )}
      </div>
    </>
  )
}

CategoriesList.propTypes = {
  categories: PropTypes.arrayOf(PropTypes.string).isRequired
}
