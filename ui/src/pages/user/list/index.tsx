import { Router, useRouter } from "next/router";

export const UserList = (): JSX.Element => {
  const route = useRouter();
  const handleCreate = () => {
    route.push("/user/create");
  };

  return (
    <div className="table-auto">
      <div className="table-row">
        <button onClick={handleCreate}>Create</button>
      </div>
    </div>
  );
};
