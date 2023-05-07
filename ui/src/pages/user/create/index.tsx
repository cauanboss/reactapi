export const getStaticProps = () => {
  return {
    props: { title: "My Title", content: "...", path: "/user/create" },
  };
};

const UserCreate = (): JSX.Element => {
  return (
    <div className="table-auto">
      <div className="row-auto">
        Name <input type="text" style={{ width: "100px" }} />
      </div>
      <div className="row-auto">
        Login <input type="text" />
      </div>
      <div className="row-auto">
        Password <input type="text" />
      </div>
      <div className="">
        <button>Save</button>
      </div>
    </div>
  );
};

export default UserCreate;
