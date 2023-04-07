import { useEffect, useRef, useState } from 'react';
import { Link } from 'react-router-dom';
import { deleteUser, getUser, getUsers } from '../../external/user';
import { ListUser } from './ListUser';

export const User = () => {
  const [userDataContext, setUserDataContext] = useState([]);
  const userRef = useRef();

  const fetchData = async () => {
    try {
      const response = await getUsers();
      const user = await response.json();
      console.log(user);
      userRef.current = user;
      setUserDataContext(userRef.current);
    } catch (error) {
      console.log(error);
    }
  };

  useEffect(() => {
    fetchData();
  }, [userRef]);

  return (
    <>
      <div>
        <Link to="/user/create">Create</Link>
      </div>
      <div>
        <ListUser userData={userDataContext} />
      </div>
    </>
  );
};
