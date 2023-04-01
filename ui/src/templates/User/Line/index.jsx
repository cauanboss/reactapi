import Prop from 'prop-types';
import { Button } from '../../../components/Button';

export const LineUser = ({ user, editEvent, deleteEvent }) => {
  return (
    <div
      key={user.id}
      style={{ display: 'flex', margin: '5px', gridAutoColumns: 'column', gridColumnGap: '10px' }}
    >
      <div>{user.name}</div>
      <div>{user.gender}</div>
      <div>{user.age}</div>
      <div>
        <Button
          onClick={() => {
            editEvent(user.id);
          }}
        >
          Editar
        </Button>
      </div>
      <div>
        <Button
          onClick={() => {
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
