import React from 'react';
import './button.scss';

function Button({ children, color }) {

  const className = `btn-el${color ? ` ${color}` : ''}`;

  return (
    <button className={className}>
      {children}
    </button>
  );
}

export default Button;
