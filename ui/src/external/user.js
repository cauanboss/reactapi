export const getUsers = async () => {
  const response = await fetch('http://localhost/api/v1/user', {
    method: 'get',
    mode: 'no-cors',
    headers: {
      'Content-Type': 'application/json',
      Accept: 'application/json',
      'Access-Control-Allow-Origin': '*',
      'Access-Control-Allow-Headers': 'Content-Type',
    },
  });
  return response;
};

export const getUser = async (id) => {
  const response = await fetch(`http://localhost/api/v1/user/${id}`, {
    method: 'get',
    mode: 'no-cors',
    headers: {
      'Content-Type': 'application/json',
      Accept: 'application/json',
      'Access-Control-Allow-Origin': '*',
      'Access-Control-Allow-Headers': 'Content-Type',
    },
  });
  return response;
};

export const deleteUser = async (id) => {
  const response = await fetch(`http://localhost/api/v1/user/${id}`, {
    method: 'delete',
    // mode: 'no-cors',
    headers: {
      'Content-Type': 'application/json',
      Accept: 'application/json',
      'Access-Control-Allow-Origin': '*',
      'Access-Control-Allow-Headers': 'Content-Type',
    },
  });
  return response;
};
