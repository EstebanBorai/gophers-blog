import React from 'react';
import { login } from '../../api/User';

function Login() {
  const [loginForm, setLoginForm] = React.useState({
    userName: '',
    password: ''
  });

  const handleChange = (event) => {
    const { target: { value, name } } = event;
    setLoginForm({
      ...loginForm,
      [name]: value
    });
  };

  const handleSubmit = async (event) => {
    event.preventDefault();

    const user = await login(loginForm);
    localStorage.setItem('songs-share:user', JSON.stringify(user));
  }

  return (
    <section>
      <h3>
        Welcome {loginForm.userName ? 'back' : 'to'}&nbsp;{loginForm.userName || 'Songs Share'}
      </h3>
      <form onSubmit={handleSubmit}>
        <input type="text" name="userName" onChange={handleChange} />
        <input type="password" name="password" onChange={handleChange} />
        <button type="submit">Login</button>
      </form>
    </section>
  );
}

export default Login;
