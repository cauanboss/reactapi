import { LineUser } from '../Line';

export const ListUser = ({ userData, editEvent, deleteEvent }) => {
  return userData.map((user) => LineUser({ user, editEvent, deleteEvent }));
};
