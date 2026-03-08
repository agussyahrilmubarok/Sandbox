import React from 'react'

export default function SkeletonThreadDetail() {
  return (
    <div className="space-y-4 mt-8">
      <div className="skeleton w-full min-h-48 rounded-md bg-base-100/30"></div>
      <div className="skeleton w-full min-h-12 rounded-md bg-base-100/30"></div>
      <div className="skeleton w-full min-h-24 rounded-md bg-base-100/30"></div>
      <div className="skeleton w-full min-h-24 rounded-md bg-base-100/30"></div>
    </div>
  )
}
