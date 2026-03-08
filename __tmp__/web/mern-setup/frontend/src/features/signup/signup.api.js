const API_URL = import.meta.env.VITE_BACKEND_URL || "http://localhost:3001/api/v1";

const signUpApi = async (user) => {
  try {
    const response = await fetch(`${API_URL}/auth/signup`, {
      method: "POST",
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
      },
      credentials: "include",
      body: JSON.stringify(user),
    });

    if (!response.ok) {
      const errorData = await response.json();
      throw new Error(errorData.message || "Sign-Up failed");
    }

    return await response.json();
  } catch (err) {
    console.error("Sign-Up error:", err.message);
    return { error: err.message };
  }
};

export default signUpApi;
