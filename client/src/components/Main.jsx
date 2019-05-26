import React from 'react';

function Main() {
  const [users, setUsers] = React.useState([]);

  React.useEffect(() => {
    const response = fetch('http://localhost:8080/api/users').then((res) => res.json()).then((u) => {
      setUsers(u);
    });
  }, []);

  return (
    <main>
      <ul>
        {users && users.map((user, index) => (
          <li key={index}>{user.firstName} - {user.lastName} - {user.userName}</li>
        ))}
      </ul>
    </main>
  );
}

export default Main;
