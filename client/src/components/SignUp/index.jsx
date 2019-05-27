import React from 'react';
import './sign-up.scss';
import Field from '../lib/Field';
import Button from '../lib/Button';

function SignUp() {
  const [formValues, setFormValues] = React.useState({
    firstName: '',
    lastName: '',
    userName: '',
    password: ''
  });

  const [isFormActive, setIsFormActive] = React.useState(false);

  const handleChange = (event) => {
    const { target: { value, name } } = event;

    setFormValues({
      ...formValues,
      [name]: value
    });
  }

  const handleMouseEnter = (event) => {
    setIsFormActive(true);
  }

  const formClassName = isFormActive ? 'active' : '';

  return (
    <section className="sign-up">
      <form
        className={formClassName}
        onMouseEnter={handleMouseEnter}
      >
        {
          isFormActive ?
          <h3>Welcome to Songs Share</h3> :
          null
        }
        <Field 
          name="firstName"
          label="Name"
          type="text"
          onChange={handleChange}
          placeholder="John"
        />
        <Field 
          name="lastName"
          label="Last"
          type="text"
          onChange={handleChange}
          placeholder="Doe"
        />
        <Field 
          name="userName"
          label="Username"
          type="text"
          onChange={handleChange}
          placeholder="johndoe"
        />
        <Field 
          name="password"
          label="Password"
          type="password"
          onChange={handleChange}
          placeholder="Your little secret"
        />
        <div className="center">
          <Button color="primary">Sign Up</Button>
          <Button>Log In</Button>
        </div>
      </form>
    </section>
  );
}

export default SignUp;
