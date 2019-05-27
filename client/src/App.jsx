import React from 'react';
import Login from './components/Login';
import Main from './components/Main';
import SignUp from './components/SignUp';

function App() {
  const loggedUser = () => {
    if (localStorage.getItem('songs-share:user') !== null) {
      return <Main />;
    } else {
      return <Login />;
    }
  }

  return <SignUp />;
}

export default App;
