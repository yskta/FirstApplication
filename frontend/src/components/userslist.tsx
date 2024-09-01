import { useQuery } from "@apollo/client";
import { GET_USERS } from "../graphql/queries";

export default function UsersList() {
  const { loading, error, data } = useQuery(GET_USERS);

  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error: {error.message}</p>;

  return (
    <ul>
        {data.map((user: any) => (
            <li key={user.id}>
                {user.username} - {user.email}
            </li>
        ))}
    </ul>
  );
}