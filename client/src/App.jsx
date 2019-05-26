import React from 'react';
import Login from './components/Login';
import Main from './components/Main';

function App() {
  const loggedUser = () => {
    if (localStorage.getItem('songs-share:user') !== null) {
      return <Main />;
    } else {
      return <Login />;
    }
  }

  return loggedUser();
}

export default App;
