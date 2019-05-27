import React from 'react';
import './field.scss';

function Field({ name, label, type, onChange, onFocus, onBlur, placeholder }) {
  const [isFocused, setIsFocused] = React.useState(false);

  const handleChange = (event) => {
    if (onChange) {
      onChange(event);
    }
  }

  const handleFocus = (event) => {
    if (onFocus) {
      onFocus(event);
    }

    setIsFocused(true);
  }

  const handleBlur = (event) => {
    if (onBlur) {
      onBlur(event);
    }

    setIsFocused(false);
  }

  let className = `input-field${isFocused ? ' focused' : ''}`;

  return (
    <div className={className}>
      <label htmlFor={name}>{label}</label>
      <input
        type={type} 
        name={name}
        placeholder={placeholder}
        onChange={handleChange}
        onFocus={handleFocus}
        onBlur={handleBlur}
      />
    </div>
  );
}

export default Field;
