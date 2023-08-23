import axios from "axios";
import { useEffect, useState } from "react";

type user = {
  name: string;
  email: string;
};

const HomeView = () => {
  const [user, setUser] = useState<user | null>(null);
  useEffect(() => {
    const url = "http://localhost:5000/user";
    axios.get(url, { withCredentials: true }).then((res) => {
      if (res.status === 200) {
        setUser(res.data.data);
      }
    });
  }, []);

  const handleLogout = async () => {
    const url = "http://localhost:5000/auth/logout";
    const response = await fetch(url, {
      method: "POST",
      credentials: "include",
    });
    const status = await response.status;
    if (status === 200) {
      console.log("logged out");
      window.location.href = "/";
    } else {
      alert("Something went wrong!");
    }
  };

  return (
    <div className="bg-gray-100 h-screen text-black">
      <div className="flex flex-col items-center justify-center h-full">
        <div className="bg-white rounded-lg shadow-lg p-5 md:p-20 mx-2">
          <div className="text-center">
            <h2 className="text-3xl font-bold">Welcome {user?.name}</h2>
            {user ? (
              <div>
                <p className="mt-2">Email: {user?.email}</p>
                <p className="mt-2">You are logged in!</p>
                <button
                  className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
                  onClick={handleLogout}
                >
                  Logout
                </button>
              </div>
            ) : (
              <a href="/login" className="mt-2 text-blue-600 hover:underline">
                Login
              </a>
            )}
          </div>
        </div>
      </div>
    </div>
  );
};

export default HomeView;
