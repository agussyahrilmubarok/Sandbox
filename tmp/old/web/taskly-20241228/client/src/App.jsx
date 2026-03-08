import { BrowserRouter, Routes, Route } from "react-router-dom";
import { ChakraProvider } from "@chakra-ui/react";
import { Toaster } from "react-hot-toast";

import { UserProvider } from "./context/UserContext";
import PrivateRoute from "./components/PrivateRoute";
import Home from "./pages/Home";
import SignUp from "./pages/SignUp";
import SignIn from "./pages/SignIn";
import Profile from "./pages/users/Profile";
import NavBar from "./components/NavBar";
import CreateTask from "./pages/tasks/CreateTask";
import UpdateTask from "./pages/tasks/UpdateTask";
import Tasks from "./pages/tasks/Tasks";
import SingleTask from "./pages/tasks/SingleTask";

export default function App() {
  return (
    <UserProvider>
      <ChakraProvider>
        <BrowserRouter>
          <NavBar />
          <Toaster position="top-center" />

          <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/signup" element={<SignUp />} />
            <Route path="/signin" element={<SignIn />} />

            <Route element={<PrivateRoute />}>
              <Route path="/profile" element={<Profile />} />
              <Route path="/create-task" element={<CreateTask />} />
              <Route path="/update-task/:taskId" element={<UpdateTask />} />
              <Route path="/tasks" element={<Tasks />} />
              <Route path="/tasks/:taskId" element={<SingleTask />} />
            </Route>
          </Routes>
        </BrowserRouter>
      </ChakraProvider>
    </UserProvider>
  );
}
