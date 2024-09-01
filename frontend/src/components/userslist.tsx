import { useQuery } from "@apollo/client";
import { GET_USERS } from "../graphql/queries";

export default function UsersList() {
  const { loading, error, data } = useQuery(GET_USERS);

  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error: {error.message}</p>;
  console.log(data.users);
  return (
    <ul>
      {data.users.map((user: { id: string; name: string; email: string }) => (
        <li key={user.id}>
          {user.name} - {user.email}
        </li>
      ))}
    </ul>
  );
}