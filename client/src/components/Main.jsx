import React from 'react';
import { getUsers } from '../api/User';

function Main() {
  const [users, setUsers] = React.useState([]);
  const [networkError, setNetworkError] = React.useState(null);

  React.useEffect(() => {
    const fetchData = async () => {
      try {
        setUsers(await getUsers());
      } catch (err) {
        setNetworkError(err);
      }
    }

    fetchData();
  }, []);

  return (
    <main>
      {
        networkError ? 
        <p>
          {JSON.stringify(networkError)}
        </p> :
        <ul>
          {users && users.map((user, index) => (
            <li key={index}>{user.firstName} - {user.lastName} - {user.userName}</li>
          ))}
        </ul>
      }
    </main>
  );
}

export default Main;
