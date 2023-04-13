import { UserList } from "./list";

export const getServerSideProps = async() => {
  return {
    props: {name: 'User'},
  }
}

const User = (a: any): JSX.Element => {
  return <UserList/>;
};

export default User;