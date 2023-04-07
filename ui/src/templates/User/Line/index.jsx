import Prop from 'prop-types';
import { useRef } from 'react';
import { Button } from '../../../components/Button';
import { deleteUser, getUser } from '../../../external/user';

export const LineUser = ({ user }) => {
  const ref = useRef(user);
  const editEvent = async (e) => {
    const response = await getUser(e);
    const user = await response.json();
    console.log('edit', user);
  };

  const deleteEvent = async (e) => {
    const del = await deleteUser(e);
    console.log('delete', e, del);
  };

  return (
    <div
      ref={ref}
      key={user.id}
      style={{ display: 'flex', margin: '5px', gridAutoColumns: 'column', gridColumnGap: '10px' }}
    >
      <div>{user.name}</div>
      <div>{user.gender}</div>
      <div>{user.age}</div>
      <div>
        <Button
          onClick={(e) => {
            editEvent(user.id);
          }}
        >
          Editar
        </Button>
      </div>
      <div>
        <Button
          onClick={(e) => {
            deleteEvent(user.id);
          }}
        >
          Deletar
        </Button>
      </div>
    </div>
  );
};

LineUser.propTypes = {
  user: Prop.object.isRequired,
  editEvent: Prop.func.isRequired,
  deleteEvent: Prop.func.isRequired,
};
