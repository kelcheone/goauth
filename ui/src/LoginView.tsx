import { useState } from "react";
const LoginView = () => {
  const [form, setForm] = useState({
    email: "",
    password: "",
  });

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setForm({ ...form, [e.target.id]: e.target.value });
  };
  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    console.log(form);

    const url = "http://localhost:5000/auth/login";
    const response = await fetch(url, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      credentials: "include",
      body: JSON.stringify(form),
    });
    const status = await response.status;
    if (status === 200) {
      console.log("logged in");
      window.location.href = "/";
    } else {
      alert("Wrong email or password!");
    }
  };

  return (
    // a form with a email and password field (tailwindcss)
    <div className="flex items-center justify-center h-screen">
      {/* the form */}
      <form
        className="bg-gray-200 shadow-md rounded px-8 pt-6 pb-8 mb-4"
        onSubmit={handleSubmit}
      >
        <div className="mb-4">
          {/* email field */}
          <label
            className="block text-gray-700 text-sm font-bold mb-2"
            htmlFor="email"
          >
            Email
          </label>
          <input
            className="bg-gray-700 shadow appearance-none border rounded w-full py-2 px-3 text-gray-200 leading-tight focus:outline-none focus:shadow-outline"
            id="email"
            type="text"
            placeholder="Email"
            onChange={handleChange}
          />
          {/* password field */}
          <label
            className="block text-gray-700 text-sm font-bold mb-2"
            htmlFor="password"
          >
            Password
          </label>
          <input
            className="bg-gray-700 shadow appearance-none border rounded w-full py-2 px-3 text-gray-200 mb-3 leading-tight focus:outline-none focus:shadow-outline"
            id="password"
            type="password"
            placeholder="Password"
            onChange={handleChange}
          />
          {/* submit button */}
          <button
            className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
            type="submit"
          >
            Sign In
          </button>

          <p className="text-center text-gray-500 text-xs">
            &copy;2021 Acme Corp. All rights reserved.
          </p>
        </div>
      </form>
    </div>
  );
};

export default LoginView;
