import { LineUser } from '../Line';

export const ListUser = ({ userData }) => {
  return userData.map((user) => LineUser({ user }));
};
