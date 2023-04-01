import { useEffect, useRef, useState } from 'react';
import { Link } from 'react-router-dom';
import { deleteUser, getUser } from '../../external/user';
import { ListUser } from './ListUser';

export const User = () => {
  const [userDataContext, setUserDataContext] = useState([]);
  const userRef = useRef();

  const fetchData = async () => {
    try {
      const response = await getUser();
      const user = await response.json();
      userRef.current = user;
      setUserDataContext(userRef.current);
    } catch (error) {
      console.log(error);
    }
  };

  useEffect(() => {
    fetchData();
  }, [userRef]);

  const editEvent = (e) => {
    console.log('edit', e);
  };

  const deleteEvent = async (e) => {
    const del = await deleteUser(e);
    console.log('delete', e, del);
  };

  return (
    <>
      <div>
        <Link to="/user/create">Create</Link>
      </div>
      <div>
        <ListUser userData={userDataContext} editEvent={editEvent} deleteEvent={deleteEvent} />
      </div>
    </>
  );
};
