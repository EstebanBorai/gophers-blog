const URL = API_URL.concat('api/users');

export async function getUsers() {
  return await fetch(URL).then((res) => res.json()).catch((err) => {
    throw new Error({
      Error: err
    });
  });
}

export async function login({ userName, password }) {
  return await fetch(URL.concat('/login'), {
    method: 'POST',
    body: JSON.stringify({
      userName,
      password
    })
  }).then((res) => res.json())
    .catch((err) => {
      throw new Error({
        Error: err
      });
    });
}
