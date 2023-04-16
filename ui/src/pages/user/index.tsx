import { UserList } from "./list";

export const getServerSideProps = () => {
  return {
    props: { name: "User" },
  };
};

const User = (): JSX.Element => {
  return <UserList />;
};

export default User;
