import React from 'react'
import PropTypes from 'prop-types'

export default function Input({ label, type, name, value, onChange, placeholder, required }) {
  return (
    <label className="form-control w-full">
      <div className="label">
        <span className="label-text">
          {label}
          {required && <span className="text-error"> *</span>}
        </span>
      </div>
      <input
        name={name}
        type={type}
        placeholder={placeholder}
        className="input input-bordered w-full"
        value={value}
        onChange={onChange}
        required={required}
      />
    </label>
  )
}

Input.propTypes = {
  label: PropTypes.string.isRequired,
  type: PropTypes.string.isRequired,
  name: PropTypes.string.isRequired,
  value: PropTypes.string.isRequired,
  onChange: PropTypes.func.isRequired,
  required: PropTypes.bool,
  placeholder: PropTypes.string
}
